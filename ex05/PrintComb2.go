package main
import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 98 ; i++ {
		for j := i + 1; j <= 99 ; j++{
			z01.PrintRune(rune(i / 10 + 48))
			z01.PrintRune(rune(i % 10 + 48))
			z01.PrintRune(rune(j / 10 + 48))
			z01.PrintRune(rune(j % 10 + 48))
			if (i != 98 || j != 99){
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
 func main() {
	PrintComb2()
}