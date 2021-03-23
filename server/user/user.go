package user

import (
	"context"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"twitter-wall-server/config"
	"twitter-wall-server/database"
)

// User is the user
type User struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Email     string    `json:"email" bson:"email"`
	Username  string    `json:"username" bson:"username"`
	Password  string    `json:"password" bsom:"password"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

// Hashes the password for the database
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Finds the user by email
func getUserbyEmail(email string) (*User, error) {
	var dbEntry User
	if err := database.DBUsers.FindOne(context.Background(), bson.M{"email": email}).Decode(&dbEntry); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &dbEntry, nil
}

// Finds the user by username
func getUserByUsername(username string) (*User, error) {
	var dbEntry User
	if err := database.DBUsers.FindOne(context.Background(), bson.M{"username": username}).Decode(&dbEntry); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &dbEntry, nil
}

// Login for user
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}

	var input LoginInput
	var userData User
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please insert valid data"})
	}

	email, err := getUserbyEmail(input.Identity)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid input!"})
	}

	user, err := getUserByUsername(input.Identity)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid input!"})
	}

	if email == nil && user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid input!"})
	}

	if email != nil {
		userData = User{
			ID:        email.ID,
			Email:     email.Email,
			Username:  email.Username,
			Password:  email.Password,
			CreatedAt: email.CreatedAt,
		}
	} else {
		userData = User{
			ID:        user.ID,
			Email:     user.Email,
			Username:  user.Username,
			Password:  user.Password,
			CreatedAt: user.CreatedAt,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(input.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid input!"})
	}

	// Creates JWT Token with claims, like username and id
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userData.Username
	claims["id"] = userData.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte(config.Conf.JWTSecret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server issue!"})
	}

	return c.JSON(fiber.Map{
		"token":    t,
		"username": userData.Username,
	})
}

// Register for user
func Register(c *fiber.Ctx) error {
	var user User
	var userData User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please insert valid data"})
	}

	if user.Email == "" || user.Password == "" || user.Username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please insert valid data"})
	}

	if err = database.DBUsers.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&userData); err != mongo.ErrNoDocuments {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Username taken!"})
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid input!"})
	}

	// stores user in database with changed password as bcrypt hash
	user.Password = hash
	user.CreatedAt = time.Now().UTC()
	if _, err = database.DBUsers.InsertOne(context.Background(), user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error in database!"})
	}

	return c.SendStatus(fiber.StatusOK)
}
