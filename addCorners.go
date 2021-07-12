package piscine

//see funktsioon kirjutab üle nurgamärgid. käivitatakse funktsioonist quadrangle.go viimasena, et nurgad oleks õiged.

func addCorners(rect [][]rune, x, y int, rune1, rune2, rune3, rune4 rune) { //võtab muutujad
	rect[y-1][x-1] = rune3 //määrab alumise parema nurga märgi
	rect[y-1][0] = rune4   //määrab alumise vasaku nurga märgi
	rect[0][x-1] = rune2   //ülemine parem
	rect[0][0] = rune1     //ülemine vasak
}

/* kontrollfunktsioon
func main() {
	x := 5
	y := 4
	rect := make([][]rune, y)
	for i := range rect {
		rect[i] = make([]rune, x)
	}
	addCorners(x, y, rect)
	for i := range rect {
		for j := range rect[i] {
			z01.PrintRune(rect[i][j])
		}
		z01.PrintRune('\n')
	}
}*/
