package main

import (
	"context"
	"flag"

	"github.com/fulltimegodev/hotel-reservation/api"
	"github.com/fulltimegodev/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27018"
const dbname = "hotel-reservation"
const userColl = "users"

var config = fiber.Config(fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
})

func main() {
	// flags
	port := flag.String("port", "3333", "The listen address of the API server")
	flag.Parse()

	// database
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		panic(err)
	}

	// handlers init
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	// versioning
	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	// routes
	app.Get("/foo", handleFoo)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	// init
	app.Listen(":" + *port)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "working just fine!"})
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "user endpoint"})
}
