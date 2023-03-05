.DEFAULT: build

build:
	@go build -o target

.PHONY: install
install:
	@go install .
