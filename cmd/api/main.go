package main

import (
	"context"
	"log"
	"time"

	"github.com/cedrickewi/hotel-reservation/internals/db"
	"github.com/cedrickewi/hotel-reservation/internals/services"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//connect to database
	client, err := db.ConnectToMongo()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	cfg := config{
		port: 8080,
		addr: ":80",
	}

	logger := zap.Must(zap.NewProduction()).Sugar()

	store := services.NewStorage(client)

	app := &application{
		store:  &store,
		config: cfg,
		logger: logger,
	}

	mux := routes(app)

	app.run(mux, app.config.port)

}
