package main

import "github.com/julienschmidt/httprouter"

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	// Define the routes and their handlers
	router.HandlerFunc("POST", "/todos", app.createHandler)
	router.HandlerFunc("GET", "/todos", app.getAllHandler)
	router.HandlerFunc("GET", "/todos/:id", app.getByIDHandler)
	router.HandlerFunc("PUT", "/todos/:id", app.updateHandler)
	router.HandlerFunc("DELETE", "/todos/:id", app.deleteHandler)

	return router
}
