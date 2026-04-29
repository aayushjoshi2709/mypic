import { routes } from "../../../common/routes";
import {
  faAdd,
  faDownload,
  faEye,
  faTrashAlt,
} from "@fortawesome/free-solid-svg-icons";
import { useDispatch } from "react-redux";
import { clearModal, setModal } from "../../../store/modal.slice";
import { ImageLoadTypes, ModalNames } from "../../../common/Constants";
import { apiClientObj } from "../../../common/apiClient";
import { setCurrentImage, setFetchImages } from "../../../store/image.slice";
import toast from "react-hot-toast";
import type { ImageInterface } from "../../../common/interfaces";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

interface ImageViewButtonsInterface {
  imgData: ImageInterface;
}

const ImageViewButtons = ({ imgData }: ImageViewButtonsInterface) => {
  const dispatch = useDispatch();

  const deleteImage = async () => {
    await apiClientObj.delete(routes.GET_SINGLE_IMAGE + imgData.id);
    dispatch(clearModal());
    dispatch(setFetchImages());
    toast.success("Image deleted successfully");
  };

  const deleteButton = (key: string) => {
    dispatch(
      setModal({
        name: ModalNames.DELETE_MODAL,
        data: {
          id: key,
          heading: "Do you want to delete this image?",
          onSubmit: deleteImage,
          type: ImageLoadTypes.IMAGE,
        },
      }),
    );
  };

  const previewButton = (id: string) => {
    dispatch(setCurrentImage({ id }));
    dispatch(
      setModal({
        name: ModalNames.PREVIEW_MODAL,
        data: {
          type: ImageLoadTypes.IMAGE,
        },
      }),
    );
  };

  const addToGroup = (id: string) => {
    dispatch(setCurrentImage({ id }));
    dispatch(
      setModal({
        name: ModalNames.ADD_IMAGE_TO_GROUP,
      }),
    );
  };

  return (
    <div className="flex gap-2 m-2">
      <button className="rounded p-2 bg-green-600 text-white">
        <FontAwesomeIcon icon={faDownload} />
      </button>
      <button
        className="rounded p-2 bg-blue-600 text-white"
        onClick={() => addToGroup(imgData.id)}
      >
        <FontAwesomeIcon icon={faAdd} />
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

export default ImageViewButtons;
