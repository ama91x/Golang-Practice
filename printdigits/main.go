// ----------------------------------------- Qustion -----------------------------------------
// Write a program that prints the decimal digits in ascending order (from 0 to 9) on a single line.
// A line is a sequence of characters preceding the end of line character ('\n').
// Allowed: github.com/01-edu/z01.PrintRune#2, --allow-builtin
// -------------------------------------------------------------------------------------------
package main

import "github.com/01-edu/z01"

func main() {
	for i := 0; i < 10; i++ {
		z01.PrintRune(rune(i + '0'))
	}
}
