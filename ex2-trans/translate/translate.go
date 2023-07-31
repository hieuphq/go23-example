package translate

import (
	"fmt"

	"github.com/dwarvesf/go23/ex2-trans/language"
)

func init() {
	fmt.Println("translate init")
}

func Print() {
	fmt.Println("Translate to " + language.EnSymbol())
}
