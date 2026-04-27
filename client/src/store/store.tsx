import { configureStore } from "@reduxjs/toolkit";
import userReducer from "./user.slice";
import imageReducer, { listenerMiddleware } from "./image.slice";
import modalReducer from "./modal.slice"
export const store = configureStore({
  reducer: {
    user: userReducer,
    image: imageReducer,
    modal: modalReducer
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().prepend(listenerMiddleware.middleware)
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
