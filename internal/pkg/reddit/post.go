package reddit

type Post struct {
	ID        string     `json:"id"`
	Subreddit string     `json:"subreddit"`
	Title     string     `json:"title"`
	Url       string     `json:"url"`
	NSFW      bool       `json:"nsfw"`
	Body      string     `json:"body"`
	Comments  []*Comment `json:"comments"`
}

type Comment struct {
	ID     string `json:"id"`
	PostID string `json:"postID"`
	Body   string `json:"body"`
}
