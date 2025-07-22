package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.isaachome.net/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMovieHanlder(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	move := data.Move{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Inception",
		Year:      2010,
		Runtime:   148,
		Genres:    []string{"Sci-Fi", "Action", " Thriller"},
		Version:   1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": move}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
