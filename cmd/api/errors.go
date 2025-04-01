package main

import (
	"net/http"
)

// errorResponse is a helper function that writes a JSON response with an error message and status code.
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// serverErrorResponse is a helper function that writes a JSON response with a server error message.
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Println(err)
	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// notFoundResponse is a helper function that writes a JSON response with a not found error message.
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// badRequestResponse is a helper function that writes a JSON response with a bad request error message.
func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}
