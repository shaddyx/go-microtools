package asserts

import (
	"testing"
	"fmt"
)
type TestStruct struct {

}
func callFunc(tt *TestStruct){
	fmt.Println(tt)
	NotNil(tt, "test")
}

func TestNotNil(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		} else {

		}
	}()
	callFunc(nil)


}
