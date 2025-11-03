package main

import (
	"log"
	"strconv"
	"strings"
)

// data
const PLAYERS = 0
const ONES_WIN_SCORE_INC = 1
const TEAMS_WIN_SCORE_INC = 2

func _getVariable(varName string) string {
	switch varName {
	case "players":
		for i, row := range varData {
			if i == 0 {
				continue
			}
			return row[PLAYERS]
		}
	case "onesWinInc":
		for i, row := range varData {
			if i == 0 {
				continue
			}
			return row[ONES_WIN_SCORE_INC]
		}
	case "teamsWinInc":
		for i, row := range varData {
			if i == 0 {
				continue
			}
			return row[TEAMS_WIN_SCORE_INC]
		}
	}
	return ""
}

func getPlayers() []string {
	players := _getVariable("players")
	return strings.Split(players, ",")
}

func getOnesWinIncrement() int {
	winIncStr := _getVariable("onesWinInc")
	winInc, err := strconv.Atoi(winIncStr)
	if err != nil {
		log.Fatal(err)
	}
	return winInc
}

func getTeamsWinIncrement() int {
	winIncStr := _getVariable("teamsWinInc")
	winInc, err := strconv.Atoi(winIncStr)
	if err != nil {
		log.Fatal(err)
	}
	return winInc
}
