import moduleDao from "../dao/module_dao";
import * as types from "./action_types";

export function loadAllModules() {
    return function(dispatch) {
        return moduleDao.getAllModules().then(modules => {
            dispatch(loadAllModulesSuccess(modules));
        }).catch(error => {
            throw(error);
        });
    };
}

export function loadModulesForRelease(release) {
    return function(dispatch) {
        return moduleDao.getAllModules(release).then(modules => {
            dispatch(loadModulesForReleaseSuccess(modules));
        }).catch(error => {
            throw(error);
        });
    };
}

export function loadAllModulesSuccess(modules) {
    return {type: types.LOAD_ALL_MODULES_SUCCESS, modules};
}

export function loadModulesForReleaseSuccess(modules) {
    return {type: types.LOAD_MODULES_FOR_RELEASE_SUCCESS, modules};
}
