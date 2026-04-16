import FormInput from "../../component/FormInput/FormInput";
import {
  RoundedButtonPrimary,
  RoundedButtonSecondary,
} from "../../component/Button/RoundedButton";
import FormWrapper from "../../component/Wrapper/FormWrapper";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router";
import toast from "react-hot-toast";
import { apiClientObj } from "../../common/apiClient";
import { routes } from "../../common/routes";
import { useSelector } from "react-redux";
import type { RootState } from "../../store/store";

const SignUp = () => {
  const navigate = useNavigate();
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const user = useSelector((state: RootState) => state.user);

  async function SignUpUser() {
    if (!name) {
      toast.error("Name is required");
      return;
    }

    if (!email) {
      toast.error("Email is required");
      return;
    }

    if (!username) {
      toast.error("Username is required");
      return;
    }

    if (!password) {
      toast.error("Password is required");
      return;
    }

    await apiClientObj.post(routes.SIGN_UP, {
      name,
      email,
      username,
      password,
    });
    toast.success("User created successfully. Please login.");
    navigate("/login");
  }

  function loginUser() {
    navigate("/login");
  }

  useEffect(() => {
    if (user) {
      navigate("/dashboard/photos");
    }
  }, [navigate, user]);

  return (
    <FormWrapper>
      <h1 className="text-6xl text-center">Sign Up</h1>
      <form className="mt-6">
        <FormInput
          name="name"
          id="name"
          type="text"
          label="email"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <FormInput
          name="email"
          id="email"
          type="text"
          label="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <FormInput
          name="username"
          id="username"
          type="text"
          label="username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <FormInput
          name="password"
          id="password"
          type="text"
          label="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <div className="flex flex-row gap-4 w-full p-8 justify-center">
          <RoundedButtonPrimary text="Sign Up" onClick={SignUpUser} />
          <RoundedButtonSecondary text="Login" onClick={loginUser} />
        </div>
      </form>
    </FormWrapper>
  );
};

export default SignUp;
