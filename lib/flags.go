package lib

import "flag"

func ApplyFlags() {
	flag.String("filtersFile", "filters.csv", "Define a custom filters file")
	flag.Int("j", 50, "Concurrent dbops count")
	flag.Bool("testing",false,`controls testing mode
	if the app is running on dry run mode following ops will not take place
		- nothing will be inserted into the db
		- filters will be applied to everything regardless of matchid`)
}
