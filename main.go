package main

import (
	"context"
	"fmt"
	"git-project-management/config"
	app "git-project-management/internal"
	"log"
	"time"

	"gitea.com/logicamp/lc"
	"github.com/danielgtaylor/huma/v2"
	humaFiber "github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	config, _ := lc.GetConfig[config.Config](&config.Config{})
	fiberApp := fiber.New()
	fiberApp.Use(cors.New())

	// Or extend your config for customization
	fiberApp.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	humaConfig := huma.DefaultConfig("Git Project Management", "1.0.0")
	humaConfig.Servers = []*huma.Server{{URL: config.BASE_URL}}
	humaConfig.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"auth": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}
	api := humaFiber.New(fiberApp, humaConfig)

	// database init ---------
	db := pg.Connect(&pg.Options{
		Addr:     config.PG_HOST,
		User:     config.PG_USER,
		Password: config.PG_PASSWORD,
		Database: config.PG_DATABASE,
	})
	defer db.Close()

	if err := db.Ping(context.Background()); err != nil {
		panic(err)
	}
	// ------------------------

	uri := "mongodb://localhost:27017"

	// Create a new client and connect to the server
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer client.Disconnect(ctx)

	// Select database and collection
	database := client.Database("w-mohamadkhani-bazargam-6628c03f-56c061745a-6f0314")
	collection := database.Collection("activity")

	// Define a filter (empty filter counts all documents)
	filter := bson.M{}

	// Get the count of documents
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of documents in the collection: %d\n", count)
	app.Setup(&api, db)

	log.Fatal(fiberApp.Listen(fmt.Sprintf(":%s", config.PORT)))
}
