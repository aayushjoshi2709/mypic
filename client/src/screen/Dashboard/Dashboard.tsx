import Card from "../../component/Card/Card";
import HeaderNavWrapper from "../../component/Wrapper/HeaderNavWrapper";
import imageData from "../../data/imagedata";

const Dashboard = () => {
  const images = imageData;
  return (
    <HeaderNavWrapper>
      <div className="flex-1 justify-center w-full">
        <main className="columns-3 gap-4 p-4 my-4">
          {images.map((img) => (
            <Card imgData={img} />
          ))}
        </main>
      </div>
    </HeaderNavWrapper>
  );
};

export default Dashboard;
