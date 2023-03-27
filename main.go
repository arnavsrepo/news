package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/barthr/newsapi"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	apiKey := viper.GetString("API_KEY")

	httpClient := &http.Client{Timeout: 30 * time.Second}
	client := newsapi.NewClient(apiKey, newsapi.WithHTTPClient(httpClient), newsapi.WithUserAgent("Arnav"))

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Minute*10))
	articles, err := client.GetTopHeadlines(ctx, &newsapi.TopHeadlineParameters{
		Country: "in",
	})
	if err != nil {
		panic(err)
	}

	for _, s := range articles.Articles {
		fmt.Printf("%+v\n\n", s)
	}
}
