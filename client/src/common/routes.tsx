const BASE_URL = import.meta.env.VITE_BASE_URL || "http://localhost:3000";
const endpoints = {
  LOGIN: "/api/v1/user/login",
  SIGN_UP: "/api/v1/user",
  DASHBOARD: "/api/v1/photos",
  USER_DETAILS: "/api/v1/user",
  CURRENT_USER: "/api/v1/user/me",
  LOGOUT: "/api/v1/user/logout",
  GET_PRESIGNED_URL: "/api/v1/presign",
  CREATE_IMAGE: "/api/v1/image",
  GET_ALL_IMAGES: "/api/v1/image",
  GET_SINGLE_IMAGE: "/api/v1/image/",
  CREATE_GROUP: "/api/v1/group",
  GET_ALL_GROUPS: "/api/v1/group",
  GET_GROUP_IMAGES: "/api/v1/group/{0}/images",
  ADD_IMAGE_TO_GROUP: "/api/v1/group/{0}/images",
  REMOVE_IMAGE_FROM_GROUP: "/api/v1/group/{1}/images/{2}",
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
  GET_PRESIGNED_URL: getFullUrl(endpoints.GET_PRESIGNED_URL),
  CREATE_IMAGE: getFullUrl(endpoints.CREATE_IMAGE),
  CREATE_GROUP: getFullUrl(endpoints.CREATE_GROUP),
  GET_ALL_IMAGES: getFullUrl(endpoints.GET_ALL_IMAGES),
  GET_GROUP_IMAGES: getFullUrl(endpoints.GET_GROUP_IMAGES),
  GET_SINGLE_IMAGE: getFullUrl(endpoints.GET_SINGLE_IMAGE),
  GET_ALL_GROUPS: getFullUrl(endpoints.GET_ALL_GROUPS),
  ADD_IMAGE_TO_GROUP: getFullUrl(endpoints.ADD_IMAGE_TO_GROUP),
  REMOVE_IMAGE_FROM_GROUP: getFullUrl(endpoints.REMOVE_IMAGE_FROM_GROUP),
};
