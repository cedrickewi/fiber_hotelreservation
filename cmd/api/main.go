package main

import (
	"log"

	"github.com/cedrickewi/hotel-reservation/internals/db"
)

func main() {
	_, err := db.ConnectToMongo()
	if err != nil {
		log.Fatal(err)
	}

}
