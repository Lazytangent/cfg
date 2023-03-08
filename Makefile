.DEFAULT: build

.PHONY: build
build:
	@mkdir -p target
	@go build -o target

.PHONY: install
install:
	@go install .

.PHONY: check
check:
	@go test ./...

.PHONY: tags
tags:
	@ctags -R .

.PHONY: fmt
fmt:
	@go fmt .
