package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(*sql.DB, error) {
	var err error
	db, err := sql.Open("postgres", "postgres://atanda0x:ethereumsol@localhost/mydb?sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		// Create model for the web service
		stmt, err := db.Prepare("CREATE TABLE WEB_URL(ID SERIAL PRIMARY KEY, URL TEXT NOT NULL);")

		if err != nil {
			log.Println(err)
		}

		res, err := stmt.Exec()
		log.Println(res)
		if err != nil {
			log.Println(err)
		}
		return db, nil
	}
}
