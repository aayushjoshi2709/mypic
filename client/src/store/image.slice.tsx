import { createSlice, createListenerMiddleware } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import { apiClientObj } from "../common/apiClient";
import { routes } from "../common/routes";
import type { ImageDataInterface, ImageInterface } from "../common/interfaces";




const initialState: ImageDataInterface = {
  images: null,
  currentImage: null,
  fetchImages: true,
};

const ImageSlice = createSlice({
  name: "image",
  initialState,
  reducers: {
    setImages: (state, action: PayloadAction<ImageInterface[]>) => {
      state.images = [...action.payload];
    },
    setCurrentImage: (state, action: PayloadAction<{id: string}>) =>{
      const newState = {
        ...state,
        currentImage: state.images?.find((image) => image.id === action.payload.id) || null
      }
      return newState
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




export const { setImages, clearImage, setCurrentImage, setFetchImages, unsetFetchImages } = ImageSlice.actions;


export const imageListenerMiddleware = createListenerMiddleware();
imageListenerMiddleware.startListening({
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
