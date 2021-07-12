package piscine

func addEdges(rect [][]rune, hor, ver rune) {
	Ylast := len(rect) - 1    // Index of last row
	Xlast := len(rect[0]) - 1 // Index of last column

	for x := range rect[0] {
		// Runes hor into top and bottom rows.
		rect[0][x], rect[Ylast][x] = hor, hor
	}

	for y := range rect {
		// Runes ver into left and right most columns.
		rect[y][0], rect[y][Xlast] = ver, ver
	}
}
