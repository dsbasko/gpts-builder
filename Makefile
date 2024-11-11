.PHONY: generate lint install-deps
.SILENT:



# Generate
generate:
	@go run $(CURDIR)/*.go


# Lint
lint:
	@$(CURDIR)/temp/bin/golangci-lint run -c .golangci.yml --path-prefix . --fix



# Dependencies
install-deps:
	@GOBIN=$(CURDIR)/temp/bin go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go mod tidy
