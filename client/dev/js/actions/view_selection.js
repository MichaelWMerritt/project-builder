import * as types from "./action_types";

export const selectView = (viewId) => {
    return {
        type: types.VIEW_SELECTED,
        payload: viewId
    }
};