package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MyMenu struct {
	name  string
	price map[string]float64
}

var in = bufio.NewReader(os.Stdin)

func Printmenu() {
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, cost := range item.price {
			fmt.Printf("\t%10s%10.2f\n", size, cost)
		}
	}
}

func Addmenu() {
	fmt.Println("enter the item")
	name, _ := in.ReadString('\n')
	menu = append(menu, MyMenu{name: name, price: make(map[string]float64)})
	fmt.Println()
}
