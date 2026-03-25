import type { PropsWithChildren } from "react";
import Header from "../Header/Header";
import SideNav from "../SideNav/SideNav";

const HeaderNavWrapper = ({ children }: PropsWithChildren) => {
  return (
    <div className="flex flex-col h-screen">
      <Header />
      <div className="flex flex-row flex-1 overflow-hidden">
        <SideNav />
        <div className="flex-10 max-h-full overflow-scroll">
          <div>{children}</div>
        </div>
      </div>
    </div>
  );
};

export default HeaderNavWrapper;
