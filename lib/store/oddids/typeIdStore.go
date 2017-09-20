package oddids

import "sync"
import (
	"../../../models/oddType"
	"strconv"
)

type typeids struct {
	sync.Mutex
	store map[string]int
}

var store *typeids

func init() {
	store = new(typeids)
	store.store = make(map[string]int)
}

func LoadAll() {
	var data = make([]oddType.Oddtype, 0)
	oddType.Model.Find(&data)
	store.Lock()
	defer store.Unlock()
	for _, d := range data {
		if d.Type == nil && d.Typeid == nil && d.Subtype == nil {
			continue
		}
		store.store[returnKey(d.Type, d.Subtype, d.Typeid)] = *d.Oddtypeid
	}
}
func Get(tp, subtp *string, tpid *int) int {
	val, ok := store.store[returnKey(tp, subtp, tpid)]
	if ok {
		return val
	}
	return 0
}
func returnKey(tp, subtp *string, tpid *int) string {
	var key string

	if tp != nil {
		key += *tp
	}
	if subtp != nil {
		key += *subtp
	}
	if tpid != nil {
		key += strconv.Itoa(*tpid)
	}
	return key

}
