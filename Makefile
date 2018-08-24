parser.linux:main.go
	GOOS=linux go build -ldflags="-w -s" -o $@ $<
.Phony:deploy
deploy:parser.linux
	rsync -vaurz --progress --remove-source-files ./$< bparser:./newpass/parser.linux
