package filters

import (
	"reflect"
	"fmt"
)

var ValidFilters = []string{
	"^-$", //closed
	"^\\d$",
}

func FilterDisableFunc(val *reflect.Value) {
	//reflect.Zero((*val).Type())
	val.SetInt(13)
	fmt.Println(val)
}

func DoFilter(filterSTR string) {

}
