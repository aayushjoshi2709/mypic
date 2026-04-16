import { useEffect } from "react";
import { apiClientObj, UNAUTHORIZED_EVENT } from "../../common/apiClient";
import { setUser } from "../../store/user.slice";
import { routes } from "../../common/routes";
import { Outlet, useNavigate } from "react-router";
import { useDispatch, useSelector } from "react-redux";
import type { RootState } from "../../store/store";
import HeaderNavWrapper from "../Wrapper/HeaderNavWrapper";

const AuthenticatedRoutes = () => {
  const user = useSelector((state: RootState) => state.user);
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
    };
    console.log("AuthenticatedRoutes: Checking authentication...");
    if (localStorage.getItem("token") !== null) {
      if (!user) {
        fetchUserData();
      }
    } else {
      navigate("/login");
    }
  }, [navigate, dispatch, user]);
  return (
    <HeaderNavWrapper>
      <Outlet />
    </HeaderNavWrapper>
  );
};

export default AuthenticatedRoutes;
