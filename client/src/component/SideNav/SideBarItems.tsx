import type { IconDefinition } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { toTitleCase } from "../../common/functions";

interface SideBarItems {
  icon: IconDefinition;
  text: string;
}
const SideBarItems = ({ icon, text }: SideBarItems) => {
  return (
    <li className="p-8 py-5 mx-4  text-xl hover:bg-blue-100  hover:rounded-full rounded-m-2">
      <FontAwesomeIcon icon={icon} />
      <span className="ml-2 font-semibold">{toTitleCase(text)}</span>
    </li>
  );
};

export default SideBarItems;
