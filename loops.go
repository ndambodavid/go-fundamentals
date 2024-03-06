package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func loops() {
	i := 80
	for { //infinite
		fmt.Println(i)
		i += 2
		break
	}

	// looping with collections

	// for key, value := range collection { ... }
	// for key := range collection { ... }
	// for _, value := range collection { ... }

	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "T-shirt", prices: map[string]float64{"small": 100.00, "medium": 150.00, "large": 200.00}},
		{name: "M-jeans", prices: map[string]float64{"small": 800.00, "medium": 900.00, "large": 1000.00}},
	}

	fmt.Println("Please select option")
	fmt.Println("[1] menu")
	fmt.Println("[2] add item")
	fmt.Println("[q] quit")
	in := bufio.NewReader(os.Stdin)
	choice, _ := in.ReadString('\n')
	choice = strings.TrimSpace(choice)

	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}

}
