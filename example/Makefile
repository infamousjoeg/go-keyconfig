run:
	go run example.go

build:
	GOOS=darwin GOARCH=amd64 go build -o bin/keyconfig-example-darwin-amd64 example.go
	GOOS=darwin GOARCH=arm64 go build -o bin/keyconfig-example-darwin-arm64 example.go
	GOOS=linux GOARCH=amd64 go build -o bin/keyconfig-example-linux-amd64 example.go
	GOOS=windows GOARCH=amd64 go build -o bin/keyconfig-example-windows-amd64.exe example.go