package translate

import (
	"github.com/andrdru/loc"
)

var T = loc.NewTranslation(map[string]loc.Locale{
	EN: loc.NewNoLocale(),
	RU: ruLocale,
})

//nolint:golint
const (
	EN = "en"
	RU = "ru"
)
