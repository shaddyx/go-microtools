package asserts

import "reflect"

func NotNil(obj interface{}, msg string){
	True(!reflect.ValueOf(obj).IsNil(), msg)
}

func True(val bool, msg string){
	if !val {
		panic(msg)
	}
}
