package main

import (
    "cms-api/routes"
    "database/sql"
    "fmt"
    "log"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
)

func connectToDB() (*sql.DB, error) {
    dsn := "root:@tcp(127.0.0.1:3306)/msf"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    fmt.Println("Berhasil terhubung ke database.")
    return db, nil
}

func main() {
    db, err := connectToDB()
    if err != nil {
        log.Fatal("Gagal koneksi ke database:", err)
    }
    defer db.Close()

    mux := routes.SetupRoutes(db)

    fs := http.FileServer(http.Dir("public"))
    mux.Handle("/", fs)

    log.Println("Server berjalan di http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}