import { useSelector } from "react-redux"
import Modal from "../Modal"
import type { RootState } from "../../../store/store";
import { routes } from "../../../common/routes";
import { useEffect, useState } from "react";
import { apiClientObj } from "../../../common/apiClient";
import type { ImageState } from "../../../store/image.slice";
import { ModalNames } from "../../../common/Constants";

export interface PreviewModalInterface{
  id: string;
}

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
    const data:PreviewModalInterface = modal.data as PreviewModalInterface;
    if(modal.name === ModalNames.PREVIEW_MODAL){
        const id = data.id;
        if(id){
            getImageData(id);
        }
    }
  }, [modal.data, modal.name])

  return (
    modal.name == ModalNames.PREVIEW_MODAL?
    <Modal>
        <div className="bg-black rounded text-white max-w-[80%]" onClick={(e) => e.stopPropagation()}>
            {
                loading?
                <p>Loading...</p>
                :
                <img src={imageData?.url}/>
            }
            
        </div>
    </Modal>:""
  )
}

export default PreviewModal