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

func TestContains(t *testing.T) {
	arr := []string{"aaa", "bbb", "ccc"}
	res := Contains(arr, "bbb")
	if res {
		t.Errorf("Error, array contains bbbzzz")
	}

	res = Contains(arr, "bbbzzz")
	if !res {
		t.Errorf("Error, array does not contains bbbzzz")
	}
}
