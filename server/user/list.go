package user

import (
	"context"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"twitter-wall-server/database"
)

type list struct {
	Owner       string    `json:"owner" bson:"owner"`
	Searchterms []search  `json:"searchterms" bson:"searchterms"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
}

type search struct {
	Name       string `json:"name" bson:"name"`
	Searchterm string `json:"searchterm" bson:"searchterm"`
}

// GetList return the lists of the user
func GetList(c *fiber.Ctx) error {
	var list list
	id := getToken(c)
	if err := database.DBLists.FindOne(context.Background(), bson.M{"owner": id}).Decode(&list); err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not found!"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server issue!"})
	}
	return c.Status(fiber.StatusOK).JSON(list)
}

// CreateList creates a list of the user
func CreateList(c *fiber.Ctx) error {
	var list list
	c.BodyParser(&list)
	if list.Searchterms == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input!"})
	}
	id := getToken(c)
	list.Owner = id
	list.CreatedAt = time.Now().UTC()
	list.UpdatedAt = time.Now().UTC()

	_, err := database.DBLists.InsertOne(context.Background(), list)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server issue!"})
	}
	database.DBLists.FindOne(context.Background(), bson.M{"owner": id}).Decode(&list)

	return c.Status(fiber.StatusOK).JSON(list)
}

// UpdateList updates the list of the user
func UpdateList(c *fiber.Ctx) error {
	var list list
	c.BodyParser(&list)
	if list.Searchterms == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input!"})
	}

	id := getToken(c)
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "searchterms", Value: list.Searchterms},
				{Key: "updatedAt", Value: time.Now().UTC()},
			},
		},
	}

	if err := database.DBLists.FindOneAndUpdate(context.Background(), bson.M{"owner": id}, update).Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not found!"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server issue!"})
	}
	database.DBLists.FindOne(context.Background(), bson.M{"owner": id}).Decode(&list)
	return c.Status(fiber.StatusOK).JSON(list)
}

// DeleteList deletes the list of the user
func DeleteList(c *fiber.Ctx) error {
	id := getToken(c)
	if _, err := database.DBLists.DeleteOne(context.Background(), bson.M{"owner": id}); err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not found!"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server issue!"})
	}
	return c.SendStatus(fiber.StatusOK)
}

// Gets the JWT Token from the request and the claimed user id
func getToken(c *fiber.Ctx) string {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}
