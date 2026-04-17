import { useDispatch, useSelector } from "react-redux";
import Card from "../../component/Card/Card";
import type { RootState } from "../../store/store";
import { useEffect } from "react";
import { setImages } from "../../store/image.slice";
import { apiClientObj } from "../../common/apiClient";
import { routes } from "../../common/routes";
import { useNavigate } from "react-router";

const Photos = () => {
  const images = useSelector((state: RootState) => state.image.images);
  const navigate = useNavigate();
  const dispatch = useDispatch();
  useEffect(() => {
    const fetchImages = async () => {
      try {
        const response = await apiClientObj.get(routes.GET_ALL_IMAGES);
        dispatch(setImages(response));
      } catch (error) {
        console.error("Error fetching images:", error);
        return [];
      }
    };
    if (images) return;
    fetchImages();
  }, [dispatch, images]);
  return (
    <div className="flex-1 justify-center w-full">
      {images && images.length > 0 ? (
        <main className="columns-3 gap-4 p-4 my-4">
          {images?.map((img, idx) => (
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
  );
};

export default Photos;
