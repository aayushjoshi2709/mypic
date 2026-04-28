import { useDispatch, useSelector } from "react-redux";
import type { RootState } from "../../store/store";
import { setFetchImages } from "../../store/image.slice";
import { useNavigate } from "react-router";
import { useEffect } from "react";
import ImageList from "../../component/ImageList/ImageList";

const Photos = () => {
  const imageData = useSelector((state: RootState) => state.image);
  const navigate = useNavigate();
  const dispatch = useDispatch();
  useEffect(()=>{
    if(imageData.images == null){
      dispatch(setFetchImages())
    }
  },
  [imageData.fetchImages, imageData.images, dispatch])
  return (
    <>
    <div className="flex-1 justify-center w-full">
      {
        imageData.images && imageData.images.length == 0 && (
        <div className="h-100 m-4 d-flex rounded-2xl content-center text-center border-2 border-dashed border-gray-400 bg-blue-100">
          <h1 className="text-4xl font-bold mb-4">No Photos Uploaded Yet</h1>
          <button
            onClick={() => navigate("/dashboard/upload")}
            className="mt-4 px-4 py-2 bg-blue-500 text-white rounded"
          >
            Upload Images Here
          </button>
        </div>
      )}
      <ImageList imageData={imageData}/>
    </div>
    </>
  );
};

export default Photos;
