package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr_sudoku := os.Args[1:]
	if len(arr_sudoku) == 1 {
		arr_sudoku = checkOne(arr_sudoku)
	}
	if !Error_len(arr_sudoku) {
		Error_console()
	} else if !Error_312(rune_sudoku(arr_sudoku)) {
		Error_console()
	} else {
		sud := rune_sudoku(arr_sudoku)
		if sudoku(&sud, 0, 0) {
			vvod_na_console(sud)
		} else {
			Error_console()
		}
	}
}

func Error_console() {
	str := "Error\n"
	for _, w := range str {
		z01.PrintRune(w)
	}
}

func Error_312(sud [9][9]rune) bool {
	for x := 0; x < 9; x++ {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if sud[i][x] == sud[j][x] && i != j && sud[i][x] != '.' {
					return false
				}
				if sud[x][i] == sud[x][j] && i != j && sud[x][i] != '.' {
					return false
				}
			}
		}
	}
	return true
}

func Error_len(arr_sudoku []string) bool {
	if len(arr_sudoku) == 9 {
		for i := 0; i < 9; i++ {
			if len(arr_sudoku[i]) != 9 {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func checkOne(args []string) []string {
	var tempString string
	tempRune := []rune(args[0])
	for x, i := range tempRune {
		if i != ' ' {
			tempString += string(i)
		}
		if i == ' ' || x == len(tempRune)-1 {
			args = append(args, tempString)
			tempString = ""
		}
	}
	return args[1:]
}

func rune_sudoku(str []string) [9][9]rune {
	var rsud [9][9]rune
	for i := 0; i < 9; i++ {
		for j, w := range str[i] {
			rsud[i][j] = w
		}
	}
	return rsud
}

func vvod_na_console(rsud [9][9]rune) {
	for _, w := range rsud {
		for _, q := range w {
			z01.PrintRune(q)
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')
	}
}

func sudoku(rsud *[9][9]rune, ru, lu int) bool {
	if ru == 8 && lu == 9 {
		return true
	}
	if lu == 9 {
		ru++
		lu = 0
	}
	if (*rsud)[ru][lu] >= '1' && (*rsud)[ru][lu] <= '9' {
		return sudoku(rsud, ru, lu+1)
	}
	for num := '1'; num <= '9'; num++ {
		if test(ru, lu, rsud, num) {
			(*rsud)[ru][lu] = num
			if sudoku(rsud, ru, lu+1) {
				return true
			}
		}
		(*rsud)[ru][lu] = '.'
	}
	return false
}

func check_vertically(lu int, rsud *[9][9]rune, num rune) bool {
	for i := 0; i < 9; i++ {
		if (*rsud)[i][lu] == num {
			return false
		}
	}
	return true
}

func check_horizontal(ru int, rsud *[9][9]rune, num rune) bool {
	for i := 0; i < 9; i++ {
		if (*rsud)[ru][i] == num {
			return false
		}
	}
	return true
}

func check_square(ru, lu int, rsud *[9][9]rune, num rune) bool {
	x := ru - ru%3
	y := lu - lu%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (*rsud)[i+x][j+y] == num {
				return false
			}
		}
	}
	return true
}

func test(ru, lu int, rsud *[9][9]rune, num rune) bool {
	if check_vertically(lu, rsud, num) && check_horizontal(ru, rsud, num) && check_square(ru, lu, rsud, num) {
		return true
	}
	return false
}
