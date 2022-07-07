package main

import (
	"bufio"
	"fmt"
	"go-amazons/amazon"
	"go-amazons/engine"
	"math/rand"
	"os"
	"strings"
	"time"
)

const INF = 0x3f3f3f3f

var (
	line  string
	board amazon.Board
	color int
	step  int
)

func main() {
	rand.Seed(time.Now().UnixNano())
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line = sc.Text()
		if line == "name?" {
			fmt.Println("name KJK")
		} else if line == "quit" {
			os.Exit(0)
		} else if strings.HasPrefix(line, "new") {
			step = 0
			words := strings.Split(line, " ")
			board = amazon.NewBoard()
			if words[1] == "black" {
				color = amazon.Black
				runSearch()
			} else {
				color = amazon.White
			}
		} else if strings.HasPrefix(line, "move") {
			words := strings.Split(line, " ")
			move := words[1]
			board[move[3]-'A'][move[2]-'A'] = board[move[1]-'A'][move[0]-'A']
			board[move[1]-'A'][move[0]-'A'] = amazon.Empty
			board[move[5]-'A'][move[4]-'A'] = amazon.Arrow
			AddRecord(int(move[0]+32), int(10-move[1]+'A'), int(move[2]+32),
				int(10-move[3]+'A'), int(move[4]+32), int(10-move[5]+'A'))
			if !board.IsGameOver() {
				runSearch()
			}
		} else if line == "end" {
			Save()
		}
	}
}

func runSearch() {
	if step < 25 {
		engine.AlphaBetaSearch(&board, 2, -INF, INF, color, color, true, true)
	} else {
		engine.AlphaBetaSearch(&board, 4, -INF, INF, color, color, true, false)
	}
	move := engine.GetBestMove()
	board.Move(move)
	fmt.Printf("move %c%c%c%c%c%c\n", move.From.Y+'A', move.From.X+'A', move.To.Y+'A', move.To.X+'A', move.Put.Y+'A', move.Put.X+'A')
	AddRecord(move.From.Y+'a', 10-move.From.X, move.To.Y+'a', 10-move.To.X, move.Put.Y+'a', 10-move.Put.X)
	step++
}
