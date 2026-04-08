import Card from "../../component/Card/Card";
import HeaderNavWrapper from "../../component/Wrapper/HeaderNavWrapper";

interface ImageDataInteface {
  src: string;
}

const Dashboard = () => {
  const images: ImageDataInteface[] = [];
  return (
    <HeaderNavWrapper>
      <div className="flex-1 justify-center w-full">
        <main className="coluxmns-3 gap-4 p-4 my-4">
          {images.map((img) => (
            <Card imgData={img} />
          ))}
        </main>
      </div>
    </HeaderNavWrapper>
  );
};

export default Dashboard;
