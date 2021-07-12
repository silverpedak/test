package piscine

import "github.com/01-edu/z01"

/*
func main() {
	x := 5
	y := 1
	QuadA(x, y)
	QuadB(x, y)
	QuadC(x, y)
	QuadD(x, y)
	QuadE(x, y)
}
*/

func QuadA(x, y int) {
	quadrangle(x, y, '-', '|', 'o', 'o', 'o', 'o')
}

func QuadB(x, y int) {
	quadrangle(x, y, '*', '*', '/', '\\', '/', '\\')
}

func QuadC(x, y int) {
	quadrangle(x, y, 'B', 'B', 'A', 'A', 'C', 'C')
}

func QuadD(x, y int) {
	quadrangle(x, y, 'B', 'B', 'A', 'C', 'C', 'A')
}

func QuadE(x, y int) {
	quadrangle(x, y, 'B', 'B', 'A', 'C', 'A', 'C')
}

func quadrangle(X, Y int, horizontal, vertical, TL, TR, BR, BL rune) {
	if !(X <= 0 || Y <= 0) { // Check if either X or Y are 0, end program if they are
		// A slice within a slice, containing runes
		rect := make([][]rune, Y) // Is of Y height
		for y := range rect {
			rect[y] = make([]rune, X) // And X width
		}

		addSpaces(rect)
		addEdges(rect, horizontal, vertical)
		addCorners(rect, X, Y, TL, TR, BR, BL)

		// Print everything out
		for y := range rect {
			for x := range rect[y] {
				z01.PrintRune(rect[y][x])
			}
			z01.PrintRune('\n')
		}
	}
}
