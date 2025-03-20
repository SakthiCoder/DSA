package main

import (
	"dsa/algorithm/searching"
	"dsa/algorithm/sorting"
	"dsa/backtraking"
	"dsa/practice"
	"dsa/queue"
	"dsa/recursion"
	"dsa/slidingwindow"
	"dsa/tree"
	"fmt"
)

func main() {

	/* Stack := &stack.Stack{}
	Stack.Push(2)
	Stack.Push(3)
	Stack.Pop()
	fmt.Println("Stack.peek()", Stack.Peek())
	Stack.Push("Sakthivel")
	Stack.Push("Ramana")
	Stack.Push("Ram")
	Stack.Pop()

	fmt.Println("isEmpty", Stack.IsEmpty()) */

	queue := &queue.Queue{}

	queue.Push(34)
	fmt.Println(queue.Peek())
	queue.Push(434)
	queue.Push(33)
	fmt.Println(queue.Peek())
	queue.Pop()
	queue.Pop()
	queue.Pop()
	fmt.Println("queue.IsEmpty()", queue.IsEmpty())
	fmt.Println(queue.Peek())

	for i := 6; i > 0; i-- {
		recursion.Example1(i)
		fmt.Println(" ")
	}

	fmt.Println("Factorial: ", recursion.Factorial(5))

	fmt.Println("Fibonacci: ", recursion.Fibonacci(6))

	fmt.Println("BinarySearch: ", searching.BinarySearch([]int{3, 5, 12, 14, 16}, 3))
	// fmt.Println("isPalindrom: ", practice.IsPalindrome("MalayalaM"))

	arr := []int{3, 1, 2, 5, 0, 10}

	sorting.BoubleSort(arr)
	fmt.Println("Bouble Sort: ", arr)

	arr = []int{3, 1, 2, 5, 0, 10}
	sorting.QuickSort(arr, 0, len(arr)-1)

	fmt.Println("Quick Sort: ", arr)

	arr = []int{3, 1, 2, 5, 0, 10}
	fmt.Println("Merge Sort Before: ", arr)

	sorting.MergeSort(arr)

	fmt.Println("Merge Sort After: ", arr)

	arr = []int{3, 1, 2, 5, 0, 10}
	sorting.SelectionSort(arr)
	fmt.Println("Selection Sort: ", arr)

	arr = []int{3, 1, 2, 5, 0, 10}
	fmt.Println("Insertion Sort Before: ", arr)
	sorting.InsertionSort(arr)
	fmt.Println("Insertion Sort After: ", arr)

	// Sliding Window
	arr = []int{10, 200, 300, 400, 500}
	fmt.Println("Sliding Window: ", slidingwindow.SumofSubArray(arr, 1, "max"))

	arr = []int{2, 3, 1, 2, 4, 3}
	target := 7

	// Get the result
	result := slidingwindow.MinSubArrayLen(target, arr)

	fmt.Println("Dynamic Sliding Window: ", result)

	arr = []int{1, 2, 1, 2, 3}
	k := 2

	// Get the result
	result = slidingwindow.LongestSubarrayWithKDistinct(arr, k)

	fmt.Println("Dynamic Sliding Window: ", result)
	// Example input
	nums := []string{"A", "B", "C"}

	// Get all permutations
	permutations := backtraking.Permute(nums)

	// Print the result
	fmt.Println("Permutations:")

	for _, perm := range permutations {
		fmt.Println(perm)
	}

	maze := [][]int{
		{1, 1, 1, 1},
		{1, 1, 0, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 1},
	}

	N := 4
	paths := backtraking.FindPaths(maze, N)
	if len(paths) > 0 {
		fmt.Println("All possible paths in lexicographical order:")
		for _, path := range paths {
			fmt.Println(path)
		}
	} else {
		fmt.Println("No path exists!")
	}

	/* paths = backtraking.FindPathsUsingLoop(maze, N)
	if len(paths) > 0 {
		fmt.Println("All possible paths in lexicographical order:")
		for _, path := range paths {
			fmt.Println(path)
		}
	} else {
		fmt.Println("No path exists!")
	} */

	/* ============= Tree ============= */

	var tree tree.BinaryTree

	tree.Insert(1)
	tree.Insert(10)
	tree.Insert(0)
	tree.Insert(9)
	tree.Insert(8)
	tree.Insert(7)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(8)
	tree.InOrderTraversal()

	arr = []int{1, 2, 3, 4}
	practice.RotateArray(arr, -1)
	fmt.Println("Rotated Array: ", arr)

	str := practice.LongestPalindrome("bbbaaabbbabcd")
	fmt.Println("Longest Palimdrome Substring : ", str)

	practice.Pattern1()
}
