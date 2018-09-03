#	-DB "nikos:nikos@tcp(18.184.217.74:3306)/ekollive"\ \
go run main.go\
	-DB "nikos:nikos@tcp(localhost:3306)/ekollive" \
	-Php "http://18.184.217.74/parseme" 	\
	-addr "18.184.193.42:1111" \
	-bar=false -live $@
