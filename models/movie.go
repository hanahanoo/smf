package models

type Movie struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Genres      string `json:"genres"`
    Duration    string `json:"duration"` 
    Artist      string `json:"artist"`
}
