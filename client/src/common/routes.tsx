const BASE_URL = import.meta.env.VITE_BASE_URL || "http://localhost:3000";
const endpoints = {
  LOGIN: "/api/v1/user/login",
  SIGN_UP: "/api/v1/user/signup",
  DASHBOARD: "/api/v1/dashboard",
};

function getFullUrl(route: string) {
  return `${BASE_URL}${route}`;
}

export const routes = {
  LOGIN: getFullUrl(endpoints.LOGIN),
  SIGN_UP: getFullUrl(endpoints.SIGN_UP),
  DASHBOARD: getFullUrl(endpoints.DASHBOARD),
};
