import {applyMiddleware, createStore} from "redux";
import rootReducer from "../reducers/root_reducer";
import thunk from "redux-thunk";

export default function configureStore() {
    return createStore(
        rootReducer,
        applyMiddleware(thunk)
    );
}
