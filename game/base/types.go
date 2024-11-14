package base

import "reflect"

func GetVarType(elem interface{}) string {
	errorInterface := reflect.TypeOf((*error)(nil)).Elem()
	if t := reflect.TypeOf(elem); t.Implements(errorInterface) {
		return elem.(error).Error()
	} else if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
