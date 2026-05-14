import { useNavigate } from "react-router";
import ImageList from "../../component/ImageList/ImageList";
import { useEffect } from "react";
import useImages from "../../customHooks/useImages";

const Photos = () => {
  const navigate = useNavigate();

  const { setGroupId, groupId, fetchImages, images } = useImages();

  useEffect(() => {
    if (!images || groupId) {
      setGroupId("");
      fetchImages(1, 10, "after");
    }
  }, [images, fetchImages, groupId, setGroupId]);

  return (
    <>
      <div className="flex-1 justify-center w-full">
        {!images ||
          (images.length == 0 && (
            <div className="h-100 m-4 d-flex rounded-2xl content-center text-center border-2 border-dashed border-gray-400 bg-blue-100">
              <h1 className="text-4xl font-bold mb-4">
                No Photos Uploaded Yet
              </h1>
              <button
                onClick={() => navigate("/dashboard/upload")}
                className="mt-4 px-4 py-2 bg-blue-500 text-white rounded"
              >
                Upload Images Here
              </button>
            </div>
          ))}

        {images ? (
          <ImageList images={images} isGroupView={false} />
        ) : (
          <p>Loading...</p>
        )}
      </div>
    </>
  );
};

export default Photos;
