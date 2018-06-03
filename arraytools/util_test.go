package arraytools

import (
	"testing"
	"fmt"
)

func TestRemove(t *testing.T) {
	arr := []string{"aaa", "bbb", "ccc"}
	a := Remove(arr, 1)
	fmt.Print(a)
}
