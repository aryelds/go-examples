package main

import "fmt"

func main() {
	cor := make(map[string]string)
	colors := map[string]string {
		"red": "#ff0000",
		"green": "#4bf785",
		"white": "#fffff",
	}

	fmt.Println(cor)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}

}
