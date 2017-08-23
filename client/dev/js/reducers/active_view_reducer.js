import * as types from "../actions/action_types";
import {BUILD_REFERENCES_VIEW_ID} from "../containers/build_references_view.jsx";

export default function (state = BUILD_REFERENCES_VIEW_ID, action) {
    switch (action.type) {
        case types.VIEW_SELECTED:
            return action.payload;
            break;
    }
    return state;
}
