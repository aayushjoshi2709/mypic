import { useSelector } from "react-redux"
import Modal from "../Modal"
import type { RootState } from "../../../store/store";
import { routes } from "../../../common/routes";
import { useEffect, useState } from "react";
import { apiClientObj } from "../../../common/apiClient";
import type { ImageState } from "../../../store/image.slice";

const PreviewModal = () => {

  const [imageData, setImageData] = useState<ImageState|null>(null);
  const [loading, setLoading] = useState<boolean>(false);
  const modal = useSelector((state: RootState) => state.modal);

  async function getImageData(id: string){
    setLoading(true)
    const imageDataObj = await apiClientObj.get(routes.GET_SINGLE_IMAGE + id)
    setImageData(imageDataObj)
    setLoading(false)
  }

  useEffect(()=>{
    if(modal.name === "PREVIEW_MODAL"){
        const id = modal.data.id;
        if(id){
            getImageData(id);
        }
    }
  }, [modal.data.id, modal.name])

  return (
    modal.name == "PREVIEW_MODAL"?
    <Modal>
        <div className="bg-black rounded text-white max-w-[80%]" onClick={(e) => e.stopPropagation()}>
            {
                loading?
                <p>Loading...</p>
                :
                <img src={routes.IMAGE_PREFIX + imageData?.key}/>
            }
            
        </div>
    </Modal>:""
  )
}

export default PreviewModal