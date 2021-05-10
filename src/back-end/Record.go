package main

import (
	"bytes"
	"encoding/csv"
	"github.com/bbalet/stopwords"
	"log"
	"os"
	"regexp"
	"strings"
)

//to hold a csv record
type Record struct {
	Filter    string
	Sentiment string
	id        int64
	Text      string
}

//pre processes the record
func(r *Record) preProcess() {
	var text string
	text = r.Text
	text = strings.ToLower(text)

	//remove usernames
	reg, _ := regexp.Compile("@[^\\s]+")
	text = reg.ReplaceAllString(text, "")

	//remove hashtags
	reg, _ = regexp.Compile("#([^\\s]+)")
	text = reg.ReplaceAllString(text, "$1")

	//remove URLS
	reg, _ = regexp.Compile("((www\\.[^\\s]+)|(https?://[^\\s]+))")
	text = reg.ReplaceAllString(text, "")

	//remove duplicate letters
	text = removeDups(text)

	//remove stop words
	text = stopwords.CleanString(text, "en", true)

	r.Text = text
}

//tokenizes a record with preprocessing
func (r *Record) tokenizeText() []string{
	r.preProcess()
	return strings.Fields(r.Text)
}

//prints all records to csv
func printToCSV(records []Record, path string) {

	var file *os.File

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		file, err = os.Create(path)
		if err != nil {
			log.Println(err)
		}
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var csvLines [][]string

	for _, v := range records {
		var csvLine []string
		csvLine = append(csvLine, v.Filter, string(v.id), v.Sentiment, v.Text)
		csvLines = append(csvLines, csvLine)
	}


	writeErr := writer.WriteAll(csvLines)
	if writeErr != nil {
		log.Println("cannot write", writeErr)
	}
}

//prints the record to the testSet csv
func (r *Record) printLineToCSV(path string) {

	var file *os.File

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		file, err = os.Create(path)
		if err != nil {
			log.Println(err)
		}
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var csvLine []string
	csvLine = append(csvLine, r.Filter, string(r.id), r.Sentiment, r.Text)

	writeErr := writer.Write(csvLine)
	if writeErr != nil {
		log.Println("cannot write", writeErr)
	} else {
		log.Println("line written")
	}
}

//removes duplicate letters
func removeDups(s string) string {
	var buf bytes.Buffer
	var last rune
	var flag bool = false

	for i, r := range s {

		if r == last && flag == false {
			buf.WriteRune(r)
			flag = true
		}

		if r != last || i == 0 {
			buf.WriteRune(r)
			last = r
			flag = false
		}

	}
	return buf.String()
}
