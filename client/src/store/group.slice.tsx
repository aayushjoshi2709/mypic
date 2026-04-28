import { createListenerMiddleware, createSlice, type PayloadAction } from "@reduxjs/toolkit"
import type { GroupDataInterface, GroupInterface } from "../common/interfaces"
import { apiClientObj } from "../common/apiClient"
import { routes } from "../common/routes"

const initialState: GroupDataInterface = {
    groups: null,
    currentGroup: null,
    fetchGroups: true
}

const GroupSlice = createSlice({
    name: "group",
    initialState: initialState,
    reducers: {
        setGroups: (state, action: PayloadAction<GroupInterface[]>) => {
            state.groups = [...action.payload];
        },
        setCurrentGroup: (state, action: PayloadAction<{id: string}>) =>{
            const newState = {
                ...state,
                currentGroup: state.groups?.find((group) => group.id === action.payload.id) || null
            }
            return newState
        },
        clearGroup: () => {
            return initialState;
        },
        setFetchGroups: (state)=>{
            state.fetchGroups = true;
        },
        unsetFetchGroups: (state)=>{
            state.fetchGroups = false;
        }
    }
})

export const {setGroups, setCurrentGroup, clearGroup, setFetchGroups, unsetFetchGroups} = GroupSlice.actions;


export const groupListenerMiddleware = createListenerMiddleware();
groupListenerMiddleware.startListening({
  actionCreator: setFetchGroups,
  effect: async (_, listenerApi) => {
    try {
      const res = await apiClientObj.get(routes.GET_ALL_GROUPS)
      listenerApi.dispatch(setGroups(res));
      listenerApi.dispatch(unsetFetchGroups());
    } catch (err) {
      console.error(err)
    }
  }
})

export default GroupSlice.reducer;
