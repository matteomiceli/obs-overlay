package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

// columns
const PLAYER_A = 0
const PLAYER_B = 1
const GAME = 2
const OUTCOME = 3

// team specific columns
const TEAM_A = 0
const TEAM_B = 1
const PLAYER = 6
const PLAYER_TEAM = 7

func getPlayer(name string, ranks []Rank) (*Rank, error) {
	for i := range ranks {
		rank := &ranks[i]
		if rank.Player == strings.TrimSpace(strings.ToLower(name)) {
			return rank, nil
		}
	}

	return &Rank{}, fmt.Errorf("No player called %s,", name)
}

func buildPlayerList(ranks *[]Rank) {
	players := getPlayers()
	slices.Sort(players)
	for _, player := range players {
		player := strings.TrimSpace(player)
		*ranks = append(*ranks, Rank{Player: player, Score: 0})
	}
}

func parseOnesData(ranks *[]Rank) *[]Rank {
	buildPlayerList(ranks)

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

		player.Score += getOnesWinIncrement()
	}

	return ranks
}

func teamLookup(teamName string, data [][]string) []string {
	playersOnTeam := []string{}
	for i, row := range data {
		if i == 0 {
			continue
		}
		playerName := row[PLAYER]
		playerTeam := row[PLAYER_TEAM]

		if playerTeam == teamName {
			playersOnTeam = append(playersOnTeam, playerName)
		}
	}

	return playersOnTeam
}

func parseTeamsData(ranks *[]Rank) *[]Rank {
	for i, row := range teamsData {
		// Skip first row
		if i == 0 {
			continue
		}
		winningTeam := strings.TrimSpace(row[OUTCOME])
		teamA := strings.TrimSpace(row[TEAM_A])
		teamB := strings.TrimSpace(row[TEAM_B])
		if winningTeam == "" {
			continue
		}
		// Team is not in the list of teams
		if winningTeam != teamA && winningTeam != teamB {
			continue
		}

		// Get members of team
		winners := teamLookup(winningTeam, teamsData)
		for _, winnerName := range winners {
			player, err := getPlayer(winnerName, *ranks)
			if err != nil {
				// No player found, create one
				newPlayer := Rank{Player: strings.ToLower(winningTeam), Score: 1}
				*ranks = append(*ranks, newPlayer)
			}

			player.Score += getTeamsWinIncrement()
		}
	}

	return ranks
}

func sortRanks(ranks []Rank) {
	sort.Slice(ranks, func(i, j int) bool {
		return ranks[i].Score > ranks[j].Score
	})
}
