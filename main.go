package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"demo/menu"
)

var in = bufio.NewReader(os.Stdin)

func main() {
loop:
	for {

		fmt.Println(("pelsea select an item in the menu"))
		fmt.Println("1) print menu")
		fmt.Println("2) add item")
		fmt.Println("3) q")
		choice, _ := in.ReadString(('\n'))

		switch strings.TrimSpace(choice) {
		case "1":
			menu.Printmenu()
		case "2":
			menu.Addmenu()
		case "3":
			{
				break loop
			}

		default:
			fmt.Println(("invalid"))

		}
	}
}
