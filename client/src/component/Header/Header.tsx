import { faUser } from "@fortawesome/free-solid-svg-icons";
import { faUpload } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Link } from "react-router";
const Header = () => {
  return (
    <header className="p-4 w-full border-b-2 border-gray-100 font-sans">
      <nav className="flex justify-between">
        <Link to="/">
          <h1 className="text-2xl font-extrabold">My Pic</h1>
        </Link>
        <ul className="flex flex-row gap-4">
          <li>
            <FontAwesomeIcon icon={faUpload} />
            Upload
          </li>
          <li>
            <FontAwesomeIcon icon={faUser} />
            Upload
          </li>
        </ul>
      </nav>
    </header>
  );
};

export default Header;
