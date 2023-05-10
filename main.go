package main

import (
	"context"
	"flag"

	"github.com/fulltimegodev/hotel-reservation/api"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27018"
const dbname = "hotel-reservation"
const userColl = "users"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	coll := client.Database(dbname).Collection(userColl)
	coll.InsertOne(ctx, map[string]string{"name": "John Doe"})

	// flags
	port := flag.String("port", ":3333", "The listen address of the API server")

	// versioning
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	// routes
	app.Get("/foo", handleFoo)
	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	// init
	app.Listen(*port)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "working just fine!"})
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "user endpoint"})
}
