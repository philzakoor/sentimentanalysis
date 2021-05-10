package main

import (
	"encoding/csv"
	"github.com/dghubble/go-twitter/twitter"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)


//reads corpus.csv a file with over 5000 hand classified test set
// return a record for the csv which includes a filter, the hand
//classified sentiment, and the tweet id
func readCorpus() []Record{

	//corpus path
	path := "corpus.csv"

	//hold record variable
	var records []Record

	//open csv
	csvfile, err := os.Open(path)
	if err != nil {
		log.Fatalln("Couldn't open CSV file", err)
	}

	//read csv
	r := csv.NewReader(csvfile)

	//loop through file
	for {

		//read csv line and put into record
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		n, err := strconv.ParseInt(rec[2], 10, 64)

		if err == nil {
			records = append(records, Record{rec[0], rec[1], n, ""})
		}
	}

	//return all csv records
	return records
}

//uses twitter api to get the tweet
//using the tweetId
//return nil if the api returns an error
func downloadTestSet(id int64) *twitter.Tweet{

	//get client
	client := getClient()


	//search twitter using filter
	tweet, resp, err := client.Statuses.Show(id, &twitter.StatusShowParams{
		ID: id,
	})

	//error handling
	if err != nil {
		log.Print(resp, err)
		return nil
	}

	return tweet
}


//downloads the tweet from twitter api
//and prints the tweet to a csv file
func (r *Record) downloadAndPrint() bool{
	tweet := downloadTestSet(r.id)
	if tweet != nil {
		r.Text = tweet.Text
		r.printLineToCSV("trainingSet.csv")
		return true
	}
	return false
}

//function to read the corpus data
// iterates through all the corpus records
// downloads and prints each record to a csv
// keeps track of how many records it's processed
//sleeps for 15mins after 180 tweets have been process due to api limits
func createTestSet() {

	records := readCorpus()
	//var cvsRecords []Record
	ctr := 0

	for _, v := range records {
		if v.downloadAndPrint() {
			ctr++
		}
		if ctr >= 180 {
			time.Sleep(15 * time.Minute)
			ctr = 0
		}
	}
}

//reads training set made from corpus with all twitter data
func getTrainingRecords() []Record {

	path := "trainingSet.csv"

	var records []Record

	csvfile, err := os.Open(path)
	if err != nil {
		log.Fatalln("Couldn't open CSV file", err)
	}

	r := csv.NewReader(csvfile)

	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		records = append(records, Record{rec[0], rec[2], 0, rec[3]})
	}

	return records
}

//sorts trained records into positive and negative categories
func getSortedSet() ([]string, []string){
	var positive []string
	var negative []string

	r := getTrainingRecords()

	for _, v := range r {
		v.preProcess()
		tempStrings := strings.Fields(v.Text)
		if v.Sentiment == "negative" {
			for _,v2 := range tempStrings {
				negative = append(negative, v2)
			}
		} else {
			for _,v2 := range tempStrings {
				positive = append(positive, v2)
			}
		}
	}

	return positive, negative
}
