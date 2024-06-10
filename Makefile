.PHONY: build clean deploy

build:
	env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o bin/stop-rds stop-rds/main.go
	env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o bin/start-rds start-rds/main.go
	
clean:
	rm -rf ./bin

deploy: clean build
	sls deploy
