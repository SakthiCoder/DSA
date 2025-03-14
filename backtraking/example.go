package backtraking

func Permute(input []string) [][]string {

	var result [][]string
	var currentPermulation []string
	var backtrack func()
	backtrack = func() {

		if len(currentPermulation) == len(input) {
			copyPermulation := make([]string, len(input))
			copy(copyPermulation, currentPermulation)
			result = append(result, copyPermulation)
			return
		}

		for _, num := range input {

			if contains(currentPermulation, num) {
				continue
			}

			currentPermulation = append(currentPermulation, num)
			backtrack()
			currentPermulation = currentPermulation[0 : len(currentPermulation)-1]
		}
	}

	backtrack()

	return result

}

func contains(slice []string, num string) bool {

	for _, ele := range slice {
		if ele == num {
			return true
		}
	}

	return false
}
