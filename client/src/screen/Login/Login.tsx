import { Link } from "react-router";
import FormInput, {
  type FormInputProps,
} from "../../component/FormInput/FormInput";
import {
  RoundedButtonPrimary,
  RoundedButtonSecondary,
} from "../../component/Button/RoundedButton";
import FormWrapper from "../../component/Wrapper/FormWrapper";

const Login = () => {
  const items: FormInputProps[] = [
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
    <FormWrapper>
      <h1 className="text-6xl text-center">Login</h1>
      <form>
        {items.map((item, idx) => {
          return <FormInput {...item} key={idx} />;
        })}
        <div className="flex flex-row gap-4 w-full p-8 justify-center">
          <RoundedButtonPrimary text="Login" />
          <Link to="/signup">
            <RoundedButtonSecondary text="Sign Up" />
          </Link>
        </div>
      </form>
    </FormWrapper>
  );
};

export default Login;
