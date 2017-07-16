package controller

import (
	"github.com/michaelwmerritt/project-builder/database"
	"encoding/json"
	"net/http"
	"log"
)

func HandleError(w http.ResponseWriter, err error) {
	status, _ := database.HandleError(err)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err.Error())
}

func HandleNotFoundError(w http.ResponseWriter, err error, id string) {
	status, notFound := database.HandleError(err)
	w.WriteHeader(status)
	var message string
	if notFound {message = id + " " + err.Error()}
	json.NewEncoder(w).Encode(message)
}

func HandleServerError(w http.ResponseWriter, logMessage string) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode("Server Error")
	log.Panic(logMessage)
}
