package reddit

import (
	"context"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func NewClient(credentials reddit.Credentials) (*reddit.Client, error) {
	client, err := reddit.NewClient(credentials)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetSubredditTopThreads(client *reddit.Client, subreddit string, limit int) ([]*reddit.Post, error) {
	options := &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: limit,
		},
		Time: "all",
	}

	posts, _, err := client.Subreddit.TopPosts(context.Background(), subreddit, options)
	return posts, err
}

func GetPostWithComments(client *reddit.Client, id string, limit int) (*Post, error) {
	thread, _, err := client.Post.Get(context.Background(), id)
	if err != nil {
		return nil, err
	}

	post := &Post{
		ID:        thread.Post.ID,
		Subreddit: thread.Post.SubredditName,
		Title:     thread.Post.Title,
		Url:       thread.Post.URL,
		NSFW:      thread.Post.NSFW,
		Body:      thread.Post.Body,
	}

	var comments []*Comment
	for i := 0; i < len(thread.Comments) || i < limit; i++ {
		c := &Comment{
			ID:     thread.Comments[i].ID,
			PostID: thread.Post.ID,
			Body:   thread.Comments[i].Body,
		}

		comments = append(comments, c)
	}

	post.Comments = comments
	return post, nil
}
