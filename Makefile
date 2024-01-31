build:
	go build -o sls-json-pusher

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o sls-json-pusher-linux-amd64

build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64  go build -o sls-json-pusher-windows-amd64.exe

build-mac-amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64  go build -o sls-json-pusher-darwin-amd64

build-mac-arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64  go build -o sls-json-pusher-darwin-arm64