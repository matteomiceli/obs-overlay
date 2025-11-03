package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var onesData [][]string
var teamsData [][]string
var varData [][]string

func main() {
	serve()
}

type Rank struct {
	Player string
	Score  int
}

func getPlayerRanks() []Rank {
	ranks := []Rank{}
	parseOnesData(&ranks)
	parseTeamsData(&ranks)
	sortRanks(ranks)
	return ranks
}

func renderView(w http.ResponseWriter, req *http.Request) {
	// Fetch data
	onesData = getSheetData("https://docs.google.com/spreadsheets/d/1c2IkaK9iFRRfE5hy8eHzn-YrDt9LSMUN32Jv51Mbt7k/export?format=csv&gid=0")
	teamsData = getSheetData("https://docs.google.com/spreadsheets/d/1c2IkaK9iFRRfE5hy8eHzn-YrDt9LSMUN32Jv51Mbt7k/export?format=csv&gid=924150328")
	varData = getSheetData("https://docs.google.com/spreadsheets/d/1c2IkaK9iFRRfE5hy8eHzn-YrDt9LSMUN32Jv51Mbt7k/export?format=csv&gid=1512517550")

	body, err := os.ReadFile("./templates/index.html")
	t := template.New("overlay")
	templ, err := t.Parse(string(body))
	if err != nil {
		log.Fatal(err)
	}
	err = templ.Execute(w, getPlayerRanks())
	if err != nil {
		log.Fatal(err)
	}
}

func serve() {
	http.HandleFunc("/", renderView)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
