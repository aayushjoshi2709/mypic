import type { ImageDataInterface } from "../../common/interfaces";
import Card from "../ImageCard/ImageCard";

const ImageList = ({ imageData }: { imageData: ImageDataInterface }) => {
  return (
    <>
      {imageData.fetchImages && <div>Loading...</div>}

      {imageData.images && imageData.images.length > 0 && (
        <main className="columns-3 gap-4 p-4 my-4">
          {imageData.images?.map((img, idx) => (
            <Card key={idx} imgData={img} />
          ))}
        </main>
      )}
    </>
  );
};

export default ImageList;
