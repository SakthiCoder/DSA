package goroutine

import (
	"fmt"
	"time"
)

func GoRoutine(char string) {
	i := 0
	for {
		fmt.Println(char, i)
		i++
		time.Sleep(2 * time.Second)
	}
}
