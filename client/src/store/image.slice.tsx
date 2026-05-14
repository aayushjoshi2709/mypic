import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import type { ImageDataInterface, ImageInterface } from "../common/interfaces";

const initialState: ImageDataInterface = {
  images: null,
  currentImage: null,
  currentPage: 0,
  currentLimit: 10,
  totalPages: 0,
};

const ImageSlice = createSlice({
  name: "image",
  initialState,
  reducers: {
    appendImages: (
      state,
      action: PayloadAction<{
        images: ImageInterface[];
        arrange: "before" | "after";
      }>,
    ) => {
      const { arrange, images } = action.payload;
      if (arrange === "before") {
        state.images = [...images, ...(state?.images ?? [])];
        if (state.images.length > state.currentLimit * 3) {
          state.images = state.images.slice(
            0,
            state.images.length - state.currentLimit,
          );
        }
      } else {
        state.images = [...(state?.images ?? []), ...images];
        if (state.images.length > state.currentLimit * 3) {
          state.images = state.images.slice(
            state.currentLimit + 1,
            state.images.length,
          );
        }
      }
    },
    setCurrentImage: (state, action: PayloadAction<{ id: string }>) => {
      const newState: ImageDataInterface = {
        ...state,
        currentImage:
          state.images?.find((image) => image.id === action.payload.id) || null,
      };
      return newState;
    },
    clearImage: () => {
      return initialState;
    },

    setTotalPages: (state, action: PayloadAction<number>) => {
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
  clearImage,
  setCurrentImage,
  setCurrentPage,
  setTotalPages,
  setCurrentLimit,
} = ImageSlice.actions;

export default ImageSlice.reducer;
