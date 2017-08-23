import React from "react";
import client from "../client";

class ModuleDao {
    static getAllModules(release) {
        let releaseQuery = "";
        if (release) {
            releaseQuery = "?releaseVersion=" + release;
        }
        let path = 'http://localhost:8080/api/v1/modules' + releaseQuery;
        return client({
            method: 'GET',
            path: path
        }).then(response => {
            return response.entity;
        }).catch(error => {
            return error;
        });
    }
}

export default ModuleDao;

