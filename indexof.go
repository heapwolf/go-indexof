package indexof

import "reflect"

func IndexOf(params ...interface{}) int {
	v := reflect.ValueOf(params[0])
	arr := reflect.ValueOf(params[1])

	if reflect.TypeOf(params[1]).Kind() != reflect.Slice {
		panic("Type Error! Second argument must be a slice")
	}

	for i := 0; i < arr.Len()-1; i++ {
		if arr.Index(i).Interface() == v.Interface() {
			return i
		}
	}
	return -1
}
