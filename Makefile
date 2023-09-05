.DEFAULT: help
.PHONY: help
help:
	@echo "Goals:"
	@echo "  build:	Build an executable into the 'target' directory."
	@echo "  install:	Install an executable into $(go env GOPATH)/bin."
	@echo "  check:	Run tests"
	@echo "  tags:		Generate tags"
	@echo "  fmt:		Run 'go fmt' and 'goimports'"

.PHONY: build
build:
	@mkdir -p target
	@go build -o target -tags static

.PHONY: install
install:
	@go -tags static install .

.PHONY: check
check:
	@go -tags static test ./...

.PHONY: tags
tags:
	@ctags -R .

.PHONY: fmt
fmt:
	@go fmt ./...
	@goimports -w .
