package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func createUserBill() bill {
	reader := bufio.NewReader(os.Stdin)
	billName := getUserInput("Please enter a bill name: ", reader)
	return createBill(billName, map[string]float64{}, 0.0)
}

func getUserInput(prompt string, reader *bufio.Reader) string {
	fmt.Print(prompt)
	billName, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("[ERROR]: %20v\n", err))
	}
	return strings.TrimSpace(billName)
}

func promptOptions(b bill) bill {
	reader := bufio.NewReader(os.Stdin)
	opt := getUserInput("How can I help you?", reader)
	fmt.Println(opt)
	return b
}

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
	output := fmt.Sprintf("Bill %v breakdown: \n", b.name)
	total := 0.0

	for itemName, price := range b.items {
		output = fmt.Sprintf("%v%-20v ... $%0.2f\n", output, itemName, price)
		total += price
	}

	output = fmt.Sprintf("%v%-20v ... $%0.2f\n%-20v ... $%0.2f", output, "Tips", b.tip, "Total", total)
	return output
}
