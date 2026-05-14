import type { ImageInterface } from "../../common/interfaces";
import Card from "../ImageCard/ImageCard";

interface ImageListInterface {
  images: ImageInterface[];
  isGroupView: boolean;
}

const ImageList = ({ images, isGroupView }: ImageListInterface) => {
  return (
    <>
      {images && images.length > 0 && (
        <main className="columns-3 gap-4 p-4 my-4">
          {images?.map((img) => (
            <Card key={img.id} imgData={img} isGroupView={isGroupView} />
          ))}
        </main>
      )}
    </>
  );
};

export default ImageList;
