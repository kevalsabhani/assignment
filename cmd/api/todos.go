package main

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/kevalsabhani/assignment/internal/data"
)

// createHandler handles the creation of a new todo item.
func (app *application) createHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Owner       string `json:"owner"`
		Priority    string `json:"priority"`
		DueDate     string `json:"due_date"`
		Completed   bool   `json:"completed"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	todo := data.Todo{
		ID:          rand.Int64N(1000000) + 1,
		Completed:   input.Completed,
		Title:       input.Title,
		Description: input.Description,
		Owner:       input.Owner,
		Priority:    data.PriorityLevel(input.Priority),
		DueDate:     input.DueDate,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}

	id, err := app.models.Todos.Create(todo)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, envelope{"id": id}, nil)
}

// getAllHandler handles the retrieval of all todo items.
func (app *application) getAllHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := app.models.Todos.GetAll()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"todos": todos}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

// getByIDHandler handles the retrieval of a todo item by its ID.
func (app *application) getByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL query parameters
	id, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	todo, err := app.models.Todos.GetByID(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if todo.ID == 0 {
		app.notFoundResponse(w, r)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"todo": todo}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// updateHandler handles the update of a todo item.
func (app *application) updateHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL query parameters
	id, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Owner       string `json:"owner"`
		Priority    string `json:"priority"`
		DueDate     string `json:"due_date"`
		Completed   bool   `json:"completed"`
	}
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Check if the todo exists
	existingTodo, err := app.models.Todos.GetByID(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if existingTodo.ID == 0 {
		app.notFoundResponse(w, r)
		return
	}

	// Update the todo with the new values
	todo := data.Todo{
		ID:          id,
		Title:       input.Title,
		Description: input.Description,
		Owner:       input.Owner,
		Priority:    data.PriorityLevel(input.Priority),
		DueDate:     input.DueDate,
		CreatedAt:   existingTodo.CreatedAt,
		UpdatedAt:   time.Now().Format(time.RFC3339),
		Completed:   input.Completed,
	}

	err = app.models.Todos.Update(todo)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusNoContent, envelope{"updated": true}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

// deleteHandler handles the deletion of a todo item.
func (app *application) deleteHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL query parameters
	id, err := app.readIDParam(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.models.Todos.Delete(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusNoContent, envelope{"deleted": true}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
