package main

import (
	"github.com/jbrukh/bayesian"
	"strings"
)

//constants for classifer
const (
	positive bayesian.Class = "positive"
	negative bayesian.Class = "negative"
)

//JSON structure for holding data for response
type Analysis struct {
	Filter string
	Sentiment string
	FrequencyWord map[string]int
}

//main function to provide twitter analysis
func analyze(s string) Analysis{

	//hold response
	var a Analysis
	a.Filter = s

	//get classifer
	classifier := getTrainedClassifer()

	//hold word sentiment numbers for calculating averages
	var pos, neg, neu int

	//hold frequency data
	m := make(map[string]int)

	//get tweets and returns them to a record structure
	t,_ := getTweets(a.Filter)
	req := tweetsToRecord(t, a.Filter)

	//loop through records
	for _, v := range req {

		//pre-process tweet data
		//add tweet words to frequency distibution
		v.preProcess()
		s := strings.Fields(v.Text)
		m = addFreqDist(s, m)

		//get sentiment for sentence and add to counters
		str, _, _ := classifier.ProbScores(s)
		if str[0] > 90 {
			pos++
		} else if str[1] > 90 {
			neg++
		} else {
			neu++
		}
	}

	//get averages of each sentiment
	avgPos := pos / len(req)
	avgNeg := neg / len(req)
	avgNeu := neu / len(req)

	//determine overall sentiment
	if avgPos > avgNeg {
		if avgPos > avgNeg {
			a.Sentiment = "Positive"
		}
	} else if avgNeg > avgPos {
		if avgNeg > avgNeu {
			a.Sentiment = "Negative"
		}
	} else {
		a.Sentiment = "Neutral"
	}

	//add frequency to response
	a.FrequencyWord = m

	return a
}

//trains classifer
func trainClassifier() *bayesian.Classifier{
	classifier := bayesian.NewClassifier(positive, negative)
	pos, neg := getSortedSet()
	classifier.Learn(pos, positive)
	classifier.Learn(neg, negative)
	return classifier
}

//returns classifer
func getTrainedClassifer() *bayesian.Classifier {
	var hasBeenTrained bool = false
	var classifer *bayesian.Classifier

	if hasBeenTrained == false {
		classifer = trainClassifier()
	}

	return classifer
}

//adds frequency of words to map
func addFreqDist (s []string, m map[string]int) map[string]int {

	fd := make(map[string]int)
	fd = m

	for _, v := range s {
		_,ok := m[v]
		if ok {
			fd[v]++
		} else {
			fd[v] = 1
		}
	}

	return fd
}