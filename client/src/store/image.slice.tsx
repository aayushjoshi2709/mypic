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
} = {
  images: null,
  currentImage: null,
};

const ImageSlice = createSlice({
  name: "image",
  initialState,
  reducers: {
    setImages: (state, action: PayloadAction<ImageState[]>) => {
      state.images = [...(state.images ?? []), ...action.payload];
    },
    addImage: (state, action: PayloadAction<ImageState>) => {
      state.images = [...(state.images ?? []), action.payload];
    },
    removeImage: (state, action: PayloadAction<string>) => {
      const imageId = action.payload;
      state.images = state.images?.filter((image) => image.id !== imageId) ?? null;
      if (state.currentImage?.id === imageId) {
        state.currentImage = null;
      }
    },
    setCurrentImage: (state, action: PayloadAction<string>) => {
      const imageId = action.payload;
      state.currentImage =
        state.images?.find((image) => image.id === imageId) ?? null;
    },
    clearImage: () => {
      return initialState;
    },
  },
});

export const { setImages, addImage, removeImage, setCurrentImage, clearImage } = ImageSlice.actions;
export default ImageSlice.reducer;
