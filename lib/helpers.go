package lib

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Intptr(i int) *int {
	tmp := i
	return &tmp
}

func Capitalize(s *string) *string {
	if s == nil {
		return s
	}
	var newStr = strings.Title(*s)
	return &newStr
}
func Get_UTC_date_time() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}
func Log_error(err ...interface{}) {
	params := []interface{}{Get_UTC_date_time()}
	fmt.Fprintln(os.Stderr, append(params, err...)...)
}
func Normalize_query(q string) string {
	return strings.Replace(strings.Replace(q, "<nil>", "null", -1), "'null'", "null", -1)
}
func Int_or_nil(data *int) interface{} {
	if nil == data {
		return nil
	}
	return *data
}
func String_or_nil(data *string) interface{} {
	if nil == data {
		return nil
	}
	return *data
}
func Float_or_nil(data *float64) interface{} {
	if nil == data {
		return nil
	}
	return *data
}
func Bool_or_nil(data *bool) interface{} {
	if nil == data {
		return nil
	}
	return *data
}
func Time_or_nil(data *time.Time) interface{} {
	if nil == data {
		return nil
	}
	return data.Format("2006-01-02 15:04:05")
}
