package main

import "fmt"

func main() {
	//ways to declare a map
	// colors := map[string]string{"red": "#FF0000", "white": "#FFFFFF", "black": "#000000"}
	// colors := make(map[string]string) // create and empty map
	// var colors map[string]string      //create and empty map

	// colors["red"] = "#FF0000" //add
	// delete(colors, "red")     //delete

	colors := map[string]string{
		"red":   "#FF0000",
		"white": "#FFFFFF",
		"black": "#000000",
	}

	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for mapKeyColor, mapValueHex := range c {
		fmt.Println("Hex code for", mapKeyColor, "is", mapValueHex)
	}
}
