.PHONY: install

all: install

install:
	CGO_ENABLED=0 go install ./cmd/...
