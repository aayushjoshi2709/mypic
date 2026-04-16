import { faDoorOpen, faUser } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Link } from "react-router";
import { useSelector } from "react-redux";
import type { RootState } from "../../store/store";
import { apiClientObj } from "../../common/apiClient";
import { routes } from "../../common/routes";

const Header = () => {
  async function logoutUser() {
    await apiClientObj.delete(routes.LOGOUT);
    localStorage.removeItem("token");
  }
  const user = useSelector((state: RootState) => state.user);
  return (
    <header className="p-4 w-full border-b-2 border-gray-100 font-sans">
      <nav className="flex justify-between">
        <Link to="/">
          <h1 className="text-2xl font-extrabold ">My Pic</h1>
        </Link>
        <ul className="flex flex-row gap-4">
          {!user && (
            <>
              <li>
                <Link to="/login">Login</Link>
              </li>
              <li>
                <Link to="/signup">Sign Up</Link>
              </li>
            </>
          )}
          {user && (
            <>
              <li>
                <FontAwesomeIcon icon={faUser} />
                {user?.username}
              </li>
              <li>
                <button onClick={logoutUser}>
                  {" "}
                  <FontAwesomeIcon icon={faDoorOpen} />
                  Logout
                </button>
              </li>
            </>
          )}
        </ul>
      </nav>
    </header>
  );
};

export default Header;
