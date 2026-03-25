import { Link } from "react-router";
import Header from "../../component/Header/Header";
import FormInput from "../../component/FormInput/FormInput";
import {
  RoundedButtonPrimary,
  RoundedButtonSecondary,
} from "../../component/Button/RoundedButton";

const SignUp = () => {
  return (
    <div className="flex flex-col h-screen">
      <Header />
      <main className="flex justify-center content-center ">
        <div className="shadow shadow-black-200 p-4 py-8 w-1/3 mt-32 rounded-2xl">
          <h1 className="text-6xl text-center">Sign Up</h1>
          <form className="mt-6">
            <FormInput name="name" id="name" type="text" label="name" />
            <FormInput name="email" id="email" type="email" label="email" />
            <FormInput
              name="username"
              id="username"
              type="text"
              label="username"
            />
            <FormInput
              name="password"
              id="password"
              type="password"
              label="password"
            />
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
