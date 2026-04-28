import { Route, Routes } from "react-router";
import "./App.css";
import Photos from "./screen/Photos/Photos";
import Home from "./screen/Home/Home";
import Login from "./screen/Login/Login";
import SignUp from "./screen/SignUp/SignUp";
import Upload from "./screen/Upload/Upload";
import AuthenticatedRoutes from "./component/AuthenticatedRoutes/AuthenticatedRoutes";
import Search from "./screen/Search/Search";
import Groups from "./screen/Groups/Groups";
import GroupPhotos from "./screen/GroupPhotos/GroupPhotos";
function App() {
  return <>
    <Routes>
      <Route path="/login" element={<Login />} />
      <Route path="/signup" element={<SignUp />} />
      <Route path="" element={<Home />} />
      <Route path="/dashboard" element={<AuthenticatedRoutes />}>
        <Route path="photos" element={<Photos />} />
        <Route path="upload" element={<Upload />} />
        <Route path="search" element={<Search />} />
        <Route path="groups" element={<Groups />} />
        <Route path="groups/:groupId" element={<GroupPhotos />} />
      </Route>
    </Routes>
  </>
}

export default App;
