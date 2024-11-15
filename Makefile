
build:
	go build -v ./...

test:
	go test -race -json -v ./... | go run  github.com/mfridman/tparse@latest -all

coverage:
	go test -v -coverprofile=.out/cover.out -covermode=atomic ./...
	go tool cover -html=.out/cover.out -o .out/cover.html

release: test
	git tag -a "v`cat ./VERSION`" -m "Release version `cat ./VERSION`"
	git push origin v`cat ./VERSION`
