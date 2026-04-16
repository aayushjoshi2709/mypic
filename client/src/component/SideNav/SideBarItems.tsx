import type { IconDefinition } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { toTitleCase } from "../../common/functions";
import { Link } from "react-router";

interface SideBarItems {
  icon: IconDefinition;
  text: string;
  link: string;
  isActive?: boolean;
}
const SideBarItems = ({ icon, text, link, isActive }: SideBarItems) => {
  return (
    <li className={`p-8 py-5 mx-4 my-2  text-xl hover:bg-blue-100  hover:rounded-full rounded-m-2 ${isActive ? "bg-blue-200 rounded-full" : ""}`}>
      <Link to={link}>
        <FontAwesomeIcon icon={icon} />
        <span className="ml-2 font-semibold">{toTitleCase(text)}</span>
      </Link>
    </li>
  );
};

export default SideBarItems;
