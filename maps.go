package main

import "fmt"

var menu = map[string]float64{
	"a": 1,
	"b": 2,
	"c": 3,
}

func getMenu() map[string]float64 {
	return menu
}

func getPrice(foodName string) float64 {
	return menu[foodName]
}

func toString() string {
	output := ""
	for key, val := range menu {
		output = fmt.Sprintf("%v%v: %v\n", output, key, val)
	}

	return output
}

func updateMenu(foodName string, newPrice float64) map[string]float64 {
	menu[foodName] = newPrice
	return menu
}
