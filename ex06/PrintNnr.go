package main
import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0{
		n *= -1
		z01.PrintRune('-') 
	}
	if n > 9 {
		PrintNbr(n / 10)
		PrintNbr(n % 10)
	} else {
		z01.PrintRune(rune(n + 48))
	}
}
