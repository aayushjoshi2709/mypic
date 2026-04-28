import { createSlice } from "@reduxjs/toolkit";
import type { PayloadAction } from "@reduxjs/toolkit";
import type { UserDataInterface } from "../common/interfaces";

const initialState = null as UserDataInterface | null;

const UserSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    setUser: (_, action: PayloadAction<UserDataInterface>) => {
      return action.payload;
    },
    clearUser: () => {
      return initialState;
    },
  },
});

export const { setUser, clearUser } = UserSlice.actions;
export default UserSlice.reducer;
