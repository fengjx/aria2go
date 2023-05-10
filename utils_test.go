package aria2go

import "testing"

func TestRandString(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandString(10))
	}
}
