package main

// function that returns result of logroithm on base of N
func logofN(input int, base int) int {
	power := 1

	for ; input > 0; input /= base {
		power++
	}

	return power
}
