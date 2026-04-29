import { useSelector } from "react-redux";
import Modal from "../Modal";
import type { RootState } from "../../../store/store";
import type { ImageLoadType } from "../../../common/Constants";
import { ImageLoadTypes, ModalNames } from "../../../common/Constants";

export interface PreviewModalDataInterface {
  type: ImageLoadType;
}

const PreviewModal = () => {
  const modal = useSelector((state: RootState) => state.modal);
  const type = modal.data?.type ?? ImageLoadTypes.IMAGE;

  const image = useSelector((state: RootState) =>
    type == ImageLoadTypes.IMAGE
      ? state.image
      : state.group.currentGroup?.imageData,
  );

  return modal.name == ModalNames.PREVIEW_MODAL ? (
    <Modal>
      <div
        className="bg-black rounded text-white max-w-[80%]"
        onClick={(e) => e.stopPropagation()}
      >
        {image?.currentImage?.url ? (
          <img src={image.currentImage.url} />
        ) : (
          <p>Loading...</p>
        )}
      </div>
    </Modal>
  ) : (
    ""
  );
};

export default PreviewModal;
