package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	horoscopes := map[string][]string{
		"Aries":       {"Today is a day to take bold actions.", "Stay cautious of overconfidence.", "A surprise opportunity awaits you."},
		"Taurus":      {"Focus on your finances today.", "A calming energy will surround you.", "Be patient; success is on the horizon."},
		"Gemini":      {"Communicate clearly with those around you.", "Your curiosity will lead to exciting discoveries.", "Expect a meaningful conversation today."},
		"Cancer":      {"Focus on home and family matters.", "Your intuition is strong—trust it.", "Take some time to relax and recharge."},
		"Leo":         {"Show your leadership skills.", "A chance for recognition is coming your way.", "Avoid unnecessary drama today."},
		"Virgo":       {"Pay attention to the details in your work.", "A helpful gesture will brighten your day.", "Find balance between perfection and practicality."},
		"Libra":       {"Seek harmony in relationships.", "An artistic pursuit will bring you joy.", "Compromise will lead to success."},
		"Scorpio":     {"Embrace your inner passion.", "A mystery will intrigue you.", "Trust is the key to progress today."},
		"Sagittarius": {"Adventure is calling you—be open to it.", "Your optimism will inspire others.", "Focus on broadening your horizons."},
		"Capricorn":   {"Hard work will pay off today.", "Stay grounded and practical.", "Set long-term goals for success."},
		"Aquarius":    {"Your innovative ideas will shine.", "Be ready for an unexpected meeting.", "Stay open-minded to new perspectives."},
		"Pisces":      {"Let your imagination guide you.", "A dream might reveal an important insight.", "Help someone in need—they'll remember it."},
	}

	fmt.Println("Daily Horoscope Generator")
	fmt.Println("==========================")
	fmt.Println("Enter your zodiac sign or type 'exit' to quit.\n")

	for {
		fmt.Print("Your zodiac sign: ")
		var sign string
		fmt.Scanln(&sign)

		if strings.ToLower(sign) == "exit" {
			fmt.Println("Goodbye! Have a great day!")
			break
		}

		sign = strings.Title(strings.ToLower(sign)) // Normalize input
		if predictions, exists := horoscopes[sign]; exists {
			randomIndex := rand.Intn(len(predictions))
			fmt.Printf("\nToday's Horoscope for %s: %s\n\n", sign, predictions[randomIndex])
		} else {
			fmt.Println("Invalid zodiac sign. Please try again.")
		}
	}
}
