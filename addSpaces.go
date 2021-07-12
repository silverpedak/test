package piscine

func addSpaces(rect [][]rune) {

	for y := range rect { // Määrab igasse võimalikku x ja y kombinatsiooni tühikud.

		for x := range rect[y] {

			rect[y][x] = ' '

		}
	}
}
