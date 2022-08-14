package main

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

// func createBill(name string) bill {
// 	return createBill(name, map[string]float64{}, 0)
// }
