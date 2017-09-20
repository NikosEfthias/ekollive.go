package filters

import (
	"reflect"
	"sync"
	"time"
	"fmt"
)

func ApplyFilters(data interface{}, filters map[string]string, waitgroup ...*sync.WaitGroup) {
	var child bool //main node or a child node
	if len(waitgroup) > 0 {
		child = true
		defer waitgroup[0].Done()
	}
	if len(filters) == 0 {
		fmt.Println("no filter")
		return
	}
	if testing {
		//reduce the throughput for testing
		time.Sleep(time.Second)
	}

	var wg = new(sync.WaitGroup)

	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		var sub = val.Field(i)
		var subType = val.Type().Field(i)
		//var valid bool
		if sub.Kind() == reflect.Ptr {
			sub = sub.Elem()
		}
		switch sub.Kind() {
		case reflect.Slice:
			//handle slices
			for ind := 0; ind < sub.Len(); ind++ {
				item := sub.Index(ind)
				if item.Kind() == reflect.Ptr {
					item = item.Elem()
				}
				switch item.Kind() {
				case reflect.Struct:
					//handle struct
					wg.Add(1)
					go ApplyFilters(item.Interface(), filters, wg)
				}
			}
			continue
		case reflect.Struct:
			//handle struct
			wg.Add(1)
			go ApplyFilters(val.Field(i).Interface(), filters, wg)
			continue
		}
		//
		//not a struct or slice, eligible for filtering
		if child && !sub.CanSet() {
			if !sub.IsValid() {
				sub = val.Field(i)
			}
			if sub.Kind()==reflect.Int{
				fmt.Println(subType.Tag.Get("filter"))
			}
		}
		tag := subType.Tag.Get("filter")
		if tag == "" || filters[tag] == "" || !sub.CanSet() {
			continue
		}
		tag = filters[tag]
		//fmt.Println("test")
		//sub.SetInt(15)
		//FilterDisableFunc(&sub)
		//fmt.Println(sub.Interface())
	}

	wg.Wait()
}
