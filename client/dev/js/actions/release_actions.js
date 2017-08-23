import releaseDao from "../dao/release_dao";
import {loadModulesForRelease} from "./module_actions";
import * as types from "./action_types";

export function loadReleases() {
    return function(dispatch) {
        return releaseDao.getAllReleases().then(releases => {
            dispatch(loadReleasesSuccess(releases));
            dispatch(loadModulesForRelease(releases[0]._id));
        }).catch(error => {
            throw(error);
        });
    };
}

export function loadReleasesSuccess(releases) {
    return {type: types.LOAD_RELEASES_SUCCESS, releases};
}