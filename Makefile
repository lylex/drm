#!make

GO_BIN ?= go
GOFMT ?= gofmt
GOIMPORTS ?= goimports
GOLINT ?= golint
SOURCE_FILES ?= ./...

export GO111MODULE := on

.PHONY: test
test:
	$(GO_BIN) test -v -failfast -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.out $(SOURCE_FILES) -run -timeout=5m

.PHONY: cover
cover: test
	$(GO_BIN) tool cover -html=coverage.out

.PHONY: lint
lint:
	$(GOLINT) ./...

.PHONY: fmt
fmt:
	find . -name '*.go' -not -wholename './vendor/*' | \
		while read -r file; do \
			$(GOFMT) -w -s "$$file"; \
			$(GOIMPORTS) -w "$$file"; \
		done

.PHONY: build
build:
	$(GO_BIN) build -o drm .

.PHONY: dist
dist:
	goreleaser release --skip-publish --snapshot --rm-dist

.PHONY: todo
todo:
	@grep \
		--exclude-dir=vendor \
		--exclude=Makefile \
		--text \
		--color \
		-nRo -E ' TODO.*' .

.DEFAULT_GOAL := build
