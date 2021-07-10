package loc

import (
	"fmt"
)

type (
	// Translation interface for locale selection
	Translation interface {
		L(locale string) Locale
	}

	// Locale interface for phrase translation
	Locale interface {
		T(key string, args ...interface{}) string
	}

	// translation implements Translation
	translation struct {
		locales map[string]Locale
	}

	// locale implements Locale
	locale struct {
		keymap map[string]string
	}

	// noLocale no-op Locale implementation for default language
	noLocale struct{}
)

var _ Translation = &translation{}
var _ Locale = &locale{}
var _ Locale = &noLocale{}

// NewTranslation new translator
func NewTranslation(languages map[string]Locale) *translation {
	return &translation{
		locales: languages,
	}
}

// NewLocale new locale
func NewLocale(m map[string]string) *locale {
	return &locale{
		keymap: m,
	}
}

// NewNoLocale new empty locale
func NewNoLocale() *noLocale {
	return &noLocale{}
}

// L get Locale
func (t *translation) L(locale string) Locale {
	var l, ok = t.locales[locale]
	if !ok {
		return &noLocale{}
	}

	return l
}

// T get key translation for current locale
func (l *locale) T(key string, args ...interface{}) string {
	var translatedKey, ok = l.keymap[key]
	if !ok {
		return fmt.Sprintf(key, args...)
	}

	return fmt.Sprintf(translatedKey, args...)
}

// T key with default language
func (n *noLocale) T(key string, args ...interface{}) string {
	return fmt.Sprintf(key, args...)
}
