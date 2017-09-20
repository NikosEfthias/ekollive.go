package filters

import (
	"sync"
	"flag"
	"os"
	"encoding/csv"
	"regexp"
	"fmt"
	"../../../lib"
	"strings"
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
	flt.filtersFile = flag.Lookup("filtersFile").Value.String()
	testing = lib.IsTesting()

	LoadAll()

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
		panic("Cannot open filters file. Create at least an empty one")
	}
	defer f.Close()
	rdr := csv.NewReader(f)
	for {
		line, err := rdr.Read()
		if nil != err {
			break
		}
		var filter, value, mid = strings.TrimSpace(line[k]), strings.TrimSpace(line[v]), strings.TrimSpace(line[matchid])
		for _, r := range ValidFilters {
			rgx := regexp.MustCompile(r)
			if rgx.MatchString(value) {
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
		localfilters[fields[0]] = fields[1]
	}
	return
}
