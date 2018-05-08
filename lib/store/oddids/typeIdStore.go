package oddids

import "sync"
import (
	"../../../models/oddType"
	"strconv"
	"../../../models"
	"strings"
	"fmt"
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
	store.Lock()
	defer store.Unlock()
	val, ok := store.store[returnKey(tp, subtp, tpid)]
	if ok {
		return val
	}
	return 0
}
func Set(o *oddType.Oddtype) int {
	if val := Get(o.Type, o.Subtype, o.Typeid); val != 0 {
		o.Oddtypeid = &val
		return val
	}

	store.Lock()
	defer store.Unlock()
	oddType.Model.Where(&oddType.Oddtype{
		Subtype: o.Subtype,
		Type:    o.Type,
		Typeid:  o.Typeid,
	}).FirstOrCreate(o)
	if o.Oddtypeid != nil {
		store.store[returnKey(o.Type, o.Subtype, o.Typeid)] = *o.Oddtypeid
		return *o.Oddtypeid
	}
	return 0
}

func SetById(odd *models.Odd) error {
	var (
		temp      string
		tempSlice []string
		err       error
	)
	if nil==odd{
			return fmt.Errorf("null odd")
	}
	store.Lock()
	defer store.Unlock()
	for k, v := range store.store {
		if v == *odd.OddTypeId {
			temp = k
			break
		}
	}
	if "" == temp {
		return fmt.Errorf("oddtypeid could not found")
	}
	tempSlice = strings.Split(temp, "|")

	odd.Type = new(string)
	odd.Subtype = new(string)
	odd.Typeid = new(int)

	*odd.Type = tempSlice[0]
	if len(tempSlice) >= 2 {
		*odd.Subtype = tempSlice[1]
	}
	if len(tempSlice) == 3 {
		*odd.Typeid, err = strconv.Atoi(tempSlice[2])
		if nil != err {
			return fmt.Errorf("error on ../lib/store/oddids , can not convert int")
		}
	}
	return nil
}

func returnKey(tp, subtp *string, tpid *int) string {
	var key string

	if tp != nil {
		key += *tp + "|"
	}
	if subtp != nil {
		key += *subtp + "|"
	}
	if tpid != nil {
		key += strconv.Itoa(*tpid)
	}
	return key

}
