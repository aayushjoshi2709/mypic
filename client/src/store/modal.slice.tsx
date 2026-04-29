import { createSlice, type PayloadAction } from "@reduxjs/toolkit";
import type { DeleteModalDataInterface } from "../component/Modal/DeleteModal/DeleteModal";
import type { PreviewModalDataInterface } from "../component/Modal/PreviewModal/PreviewModal";

export interface Modal {
  data?: DeleteModalDataInterface | PreviewModalDataInterface | null;
  name: string;
}

const initialState: Modal = {
  data: null,
  name: "",
};

const ModalSlice = createSlice({
  name: "modal",
  initialState,
  reducers: {
    setModal: (state, action: PayloadAction<Modal>) => {
      state = action.payload;
      return state;
    },
    clearModal: () => {
      return initialState;
    },
  },
});

export const { setModal, clearModal } = ModalSlice.actions;

export default ModalSlice.reducer;
