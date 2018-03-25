package filters

import (
	"reflect"
	"time"
)

func ApplyFilters(data interface{}, filters map[string]string) {

	if len(filters) == 0 {
		return
	}
	if testing {
		//reduce the throughput for testing
		time.Sleep(time.Second)
	}
	var value = reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr && (value.Elem().Kind() == reflect.Struct || value.Elem().Kind() == reflect.Slice) {
		value = value.Elem()
	}
	switch value.Kind() {
	case reflect.Struct:
		loopStruct(&value, filters)
	case reflect.Slice:
		loopSlice(&value, filters)
	}
}

func loopStruct(val *reflect.Value, filters map[string]string) {
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch {
		case field.Kind() == reflect.Struct:
			loopStruct(&field, filters)
		case field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct:
			var elem = field.Elem()
			loopStruct(&elem, filters)
		case field.Kind() == reflect.Slice:
			loopSlice(&field, filters)
		case field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Slice:
			var elem = field.Elem()
			loopSlice(&elem, filters)
		default:
			var tag = reflect.TypeOf(val.Interface()).Field(i).Tag.Get("filter")
			flt, ok := filters[tag]
			if !ok {
				continue
			}
			_ = flt
			//fmt.Println(field.CanSet(), field.Kind(), reflect.TypeOf(val.Interface()).Field(i).Tag.Get("filter"), reflect.TypeOf(val.Interface()).Field(i).Tag.Get("json"))
			filterField(&field, flt)

		}

	}
}
func loopSlice(val *reflect.Value, filters map[string]string) {
	for i := 0; i < val.Len(); i++ {
		var field = val.Index(i)
		switch {
		case field.Kind() == reflect.Struct:
			loopStruct(&field, filters)
		case field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct:
			var elem = field.Elem()
			loopStruct(&elem, filters)
		case field.Kind() == reflect.Slice:
			loopSlice(&field, filters)
		case field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Slice:
			var elem = field.Elem()
			loopSlice(&elem, filters)
		}
	}
}
func filterField(field *reflect.Value, filter string) {
	//var fld reflect.Value
	for r, f := range Filter {
		if !r.MatchString(filter) {
			continue
		}
		f(field, filter)
	}
}
