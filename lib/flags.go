package lib

import (
	"flag"
)

var (
	J           *int
	Testing     *bool
	Port        *string
	DB          *string
	ProxyURL    *string
	BAR         *bool
	DumpTags    *bool
	Profile     *bool
	DB2         *string
	LockOdds    *bool
	Time        *int
	DisableMeta *bool
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
	BAR = flag.Bool("b", false, "Diplay the limiter and goroutine numbers")
	ProxyURL = flag.String("addr", "localhost:8080", "betradar Proxy url to use to get live data")
	DB = flag.String("DB", "root:@tcp(127.0.0.1:3306)/test", "DB address to use")
	DB2 = flag.String("DB2", "root:@tcp(127.0.0.1:3306)/test", "second db address to use")
	Port = flag.String("PORT", "9090", "Port number to listen on")
	J = flag.Int("j", 50, "Concurrent dbops count")
	Time = flag.Int("timeout", 60, "timeout for goroutines")
	DisableMeta = flag.Bool("disableMeta", false, " disable printing meta")
	LockOdds = flag.Bool("lockodds", false, "Whether to use mutex on odd inserts or not, When db cannot handle a lot of inserts use this.")
	Testing = flag.Bool("testing", false, `controls testing mode
	if the app is running on dry run mode following ops will not take place
		- nothing will be inserted into the db
		- filters will be applied to everything regardless of matchid`)
}
