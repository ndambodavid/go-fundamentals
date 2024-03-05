package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func demo() {
	var a string
	a = "maiko"
	fmt.Println(a)

	var z int = 3
	fmt.Println(float64(z))
}

func arrays() {
	var arr [4]string
	fmt.Println(arr)

	arr = [4]string{"Coffee", "Expresso", "Tea", "Cappucino"}
	fmt.Println(arr)
}

func slice() {
	var s []int
	fmt.Println(s)
	s = []int{1, 2, 3, 4}

	fmt.Println(s[1])
	s[1] = 50
	fmt.Println(s)

	s = append(s, 5, 6, 7, 8)
	fmt.Println(s)

	s = slices.Delete(s, 2, 6)
	fmt.Println(s)
}

func maps() {
	//var m map[string]int
	//fmt.Println(m)

	var ms map[string][]string
	fmt.Println(ms)

	ms = map[string][]string{
		"soft": {"Soda", "Energy", "Tea"},
		"hard": {"Whisky", "Gin", "B&W"},
	}
	fmt.Println(ms)

	delete(ms, "soft")
	fmt.Println(ms)
}

func structs() {
	fmt.Println("Please select option")
	fmt.Println("[1] menu")
	in := bufio.NewReader(os.Stdin)
	choice, _ := in.ReadString('\n')
	choice = strings.TrimSpace(choice)

	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "T-shirt", prices: map[string]float64{"small": 100.00, "medium": 150.00, "large": 200.00}},
		{name: "M-jeans", prices: map[string]float64{"small": 800.00, "medium": 900.00, "large": 1000.00}},
	}

	fmt.Println(menu)
}
