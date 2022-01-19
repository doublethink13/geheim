usage:
	@echo "Usage:"
	@echo "	make usage (default)"
	@echo "	make lint"
	@echo "	make build_and_test"
	@echo "	make test_all"
	@echo "	make test_cov"
	@echo "	make release_binary_github"

lint:
	@golangci-lint linters
	@golangci-lint run ./...

build_and_test:
	@go build .
	@./geheim
	@rm -rf geheim

test_all:
	@go test -coverprofile=coverage.out ./... -v
	@go tool cover -func=coverage.out

test_cov:
	@go get github.com/boumenot/gocover-cobertura
	@go get github.com/gorilla/mux
	@echo ""
	@go test ./... -coverprofile=coverage.txt -covermode count
	@gocover-cobertura < coverage.txt > coverage.xml
	@go mod tidy

benchmarks:
	@GEHEIM_LOG_LEVEL=0 && go test -bench=. -run=".*_benchmark_.*"
		
release_binary_github:
	@go install github.com/goreleaser/goreleaser@latest
	@goreleaser check
	@goreleaser release
