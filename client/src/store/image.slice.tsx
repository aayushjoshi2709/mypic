import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import type { ImageDataInterface, ImageInterface } from "../common/interfaces";

const initialState: ImageDataInterface = {
  images: null,
  currentImage: null,
  currentPage: 0,
  currentLimit: 10,
  totalPages: null,
};

const ImageSlice = createSlice({
  name: "image",
  initialState,
  reducers: {
    appendImages: (state, action: PayloadAction<ImageInterface[]>) => {
      state.images = [...(state?.images ?? []), ...action.payload];
    },
    setCurrentImage: (state, action: PayloadAction<string>) => {
      const newState: ImageDataInterface = {
        ...state,
        currentImage:
          state.images?.find((image) => image.id === action.payload) || null,
      };
      return newState;
    },
    clearImages: () => {
      return initialState;
    },

    setTotalPages: (state, action: PayloadAction<number | null>) => {
      const newState: ImageDataInterface = {
        ...state,
        totalPages: action.payload,
      };
      return newState;
    },

    setCurrentPage: (state, action: PayloadAction<number>) => {
      const newState: ImageDataInterface = {
        ...state,
        currentPage: action.payload,
      };
      return newState;
    },
    setCurrentLimit: (state, action: PayloadAction<number>) => {
      const newState: ImageDataInterface = {
        ...state,
        currentLimit: action.payload,
      };
      return newState;
    },
  },
});

export const {
  appendImages,
  clearImages,
  setCurrentImage,
  setCurrentPage,
  setTotalPages,
  setCurrentLimit,
} = ImageSlice.actions;

export default ImageSlice.reducer;
