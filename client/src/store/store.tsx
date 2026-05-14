import { configureStore } from "@reduxjs/toolkit";
import userReducer from "./user.slice";
import groupReducer, { groupListenerMiddleware } from "./group.slice";
import imageReducer from "./image.slice";
import modalReducer from "./modal.slice";
export const store = configureStore({
  reducer: {
    user: userReducer,
    image: imageReducer,
    modal: modalReducer,
    group: groupReducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().prepend(groupListenerMiddleware.middleware),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
