
build:
	go build -v ./...

test:
	go test -race -v ./...

coverage:
	go test -v -coverprofile=.out/cover.out -covermode=atomic ./...
	go tool cover -html=.out/cover.out -o .out/cover.html
