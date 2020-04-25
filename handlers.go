package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func ConfirmHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	key := vars["id"]

	user, err := FindUserKey(key)
	if err != nil || user == nil {
		sendError(w, notFound)
		return
	}
	if user.Key == key && user.Confirmed {
		response := requestResponse{
			Email:   user.Email,
			Message: "email address has already been confirmed",
		}
		sendJSON(w, response)
		return
	}

	if user.Key == key {
		user.Confirmed = true
		user.Update()
	}
	response := requestResponse{
		Email:   user.Email,
		Message: "email address is now confirmed",
	}
	sendJSON(w, response)
}

func UnsubscribeHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	email := r.URL.Query().Get("email")

	user, err := FindUser(email)
	if err != nil || user == nil {
		sendError(w, notFound)
		return
	}

	user.Delete()

	status := &requestResponse{
		Email:   user.Email,
		Message: "email has been unsubscribed",
	}

	sendJSON(w, status)
}

func ResendHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	email := r.URL.Query().Get("email")

	user, err := FindUser(email)
	if err != nil || user == nil {
		sendError(w, notFound)
		return
	}
	status := &requestResponse{
		Email: user.Email,
	}
	if user.Confirmed {
		status.Message = "email has already been confirmed"
	} else {
		status.Message = "email has not been confirmed yet"
	}
	sendJSON(w, status)
}

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	email := r.URL.Query().Get("email")

	user, err := FindUser(email)
	if err != nil || user == nil {
		sendError(w, notFound)
		return
	}
	status := &requestResponse{
		Email: user.Email,
	}
	if user.Confirmed {
		status.Message = "email has been confirmed"
	} else {
		status.Message = "email has not been confirmed yet " + user.Key
	}
	sendJSON(w, status)
}

func SendHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	user, err := FindUser("info@socialeck.com")
	if err != nil || user == nil {
		sendError(w, notFound)
		return
	}

	err = SendEmail(user)
	if err != nil {
		sendError(w, err)
		return
	}
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req requestJSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		sendError(w, err)
		return
	}

	user, err := FindUser(req.Email)
	if err != nil {
		sendError(w, notFound)
		return
	}

	if user != nil {
		response := &requestResponse{
			Email:   user.Email,
			Message: "this email need to be confirmed",
		}
		sendJSON(w, response)
		return
	}

	user = &User{
		Email: req.Email,
		Key:   RandomString(32),
		Sent:  0,
	}

	user.Create()

	response := &requestResponse{
		Email:   user.Email,
		Message: "check email",
	}

	sendJSON(w, response)
}

func sendJSON(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}

func sendError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&errorResponse{err.Error()})
}
