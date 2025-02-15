package fuzztestingingo

import "testing"

func FuzzSafeDivision(f *testing.F) {
	f.Fuzz(func(t *testing.T, a int, b int) {
		SafeDivision(a, b)
	})
}
