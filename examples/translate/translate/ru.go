package translate

import (
	"github.com/andrdru/loc"
)

var ruLocale = loc.NewLocale(ru)

var ru = map[string]string{
	Hello:     "Привет, %s!",
	HowAreYou: "Как дела?",
}
