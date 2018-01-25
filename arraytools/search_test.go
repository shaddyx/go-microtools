package arraytools

import "testing"

func TestFind(t *testing.T) {
	arr := []string{"aaa", "bbb", "ccc"}
	res := Find(arr, "bbb")
	if res != 1 {
		t.Errorf("Position was incorrect, got: %d, want: %d.", res, 1)
	}

	res = Find(arr, "bbbzzz")
	if res != -1 {
		t.Errorf("Position was incorrect, got: %d, want: %d.", res, 1)
	}
}
