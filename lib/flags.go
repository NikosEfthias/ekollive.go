package lib

import (
	"flag"
	"log"
)

var (
	FiltersFile         *string
	J                   *int
	Testing             *bool
	Port                *string
	DB                  *string
	Key                 *string
	BetradarURL         *string
	BetradarBookmakerId *string
)

func init() {
	if !flag.Parsed() {
		applyFlags()
		flag.Parse()
		if *Key == "" || *BetradarBookmakerId == "" {
			log.Fatalln("No betradar key or bookmakerid is present")
		}
	}

}
func applyFlags() {
	BetradarBookmakerId = flag.String("BookmakerID", "", "Set betradar bookmaker id")
	BetradarURL = flag.String("BetradarURL", "liveoddstest.betradar.com:1984", "betradar url to use to get live data")
	Key = flag.String("KEY", "", "Betradar key to use")
	DB = flag.String("DB", "root:@tcp(127.0.0.1:3306)/test", "DB address to use")
	FiltersFile = flag.String("filtersFile", "filters.csv", "Define a custom filters file")
	Port = flag.String("PORT", "9090", "Port number to listen on")
	J = flag.Int("j", 50, "Concurrent dbops count")
	Testing = flag.Bool("testing", false, `controls testing mode
	if the app is running on dry run mode following ops will not take place
		- nothing will be inserted into the db
		- filters will be applied to everything regardless of matchid`)
}
