package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func createBill(name string, items map[string]float64, tip float64) bill {
	newBill := bill{
		name:  name,
		items: items,
		tip:   tip,
	}
	return newBill
}

//	func createBill(name string) bill {
//		return createBill(name, map[string]float64{}, 0)
//	}
func (b *bill) updateItems(itemName string, price float64) *bill {
	b.items[itemName] = price
	return b
}

// Go uses pass by value for struct too, hence use pointer of struct instead
// For struct, pointers are automatically dereferenced.
func (b *bill) updateTip(newTip float64) {
	b.tip = newTip
}

func (b *bill) toString() string {
	output := "Bill breakdown: \n"
	total := 0.0

	for itemName, price := range b.items {
		output = fmt.Sprintf("%v%-20v ... $%0.2f\n", output, itemName, price)
		total += price
	}

	output = fmt.Sprintf("%v%-20v ... $%0.2f\n%-20v ... $%0.2f", output, "Tips", b.tip, "Total", total)
	return output
}
