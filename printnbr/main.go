// Write a function that prints an int passed in parameter. All possible values of type int have to go through.
// You cannot convert to int64.
// piscine.PrintNbr(-123)
// piscine.PrintNbr(0)
// piscine.PrintNbr(123)
// $ go run .
// -1230123
// $

package main

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	if n > 10 {
		PrintNbr(n / 10)
	}

	z01.PrintRune(rune(n%10 + '0'))
}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	PrintNbr(323)
	z01.PrintRune('\n')
}
