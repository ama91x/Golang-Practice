// ----------------------------------------- Qustion -----------------------------------------
// Write a program that prints the Latin alphabet in lowercase on a single line.
// A line is a sequence of characters preceding the end of line character ('\n').
// Allowed: github.com/01-edu/z01.PrintRune#2, --allow-builtin
// -------------------------------------------------------------------------------------------

package main

import "github.com/01-edu/z01"

func main() {
	for i := 'a'; i <= 'z'; i++ {
		z01.PrintRune(i)
	}
}
