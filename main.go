package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type mystr struct {
		name  string
		price map[string]float64
	}
	menu := []mystr{
		{name: "coffee", price: map[string]float64{"small": 10, "medium": 20, "large": 30}},
		{name: "tea", price: map[string]float64{"small": 10, "medium": 20, "large": 30}},
	}
loop:
	for {
		in := bufio.NewReader(os.Stdin)
		fmt.Println(("pelsea select an item in the menu"))
		fmt.Println("1) print menu")
		fmt.Println("2) add item")
		fmt.Println("3) q")
		choice, _ := in.ReadString(('\n'))

		switch strings.TrimSpace(choice) {
		case "1":
			for _, item := range menu {
				fmt.Println(item.name)
				fmt.Println(strings.Repeat("-", 10))
				for size, cost := range item.price {
					fmt.Printf("\t%10s%10.2f\n", size, cost)
				}
			}
		case "2":
			{
				fmt.Println("enter the item")
				name, _ := in.ReadString('\n')
				menu = append(menu, mystr{name: name, price: make(map[string]float64)})
				fmt.Println()
			}
		case "3":
			{
				break loop
			}

		default:
			fmt.Println(("invalid"))

		}
	}
}
