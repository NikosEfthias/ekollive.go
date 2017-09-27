package filters

import (
	"sync"
	"../../../lib"
	"strings"
	"errors"
	filterModel "../../../models/filters"
	"fmt"
	"os"
)

type filters struct {
	sync.Mutex
	filters map[string]string
}

var flt *filters
var testing bool

func Init() {
	flt = new(filters)
	flt.filters = make(map[string]string)
	testing = *lib.Testing

	LoadAll()

}
func Add(filter string) error {
	var filterParts = strings.Split(filter, ",")
	if len(filterParts) != 3 {
		return errors.New("filter format must be matchid,filterkey,filtervalue")
	}
	var mid, key, val = strings.TrimSpace(filterParts[0]), strings.TrimSpace(filterParts[1]), strings.TrimSpace(filterParts[2])
	var filterOk bool
	for reg, _ := range Filter {
		if reg.MatchString(val) {
			filterOk = true
		}
	}
	if !filterOk {
		return errors.New("Unsupported filter : [ " + val + " ]")
	}
	flt.Lock()
	flt.filters[mid] += key + "=" + val + ";"
	flt.Unlock()
	currFilter := &filterModel.Filter{
		Matchid: &mid,
		Filter:  &val,
		Key:     &key,
	}
	filterModel.Model.
		Where(&filterModel.Filter{Matchid: &mid, Key: &key}).
		Assign(currFilter).FirstOrCreate(&filterModel.Filter{})
	return nil
}
func LoadAll() {
	var filtersValid bool
	_ = filtersValid
	flt.Lock()
	defer flt.Unlock()
	var prevFilters = new([]*filterModel.Filter)
	filterModel.Model.Find(prevFilters)
	for _, f := range *prevFilters {

		for r := range Filter {
			if r.MatchString(*f.Filter) {
				filtersValid = true
				break
			}
		}
		if !filtersValid {
			fmt.Printf("invalid filter!! key=<%s> value=<%s>", *f.Key, *f.Filter)
			os.Exit(-1)
		}
		flt.filters[*f.Matchid] += *f.Key + "=" + *f.Filter + ";"
	}
}

func GetFiltersByMatchId(matchid string) (localfilters map[string]string) {
	localfilters = make(map[string]string)
	if filters, ok := flt.filters["*"]; ok {
		//all matches filter
		for _, filter := range strings.Split(filters, ";") {
			fields := strings.Split(filter, "=")
			if len(fields) < 2 {
				continue
			}
			localfilters[fields[0]] = fields[1]
		}
	}
	filters, ok := flt.filters[matchid]
	if !ok {
		return
	}
	for _, filter := range strings.Split(filters, ";") {
		//if there's specific filter override the * filter
		fields := strings.Split(filter, "=")
		if len(fields) < 2 {
			continue
		}
		localfilters[fields[0]] = fields[1]
	}
	return
}
