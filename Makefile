default: install

.PHONY: install
install:
	@go install

.PHONY: install-race
install-race:
	@go install -race

.PHONY: run
run:
	@go run main.go

.PHONY: fmt
fmt:
	@goimports -w .

.PHONY: test
test: fmt
	@go test ./... -timeout=10s
