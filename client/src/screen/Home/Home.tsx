import { Link } from "react-router";
import Header from "../../component/Header/Header";
import RoundedButton from "../../component/Button/RoundedButton";

const Home = () => {
  return (
    <>
      <Header />
      <div className="max-h-full overflow-scroll">
        <main className="text-center  h-screen w-full d-flex content-center">
          <h1 className="text-6xl">A Smarter Home for Your Memories</h1>
          <h2 className="mt-4 text-xl">
            Bring all your photos and videos together in one secure place. Enjoy
            unlimited storage, effortless organization, and access from
            anywhere..
          </h2>
          <div>
            <Link to="/login">
              <RoundedButton
                text="Get Started"
                classNames={[
                  "hover:bg-green-600",
                  "bg-green-500",
                  "mt-8",
                  "px-8",
                  "py-4",
                  "text-xl"
                ]}
              />
            </Link>
          </div>
        </main>
      </div>
    </>
  );
};

export default Home;
