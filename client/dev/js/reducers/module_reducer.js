import * as types from "../actions/action_types";
import initialState from "./initial_state";

export default function moduleReducer(state = initialState.modules, action) {
    switch(action.type) {
        case types.LOAD_ALL_MODULES_SUCCESS:
        case types.LOAD_MODULES_FOR_RELEASE_SUCCESS:
            return action.modules;
        default:
            return state;
    }
}

