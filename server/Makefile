.PHONY: all
all:
	env GOOS=darwin GOARCH=amd64 go build -tags=jsoniter -o ./bin/macos/  . &&\
	cp -R ./app ./bin/macos/ &&\
	env GOOS=windows GOARCH=amd64 go build -tags=jsoniter -o ./bin/windows/ . && \
	cp -R ./app ./bin/windows/ &&\
	env GOOS=linux GOARCH=amd64 go build -tags=jsoniter -o ./bin/ubuntu/ .  && \
	cp -R ./app ./bin/ubuntu/