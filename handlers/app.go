package handlers

import (
	"database/sql"
	"net/http"
	"strings"

	"server/repository"
)

type App struct {
	userHandler    *UserHandler
	testHandler    *TestHandler
	commentHandler *CommentHandler
}

func NewApp(db *sql.DB) *App {
	return &App{
		userHandler:    NewUserHandler(repository.NewUserRepository(db)),
		testHandler:    NewTestHandler(repository.NewTestRepository(db)),
		commentHandler: NewCommentHandler(repository.NewCommentRepository(db)),
	}
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api")

	switch {
	case strings.HasPrefix(path, "/users"):
		switch r.Method {
		case http.MethodPost:
			app.userHandler.Create(w, r)
		case http.MethodGet:
			app.userHandler.Get(w, r)
		case http.MethodPut:
			app.userHandler.Update(w, r)
		default:
			respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}

	case strings.HasPrefix(path, "/tests"):
		if strings.Contains(path, "/user") {
			app.testHandler.GetByUser(w, r)
			return
		}
		switch r.Method {
		case http.MethodPost:
			app.testHandler.Create(w, r)
		case http.MethodGet:
			app.testHandler.Get(w, r)
		case http.MethodPut:
			app.testHandler.Update(w, r)
		default:
			respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}

	case strings.HasPrefix(path, "/comments"):
		switch r.Method {
		case http.MethodPost:
			app.commentHandler.Create(w, r)
		case http.MethodGet:
			app.commentHandler.GetByTest(w, r)
		case http.MethodPut:
			app.commentHandler.Update(w, r)
		default:
			respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}

	default:
		respondWithError(w, http.StatusNotFound, "Not found")
	}
}
