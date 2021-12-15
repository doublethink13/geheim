usage:
	@echo "Usage:"
	@echo "	make usage (default)"
	@echo "	make test_all"

test_all:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out

test_cov:
	@go get github.com/boumenot/gocover-cobertura
	@go get github.com/gorilla/mux
	@echo ""
	@go test ./... -coverprofile=coverage.txt -covermode count
	@gocover-cobertura < coverage.txt > coverage.xml
	@go mod tidy
