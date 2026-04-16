import { useEffect } from "react";
import Card from "../../component/Card/Card";
import HeaderNavWrapper from "../../component/Wrapper/HeaderNavWrapper";
import { useSelector } from "react-redux";
import type { RootState } from "../../store/store";
import { useNavigate } from "react-router";

interface ImageDataInteface {
  src: string;
}

const Dashboard = () => {
  const navigate = useNavigate();
  const user = useSelector((state: RootState) => state.user);
  useEffect(() => {
    if (!user) {
      navigate("/login");
    }
  }, [user, navigate]);
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
