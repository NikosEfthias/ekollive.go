package lib

import (
	"flag"
)

var (
	FiltersFile *string
	J           *int
	Testing     *bool
	Port        *string
	DB          *string
	ProxyURL    *string
	BAR         *bool
	DumpTags    *bool
	Profile     *bool
)

func init() {
	if !flag.Parsed() {
		applyFlags()
		flag.Parse()
	}

}
func applyFlags() {
	Profile = flag.Bool("profile", false, "cpu profiling")
	DumpTags = flag.Bool("dt", false, "Dump raw xml tags to the stdout")
	BAR = flag.Bool("BAR", false, "Diplay the limiter and goroutine numbers")
	ProxyURL = flag.String("addr", "localhost:8080", "betradar Proxy url to use to get live data")
	DB = flag.String("DB", "root:@tcp(127.0.0.1:3306)/test", "DB address to use")
	FiltersFile = flag.String("filtersFile", "filters.csv", "Define a custom filters file")
	Port = flag.String("PORT", "9090", "Port number to listen on")
	J = flag.Int("j", 50, "Concurrent dbops count")
	Testing = flag.Bool("testing", false, `controls testing mode
	if the app is running on dry run mode following ops will not take place
		- nothing will be inserted into the db
		- filters will be applied to everything regardless of matchid`)
}
