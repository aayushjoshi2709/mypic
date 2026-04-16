import { Route, Routes, useNavigate } from "react-router";
import "./App.css";
import Dashboard from "./screen/Dashboard/Dashboard";
import Home from "./screen/Home/Home";
import Login from "./screen/Login/Login";
import SignUp from "./screen/SignUp/SignUp";
import { apiClientObj, UNAUTHORIZED_EVENT } from "./common/apiClient";
import { useEffect } from "react";
import { routes } from "./common/routes";
import { useDispatch } from "react-redux";
import { setUser } from "./store/user.slice";

function App() {
  const navigate = useNavigate();
  const dispatch = useDispatch();

  useEffect(() => {
    const handleUnauthorized = () => navigate("/login");
    window.addEventListener(UNAUTHORIZED_EVENT, handleUnauthorized);
    return () =>
      window.removeEventListener(UNAUTHORIZED_EVENT, handleUnauthorized);
  }, [navigate]);

  useEffect(() => {
    const fetchUserData = async () => {
      const userData = await apiClientObj.get(routes.CURRENT_USER);
      dispatch(setUser(userData));
      if (userData) {
        navigate("/dashboard");
      }
    };
    if (localStorage.getItem("token") !== null) {
      fetchUserData();
    } else {
      navigate("/login");
    }
  }, [navigate, dispatch]);

  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/dashboard" element={<Dashboard />} />
      <Route path="/login" element={<Login />} />
      <Route path="/signup" element={<SignUp />} />
    </Routes>
  );
}

export default App;
