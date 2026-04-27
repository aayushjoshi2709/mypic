import { createSlice, createListenerMiddleware } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import type { UserState } from "./user.slice";
import { apiClientObj } from "../common/apiClient";
import { routes } from "../common/routes";


export interface ImageState {
  id: string;
  url: string;
  createdAt: string;
  updatedAt: string;
  user: Partial<UserState>;
}

const initialState: {
  images: ImageState[] | null;
  currentImage: ImageState | null;
  fetchImages: boolean;
} = {
  images: null,
  currentImage: null,
  fetchImages: true,
};

const ImageSlice = createSlice({
  name: "image",
  initialState,
  reducers: {
    setImages: (state, action: PayloadAction<ImageState[]>) => {
      state.images = [...action.payload];
    },
    clearImage: () => {
      return initialState;
    },
    setFetchImages: (state)=>{
      state.fetchImages = true;
    },
    unsetFetchImages: (state)=>{
      state.fetchImages = false;
    }
  },
});




export const { setImages, clearImage, setFetchImages, unsetFetchImages } = ImageSlice.actions;


export const listenerMiddleware = createListenerMiddleware();
listenerMiddleware.startListening({
  actionCreator: setFetchImages,
  effect: async (_, listenerApi) => {
    try {
      const res = await apiClientObj.get(routes.GET_ALL_IMAGES)
      listenerApi.dispatch(setImages(res));
      listenerApi.dispatch(unsetFetchImages());
    } catch (err) {
      console.error(err)
    }
  }
})


export default ImageSlice.reducer;
