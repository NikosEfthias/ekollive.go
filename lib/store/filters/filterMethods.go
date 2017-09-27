package filters

import (
	"reflect"
	"regexp"
	"strconv"
)

var Filter = map[*regexp.Regexp]func(*reflect.Value, string){
	regexp.MustCompile("^-$"):  filterDisableFunc,
	regexp.MustCompile("^\\d+$"): setInt,
}

func filterDisableFunc(val *reflect.Value, _ string) {
	if val.CanSet() {
		val.Set(reflect.Zero(val.Type()))
	}
}
func setInt(val *reflect.Value, filter string) {
	s, _ := strconv.Atoi(filter)

	switch val.Kind() {
	case reflect.Ptr:
		elem:=val.Type().Elem().Kind()
		_=elem
		if val.Elem().CanSet() && val.Elem().Kind() == reflect.Int {
			val.Elem().SetInt(int64(s))
		} else if val.CanSet() && val.Type().Elem().Kind() == reflect.Int {
			val.Set(reflect.New(val.Type().Elem()))
			val.Elem().SetInt(int64(s))
		}

	case reflect.Int:
		if val.CanSet() {
			val.SetInt(int64(s))
		}

	}
}
