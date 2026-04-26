import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { routes } from "../../common/routes";
import { faDownload, faEye, faTrashAlt } from "@fortawesome/free-solid-svg-icons";
import { useDispatch } from "react-redux";
import { setModal } from "../../store/modal.slice";

interface CardProps {
  imgData: {
    key: string;
  };
}



const Card = ({ imgData }: CardProps) => {
  const dispatch = useDispatch();
  const deleteButton = (key: string) => {
      dispatch(setModal({
        name: "DELETE_MODAL",
        data:{
          key: key,
          heading: "Do you want to delete this image?"
        }
      }));
  }
  const previewButton  = (key: string) => {
      dispatch(setModal({
        name: "PREVIEW_MODAL",
        data:{
          key: key
        }
      }));
  }
  return (
    <div className="group break-inside-avoid mb-4 relative">
      <img
        className="rounded-sm w-full hover:shadow-xl min-h-[200px] h-auto block"
        src={routes.IMAGE_PREFIX + imgData.key}
      />
      <div className="absolute bottom-0 right-0 hidden group-hover:flex w-full bg-black/60 justify-end">
        <div className="flex gap-2 m-2">
            <button className="rounded p-2 bg-green-600 text-white">
              <FontAwesomeIcon icon={faDownload} />
            </button>
            <button className="rounded p-2 bg-yellow-600 text-white" onClick={()=> previewButton(imgData.key)}>
              <FontAwesomeIcon icon={faEye}/>
            </button>
            <button className="rounded p-2 bg-red-600 text-white" onClick={ () => deleteButton(imgData.key)}>
              <FontAwesomeIcon icon={faTrashAlt}/>
            </button>
        </div>
      </div>
    </div>
  );
};

export default Card;
