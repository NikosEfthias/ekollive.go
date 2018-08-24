package lib

import (
	"flag"
	"fmt"
	"os"
)

var (
	J            *int
	Port         *string
	Lang_api_url *string
	DB           *string
	PhpPostADDR  *string
	Profile      *bool
	Bar          *bool
	IsLive       *bool
	print        *bool
	Key          *string
	Pass         *string
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
			"\nlang api addr=>", *Lang_api_url,
			"\ndb=>", *DB,
			"\nbar=>", *Bar,
			"\nphppostaddr=>", *PhpPostADDR,
			"\nProfile=>", *Profile,
			"\nKey=>", *Key,
			"\nPass=>", *Pass,
		)
		os.Exit(0)
	}

}
func applyFlags() {
	print = flag.Bool("print", false, "print current flags and die")
	IsLive = flag.Bool("live", false, "live or prematches")
	Bar = flag.Bool("bar", true, "activate good/bad ratio display mode")
	Profile = flag.Bool("prof", false, "activate cpu profiling")
	Lang_api_url = flag.String("addr", "localhost:8088", "betconstruct lang api url to use to get live data")
	// Lang_api_url = flag.String("addr", "translations-stream.betconstruct.com:8088", "betconstruct lang api url to use to get live data")
	DB = flag.String("DB", "root:root@tcp(localhost:3306)/test", "DB address to use")
	Port = flag.String("PORT", "9090", "Port number to listen on")
	Key = flag.String("key", "", "betconstruct key to use")
	Pass = flag.String("pass", "", "betconstruct password")
	J = flag.Int("j", 100, "Concurrent dbops count")
	PhpPostADDR = flag.String("Php", "http://localhost/parseme", "php address to use for sending data")
}
