package main

import "fmt"

func main() {
	// var colors map[string]string
	colors := make(map[string]string)
	// colors := map[string]string{
	// 	"red":   "#ff000",
	// 	"green": "#f12345",
	// }
	colors["white"] = "#ff0000"

	delete(colors, "white")

	fmt.Println(colors)
}
