import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import type { UserInterface } from "../common/interfaces";

const initialState = null as UserInterface | null;

const UserSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    setUser: (_, action: PayloadAction<UserInterface>) => {
      return action.payload;
    },
    clearUser: () => {
      return initialState;
    },
  },
});

export const { setUser, clearUser } = UserSlice.actions;
export default UserSlice.reducer;
