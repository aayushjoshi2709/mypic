import {
  createListenerMiddleware,
  createSlice,
  type PayloadAction,
} from "@reduxjs/toolkit";
import type {
  GroupDataInterface,
  GroupInterface,
  ImageInterface,
} from "../common/interfaces";
import { apiClientObj } from "../common/apiClient";
import { routes } from "../common/routes";

const initialState: GroupDataInterface = {
  groups: null,
  currentGroup: null,
  fetchGroups: true,
};

const GroupSlice = createSlice({
  name: "group",
  initialState: initialState,
  reducers: {
    setGroups: (state, action: PayloadAction<GroupInterface[]>) => {
      state.groups = [...action.payload];
    },
    setCurrentGroup: (state, action: PayloadAction<{ id: string }>) => {
      const selectedGroup =
        state.groups?.find((group) => group.id === action.payload.id) || null;
      const currentGroup: GroupInterface = {
        ...selectedGroup,
        imageData: {
          images: null,
          currentImage: null,
          fetchImages: true,
        },
        userData: [],
      } as GroupInterface;
      const newState = {
        ...state,
        currentGroup: currentGroup,
      };
      return newState;
    },
    clearGroup: () => {
      return initialState;
    },
    setFetchGroups: (state) => {
      state.fetchGroups = true;
    },
    unsetFetchGroups: (state) => {
      state.fetchGroups = false;
    },
    setGroupImageData: (state, action: PayloadAction<ImageInterface[]>) => {
      if (state.currentGroup) {
        state.currentGroup.imageData = {
          images: [...action.payload],
          currentImage: null,
          fetchImages: false,
        };
      }
      return state;
    },
    setCurrentGroupImage: (state, action: PayloadAction<{ id: string }>) => {
      const allImages = state.currentGroup?.imageData.images ?? [];
      const currentImage =
        allImages.find((image) => image.id === action.payload.id) ?? null;
      if (state.currentGroup) {
        state.currentGroup.imageData = {
          images: allImages,
          currentImage: currentImage,
          fetchImages: false,
        };
      }
      return state;
    },
    removeCurrentGroupImage: (state, action: PayloadAction<{ id: string }>) => {
      const allImages = state.currentGroup?.imageData.images ?? [];
      const newImages =
        allImages.filter((image) => image.id !== action.payload.id) ?? null;
      if (state.currentGroup) {
        state.currentGroup.imageData = {
          images: newImages,
          currentImage: null,
          fetchImages: false,
        };
      }
      return state;
    }
  },
});

export const {
  setGroups,
  setCurrentGroup,
  clearGroup,
  setFetchGroups,
  unsetFetchGroups,
  setCurrentGroupImage,
  setGroupImageData,
  removeCurrentGroupImage,
} = GroupSlice.actions;

export const groupListenerMiddleware = createListenerMiddleware();
groupListenerMiddleware.startListening({
  actionCreator: setFetchGroups,
  effect: async (_, listenerApi) => {
    try {
      const res = await apiClientObj.get(routes.GET_ALL_GROUPS);
      listenerApi.dispatch(setGroups(res));
      listenerApi.dispatch(unsetFetchGroups());
    } catch (err) {
      console.error(err);
    }
  },
});

export default GroupSlice.reducer;
