package fuzztestingingo

func SafeDivision(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}
