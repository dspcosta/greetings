package greetings

import (
	"errors"
	"strings"
)

var greetingsHello = map[string]string{
	"portuguese": "Olá, tudo bem?",
	"english":    "Hello, how are you?",
	"dutch":      "Hallo, hoe gaat het?",
	"german":     "Hallo, wie geht es dir?",
	"italian":    "Ciao, come va?",
}

// GreetingsOptions Language to be used.
type GreetingsOptions struct {
	Language string
}

func (greetings GreetingsOptions) Hello() (msg string, err error) {
	if ch, ok := greetingsHello[greetings.Language]; ok {
		return ch, nil
	} else {
		return "", errors.New("language option not found")
	}
}

// HelloCustom method to generate the greeting message with the custom name provided in the end of phrase.
// Parameters:
// 	greetings GreetingsOptions - The language to be used.
// 	name string - The custom name to be used.
// Returns:
//	msg string - Greetings message.
//	err error - Error
func (greetings GreetingsOptions) HelloCustom(name string) (msg string, err error) {
	selection := strings.ToLower(greetings.Language)
	if ch, ok := greetingsHello[selection]; ok {
		return ch + " " + name, nil
	} else {
		return "", errors.New("language option not found")
	}
}

func Hello(language string) (choice string, err error) {
	if ch, ok := greetingsHello[language]; ok {
		return ch, nil
	} else {
		return "", errors.New("language option not found")
	}
}
