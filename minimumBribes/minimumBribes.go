package minimumBribes

import (
	"fmt"
	"strconv"
)

type numOfBribes int

// Complete the minimumBribes function below.
func minimumBribes(q []int32) string {
	var bribes int
	var sorted bool
	swaps := make(map[int32]numOfBribes) // records the number of jumps for each person
	for !sorted {
		fmt.Println("Q = ", q, "bribes =", bribes, "swaps = ", swaps)
		for i := len(q) - 1; i-1 > -1; i-- {
			if i == 1 && q[0] == 1 {
				sorted = true
			}

			// check if curr person need to be swapped
			if q[i] < q[i-1] {
				if swaps[q[i-1]] == 2 {
					return "Too chaotic"
				}
				q[i], q[i-1] = q[i-1], q[i]
				bribes++
				swaps[q[i]]++
				break
			}
		}
	}
	return strconv.Itoa(bribes)
}
