import { createSlice, type PayloadAction } from "@reduxjs/toolkit"

export interface Modal{
    data: {
        heading?: string,
        key?: string
    },
    name: string | null
}

const initialState: Modal = {
    data: {},
    name: null
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
