package main

import (
	"net/http"
	"database/sql"
)

var db *sql.DB
func main() {
	db = initDB("storage.db")
	migrate(db)

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/view/db/", viewDbHandler)
	http.HandleFunc("/edit/db/", editDbHandler)
	http.HandleFunc("/save/db/", saveDbHandler)
	http.ListenAndServe(":8080", nil)
}


