package lib

import (
	"flag"
	"strconv"
)

func IsTesting() bool {
	var testing bool
	testing, _ = strconv.ParseBool(flag.Lookup("testing").Value.String())
	return testing
}
