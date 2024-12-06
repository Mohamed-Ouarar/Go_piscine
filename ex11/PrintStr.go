package main
import(
	"github.com/01-edu/z01"

)

func PrintStr(s string) {
	for _, character := range(s){
		z01.PrintRune(character)
	}
	z01.PrintRune('\n')
}