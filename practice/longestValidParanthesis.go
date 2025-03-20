package practice

import "dsa/stack"

func LongestValidParanthesis(str string) {
	var Stack stack.Stack

	// maxLen := 0
	for i := 0; i < len(str); i++ {
		if Stack.IsEmpty() || str[i] == '(' {
			Stack.Push(str[i])
		} else {
			Stack.Pop()

		}
	}

}
