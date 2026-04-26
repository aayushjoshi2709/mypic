import { useDispatch } from "react-redux"
import { clearModal } from "../../store/modal.slice";

const Modal = ({children}: {children:React.ReactNode}) => {
  const dispatch = useDispatch();

  return (
    <div className="h-screen w-screen fixed z-10 bg-black/90 flex items-center justify-center" onClick={
        ()=>{
            dispatch(clearModal())
        }
    }>
        {children}
    </div>
  )
}

export default Modal