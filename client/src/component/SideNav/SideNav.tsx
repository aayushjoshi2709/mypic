import {
  faImage,
  faPerson,
  faSearch,
  faUpload,
  faUserGroup,
} from "@fortawesome/free-solid-svg-icons";
import SideBarItems from "./SideBarItems";
import { useLocation } from "react-router";

const SideNav = () => {
  const location = useLocation();
  
  const menuItems: SideBarItems[] = [
    {
      icon: faUpload,
      text: "upload",
      link: "/dashboard/upload",
      isActive: location.pathname === "/dashboard/upload",
    },
    {
      icon: faImage,
      text: "photos",
      link: "/dashboard/photos",
      isActive: location.pathname === "/dashboard/photos",
    },
    {
      icon: faSearch,
      text: "search",
      link: "/dashboard/search",
      isActive: location.pathname === "/dashboard/search",
    },
    {
      icon: faUserGroup,
      text: "groups",
      link: "/dashboard/groups",
      isActive: location.pathname === "/dashboard/groups",
    },
    {
      icon: faPerson,
      text: "persons",
      link: "/dashboard/persons",
      isActive: location.pathname === "/dashboard/persons",
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
