import { Link } from "react-router";
import Header from "../../component/Header/Header";
import FormInput, {
  type FormInputProps,
} from "../../component/FormInput/FormInput";
import {
  RoundedButtonPrimary,
  RoundedButtonSecondary,
} from "../../component/Button/RoundedButton";

const SignUp = () => {
  const items: FormInputProps[] = [
    {
      name: "name",
      id: "name",
      type: "text",
      label: "email",
    },
    {
      name: "email",
      id: "email",
      type: "email",
      label: "email",
    },
    {
      name: "username",
      id: "username",
      type: "text",
      label: "username",
    },
    {
      name: "password",
      id: "password",
      type: "password",
      label: "password",
    },
  ];
  return (
    <div className="flex flex-col h-screen">
      <Header />
      <main className="flex justify-center content-center ">
        <div className="shadow shadow-black-200 p-4 py-8 w-1/3 mt-32 rounded-2xl">
          <h1 className="text-6xl text-center">Sign Up</h1>
          <form className="mt-6">
            {items.map((item, idx) => {
              return <FormInput {...item} key={idx} />;
            })}
            <div className="flex flex-row gap-4 w-full p-8 justify-center">
              <RoundedButtonPrimary text="Sign Up" />
              <Link to="/login">
                <RoundedButtonSecondary text="Login" />
              </Link>
            </div>
          </form>
        </div>
      </main>
    </div>
  );
};

export default SignUp;
