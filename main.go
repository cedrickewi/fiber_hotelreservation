package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/cedrickewi/hotel-reservation/api"
	"github.com/cedrickewi/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userColl = "users"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	usercollection := client.Database(dbname).Collection(userColl)

	user := types.User{
		FirstName: "John",
		LastName:  "paul",
	}
	res, err := usercollection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("user inserted", res)

	listenAddr := flag.String("listenAddr", ":5000", "The listen addr of the api server")
	flag.Parse()
	app := fiber.New()
	apiV1 := app.Group("api/v1")

	apiV1.Get("/user", api.HandleGetUsers)
	apiV1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenAddr)
}
