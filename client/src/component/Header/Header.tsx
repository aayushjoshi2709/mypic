import { faHouse } from "@fortawesome/free-solid-svg-icons";
import { faUpload } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
const Header = () => {
  return (
    <header className="p-4 w-full bg-amber-200 font-sans">
      <nav className="flex justify-between">
        <h1 className="text-xl">My Pic</h1>
        <ul className="flex flex-row gap-4">
          <li>
            <FontAwesomeIcon icon={faHouse} />
            Home
          </li>
          <li>
            <FontAwesomeIcon icon={faUpload} />
            Upload
          </li>
        </ul>
      </nav>
    </header>
  );
};

export default Header;
