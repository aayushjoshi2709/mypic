import { useSelector } from "react-redux"
import Modal from "../Modal"
import type { RootState } from "../../../store/store";
import { ModalNames } from "../../../common/Constants";

const PreviewModal = () => {
  const modal = useSelector((state: RootState) => state.modal);
  const image = useSelector((state: RootState) => state.image)

  return (
    modal.name == ModalNames.PREVIEW_MODAL && image.currentImage?
    <Modal>
        <div className="bg-black rounded text-white max-w-[80%]" onClick={(e) => e.stopPropagation()}>
            <img src={image.currentImage.url}/>
        </div>
    </Modal>:""
  )
}

export default PreviewModal