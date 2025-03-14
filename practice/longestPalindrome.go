package practice

import "fmt"

func LongestPalindrome(str string) string {

	if len(str) <= 1 {
		return str
	}

	maxStr := string(str[0])

	for i := 0; i < len(str)-1; i++ {

		odd := ExpandFromCenter(str, i, i)
		even := ExpandFromCenter(str, i, i+1)

		if len(odd) > len(maxStr) {
			maxStr = odd
		}

		if len(even) > len(maxStr) {
			maxStr = even
		}
	}

	return maxStr
}

func ExpandFromCenter(str string, left, right int) string {

	fmt.Println("str[left] == str[right]", str[left] == str[right])
	fmt.Println("len right: ", right < len(str))
	fmt.Println("len Left : ", left >= 0)
	for left >= 0 && right < len(str) && str[left] == str[right] {
		left--
		right++
	}

	return str[left+1 : right]
}
