package lib

import "flag"

var (
	FiltersFile *string
	J           *int
	Testing     *bool
)

func init() {
	if !flag.IsParsed() {
		flag.Parse()
	}

}
func applyFlags() {
	FiltersFile = flag.String("filtersFile", "filters.csv", "Define a custom filters file")
	J = flag.Int("j", 50, "Concurrent dbops count")
	Testing = flag.Bool("testing", false, `controls testing mode
	if the app is running on dry run mode following ops will not take place
		- nothing will be inserted into the db
		- filters will be applied to everything regardless of matchid`)
}
