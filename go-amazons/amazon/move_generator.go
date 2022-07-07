package amazon

func (board Board) Generator(color int) []Move {
	var moves []Move
	allChess := board.getAllChess(color)
	for _, chess := range allChess {
		for j := 0; j < 8; j++ {
			x, y := chess.X+dir[j][0], chess.Y+dir[j][1]
			for legal(x, y) && board[x][y] == Empty {
				for k := 0; k < 8; k++ {
					ax, ay := x+dir[k][0], y+dir[k][1]
					for legal(ax, ay) && (board[ax][ay] == Empty || ax == chess.X && ay == chess.Y) {
						moves = append(moves, Move{
							From: Position{chess.X, chess.Y},
							To:   Position{x, y},
							Put:  Position{ax, ay},
						})
						ax += dir[k][0]
						ay += dir[k][1]
					}
				}
				x += dir[j][0]
				y += dir[j][1]
			}
		}
	}
	return moves
}
