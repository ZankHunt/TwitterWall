package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v2"

	"twitter-wall-server/config"
	"twitter-wall-server/database"
	"twitter-wall-server/twitter"
	"twitter-wall-server/user"
)

func main() {
	// Creates a new Config
	config.NewConfig()
	database.ConnectDB(config.Conf.MongoURL)

	app := fiber.New()
	app.Use(logger.New())

	// Hosts the client
	app.Static("/", "../client/public")

	// user /login & /register
	u := app.Group("/user")
	u.Post("/login", user.Login)
	u.Post("/register", user.Register)

	// /search
	app.Get("/search", twitter.GetSearch)

	// now follows restricted area. the user needs a JWT token
	app.Use(jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.SendStatus(fiber.StatusUnauthorized)
		},
		SigningKey: []byte(config.Conf.JWTSecret),
	}))

	// list of the user
	app.Get("/lists", user.GetList)
	app.Post("/lists", user.CreateList)
	app.Put("/lists", user.UpdateList)
	app.Delete("/lists", user.DeleteList)

	// Starts the server with https / TLS. When there is an issue, will stop the program
	log.Fatal(app.ListenTLS(":443", "./certs/tls.crt", "./certs/tls.key"))
}
