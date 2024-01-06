package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"os"

	"github.com/mattn/go-mastodon"
	"github.com/samber/lo"
)

type YouTubeVideo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func main() {
	videos := make([]YouTubeVideo, 0, 200)
	f := lo.Must(os.Open("ost.json"))
	defer f.Close()
	lo.Must0(json.NewDecoder(f).Decode(&videos))

	mstdn := mastodon.NewClient(&mastodon.Config{
		Server:       "https://oyasumi.space",
		ClientID:     loadEnv("MASTODON_OST_CLIENT_ID"),
		ClientSecret: loadEnv("MASTODON_OST_CLIENT_SECRET"),
		AccessToken:  loadEnv("MASTODON_OST_ACCESS_TOKEN"),
	})

	video := videos[rand.Intn(len(videos))]
	lo.Must(mstdn.PostStatus(
		context.Background(),
		&mastodon.Toot{
			Status: video.Title + "\n" + "https://youtu.be/" + video.ID,
		},
	))
}

func loadEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(key + " is empty")
	}
	return v
}
