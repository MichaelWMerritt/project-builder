import * as types from "../actions/action_types";
import initialState from "./initial_state";

export function buildReferencesReducer(state = initialState.buildReferences, action) {
    switch(action.type) {
        case types.LOAD_ALL_BUILD_REFERENCES_SUCCESS:
            return action.buildReferences;
        default:
            return state;
    }
}

export function buildFormStateReducer(state = initialState.buildFormState, action) {
    switch(action.type) {
        case types.BUILD_FORM_STATE:
            return action.buildFormState;
        default:
            return state;
    }
}
