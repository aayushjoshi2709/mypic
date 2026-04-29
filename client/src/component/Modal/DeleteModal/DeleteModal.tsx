import { useDispatch, useSelector } from "react-redux";
import {
  RoundedButtonPrimary,
  RoundedButtonSecondary,
} from "../../Button/RoundedButton";
import Modal from "../Modal";
import type { RootState } from "../../../store/store";
import { ModalNames } from "../../../common/Constants";
import type { ImageLoadType } from "../../../common/Constants";
import { clearModal } from "../../../store/modal.slice";

export interface DeleteModalDataInterface {
  heading: string;
  id: string;
  type: ImageLoadType;
  onSubmit: () => void;
}

const DeleteModal = () => {
  const dispatch = useDispatch();
  const modal = useSelector((state: RootState) => state.modal);
  const data: DeleteModalDataInterface = modal.data as DeleteModalDataInterface;
  return modal.name == ModalNames.DELETE_MODAL ? (
    <Modal>
      <div
        className="bg-black rounded text-white w-[50%]"
        onClick={(e) => e.stopPropagation()}
      >
        <h3 className="p-2 text-center">{data.heading}</h3>
        <hr />
        <div className="p-4 mt-4 flex w-full gap-4 items-center justify-center">
          <RoundedButtonPrimary
            text="Yes"
            onClick={() => {
              data.onSubmit();
            }}
          />
          <RoundedButtonSecondary
            text="No"
            onClick={() => {
              dispatch(clearModal());
            }}
          />
        </div>
      </div>
    </Modal>
  ) : (
    ""
  );
};

export default DeleteModal;
