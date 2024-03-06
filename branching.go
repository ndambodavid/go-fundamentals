package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func branchs() {

	// switch test expression {
	//		case expression1:
	//			statements
	//		case expression2, expression3:
	//			statements
	//		default:
	//			statements

	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "T-shirt", prices: map[string]float64{"small": 100.00, "medium": 150.00, "large": 200.00}},
		{name: "M-jeans", prices: map[string]float64{"small": 800.00, "medium": 900.00, "large": 1000.00}},
	}

loop:
	for {
		fmt.Println("Please select option")
		fmt.Println("[1] menu")
		fmt.Println("[2] add item")
		fmt.Println("[q] quit")
		in := bufio.NewReader(os.Stdin)
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {

		case "1":

			for _, item := range menu {
				fmt.Println(item.name)
				fmt.Println(strings.Repeat("-", 10))
				for size, price := range item.prices {
					fmt.Printf("\t%10s%10.2f\n", size, price)
				}
			}
		case "2":
			fmt.Println("Please enter the name of the new item:")
			name, _ := in.ReadString('\n')
			menu = append(menu, menuItem{name: name, prices: make(map[string]float64)})

		case "q":
			break loop

		default:
			fmt.Println("Unknown option")
		}
	}

}
