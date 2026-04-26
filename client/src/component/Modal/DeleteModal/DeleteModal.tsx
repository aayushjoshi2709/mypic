import { useDispatch, useSelector } from "react-redux"
import { RoundedButtonPrimary, RoundedButtonSecondary } from "../../Button/RoundedButton"
import Modal from "../Modal"
import type { RootState } from "../../../store/store";
import { clearModal } from "../../../store/modal.slice";

const DeleteModal = () => {

  const modal = useSelector((state: RootState) => state.modal);
  const dispatch = useDispatch();

  return (
    modal.name == "DELETE_MODAL"?
    <Modal>
        <div className="bg-black rounded text-white w-[50%]" onClick={(e) => e.stopPropagation()} >
            <h3 className="p-2 text-center">{modal.data?.heading}</h3>
            <hr/>
            <div className="p-4 mt-4 flex w-full gap-4 items-center justify-center">
                <RoundedButtonPrimary text="Yes" onClick={()=>{}}/>
                <RoundedButtonSecondary text="No" onClick={()=>{
                    dispatch(clearModal())
                }}/>
            </div>
        </div>
    </Modal>:""
  )
}

export default DeleteModal