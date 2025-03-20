package practice

import "fmt"

func Pattern1() {
	fmt.Println()

	k := 3
	n := 7
	for row := 0; row <= n; row++ {
		for col := 0; col <= n; col++ {
			if (row*col-1)%2 == 1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		k--
		fmt.Println()
	}
}
