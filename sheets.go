package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"time"
)

func getSheetData(url string) [][]string {
	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	csvReader := csv.NewReader(res.Body)
	content, err := csvReader.ReadAll()

	fmt.Printf("Fetching sheet %s\nelapsed: %s\n", url, time.Since(start))
	return content
}
