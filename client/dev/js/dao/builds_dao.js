import React from "react";
import client from "../client";

class BuildsDao {
    static getAllBuildReferences() {
        let path = 'http://localhost:8080/api/v1/builds';
        return client({
            method: 'GET',
            path: path
        }).then(response => response.entity)
            .catch(error => error);
    }

    static createBuild(build) {
        let path = 'http://localhost:8080/api/v1/builds';
        return client({
            method: 'POST',
            path: path,
            entity: JSON.stringify(build)
            // headers: {
            //     'Content-Type': 'application/json',
            //     "Access-Control-Allow-Origin": "*",
            //     "Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept"
            // }
        }).then(response => {
            return response.entity
        }).catch(error => {
            return error;
        });
    }
}

export default BuildsDao;
