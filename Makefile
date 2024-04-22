.PHONY: build
build:
	go build -o app cmd/main.go 

.PHONY: test
test:
	go test ./... -race -cover