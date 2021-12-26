package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed input-sample.txt
var sample string

type Game struct {
	p1_score   int
	p1_pos     int
	p2_score   int
	p2_pos     int
	p1_turn    bool
	dimensions int
}

type Wins struct {
	p1_wins int
	p2_wins int
}

type Player struct {
	score int
	pos   int
}

func getInput(input string) ([]Player, Game, Wins) {
	var players []Player
	player_input := strings.Split(input, "\n")
	for _, player := range player_input {
		start_pos_value := strings.Split(player, ": ")
		start_pos, _ := strconv.Atoi(start_pos_value[1])
		players = append(players, Player{score: 0, pos: start_pos})
	}
	game := Game{p1_score: 0, p1_pos: players[0].pos,
		p2_score: 0, p2_pos: players[1].pos,
		p1_turn: true, dimensions: 1}
	wins := Wins{p1_wins: 0, p2_wins: 0}
	return players, game, wins
}

func rollDeterministicDice(value *int) {
	if *value >= 100 {
		*value = 1
	} else {
		*value++
	}
}

func calcScore(total_roll int) int {
	value := total_roll % 10
	if value == 0 {
		value = 10
	}
	return value
}

func DeterministicDice(input string) int {
	rolls := 0
	dice_value := 0
	final_score := 1000
	players, _, _ := getInput(input)
	win := false
	for !win {
		for i, _ := range players {
			total_roll := 0
			for turn := 1; turn <= 3; turn++ {
				rollDeterministicDice(&dice_value)
				rolls++
				total_roll += dice_value
			}
			players[i].pos = calcScore(total_roll + players[i].pos)
			players[i].score += players[i].pos
			if players[i].score >= final_score {
				if i == 0 {
					return rolls * players[1].score
				} else {
					return rolls * players[0].score
				}
			}
		}
	}
	return rolls
}

func calcDimension(roll int) int {
	switch roll {
	case 9: // 333x1
		return 1
	case 8: // 332x3
		return 3
	case 7: // 322x3 331x3
		return 6
	case 6: // 123x6 222x1
		return 7
	case 5: // 122x3 113x3
		return 6
	case 4: // 112x3
		return 3
	case 3: // 111x1
		return 1
	default:
		fmt.Println("error")
		return 0
	}
}

func calcTurn(pos, score *int, roll int) bool {
	*pos = calcScore(*pos + roll)
	*score += *pos
	if *score >= 21 {
		return true
	}
	return false
}

func calcGame(game Game, roll int) (bool, Game) {
	new_game := game
	win := false
	if game.p1_turn {
		win = calcTurn(&new_game.p1_pos, &new_game.p1_score, roll)

	} else {
		win = calcTurn(&new_game.p2_pos, &new_game.p2_score, roll)
	}
	return win, new_game
}

func rollDiracDice(game Game, wins *Wins) {
	// Check for each roll
	for roll := 3; roll <= 9; roll++ {
		dimensions := calcDimension(roll)
		if win, new_game := calcGame(game, roll); win {
			if game.p1_turn {
				(*wins).p1_wins += (dimensions * game.dimensions)
			} else {
				(*wins).p2_wins += (dimensions * game.dimensions)
			}
		} else {
			// swap turns
			new_game.p1_turn = !new_game.p1_turn
			new_game.dimensions *= dimensions
			rollDiracDice(new_game, wins)
		}
	}
}

func DiracDice(input string) int {
	_, game, wins := getInput(input)
	rollDiracDice(game, &wins)
	if wins.p1_wins > wins.p2_wins {
		return wins.p1_wins
	}
	return wins.p2_wins
}

func main() {
	fmt.Println(DeterministicDice(input))
	fmt.Println(DiracDice(input))
}
