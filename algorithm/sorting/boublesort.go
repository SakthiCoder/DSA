package sorting

func BoubleSort(InputArray []int) {
	for i := 0; i < len(InputArray)-1; i++ {

		for j := i + 1; j < len(InputArray); j++ {

			if InputArray[i] > InputArray[j] {
				InputArray[i], InputArray[j] = InputArray[j], InputArray[i]
			}

		}
	}
}
