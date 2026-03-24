import type { PropsWithChildren } from "react";
import Header from "../Header/Header";
import Footer from "../Footer/Footer";
import styles from "./HeaderFooter.module.css";

const HeaderFooterWrapper = ({ children }: PropsWithChildren) => {
  return (
    <div className={styles.mainContainer}>
      <Header />
      {children}
      <Footer />
    </div>
  );
};

export default HeaderFooterWrapper;
