package engine

import (
	"go-amazons/amazon"
)

var bestMove amazon.Move

func AlphaBetaSearch(board *amazon.Board, depth, alpha, beta, cur, color int, first, before bool) int {
	if depth == 0 || board.IsGameOver() {
		var result int
		if before {
			result = board.ValueBefore50()
		} else {
			result = board.Value()
		}
		if color == amazon.Black {
			return result
		} else {
			return -result
		}
	}
	moves := board.Generator(cur)
	var value int
	for _, move := range moves {
		board.Move(move)
		value = -AlphaBetaSearch(board, depth-1, -beta, -alpha, amazon.Rev(cur), color, false, before)
		board.ReMove(move)
		if alpha < value {
			alpha = value
			if first {
				bestMove = move
			}
		}
		if beta <= alpha {
			return alpha
		}
	}
	return alpha
}

func GetBestMove() amazon.Move {
	return bestMove
}
