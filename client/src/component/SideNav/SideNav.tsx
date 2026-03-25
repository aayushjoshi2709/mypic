import {
  faImage,
  faPerson,
  faSearch,
  faUpload,
  faUserGroup,
} from "@fortawesome/free-solid-svg-icons";
import SideBarItems from "./SideBarItems";

const SideNav = () => {
  const menuItems: SideBarItems[] = [
    {
      icon: faUpload,
      text: "upload",
    },
    {
      icon: faImage,
      text: "photos",
    },
    {
      icon: faSearch,
      text: "search",
    },
    {
      icon: faUserGroup,
      text: "groups",
    },
    {
      icon: faPerson,
      text: "persons",
    },
  ];
  return (
    <aside className="py-4 flex-2 mr-2 shadow">
      <ul>
        {menuItems.map((item, idx) => {
          return <SideBarItems {...item} key={idx} />;
        })}
      </ul>
    </aside>
  );
};

export default SideNav;
