package hungarian

// Solve solves assignment problem using Hungarian algorithm.
// Returns assignment and total cost.
func Solve(cost [][]int) ([]int, int) {
	n := len(cost)
	if n == 0 {
		return []int{}, 0
	}

	// TODO: implement Hungarian algorithm
	// Temporary stub - assigns each row to same column
	assignment := make([]int, n)
	total := 0
	for i := 0; i < n; i++ {
		assignment[i] = i
		total += cost[i][i]
	}

	return assignment, total
}
