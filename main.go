package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	//onesData := getSheetData("https://docs.google.com/spreadsheets/d/1c2IkaK9iFRRfE5hy8eHzn-YrDt9LSMUN32Jv51Mbt7k/export?format=csv&gid=0")
	//teamsData := getSheetData("https://docs.google.com/spreadsheets/d/1c2IkaK9iFRRfE5hy8eHzn-YrDt9LSMUN32Jv51Mbt7k/export?format=csv&gid=924150328")
	//fmt.Println(teamsData)
	serve()
}

type Rank struct {
	Player string
	Score  int
}

func renderView(w http.ResponseWriter, req *http.Request) {

	body, err := os.ReadFile("./templates/index.html")
	t := template.New("overlay")
	templ, err := t.Parse(string(body))
	if err != nil {
		log.Fatal(err)
	}
	templ.Execute(w, Rank{Player: "matteo", Score: 2})
}

func serve() {
	http.HandleFunc("/", renderView)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
