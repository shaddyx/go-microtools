package arraytools

import "reflect"

//
//
//
func Find(s interface{}, elem interface{}) int {
	arrV := reflect.ValueOf(s)
	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {

			// XXX - panics if slice element points to an unexported struct field
			// see https://golang.org/pkg/reflect/#Value.Interface
			if arrV.Index(i).Interface() == elem {
				return i
			}
		}
	}

	return -1
}

func Contains(s interface{}, elem interface{}) bool {
	return Find(s, elem) != -1
}

