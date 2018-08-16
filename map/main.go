package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":    "#ff0000",
		"blue":   "#008000",
		"yellow": "#ff0000",
		"white":  "#ffffff",
	}

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, val := range m {
		fmt.Printf("The color %v has a hex value of %v\n", key, val)
	}
}
