package practice

import "fmt"

func Pattern1() {
	fmt.Println()

	k := 3
	n := 7
	for row := 1; row <= n; row++ {

		for k := 1; k < n; k++ {
			
		}

		for col := 1; col <= n; col++ {
			// fmt.Print("* ")
		}
		k--
		fmt.Println()
	}
}
