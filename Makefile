#!make

SOURCE_FILES?=./...

export GO111MODULE := on

.PHONY: test
test:
	go test -v -failfast -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.out $(SOURCE_FILES) -run -timeout=5m

.PHONY: cover
cover: test
	go tool cover -html=coverage.out

.PHONY: fmt
fmt:
	find . -name '*.go' -not -wholename './vendor/*' | \
		while read -r file; \
			do gofmt -w -s "$$file"; \
			goimports -w "$$file"; \
		done

.PHONY: build
build:
	go build -o drm .

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
