package main

import (
	"fmt"
	"strings"
)

const PLAYER_A = 0
const PLAYER_B = 1
const GAME = 2
const OUTCOME = 3

func getPlayer(name string, ranks []Rank) (*Rank, error) {
	for i := range ranks {
		rank := &ranks[i]
		if rank.Player == strings.TrimSpace(strings.ToLower(name)) {
			return rank, nil
		}
	}

	return &Rank{}, fmt.Errorf("No player called %s,", name)
}

func buildPlayerList(data [][]string, ranks *[]Rank) {
	players := []string{}
	for i, row := range data {
		if i == 0 {
			continue
		}

		player := strings.TrimSpace(row[PLAYER_A])

		newPlayer := true
		for _, p := range players {
			if player == p {
				newPlayer = false
			}
		}
		if newPlayer {
			players = append(players, player)
			*ranks = append(*ranks, Rank{Player: player, Score: 0})
		}
	}
}

func parseOnesData(ranks *[]Rank) *[]Rank {
	onesData := getSheetData("https://docs.google.com/spreadsheets/d/1c2IkaK9iFRRfE5hy8eHzn-YrDt9LSMUN32Jv51Mbt7k/export?format=csv&gid=0")
	buildPlayerList(onesData, ranks)

	for i, row := range onesData {
		// Skip first row
		if i == 0 {
			continue
		}
		winner := strings.TrimSpace(row[OUTCOME])
		playerA := strings.TrimSpace(row[PLAYER_A])
		playerB := strings.TrimSpace(row[PLAYER_B])
		if winner == "" {
			continue
		}
		// Player is not in the list of players
		if winner != playerA && winner != playerB {
			continue
		}
		player, err := getPlayer(winner, *ranks)
		if err != nil {
			// No player found, create one
			newPlayer := Rank{Player: strings.ToLower(winner), Score: 1}
			*ranks = append(*ranks, newPlayer)
		}

		player.Score += 1
		fmt.Println(player)
	}

	fmt.Println(ranks)
	return ranks
}
