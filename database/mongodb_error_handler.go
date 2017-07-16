package database

import (
	"gopkg.in/mgo.v2"
	"net/http"
)
/**
err cannot be nil
returns http status, message, and boolean describing if error is a not found error
 */
func HandleError(err error) (int, bool) {
	switch err {
		case mgo.ErrNotFound:
			return http.StatusNotFound, true
		default:
			return http.StatusBadRequest, false
	}
}