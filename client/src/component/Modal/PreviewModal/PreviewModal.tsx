import { useSelector } from "react-redux"
import Modal from "../Modal"
import type { RootState } from "../../../store/store";
import { routes } from "../../../common/routes";

const PreviewModal = () => {

  const modal = useSelector((state: RootState) => state.modal);

  return (
    modal.name == "PREVIEW_MODAL"?
    <Modal>
        <div className="bg-black rounded text-white w-[50%]" onClick={(e) => e.stopPropagation()}>
            <img src={routes.IMAGE_PREFIX + modal.data.key}/>
        </div>
    </Modal>:""
  )
}

export default PreviewModal