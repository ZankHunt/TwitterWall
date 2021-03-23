package twitter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"twitter-wall-server/config"
	"twitter-wall-server/database"
)

// GetSearch searches on the Twitter API /search/recent
func GetSearch(c *fiber.Ctx) error {
	// Gets the query from the /search?query=
	query := url.QueryEscape(c.Query("query"))
	dbEntry := DatabaseEntry{}
	frontEndData := []FrontEndResponse{}

	err := database.DBTweets.FindOne(context.Background(), bson.M{"query": query}).Decode(&dbEntry)
	if err != nil && err != mongo.ErrNoDocuments {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if err == mongo.ErrNoDocuments { // No Document
		frontEndData, _ = newSearch(c, query, frontEndData)
	} else { // Document is not older than 5 min
		frontEndData = dbEntry.FrontEndResponse
	}

	// Send the data to the client as JSON
	return c.Status(fiber.StatusOK).JSON(frontEndData)
}

// SearchRequest creates a new searchRequest to the Twitter API with the search query and options and tokens
func searchRequest(query string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/recent?query=%s%s", config.Conf.APIURL, query, config.Conf.Options), nil)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Conf.BearerToken))
	return request
}

func newSearch(c *fiber.Ctx, query string, frontEndData []FrontEndResponse) ([]FrontEndResponse, error) {
	// Creates a http Client with a timeout of 1 min
	client := http.Client{
		Timeout: time.Minute,
	}

	// http Client sends a search request with the the query. client Do closes the body
	response, err := client.Do(searchRequest(query))
	if err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if response.StatusCode != fiber.StatusOK {
		return nil, c.Status(response.StatusCode).JSON(response.Body)
	}

	// Decodes the response from the Twitter API to JSON
	search := TweetResponseSearch{}
	if err := json.NewDecoder(response.Body).Decode(&search); err != nil {
		return nil, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Copies needed values from the response
	data := TweetResponse{
		Data:     search.Data,
		Includes: search.Includes,
	}

	// Gets the needed info from the JSON response
	frontEndData = createFrontEndJSON(data)

	if len(frontEndData) > 0 {
		dbEntry := DatabaseEntry{
			Expiration:       time.Now().Add(5 * time.Minute).UTC(),
			Query:            query,
			FrontEndResponse: frontEndData,
		}
		_, err = database.DBTweets.InsertOne(context.Background(), dbEntry)
		if err != nil {
			return nil, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}
	return frontEndData, nil
}
