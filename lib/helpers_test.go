package lib

import (
	"testing"
)

func Test_Prepare_on_duplicate_key_updates(t *testing.T) {
	var str = "tst"
	var nilptr *string
	testCases := map[string]string{
		Prepare_on_duplicate_key_updates(map[string]interface{}{
			"1": nil,
			"2": "hey",
			"3": &str,
			"4": nilptr,
			"5": 14,
			"6": 12.234,
		}): "on duplicate key update `2`='hey', `3`='tst', `5`=14, `6`=12.234",
		Prepare_on_duplicate_key_updates(map[string]interface{}{
			"1": nil,
		}): "",
	}
	for actual, expected := range testCases {
		if actual != expected {
			t.Fatalf("\nactual => %s\nexpected=>%s", actual, expected)
		}
	}
}
func Test_Calculate_live_match_status(t *testing.T) {
	var z = 0
	var o = 1
	var tw = 2
	_, _ = z, tw
	var tr = true
	cases := []map[interface{}]interface{}{
		{1: Calculate_live_status(&o, &o, &tr, &tr)},
		{0: Calculate_live_status(nil, &o, &tr, &tr)},
		{0: Calculate_live_status(&tw, &o, &tr, &tr)},
	}
	for _, cs := range cases {
		for expected, actual := range cs {
			if actual != expected {
				t.Fatalf("\nexpected=>%v\ngot=>%v", expected, actual)
			}
		}
	}
}
