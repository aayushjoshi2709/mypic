import { createSlice, type PayloadAction } from "@reduxjs/toolkit";
import type { GroupDataInterface, GroupInterface } from "../common/interfaces";

const initialState: GroupDataInterface = {
  groups: null,
  currentGroup: null,
  currentPage: 0,
  currentLimit: 10,
  totalPages: null,
};

const GroupSlice = createSlice({
  name: "group",
  initialState: initialState,
  reducers: {
    setGroups: (state, action: PayloadAction<GroupInterface[]>) => {
      state.groups = [...action.payload];
    },
    setCurrentPage: (state, action: PayloadAction<number>) => {
      state.currentPage = action.payload;
    },
    setCurrentLimit: (state, action: PayloadAction<number>) => {
      state.currentLimit = action.payload;
    },
    setTotalPages: (state, action: PayloadAction<number | null>) => {
      state.totalPages = action.payload;
    },
    setCurrentGroup: (state, action: PayloadAction<string | null>) => {
      state.currentGroup =
        (state.groups?.find(
          (group) => group.id === action.payload,
        ) as GroupInterface) || null;
    },
    clearGroups: () => {
      return initialState;
    },
    addGroup: (state, action: PayloadAction<GroupInterface>) => {
      if (state.groups) {
        state.groups.concat(action.payload);
      } else {
        state.groups = [action.payload];
      }
    },
  },
});

export const {
  setGroups,
  setCurrentPage,
  setCurrentLimit,
  setTotalPages,
  setCurrentGroup,
  addGroup,
  clearGroups,
} = GroupSlice.actions;

export default GroupSlice.reducer;
