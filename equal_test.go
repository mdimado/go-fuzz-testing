package fuzztestingingo

import "testing"

func TestEqual(t *testing.T) {
	if !Equal([]byte{'f', 'u', 'u', 'z'}, []byte{'f', 'u', 'u', 'z'}) {
		t.Error("expected True got False")
	}
}
