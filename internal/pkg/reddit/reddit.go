package reddit

import (
	"context"
	"fmt"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"log"
)

func NewClient(credentials reddit.Credentials) (*reddit.Client, error) {
	client, err := reddit.NewClient(credentials)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetSubredditTopPosts(client *reddit.Client, subreddit string, limit int) {
	options := &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: limit,
		},
		Time: "all",
	}

	posts, _, err := client.Subreddit.TopPosts(context.Background(), subreddit, options)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Received %d posts.\n", len(posts))
}
