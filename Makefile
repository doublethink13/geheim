usage:
	@echo "Usage:"
	@echo "	make usage (default)"
	@echo "	make test_all"

test_all:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out
