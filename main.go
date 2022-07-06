package main

import (
	"github.com/daniarlert/reddlet/internal/pkg/reddit"
	"github.com/spf13/viper"
	goreddit "github.com/vartanbeno/go-reddit/v2/reddit"
	"log"
)

type Config struct {
	RedditID       string `mapstructure:"REDDIT_ID"`
	RedditSecret   string `mapstructure:"REDDIT_SECRET"`
	RedditUsername string `mapstructure:"REDDIT_USERNAME"`
	RedditPassword string `mapstructure:"REDDIT_PASSWORD"`
}

func main() {
	var config Config

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error while reading config file: %v\n", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("error while unmarshaling config: %v\n", err)
	}

	credentials := goreddit.Credentials{
		ID:       config.RedditID,
		Secret:   config.RedditSecret,
		Username: config.RedditUsername,
		Password: config.RedditPassword,
	}

	client, err := reddit.NewClient(credentials)
	if err != nil {
		log.Fatalf("error creating a new reddit client: %v\n", err)
	}

	reddit.GetSubredditTopPosts(client, "golang", 5)

	// cmd.Execute()
}
