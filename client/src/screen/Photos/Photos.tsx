import { useDispatch, useSelector } from "react-redux";
import Card from "../../component/Card/Card";
import type { RootState } from "../../store/store";
import { useEffect } from "react";
import { setImages } from "../../store/image.slice";
import { apiClientObj } from "../../common/apiClient";
import { routes } from "../../common/routes";

const Photos = () => {
  const images = useSelector((state: RootState) => state.image.images);
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
      <main className="columns-3 gap-4 p-4 my-4">
        {images?.map((img) => (
          <Card imgData={img} />
        ))}
      </main>
    </div>
  );
};

export default Photos;
