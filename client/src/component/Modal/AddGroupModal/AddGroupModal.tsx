import { useDispatch, useSelector } from "react-redux"
import Modal from "../Modal"
import type { RootState } from "../../../store/store";
import { ModalNames } from "../../../common/Constants";
import { useEffect, useState } from "react";
import { apiClientObj } from "../../../common/apiClient";
import { routes } from "../../../common/routes";
import { clearModal } from "../../../store/modal.slice";
import toast from "react-hot-toast";
import useImageDialog from "../../../customHooks/useImageDialog";
import { setFetchGroups } from "../../../store/group.slice";
import { RoundedButtonSecondary } from "../../Button/RoundedButton";

export interface DeleteModalInterface{
  heading: string,
  id: string,
  onSubmit: ()=>void
}

const AddGroupModal = () => {
  const [name, setName] = useState<string>("");
  const [error, setError] = useState<string>("");
  const [loading, setLoading] = useState<boolean>(false);
  const dispatch = useDispatch();
  const modal = useSelector((state: RootState) => state.modal);

  const [openFileDialog, s3Key, originalFileName, uploadInProgress] = useImageDialog();
  
  useEffect(()=>{
    if(s3Key && originalFileName){
       setLoading(false);
    }
  },[dispatch, originalFileName, s3Key])

  function validateInput(){
    setError("");
    if(name.length == 0){
        setError("Enter valid name");
        return false;
    }

    if(s3Key == null){
        setError("Please upload a group image");
        return false;
    }
    return true;
  }
  
  async function createGroup(){
    if(validateInput()){
        setLoading(true);
        await apiClientObj.post(
            routes.CREATE_GROUP,
            {
                name: name,
                imageKey: s3Key
            }
        )
        setLoading(false);
        dispatch(setFetchGroups());
        toast.success("Group created successfully...")
        dispatch(clearModal());
    }
  }

  return (
    modal.name == ModalNames.ADD_GROUP?
    <Modal>
        <div className="bg-white rounded w-fit" onClick={(e) => e.stopPropagation()} >
            <div className="p-4 px-10 mt-4 flex flex-col w-full gap-4 items-center justify-center">
                <label htmlFor="name" className=" ">What you want to name your group?</label>
                <input className="border-1 rounded p-2 w-full" type="text" id="name" value={name} onChange={(e) => setName(e.target.value)}/>
                <div className="flex gap-2 justify-around items-center  my-4 w-full">
                    <label htmlFor="name" className="">Add group image?</label>
                    <RoundedButtonSecondary
                      text="Upload"
                      onClick={()=>
                        openFileDialog()
                      }
                    />
                </div>
                {error?<p className="sm text-red-700">{error}</p>:""}
                <input
                  className=" px-8 py-2 rounded-full text-white bg-green-700 hover:bg-green-800 border-1 border-white"
                  disabled={uploadInProgress}
                  type="button" value={loading?"Adding Group":(uploadInProgress?"Uploading Image":"Add")} onClick={()=>{
                  createGroup()
                }}/>
            </div>
        </div>
    </Modal>:""
  )
}

export default AddGroupModal