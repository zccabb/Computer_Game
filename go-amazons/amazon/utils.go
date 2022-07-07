package amazon

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func legal(x, y int) bool {
	return x >= 0 && y >= 0 && x < 10 && y < 10
}

func Rev(color int) int {
	if color == Black {
		return White
	}
	return Black
}
