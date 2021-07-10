package main

import (
	"fmt"

	t "github.com/andrdru/loc/examples/translate/translate"
)

func main() {
	var en = t.T.L(t.EN)
	var ru = t.T.L(t.RU)

	fmt.Println(en.T(t.Hello, "Vasya"), en.T(t.HowAreYou))
	fmt.Println(ru.T(t.Hello, "Вася"), ru.T(t.HowAreYou))
}
