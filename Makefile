usage:
	@echo "Usage:"
	@echo "	make usage (default)"
	@echo "	make github_install_dependencies"
	@echo "	make test_all"
	@echo "	make test_cov"
	@echo "	make release_binary"

github_install_dependencies:
	@go get -v golang.org/x/tools/gopls@latest && \
		go get -v github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest && \
		go get -v github.com/ramya-rao-a/go-outline@latest && \
		go get -v github.com/go-delve/delve/cmd/dlv@master && \
		go get -v github.com/go-delve/delve/cmd/dlv@latest && \
		go get -v honnef.co/go/tools/cmd/staticcheck@latest

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

release_binary:
	@go install github.com/goreleaser/goreleaser@latest
	@goreleaser check
	@goreleaser release
	@go mod tidy
