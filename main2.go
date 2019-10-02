package main

import (
	"fmt"
)

func main() {
	/*
		var locale, greeting string
		var languages = [4]string{"en", "es", "de", "fr"}
		locale = languages[1]
	*/
	var greeting string
	/*var languages = [4]string{"en", "es", "de", "fr"}*/
	/*
		if locale == "en" {
			greeting = "Hello"
		} else if locale == "es" {
			greeting = "Hola"
		} else if locale == "de" {
			greeting = "Guten Tag"
		} else {
			greeting = "Yo"
		}
	*/

	var lan string
	fmt.Printf("Please Type your language\n")
	fmt.Scanln(&lan)

	if lan == "english" {
		lan = "en"
	} else if lan == "espanol" {
		lan = "es"
	} else if lan == "deutsch" {
		lan = "de"
	} else if lan == "korean" {
		lan = "ko"
	} else {
		lan = "yo"
	}

	switch lan {
	case "en":
		greeting = "Hello"
	case "es":
		greeting = "Hola"
	case "de":
		greeting = "Guten Tag"
	case "fr":
		greeting = "Bonjour"
	case "ko":
		greeting = "Annyeonghaseyo"
	default:
		greeting = "Yo"
	}

	fmt.Printf(greeting + ", Go!\n")
}
