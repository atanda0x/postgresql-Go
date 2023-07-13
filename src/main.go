package main

import (
	"log"

	"github.com/atanda0x/postgresql-Go/src/models"
)

func main() {
	db, err := models.InitDB()
	if err != nil {
		log.Println(db)
	}
}
