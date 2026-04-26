import { useDispatch } from "react-redux"
import { clearModal } from "../../store/modal.slice";

const Modal = ({children}: {children:React.ReactNode}) => {
  const dispatch = useDispatch();

  return (
    <div className="absolute top-0 left-0 h-screen w-screen fixed z-10 bg-black/70 flex items-center justify-center" onClick={
        ()=>{
            dispatch(clearModal())
        }
    }>
        {children}
    </div>
  )
}

export default Modal