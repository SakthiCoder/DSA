package backtraking

import "fmt"

func FindPaths(InputMaze [][]int, GridSize int) []string {
	var PossiblePaths []string

	if InputMaze[0][0] == 0 {
		return PossiblePaths
	}

	visited := make([][]bool, GridSize)
	for i := range visited {
		visited[i] = make([]bool, GridSize)
	}

	findPathsUtil(InputMaze, 0, 0, GridSize, "", visited, &PossiblePaths)

	return PossiblePaths
}

func findPathsUtil(InputMaze [][]int, Row int, Column int, GridSize int, Path string, Visited [][]bool, PossiblePaths *[]string) {

	if Row == GridSize-1 && Column == GridSize-1 {
		*PossiblePaths = append(*PossiblePaths, Path)
		return
	}

	Visited[Row][Column] = true

	// Move To the Down
	if isValidMove(InputMaze, Row+1, Column, GridSize) && !Visited[Row+1][Column] {
		findPathsUtil(InputMaze, Row+1, Column, GridSize, Path+"D", Visited, PossiblePaths)
	}

	if isValidMove(InputMaze, Row, Column-1, GridSize) && !Visited[Row][Column-1] {
		findPathsUtil(InputMaze, Row, Column-1, GridSize, Path+"L", Visited, PossiblePaths)
	}

	if isValidMove(InputMaze, Row, Column+1, GridSize) && !Visited[Row][Column+1] {
		findPathsUtil(InputMaze, Row, Column+1, GridSize, Path+"R", Visited, PossiblePaths)
	}

	if isValidMove(InputMaze, Row-1, Column, GridSize) && !Visited[Row-1][Column] {
		findPathsUtil(InputMaze, Row-1, Column, GridSize, Path+"U", Visited, PossiblePaths)
	}

	// to BackTrack Undo the Current Path Because Try in Another Way To Find Path
	Visited[Row][Column] = false

}

type State struct {
	Row, Column int
	Path        string
}

func FindPathsUsingLoop(InputMaze [][]int, GridSize int) []string {
	var PossiblePaths []string

	// If the starting point is blocked, return no path
	if InputMaze[0][0] == 0 {
		return PossiblePaths
	}

	visited := make([][]bool, GridSize)
	for i := range visited {
		visited[i] = make([]bool, GridSize)
	}

	// Stack to simulate recursion (using row, column, and path)
	stack := []State{{Row: 0, Column: 0, Path: ""}}
	visited[0][0] = true

	// Loop to simulate the DFS traversal
	for len(stack) > 0 {
		fmt.Println("i")
		// Pop the top element from the stack
		currentState := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// If we've reached the destination, append the path
		if currentState.Row == GridSize-1 && currentState.Column == GridSize-1 {
			PossiblePaths = append(PossiblePaths, currentState.Path)
			continue
		}

		// Try moving to Down, Left, Right, and Up directions (in this order for lexicographical order)
		// Down
		if isValidMove(InputMaze, currentState.Row+1, currentState.Column, GridSize) && !visited[currentState.Row+1][currentState.Column] {
			visited[currentState.Row+1][currentState.Column] = true
			stack = append(stack, State{Row: currentState.Row + 1, Column: currentState.Column, Path: currentState.Path + "D"})
		}
		// Left
		if isValidMove(InputMaze, currentState.Row, currentState.Column-1, GridSize) && !visited[currentState.Row][currentState.Column-1] {
			visited[currentState.Row][currentState.Column-1] = true
			stack = append(stack, State{Row: currentState.Row, Column: currentState.Column - 1, Path: currentState.Path + "L"})
		}
		// Right
		if isValidMove(InputMaze, currentState.Row, currentState.Column+1, GridSize) && !visited[currentState.Row][currentState.Column+1] {
			visited[currentState.Row][currentState.Column+1] = true
			stack = append(stack, State{Row: currentState.Row, Column: currentState.Column + 1, Path: currentState.Path + "R"})
		}
		// Up
		if isValidMove(InputMaze, currentState.Row-1, currentState.Column, GridSize) && !visited[currentState.Row-1][currentState.Column] {
			visited[currentState.Row-1][currentState.Column] = true
			stack = append(stack, State{Row: currentState.Row - 1, Column: currentState.Column, Path: currentState.Path + "U"})
		}

		// Backtrack: Unmark the current cell as visited **after** all possible moves have been explored.
		// We should only unmark it once we pop it off the stack, meaning we have fully explored all paths from that cell.
	}

	return PossiblePaths
}

func isValidMove(InputMaze [][]int, Row int, Column int, GridSize int) bool {
	// Check bounds and if the cell is open (1)
	return Row >= 0 && Column >= 0 && Row < GridSize && Column < GridSize && InputMaze[Row][Column] == 1
}
