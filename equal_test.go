package fuzztestingingo

import "testing"

func TestEqual(t *testing.T) {
	if !Equal([]byte{'f', 'u', 'u', 'z'}, []byte{'f', 'u', 'u', 'z'}) {
		t.Error("expected True got False")
	}
}

func FuzzEqual(f *testing.F) {
	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		Equal(a, b)
	})
}
