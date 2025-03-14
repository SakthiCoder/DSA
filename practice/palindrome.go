package practice

func IsPalindrome(str string) bool {

	if len(str) == 0 || len(str) == 1 {
		return false
	}

	left := 0
	right := len(str) - 1

	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}

	return true
}
