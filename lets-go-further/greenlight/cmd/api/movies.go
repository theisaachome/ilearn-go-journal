package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.isaachome.net/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Create a new movie")
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
