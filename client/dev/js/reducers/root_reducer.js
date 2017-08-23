import {combineReducers} from "redux";
import UserReducer from "./reducers-users";
import ActiveUserReducer from "./reducer-active-user";
import ActiveViewReducer from "./active_view_reducer";
import ReleaseReducer from "./release_reducer";
import ModuleReducer from "./module_reducer";
import {buildFormStateReducer, buildReferencesReducer} from "./build_reducer";

const rootReducer = combineReducers({
    users: UserReducer,
    activeUser: ActiveUserReducer,
    activeView: ActiveViewReducer,
    releases: ReleaseReducer,
    modules: ModuleReducer,
    buildFormState: buildFormStateReducer,
    buildReferences: buildReferencesReducer
});

export default rootReducer;