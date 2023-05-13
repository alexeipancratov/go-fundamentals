package menu

import (
	"fmt"
	"strings"
)

func Demo() {
loop:
	for {
		fmt.Println("Please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) quit")
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			Print()
		case "2":
			err := Add()
			if err != nil {
				fmt.Println(fmt.Errorf("invalid input: %w", err))
			}
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}
	}
}
