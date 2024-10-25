build-PingFunction:
	GOOS=linux GOARCH=amd64 go build -o main handlers/ping.go

build-HelloFunction:
	GOOS=linux GOARCH=amd64 go build -o main handlers/hello.go

build-AddFunction:
	GOOS=linux GOARCH=amd64 go build -o main handlers/add.go

# General build command for all functions (optional)
build-all: build-PingFunction build-HelloFunction build-AddFunction
