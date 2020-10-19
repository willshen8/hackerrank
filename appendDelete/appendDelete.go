package appendDelete

import (
	"math"
)

func appendAndDelete(s string, t string, k int32) bool {
	lengthDifference := len(s) - len(t)
	firstIndexDiff := firstDifferentIndexPosition(s, t)
	if math.Abs(float64(lengthDifference)) > float64(k) {
		return false
	}
	if firstIndexDiff == -1 {
		if int(k)-len(t)-len(s) < 0 {
			return true
		}
		if (k-int32(math.Abs(float64(lengthDifference))))%2 != 0 {
			return false
		}
	} else {
		if lengthDifference > 0 && (int(k)-lengthDifference)%2 != 0 {
			return false
		} else if lengthDifference < 0 && (k-int32(math.Abs(float64(lengthDifference))))%2 != 0 {
			return false
		} else if lengthDifference == 0 && (((len(s) - firstIndexDiff) * 2) > int(k)) {
			return false
		}
	}
	return true
}

// firstDifferentIndexPosition returns the first index of the two strings that is different, otherwise -1
func firstDifferentIndexPosition(a, b string) int {
	minimum := math.Min(float64(len(a)), float64(len(b)))
	for i := 0; i < int(minimum); i++ {
		if a[i] != b[i] {
			return i
		}
	}
	return -1
}
