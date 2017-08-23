import React from "react";
import client from "../client";

class ReleaseDao {
    static getAllReleases() {
        return client({
            method: 'GET',
            path: 'http://localhost:8080/api/v1/releases'
        }).then(response => {
            return response.entity;
        }).catch(error => {
            return error;
        });
    }
}

export default ReleaseDao;
