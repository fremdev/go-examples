package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	colors["white"] = "#ffffff"

	printMap(colors)

	// Examples of another way of creating maps
	// var colors1 = map[string]string
	// colors2 := make(map[string]string)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println(color + ": " + hex)
	}
}
