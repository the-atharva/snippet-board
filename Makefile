run:
	go run $(shell find cmd/web -name '*.go' ! -name '*_test.go')

test:
	go test -v ./...
