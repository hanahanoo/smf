package handlers

import (
    "cms-api/models"
    "database/sql"
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
)

func GetMovies(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    rows, err := db.Query("SELECT id, title, description, genres, duration, artist FROM movies")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var movies []models.Movie
    for rows.Next() {
        var m models.Movie
        rows.Scan(&m.ID, &m.Title, &m.Description, &m.Genres, &m.Duration, &m.Artist)
        movies = append(movies, m)
    }
    json.NewEncoder(w).Encode(movies)
}

func GetMovieByID(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    id := extractID(r.URL.Path)
    var m models.Movie
    err := db.QueryRow("SELECT id, title, description, genres, duration, artist FROM movies WHERE id = ?", id).
        Scan(&m.ID, &m.Title, &m.Description, &m.Genres, &m.Duration, &m.Artist)
    if err != nil {
        http.NotFound(w, r)
        return
    }
    json.NewEncoder(w).Encode(m)
}

func CreateMovie(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    var m models.Movie
    json.NewDecoder(r.Body).Decode(&m)
    result, err := db.Exec("INSERT INTO movies (title, description, genres, duration, artist) VALUES (?, ?, ?, ?, ?)",
        m.Title, &m.Description, &m.Genres, &m.Duration, &m.Artist)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    id, _ := result.LastInsertId()
    m.ID = int(id)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(m)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    id := extractID(r.URL.Path)
    var m models.Movie
    json.NewDecoder(r.Body).Decode(&m)
    _, err := db.Exec("UPDATE movies SET title = ?, description = ?, genres = ?, duration = ?, artist = ? WHERE id = ?",
        m.Title, &m.Description, &m.Genres, &m.Duration, &m.Artist, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    m.ID = id
    json.NewEncoder(w).Encode(m)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    id := extractID(r.URL.Path)
    _, err := db.Exec("DELETE FROM movies WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func extractID(path string) int {
    parts := strings.Split(path, "/")
    idStr := parts[len(parts)-1]
    id, _ := strconv.Atoi(idStr)
    return id
}
