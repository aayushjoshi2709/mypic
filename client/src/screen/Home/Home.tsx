import { Link } from "react-router";
import Header from "../../component/Header/Header";

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
              <button className="p-4 px-8 mt-8 bg-green-500 text-white font-semibold rounded-full">
                Get Started
              </button>
            </Link>
          </div>
        </main>
      </div>
    </>
  );
};

export default Home;
