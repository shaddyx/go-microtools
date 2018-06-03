package arraytools

import "reflect"

func Remove(s interface{}, index int) interface{}{
	arrV := reflect.ValueOf(s)
	counter := 0
	if arrV.Kind() == reflect.Slice {
		var interfaceSlice []interface{} = make([]interface{}, arrV.Len() - 1)
		for i := 0; i<arrV.Len(); i++ {
			if i == index {
				continue
			}
			interfaceSlice[counter] = arrV.Index(i).Interface()
			counter ++
		}
		return interfaceSlice
	}
	panic("Argument is not a slice")

}

