import { configureStore } from "@reduxjs/toolkit";
import userReducer from "./user.slice";
import imageReducer from "./image.slice";
export const store = configureStore({
  reducer: {
    user: userReducer,
    image: imageReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
