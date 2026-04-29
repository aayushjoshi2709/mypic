import type { ImageDataInterface } from "../../common/interfaces";
import Card from "../ImageCard/ImageCard";

interface ImageListInterface {
  imageData: ImageDataInterface;
  isGroupView: boolean;
}

const ImageList = ({ imageData, isGroupView }: ImageListInterface) => {
  return (
    <>
      {imageData.fetchImages && <div>Loading...</div>}

      {imageData.images && imageData.images.length > 0 && (
        <main className="columns-3 gap-4 p-4 my-4">
          {imageData.images?.map((img) => (
            <Card key={img.id} imgData={img} isGroupView={isGroupView} />
          ))}
        </main>
      )}
    </>
  );
};

export default ImageList;
