package routes

import (
    "cms-api/handlers"
    "database/sql"
    "net/http"
)

func SetupRoutes(db *sql.DB) *http.ServeMux {
    mux := http.NewServeMux()

    mux.HandleFunc("/api/movies", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            handlers.GetMovies(w, r, db)
        case "POST":
            handlers.CreateMovie(w, r, db)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    mux.HandleFunc("/api/movies/", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            handlers.GetMovieByID(w, r, db)
        case "PUT":
            handlers.UpdateMovie(w, r, db)
        case "DELETE":
            handlers.DeleteMovie(w, r, db)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    return mux
}
