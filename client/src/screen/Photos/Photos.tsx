import { useDispatch, useSelector } from "react-redux";
import Card from "../../component/Card/Card";
import type { RootState } from "../../store/store";
import { setFetchImages } from "../../store/image.slice";
import { useNavigate } from "react-router";
import { useEffect } from "react";

const Photos = () => {
  const image = useSelector((state: RootState) => state.image);
  const navigate = useNavigate();
  const dispatch = useDispatch();
  useEffect(()=>{
    if(image.images == null){
      dispatch(setFetchImages())
    }
  },
  [image.fetchImages, image.images, dispatch])
  return (
    <>
    <div className="flex-1 justify-center w-full">
      {image.fetchImages ? (
        <div>Loading...</div>
      ) : image.images && image.images.length > 0 ? (
        <main className="columns-3 gap-4 p-4 my-4">
          {image.images?.map((img, idx) => (
            <Card key={idx} imgData={img} />
          ))}
        </main>
      ) : (
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
    </div>
    </>
  );
};

export default Photos;
