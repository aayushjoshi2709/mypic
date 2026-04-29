import { routes } from "../../../common/routes";
import {
  faDownload,
  faEye,
  faTrashAlt,
} from "@fortawesome/free-solid-svg-icons";
import { useDispatch, useSelector } from "react-redux";
import { clearModal, setModal } from "../../../store/modal.slice";
import { ImageLoadTypes, ModalNames } from "../../../common/Constants";
import { apiClientObj } from "../../../common/apiClient";
import toast from "react-hot-toast";
import type { ImageInterface } from "../../../common/interfaces";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  removeCurrentGroupImage,
  setCurrentGroupImage,
} from "../../../store/group.slice";
import type { RootState } from "../../../store/store";

interface ImageViewButtonsInterface {
  imgData: ImageInterface;
}

const GroupViewButtons = ({ imgData }: ImageViewButtonsInterface) => {
  const dispatch = useDispatch();
  const currentGroup = useSelector(
    (state: RootState) => state.group.currentGroup,
  );

  const removeImage = async () => {
    if (currentGroup && currentGroup.id) {
      await apiClientObj.delete(
        routes.REMOVE_IMAGE_FROM_GROUP.replace("{1}", currentGroup.id).replace(
          "{2}",
          imgData.id,
        ),
      );
      dispatch(clearModal());
      dispatch(removeCurrentGroupImage({ id: imgData.id }));
      toast.success("Image removed from group successfully");
    }
  };

  const deleteButton = (key: string) => {
    dispatch(
      setModal({
        name: ModalNames.DELETE_MODAL,
        data: {
          id: key,
          heading: "Do you want to remove this image from this group?",
          onSubmit: removeImage,
          type: ImageLoadTypes.GROUP,
        },
      }),
    );
  };

  const previewButton = (id: string) => {
    dispatch(setCurrentGroupImage({ id }));
    dispatch(
      setModal({
        name: ModalNames.PREVIEW_MODAL,
        data: {
          type: ImageLoadTypes.GROUP,
        },
      }),
    );
  };

  return (
    <div className="flex gap-2 m-2">
      <button className="rounded p-2 bg-green-600 text-white">
        <FontAwesomeIcon icon={faDownload} />
      </button>
      <button
        className="rounded p-2 bg-yellow-600 text-white"
        onClick={() => previewButton(imgData.id)}
      >
        <FontAwesomeIcon icon={faEye} />
      </button>
      <button
        className="rounded p-2 bg-red-600 text-white"
        onClick={() => deleteButton(imgData.id)}
      >
        <FontAwesomeIcon icon={faTrashAlt} />
      </button>
    </div>
  );
};

export default GroupViewButtons;
