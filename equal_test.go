package fuzztestingingo

import "testing"

func TestEqual(t *testing.T) {
	if !Equal([]byte{'f', 'u', 'u', 'z'}, []byte{'f', 'u', 'u', 'z'}) {
		t.Error("expected True got False")
	}
	if Equal([]byte{'f', 'u', 'z', 'z'}, []byte{'f', 'u', 'u', 'z'}) {
		t.Error("expected False got True")
	}
}

func FuzzEqual(f *testing.F) {
	//seed inputs
	f.Add([]byte{'f', 'u', 'z', 'z'}, []byte{'f', 'u', 'z', 'z'})
	f.Add([]byte{'a', 'b', 'c'}, []byte{'a', 'b', 'd'})

	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		//self-equality check
		if !Equal(a, a) {
			t.Errorf("Expected Equal(%v, %v) to be true, got false", a, a)
		}

		//symmetry check
		if Equal(a, b) != Equal(b, a) {
			t.Errorf("Symmetry violated: Equal(%v, %v) != Equal(%v, %v)", a, b, b, a)
		}

		//length check
		if len(a) != len(b) && Equal(a, b) {
			t.Errorf("Expected Equal(%v, %v) to be false due to length mismatch", a, b)
		}
	})
}
