import * as types from "../actions/action_types";
import initialState from "./initial_state";

export default function releaseReducer(state = initialState.releases, action) {
    switch(action.type) {
        case types.LOAD_RELEASES_SUCCESS:
            return action.releases;
        default:
            return state;
    }
}
