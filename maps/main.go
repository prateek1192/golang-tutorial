package main

import "fmt"

func main(){

	// Method of declaring maps

	// Method 1
	var colors1 map[string]string

	// Method 2
	colors2 := map[string]string{
		"red": "#ff0000",
		"green" : "#4bf745",
	}

	colors3 := make(map[string]string)

	colors3["white"] = "#ffffff"

	fmt.Println(colors1)
	fmt.Println(colors2)
	fmt.Println(colors3)

	delete(colors2, "green")
	fmt.Println(colors2)

	colors4 := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors4)

}

func printMap(c map[string]string){
	for key, value := range c {
		fmt.Println("Hex code for", key, "is ", value)
	}
}