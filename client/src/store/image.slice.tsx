import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import type { UserState } from "./user.slice";

export interface ImageState {
  id: string;
  key: string;
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
export default ImageSlice.reducer;
