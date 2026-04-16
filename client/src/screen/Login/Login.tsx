import FormInput from "../../component/FormInput/FormInput";
import {
  RoundedButtonPrimary,
  RoundedButtonSecondary,
} from "../../component/Button/RoundedButton";
import FormWrapper from "../../component/Wrapper/FormWrapper";
import { useEffect, useState } from "react";
import { routes } from "../../common/routes";
import { apiClientObj } from "../../common/apiClient";
import { useNavigate } from "react-router";
import toast from "react-hot-toast";
import { useSelector } from "react-redux";
import type { RootState } from "../../store/store";

const Login = () => {
  const navigate = useNavigate();

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const user = useSelector((state: RootState) => state.user);

  async function loginUser() {
    if (!username) {
      toast.error("Username is required");
      return;
    }

    if (!password) {
      toast.error("Password is required");
      return;
    }

    const res = await apiClientObj.post(routes.LOGIN, {
      username,
      password,
    });
    const token = res.token;
    localStorage.setItem("token", token);
    toast.success("Logged in successfully");
    navigate("/dashboard");
  }

  async function signUpUser() {
    navigate("/signup");
  }

  useEffect(() => {
    if (user) {
      navigate("/dashboard");
    }
  }, [navigate, user]);

  return (
    <FormWrapper>
      <h1 className="text-6xl text-center">Login</h1>
      <form>
        <FormInput
          name="username"
          id="username"
          type="text"
          label="username"
          value={username}
          required={true}
          onChange={(e) => setUsername(e.target.value)}
        />
        <FormInput
          name="password"
          id="password"
          type="password"
          label="password"
          value={password}
          required={true}
          onChange={(e) => setPassword(e.target.value)}
        />
        <div className="flex flex-row gap-4 w-full p-8 justify-center">
          <RoundedButtonPrimary text="Login" onClick={loginUser} />
          <RoundedButtonSecondary text="Sign Up" onClick={signUpUser} />
        </div>
      </form>
    </FormWrapper>
  );
};

export default Login;
