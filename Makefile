export GOPATH := $(shell go env GOPATH)
export GOBIN := $(PWD)/bin
export PATH := $(GOBIN):$(PATH)

.PHONY: commit-checker
commit-checker:
	go build -o ./tools/commit-msg ./cmd/main.go
	rm -f .git/hooks/commit-msg
	ln -s ../../tools/commit-msg .git/hooks/commit-msg

.PHONY: help
help:
	@printf "\tcommit-checker:\tget commit message checker\n"
