build:
	go build -o sls-json-pusher

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o sls-json-pusher