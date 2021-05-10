package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

//returns the client for twitter api
//uses twitter package
//uses oauth01 package
func getClient() *twitter.Client {

	//My twitter api tokens
	config := oauth1.NewConfig("","")
	token := oauth1.NewToken("","")
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)
	//boolean if client is set
	var isSet bool
	//client variable to be returned
	var client *twitter.Client

	// Twitter client
	if !isSet {
		client := twitter.NewClient(httpClient)
		isSet = true
		return client
	}

	return client
}
