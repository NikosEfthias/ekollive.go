package conf

//Conf file
var Conf = map[string]string{
	//Don't forget to forward port
	//ssh -L 8080:liveoddstest.betradar.com:1984 <user@server>
	"betradar--url":        "localhost:8080",
	"betradar-prod-url":    "liveoddstest.betradar.com:1984",
	"betradar-bookmakerid": "5098",
	"betradar-key":         "Ea5VF7kFv",
	//"betradar-db":"ekol:7PrqE6suJLkYtzse@tcp(188.166.90.249:3306)/livedata",
	"betradar-db": "root:@tcp(127.0.0.1:3306)/ekollive",
}
