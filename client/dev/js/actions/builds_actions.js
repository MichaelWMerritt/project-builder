import buildsDao from "../dao/builds_dao";
import * as types from "./action_types";

export function updateBuildFormState(buildFormState) {
    return {type: types.BUILD_FORM_STATE, buildFormState};
}

export function loadAllBuildReferences() {
    return function(dispatch) {
        return buildsDao.getAllBuildReferences().then(buildReferences => {
            dispatch(loadAllBuildReferencesSuccess(buildReferences));
        }).catch(error => {
            throw(error);
        });
    };
}

export function loadAllBuildReferencesSuccess(buildReferences) {
    return {type: types.LOAD_ALL_BUILD_REFERENCES_SUCCESS, buildReferences};
}

export function createBuild(build) {
    return function() {
        return buildsDao.createBuild(build).then(() => {
            loadAllBuildReferences();
        }).catch(error => {
            throw(error)
        });
    }
}