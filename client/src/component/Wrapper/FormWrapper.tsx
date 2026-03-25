import { type PropsWithChildren } from "react";
import Header from "../Header/Header";

const FormWrapper = ({ children }: PropsWithChildren) => {
  return (
    <div className="flex flex-col h-screen">
      <Header />
      <main className="flex justify-center items-center h-full">
        <div className="shadow shadow-black-200 p-4 py-8 w-1/3 rounded-2xl">
          {children}
        </div>
      </main>
    </div>
  );
};

export default FormWrapper;
