.PHONY: test
test: vet
	go test -v ./...

.PHONY: fmt
fmt:
	@echo "Running gofmt on all sources ..."
	@gofmt -s -l -w .

.PHONY: fmtcheck
fmtcheck:
	@bash -c "diff -u <(echo -n) <(gofmt -d .)"

.PHONY: vet
vet:
	go vet ./...

.PHONY: examples
examples:
	@echo "Running example code ...."
	@echo "\n--> examples/collection ...\n"
	@go run examples/collection/main.go
	@echo "\n--> examples/compare ...\n"
	@go run examples/compare/main.go
	@echo "\n--> examples/newversion ..\n"
	@go run examples/newversion/main.go
	@echo "\n--> examples/series ...\n"
	@go run examples/series/main.go
