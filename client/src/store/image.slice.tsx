import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";

export interface ImageState {
  id: string;
  url: string;
  createdAt: string;
  updatedAt: string;
}

const initialState: ImageState = {
  id: "",
  url: "",
  createdAt: "",
  updatedAt: "",
};

const ImageSlice = createSlice({
  name: "image",
  initialState,
  reducers: {
    setImage: (state, action: PayloadAction<ImageState>) => {
      return action.payload;
    },
    clearImage: () => {
      return initialState;
    },
  },
});

export const { setImage, clearImage } = ImageSlice.actions;
export default ImageSlice.reducer;
