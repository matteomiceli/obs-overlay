package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var onesData [][]string
var teamsData [][]string
var varData [][]string

//go:embed templates/index.html
var indexHtml string

func main() {
	serve()
}

type Rank struct {
	Player string
	Score  int
}
type PageData struct {
	Ranks   []Rank
	Refresh template.HTMLAttr
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

	t := template.New("overlay")
	templ, err := t.Parse(indexHtml)
	if err != nil {
		log.Fatal(err)
	}
	err = templ.Execute(w, PageData{
		Ranks:   getPlayerRanks(),
		Refresh: template.HTMLAttr(fmt.Sprintf(`content="%s"`, getRefreshTime())),
	})
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
