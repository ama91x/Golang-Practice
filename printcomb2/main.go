package main

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for x := 0; x <= 9; x++ {
				for y := 0; y <= 9; y++ {
					if (i < x) || (i == x && j < y) {
						z01.PrintRune(rune(i + '0'))
						z01.PrintRune(rune(j + '0'))
						z01.PrintRune(' ')
						z01.PrintRune(rune(x + '0'))
						z01.PrintRune(rune(y + '0'))
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}

func main() {
	PrintComb2()
}
