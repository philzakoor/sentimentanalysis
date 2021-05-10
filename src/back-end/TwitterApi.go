package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"log"
)

//Searches twitter for tweets using twitter API and returns tweets
//uses twitter package
func getTweets(filter string) ([]twitter.Tweet, string) {
	//get client
	client := getClient()

	//search twitter using filter
	tweets, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Count: 5,
		Query: filter,
		Lang: "en",
	})

	//error handling
	if err != nil {
		log.Print(resp, err)
		return nil, ""
	}

	//return tweets
	return tweets.Statuses, filter
}

//transfers tweets to records for analysis
func tweetsToRecord (t []twitter.Tweet, f string) []Record {
	var r []Record

	for _, v := range t {
		tempR := Record{Text: v.Text, id: v.ID,Filter: f}
		r = append(r, tempR)
	}

	return r
}


