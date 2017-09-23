package filters

import (
	"time"
	"fmt"
	"reflect"
	"sync"
)

func ApplyFilters(data interface{}, filters map[string]string) {

	if len(filters) == 0 {
		fmt.Println("no filter")
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
	var wg sync.WaitGroup
	switch value.Kind() {
	case reflect.Struct:
		wg.Add(1)
		go loopStruct(&value, filters, &wg)
	case reflect.Slice:
		wg.Add(1)
		go loopSlice(&value, filters, &wg)
	}
	wg.Wait()
}

func loopStruct(val *reflect.Value, filters map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch {
		case field.Kind() == reflect.Struct:
			wg.Add(1)
			go loopStruct(&field, filters, wg)
		case field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct:
			wg.Add(1)
			var elem = field.Elem()
			go loopStruct(&elem, filters, wg)
		case field.Kind() == reflect.Slice:
			wg.Add(1)
			go loopSlice(&field, filters, wg)
		case field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Slice:
			wg.Add(1)
			var elem = field.Elem()
			go loopSlice(&elem, filters, wg)
		default:
			var tag = reflect.TypeOf(val.Interface()).Field(i).Tag.Get("filter")
			flt, ok := filters[tag];
			if !ok {
				continue
			}
			_ = flt
			//fmt.Println(field.CanSet(), field.Kind(), reflect.TypeOf(val.Interface()).Field(i).Tag.Get("filter"), reflect.TypeOf(val.Interface()).Field(i).Tag.Get("json"))
			filterField(&field, flt)

		}

	}
}
func loopSlice(val *reflect.Value, filters map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < val.Len(); i++ {
		var field = val.Index(i)
		switch {
		case field.Kind() == reflect.Struct:
			wg.Add(1)
			go loopStruct(&field, filters, wg)
		case field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct:
			wg.Add(1)
			var elem = field.Elem()
			go loopStruct(&elem, filters, wg)
		case field.Kind() == reflect.Slice:
			wg.Add(1)
			go loopSlice(&field, filters, wg)
		case field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Slice:
			wg.Add(1)
			var elem = field.Elem()
			go loopSlice(&elem, filters, wg)
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
