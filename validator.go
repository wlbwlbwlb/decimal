package decimal

import "reflect"

func CustomFunc(field reflect.Value) interface{} {
	if val, ok := field.Interface().(Decimal); ok {
		f, _ := val.Float64()
		return f
	}
	return nil
}
