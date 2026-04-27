import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { routes } from "../../common/routes";
import { faDownload, faEye, faTrashAlt } from "@fortawesome/free-solid-svg-icons";
import { useDispatch } from "react-redux";
import { clearModal, setModal } from "../../store/modal.slice";
import { ModalNames } from "../../common/Constants";
import DeleteModal from "../Modal/DeleteModal/DeleteModal";
import PreviewModal from "../Modal/PreviewModal/PreviewModal";
import { apiClientObj } from "../../common/apiClient";
import { setCurrentImage, setFetchImages } from "../../store/image.slice";
import toast from "react-hot-toast";
import type { ImageDataInterface } from "../../common/interFaces";

interface CardProps {
  imgData: ImageDataInterface
}

const Card = ({ imgData }: CardProps) => {
  const dispatch = useDispatch();
  const deleteImage = async () => {
    await apiClientObj.delete(routes.GET_SINGLE_IMAGE + imgData.id);
    dispatch(clearModal())
    dispatch(setFetchImages())
    toast.success("Image deleted successfully")
  }
  const deleteButton = (key: string) => {
      dispatch(setModal({
        name: ModalNames.DELETE_MODAL,
        data: {
          id: key,
          heading: "Do you want to delete this image?",
          onSubmit: deleteImage
        }
      }));
  }
  const previewButton  = (id: string) => {
      dispatch(setCurrentImage({id}))
      dispatch(setModal({
        name: ModalNames.PREVIEW_MODAL
      }));
  }
  return <>
    <DeleteModal/>
    <PreviewModal/>
    <div className="group break-inside-avoid mb-4 relative">
      <img
        className="rounded-sm w-full hover:shadow-xl min-h-[200px] h-auto block"
        src={imgData.url}
      />
      <div className="absolute bottom-0 right-0 hidden group-hover:flex w-full bg-black/60 justify-end">
        <div className="flex gap-2 m-2">
            <button className="rounded p-2 bg-green-600 text-white">
              <FontAwesomeIcon icon={faDownload} />
            </button>
            <button className="rounded p-2 bg-yellow-600 text-white" onClick={()=> previewButton(imgData.id)}>
              <FontAwesomeIcon icon={faEye}/>
            </button>
            <button className="rounded p-2 bg-red-600 text-white" onClick={ () => deleteButton(imgData.id)}>
              <FontAwesomeIcon icon={faTrashAlt}/>
            </button>
        </div>
      </div>

    </div>
    </>
};

export default Card;
