import Card from "../../component/Card/Card";
interface ImageDataInteface {
  src: string;
}

const Photos = () => {
  const images: ImageDataInteface[] = [];
  return (
    <div className="flex-1 justify-center w-full">
      <main className="coluxmns-3 gap-4 p-4 my-4">
        {images.map((img) => (
          <Card imgData={img} />
        ))}
      </main>
    </div>
  );
};

export default Photos;
