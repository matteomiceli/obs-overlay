package main

import (
	"encoding/csv"
	"log"
	"net/http"
)

func getSheetData(url string) [][]string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	csvReader := csv.NewReader(res.Body)
	content, err := csvReader.ReadAll()
	return content
}
