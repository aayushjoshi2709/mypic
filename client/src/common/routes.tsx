const BASE_URL = import.meta.env.VITE_BASE_URL || "http://localhost:3000";
const endpoints = {
  LOGIN: "/api/v1/user/login",
  SIGN_UP: "/api/v1/user",
  DASHBOARD: "/api/v1/photos",
  USER_DETAILS: "/api/v1/user",
  CURRENT_USER: "/api/v1/user/me",
  LOGOUT: "/api/v1/user/logout",
};

function getFullUrl(route: string) {
  return `${BASE_URL}${route}`;
}

export const routes = {
  LOGIN: getFullUrl(endpoints.LOGIN),
  SIGN_UP: getFullUrl(endpoints.SIGN_UP),
  DASHBOARD: getFullUrl(endpoints.DASHBOARD),
  USER_DETAILS: getFullUrl(endpoints.USER_DETAILS),
  CURRENT_USER: getFullUrl(endpoints.CURRENT_USER),
  LOGOUT: getFullUrl(endpoints.LOGOUT),
};
