.PHONY: test
test:
	go test -v -race ./...

.PHONY: benchmark
benchmark:
	go test -run=none -bench=. -benchmem ./...

.PHONY: coverage
coverage:
	go test -covermode=atomic -coverprofile=cover.out ./...
	go tool cover -html=ccover.out -o cover.html

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	go vet ./...
