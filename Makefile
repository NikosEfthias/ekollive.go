lang.linux:main.go
	GOOS=linux go build -ldflags="-w -s" -o $@ $^
.Phony:deploy
deploy:lang.linux
	rsync -vaurz --progress $< bproxy:./
