import { Route, Routes } from "react-router";
import "./App.css";
import Dashboard from "./screen/Dashboard/Dashboard";
import Home from "./screen/Home/Home";
import Login from "./screen/Login/Login";
import SignUp from "./screen/SignUp/SignUp";

function App() {
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
