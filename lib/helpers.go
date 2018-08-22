package lib

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
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

func Time_or_nil(data *string) interface{} {
	if nil == data {
		return nil
	}
	tm := strings.Split(*data, ".")[0]
	if strings.HasSuffix(tm, "Z") {
		sp := strings.Split(tm, ":")
		return strings.Join(sp[:len(sp)-1], ":")
	}
	return tm
}

type Side int

const (
	Side_home Side = iota
	Side_away
)

func Split_score_fields(data *string, side Side) *int {
	if nil == data {
		return nil
	}
	d := strings.Split(*data, ":")
	if len(d) < 2 {
		Log_error("invalid score field", *data)
		return nil
	}
	i__field, err := strconv.Atoi(d[side])
	if nil != err {
		Log_error("invalid score field", *data)
		return nil
	}
	return &i__field
}
func Prepare_on_duplicate_key_updates(fields map[string]interface{}) string {
	var product []string
	for k, v := range fields {
		if v == nil {
			continue
		}
		var current string
		val := reflect.ValueOf(v)
		if v == reflect.Zero(val.Type()).Interface() {
			continue
		} else if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		if val.Kind() == reflect.String {
			current = fmt.Sprintf("`%s`='%v'", k, val.Interface())
		} else {
			current = fmt.Sprintf("`%s`=%v", k, val.Interface())
		}
		product = append(product, current)
	}
	if len(product) > 0 {
		return "on duplicate key update " + strings.Join(product, ", ")
	}
	return ""
}
func Calculate_live_status(mtc_status *int, live_status *int, is_live *bool, is_started *bool) int {
	if nil == mtc_status || nil == live_status || nil == is_live || nil == is_started {
		return 0
	}
	var current uint64
	const (
		mts_0 uint64 = 1 << iota
		mts_1
		mts_2
		ls_1
		ls_2
		live
		started
	)
	switch *mtc_status {
	case 0:
		current |= mts_0
	case 1:
		current |= mts_1
	case 2:
		current |= mts_2
	}
	switch *live_status {
	case 1:
		current |= ls_1
	case 2:
		current |= ls_2
	}
	if *is_live {
		current |= live
	}
	if *is_started {
		current |= started
	}
	switch current {
	case mts_0 | ls_1:
		fallthrough
	case mts_0 | ls_1 | live:
		fallthrough
	case mts_2 | ls_2:
		return 0
	case mts_1 | ls_1 | live | started:
		return 1
	default:
		return 0
	}
	return 0
}
func Calculate_live_match_status(mtc_status *int, period *int) *string {
	if nil == period || nil == mtc_status {
		return nil
	}
	var cases = map[int]string{
		0: "not_started",
		1: fmt.Sprint(*period),
		2: "stopped",
		3: "not_started",
	}
	v, ok := cases[*mtc_status]
	if !ok {
		return nil
	}
	return &v
}
func Bool_to_int(n *bool) int {
	if nil == n {
		return 0
	}
	switch *n {
	case true:
		return 1
	case false:
		return 0
	}
	return 0
}
