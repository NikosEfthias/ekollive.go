package lib

import (
	"flag"
	"fmt"
	"os"
)

var (
	J           *int
	Port        *string
	ProxyURL    *string
	DB          *string
	PhpPostADDR *string
	Profile     *bool
	Bar         *bool
	IsLive      *bool
	print       *bool
)

func init() {
	if !flag.Parsed() {
		applyFlags()
		flag.Parse()
	}
	if *print {
		fmt.Println(
			"j=>", *J,
			"\nislive=>", *IsLive,
			"\nport=>", *Port,
			"\nproxyurl=>", *ProxyURL,
			"\ndb=>", *DB,
			"\nbar=>", *Bar,
			"\nphppostaddr=>", *PhpPostADDR,
			"\nProfile=>", *Profile,
		)
		os.Exit(0)
	}

}
func applyFlags() {
	print = flag.Bool("print", false, "print current flags and die")
	IsLive = flag.Bool("live", false, "live or prematches")
	Bar = flag.Bool("bar", true, "activate good/bad ratio display mode")
	Profile = flag.Bool("prof", false, "activate cpu profiling")
	ProxyURL = flag.String("addr", "localhost:1111", "betconstruct Proxy url to use to get live data")
	DB = flag.String("DB", "root:root@tcp(localhost:3306)/test", "DB address to use")
	Port = flag.String("PORT", "9090", "Port number to listen on")
	J = flag.Int("j", 100, "Concurrent dbops count")
	PhpPostADDR = flag.String("Php", "http://localhost/parseme", "php address to use for sending data")
}
