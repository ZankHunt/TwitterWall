package twitter

import (
	"time"
)

// DatabaseEntry is the type for the MongoDB
type DatabaseEntry struct {
	Expiration       time.Time          `json:"expiration,omitempty" bson:"expiration,omitempty"`
	Query            string             `json:"query,omitempty" bson:"query,omitempty"`
	FrontEndResponse []FrontEndResponse `json:"frontEndResponses,omitempty" bson:"frontEndResponses,omitempty"`
}

// FrontEndResponse has the types for the front end
type FrontEndResponse struct {
	AuthorID   string     `json:"authorID"`
	Username   string     `json:"username"`
	Name       string     `json:"name"`
	Date       time.Time  `json:"date"`
	Text       string     `json:"text"`
	ProfilePic string     `json:"profilePicURL"`
	Media      []Media    `json:"medias,omitempty"`
	RefTweet   []RefTweet `json:"refTweet,omitempty"`
	Likes      int        `json:"likes"`
	Retweets   int        `json:"retweets"`
	Quotes     int        `json:"quotes"`
}

// RefTweet types for the front end
type RefTweet struct {
	AuthorID   string    `json:"authorID,omitempty"`
	Username   string    `json:"username,omitempty"`
	Name       string    `json:"name,omitempty"`
	Date       time.Time `json:"date,omitempty"`
	Text       string    `json:"text,omitempty"`
	ProfilePic string    `json:"profilePicURL,omitempty"`
}

// Media types for the front end
type Media struct {
	Width           int    `json:"width,omitempty"`
	PreviewImageURL string `json:"previewImageURL,omitempty"`
	Type            string `json:"type,omitempty"`
	Height          int    `json:"height,omitempty"`
	URL             string `json:"url,omitempty"`
}

// TweetResponseSearch types for Twitter API
type TweetResponseSearch struct {
	Data     []Data   `json:"data"`
	Includes Includes `json:"includes"`
	Errors   []Error  `json:"errors,omitempty"`
	Meta     Meta     `json:"meta,omitempty"`
}

// TweetResponseStream type for Twitter API
type TweetResponseStream struct {
	Data     Data     `json:"data"`
	Includes Includes `json:"includes"`
	Errors   []Error  `json:"errors,omitempty"`
	Rules    []Rule   `json:"matching_rules,omitempty"`
}

// TweetResponse types from the Twitter API for backend
type TweetResponse struct {
	Data     []Data   `json:"data"`
	Includes Includes `json:"includes"`
}

// Data types for Twitter API
type Data struct {
	ID                string            `json:"id"`
	CreatedAt         time.Time         `json:"created_at"`
	PublicMetrics     PublicMetrics     `json:"public_metrics"`
	AuthorID          string            `json:"author_id"`
	InReplyToUserID   string            `json:"in_reply_to_user_id,omitempty"`
	PossiblySensitive bool              `json:"possibly_sensitive"`
	Text              string            `json:"text"`
	ReferencedTweets  []ReferencedTweet `json:"referenced_tweets,omitempty"`
	Attachments       Attachments       `json:"attachments,omitempty"`
}

// PublicMetrics types for the Twitter API
type PublicMetrics struct {
	RetweetCount int `json:"retweet_count"`
	ReplyCount   int `json:"reply_count"`
	LikeCount    int `json:"like_count"`
	QuoteCount   int `json:"quote_count"`
}

// ReferencedTweet types for Twitter API
type ReferencedTweet struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// Attachments types for Twitter API
type Attachments struct {
	MediaKeys []string `json:"media_keys"`
}

// Includes types for the Twitter API
type Includes struct {
	Users  []User       `json:"users"`
	Tweets []Tweet      `json:"tweets"`
	Media  []TweetMedia `json:"media"`
}

// User types for the Twitter API
type User struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	ProfileImageURL string `json:"profile_image_url"`
	ID              string `json:"id"`
}

// Tweet types for the Twitter API
type Tweet struct {
	ID                string            `json:"id"`
	CreatedAt         time.Time         `json:"created_at"`
	PublicMetrics     PublicMetrics     `json:"public_metrics"`
	AuthorID          string            `json:"author_id"`
	InReplyToUserID   string            `json:"in_reply_to_user_id,omitempty"`
	PossiblySensitive bool              `json:"possibly_sensitive"`
	Text              string            `json:"text"`
	ReferencedTweets  []ReferencedTweet `json:"referenced_tweets,omitempty"`
	Attachments       Attachments       `json:"attachments,omitempty"`
}

// TweetMedia types for the Twitter API
type TweetMedia struct {
	Width           int    `json:"width,omitempty"`
	PreviewImageURL string `json:"preview_image_url,omitempty"`
	Type            string `json:"type,omitempty"`
	Height          int    `json:"height,omitempty"`
	URL             string `json:"url,omitempty"`
	MediaKey        string `json:"media_key,omitempty"`
}

// Error types for the Twitter API
type Error struct {
	Value        string `json:"value"`
	Parameter    string `json:"parameter"`
	ResourceType string `json:"resource_type"`
	Section      string `json:"section"`
	Title        string `json:"title"`
	Detail       string `json:"detail"`
	Type         string `json:"type"`
}

// Meta types for the Twitter API
type Meta struct {
	NewestID    string `json:"newest_id"`
	OldestID    string `json:"oldest_id"`
	ResultCount int    `json:"result_count"`
	NextToken   string `json:"next_token"`
}

// Rule types for the Twitter API
type Rule struct {
	ID  int    `json:"id"`
	Tag string `json:"tag"`
}
