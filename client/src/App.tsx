import { Route, Routes, useNavigate } from "react-router";
import "./App.css";
import Dashboard from "./screen/Dashboard/Dashboard";
import Home from "./screen/Home/Home";
import Login from "./screen/Login/Login";
import SignUp from "./screen/SignUp/SignUp";
import { UNAUTHORIZED_EVENT } from "./common/apiClient";
import { useEffect } from "react";

function App() {
  const navigate = useNavigate();

  useEffect(() => {
    const handleUnauthorized = () => navigate("/login");
    window.addEventListener(UNAUTHORIZED_EVENT, handleUnauthorized);
    return () =>
      window.removeEventListener(UNAUTHORIZED_EVENT, handleUnauthorized);
  }, [navigate]);
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
