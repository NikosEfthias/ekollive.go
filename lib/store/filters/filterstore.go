package filters

import (
	"sync"
	"os"
	"encoding/csv"
	"fmt"
	"../../../lib"
	"strings"
	"errors"
)

type filters struct {
	sync.Mutex
	filtersFile string
	filters     map[string]string
}

var flt *filters
var testing bool

func Init() {
	flt = new(filters)
	flt.filters = make(map[string]string)
	flt.filtersFile = *lib.FiltersFile
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
	f, err := os.OpenFile(flt.filtersFile, os.O_APPEND|os.O_WRONLY, 0644)
	if nil != err {
		return err
	}
	defer f.Close()
	f.WriteString("\n" + strings.Join([]string{mid, key, val}, ","))
	return nil
}
func LoadAll() {
	const (
		matchid = iota
		k
		v
	)
	var filtersValid bool
	flt.Lock()
	defer flt.Unlock()
	f, err := os.Open(flt.filtersFile)
	if nil != err {
		f, err = os.Create(flt.filtersFile)
		if nil != err {
			panic(err)
		}
		//panic("Cannot open filters file. Create at least an empty one")
	}
	defer f.Close()
	rdr := csv.NewReader(f)
	for {
		line, err := rdr.Read()
		if nil != err {
			break
		}
		var filter, value, mid = strings.TrimSpace(line[k]), strings.TrimSpace(line[v]), strings.TrimSpace(line[matchid])
		for r, _ := range Filter {
			if r.MatchString(value) {
				filtersValid = true
				break
			}
		}
		if !filtersValid {
			fmt.Printf("invalid filter!! key=<%s> value=<%s>", filter, value)
			os.Exit(-1)
		}
		flt.filters[mid] += filter + "=" + value + ";"
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
