package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config returns a configuration for the API
type Config struct {
	MongoURL    string
	BearerToken string
	APIURL      string
	Options     string
	JWTSecret   string
}

// Conf is the Config
var Conf Config

// NewConfig creates a new config
func NewConfig() {
	// Loads environment variables e.g. PORT, BEARER_TOKEN
	// BEARER_TOKEN from the Twitter API is required
	// When .env does not exist, will exit the program
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	Conf = Config{
		MongoURL:    os.Getenv("MONGO_URL"),
		BearerToken: os.Getenv("BEARER_TOKEN"),
		APIURL:      "https://api.twitter.com/2/tweets/search",
		Options:     "&expansions=author_id,referenced_tweets.id,attachments.media_keys,referenced_tweets.id.author_id&tweet.fields=attachments,in_reply_to_user_id,public_metrics,possibly_sensitive,referenced_tweets,created_at&user.fields=profile_image_url&media.fields=height,preview_image_url,public_metrics,width,url",
		JWTSecret:   os.Getenv("JWT_SECRET"),
	}
	log.Println("Config loaded!")
}
