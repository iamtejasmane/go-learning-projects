package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#f12345",
		"white": "#fffff",
	}
	printMap(colors)

	// fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
