import { useDispatch, useSelector } from "react-redux";
import type { RootState } from "../../store/store";
import { useNavigate, useParams } from "react-router";
import { useEffect } from "react";
import ImageList from "../../component/ImageList/ImageList";
import { setCurrentGroup, setGroupImageData } from "../../store/group.slice";
import { apiClientObj } from "../../common/apiClient";
import { routes } from "../../common/routes";
const GroupPhotos = () => {
  const params = useParams();
  const navigate = useNavigate();
  const dispatch = useDispatch();

  const {groupId} = params;
  
  useEffect(()=>{
    if(groupId){
        dispatch(setCurrentGroup({id: groupId}))
    }
  },[groupId, dispatch])

  const currentGroup = useSelector((state: RootState) => state.group.currentGroup);
  const imageData = currentGroup?.imageData;
  

  
  useEffect(()=>{
    async function fetchGroupImageData(){
        if(groupId){
            const res = await apiClientObj.get(routes.GET_GROUP_IMAGES.replace("{}", groupId))
            dispatch(setGroupImageData(res))
        }
    }
    if(imageData?.images == null){
        fetchGroupImageData();
    }
  },
  [dispatch, groupId, imageData?.images])


  return (

    <>
        {imageData &&
            <div className="flex-1 justify-center w-full">
                
                    <div className="m-4 p-8 d-flex rounded-2xl content-center text-center border-2 border-dashed border-gray-400 bg-blue-100">
                    <h1 className="text-4xl font-bold mb-4">{currentGroup.name}</h1>

                    {
                        imageData.images && imageData.images.length == 0 && (
                            <h2 className="mb-2">No photos yet...</h2>
                        )
                    }
                    <h2> Add some photos to group form photos tab</h2>
                    <button
                        onClick={() => navigate("/dashboard/photos")}
                        className="mt-4 px-4 py-2 bg-blue-500 text-white rounded"
                    >
                        Goto photos Tab
                    </button>
                    </div>
                
                <ImageList imageData={imageData}/>
            </div>
        }
    </>
  );
};

export default GroupPhotos;
