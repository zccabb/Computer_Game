package amazon

func (board *Board) IsGameOver() bool {
	blackOver, blackPositions := true, board.getChessAround(Black)
	for _, position := range blackPositions {
		if board[position.X][position.Y] == Empty {
			blackOver = false
			break
		}
	}
	whiteOver, whitePositions := true, board.getChessAround(White)
	for _, position := range whitePositions {
		if board[position.X][position.Y] == Empty {
			whiteOver = false
			break
		}
	}
	return blackOver || whiteOver
}

func (board *Board) freedom(color int) int {
	freedom := 0
	allChess := board.getAllChess(color)
	for _, chess := range allChess {
		for k := 0; k < 8; k++ {
			x, y := chess.X+dir[k][0], chess.Y+dir[k][1]
			for legal(x, y) && board[x][y] == Empty {
				freedom++
				x += dir[k][0]
				y += dir[k][1]
			}
		}
	}
	return freedom
}

func (board *Board) score(color int) int {
	score, positions := 0, board.getChessAround(color)
	for _, position := range positions {
		if !legal(position.X, position.Y) {
			score += 3
		} else if board[position.X][position.Y] == Arrow {
			score += 4
		} else if board[position.X][position.Y] == color {
			score += 5
		} else if board[position.X][position.Y] != Empty {
			score += 3
		}
	}
	return score
}

func (board *Board) getAllChess(color int) []Position {
	var positions []Position
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if board[i][j] == color {
				positions = append(positions, Position{i, j})
			}
		}
	}
	return positions
}

func (board *Board) getChessAround(color int) []Position {
	var positions []Position
	chess := board.getAllChess(color)
	for _, chess := range chess {
		for k := 0; k < 8; k++ {
			x, y := chess.X+dir[k][0], chess.Y+dir[k][1]
			if legal(x, y) {
				positions = append(positions, Position{x, y})
			}
		}
	}
	return positions
}

func (board *Board) looseness(color int) int {
	l := 0
	chess := board.getAllChess(color)
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			l += abs(chess[i].X-chess[j].X) + abs(chess[i].Y-chess[j].Y)
		}
	}
	return l
}

func (board *Board) Value() int {
	return (board.freedom(Black)-board.freedom(White))*10 - board.score(Black) + board.score(White)
}

func (board *Board) ValueBefore50() int {
	return (board.freedom(Black)-board.freedom(White))*10 - board.score(Black) + board.score(White) + board.looseness(Black) - board.looseness(White)
}
