import Card from "../../component/Card/Card";
import HeaderFooterWrapper from "../../component/Wrapper/HeaderFooterWrapper";
import imageData from "../../data/imagedata";

const Home = () => {
  const images = imageData;
  return (
    <HeaderFooterWrapper>
      <div className="flex-1 justify-center w-full">
        <main className="columns-3 gap-4 p-4">
          {images.map((img) => (
            <Card imgData={img} />
          ))}
        </main>
      </div>
    </HeaderFooterWrapper>
  );
};

export default Home;
