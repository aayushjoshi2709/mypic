
import { useDispatch, useSelector } from "react-redux"
import Modal from "../Modal"
import type { RootState } from "../../../store/store";
import { ModalNames } from "../../../common/Constants";
import { apiClientObj } from "../../../common/apiClient";
import GroupCard from "../../GroupCard/GroupCard";
import { useEffect } from "react";
import { setFetchGroups } from "../../../store/group.slice";
import { routes } from "../../../common/routes";
import toast from "react-hot-toast";
import { clearModal } from "../../../store/modal.slice";

export interface DeleteModalInterface{
  heading: string,
  id: string,
  onSubmit: ()=>void
}

const AddGroupModal = () => {
  const modal = useSelector((state: RootState) => state.modal);
  const currentImageId = useSelector((state: RootState) => state.image.currentImage?.id);
  const group = useSelector((state: RootState) => state.group);
  const dispatch = useDispatch();

  useEffect(()=>{
    if(group.groups == null){
        dispatch(setFetchGroups())
    }
  },[dispatch, group.groups])
  
  async function addImageToGroup(groupId: string){
    console.log("groupid", groupId)
    console.log(currentImageId)
    if(currentImageId){
        await apiClientObj.post(
            routes.ADD_IMAGE_TO_GROUP.replace("{}", groupId),
            {
                imageId: currentImageId
            }
        );
        toast.success("Image successfully added to group")
        dispatch(clearModal())
    }
  }

  return (
    modal.name == ModalNames.ADD_IMAGE_TO_GROUP?
    <Modal>
        <div className="bg-white rounded w-[60%] " onClick={(e) => e.stopPropagation()} >
            <div className=" p-2  mt-2 flex flex-col w-full gap-4 items-center justify-center">
                <label htmlFor="name" className="text-2xl">Select the group where you want to add this image?</label>
                <hr className="text-black w-full"/>
                <main className="p-4 flex flex-wrap justify-center gap-5 h-[500px] overflow-y-scroll">
                    {group.groups?.map((group)=>{
                        return <GroupCard groupData={group} onClick={addImageToGroup}/>
                    })}
                </main>
            </div>
        </div>
    </Modal>:""
  )
}

export default AddGroupModal

