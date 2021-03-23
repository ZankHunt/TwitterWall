package twitter

// GetJSONResponse creats a prepared JSON Response for the frontend
func createFrontEndJSON(input TweetResponse) []FrontEndResponse {
	// New output array
	output := []FrontEndResponse{}
	for _, val := range input.Data {
		if !val.PossiblySensitive {
			// New FrontEndResponse
			response := FrontEndResponse{}

			// Basic info
			response.AuthorID = val.AuthorID
			response.Date = val.CreatedAt
			response.Likes = val.PublicMetrics.LikeCount
			response.Retweets = val.PublicMetrics.RetweetCount
			response.Quotes = val.PublicMetrics.QuoteCount + val.PublicMetrics.ReplyCount
			response.Text = val.Text

			// Gets the media key from twitter API
			mediaKeys := val.Attachments.MediaKeys

			// When they are not null, will copy the values from twitter api
			if mediaKeys != nil && len(mediaKeys) > 0 {
				response.Media = make([]Media, len(mediaKeys))
				for _, media := range input.Includes.Media {
					for j := range mediaKeys {
						if mediaKeys[j] == media.MediaKey {
							response.Media[j].Height = media.Height
							if media.PreviewImageURL != "" {
								response.Media[j].PreviewImageURL = media.PreviewImageURL + ":small"
							}
							response.Media[j].Type = media.Type
							if media.URL != "" {
								response.Media[j].URL = media.URL + ":small"
							}
							response.Media[j].Width = media.Width
						}
					}
				}
			}

			// Gets the refered tweets from Twitter API
			refTweets := val.ReferencedTweets

			// When they are not null and no media keys, copies only the first tweet
			if refTweets != nil && len(refTweets) > 0 && len(mediaKeys) == 0 {
				for _, ref := range input.Includes.Tweets {
					if !ref.PossiblySensitive && refTweets[0].ID == ref.ID {
						response.RefTweet = make([]RefTweet, 1)
						response.RefTweet[0].AuthorID = ref.AuthorID
						response.RefTweet[0].Date = ref.CreatedAt
						response.RefTweet[0].Text = ref.Text
					}
				}
			}

			// Gets the user info for the tweet and refered tweet
			for _, user := range input.Includes.Users {
				if user.ID == response.AuthorID {
					response.Name = user.Name
					response.Username = "@" + user.Username
					response.ProfilePic = user.ProfileImageURL
				}

				if len(response.RefTweet) > 0 && user.ID == response.RefTweet[0].AuthorID {
					response.RefTweet[0].Name = user.Name
					response.RefTweet[0].Username = "@" + user.Username
					response.RefTweet[0].ProfilePic = user.ProfileImageURL
				}
			}

			// Appends FrontEndResponse to the output array
			output = append(output, response)
		}
	}
	return output
}
