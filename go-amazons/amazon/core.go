package amazon

const (
	Empty = iota
	Black
	White
	Arrow
)

var (
	dir = [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
	}
)

type Position struct {
	X int
	Y int
}

type Move struct {
	From Position
	To   Position
	Put  Position
}

type Board [10][10]int

func NewBoard() Board {
	var board [10][10]int

	board[0][3] = White
	board[0][6] = White
	board[3][0] = White
	board[3][9] = White

	board[6][0] = Black
	board[6][9] = Black
	board[9][3] = Black
	board[9][6] = Black

	return board
}

func (board *Board) Move(move Move) {
	board[move.To.X][move.To.Y] = board[move.From.X][move.From.Y]
	board[move.From.X][move.From.Y] = Empty
	board[move.Put.X][move.Put.Y] = Arrow
}

func (board *Board) ReMove(move Move) {
	board[move.From.X][move.From.Y] = board[move.To.X][move.To.Y]
	board[move.To.X][move.To.Y] = Empty
	if move.From.X != move.Put.X || move.From.Y != move.Put.Y {
		board[move.Put.X][move.Put.Y] = Empty
	}
}
