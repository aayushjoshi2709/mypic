import { createSlice, type PayloadAction } from "@reduxjs/toolkit"
import type { DeleteModalInterface } from "../component/Modal/DeleteModal/DeleteModal"

export interface Modal{
    data?: DeleteModalInterface | null,
    name: string
}

const initialState: Modal = {
    data: null,
    name: ""
}

const ModalSlice = createSlice({
    name: "modal",
    initialState,
    reducers:{
        setModal: (state, action: PayloadAction<Modal>)=>{
            state = action.payload
            return state;
        },
        clearModal: ()=>{
            return initialState
        }
    }
})

export const {setModal, clearModal} = ModalSlice.actions

export default ModalSlice.reducer
