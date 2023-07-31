package language

import "fmt"

var f = func() string {
	fmt.Println("variable f initialized")
	return "test"
}()

func init() {
	fmt.Println("language init")
}

func EnSymbol() string {
	return "En"
}
