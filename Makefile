.PHONY: all clean help

all: linter	# build all binary

clean:		# clean-up the environment

help:		# show this message
	@printf "Usage: make [OPTION]\n"
	@printf "\n"
	@perl -nle 'print $$& if m{^[\w-]+:.*?#.*$$}' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?#"} {printf "    %-18s %s\n", $$1, $$2}'

GO      := go
GOFMT   := $(GO)fmt -w -s
GOFLAG  := -ldflags="-s -w"
GOTEST  := $(GO) test -race -cover -failfast -timeout 2s
GOBENCH := $(GO) test -bench=. -cover -failfast -benchmem

linter:
	$(GOFMT) $(shell find . -name '*.go')
	$(GOTEST)
	$(GOBENCH)

$(BIN): linter

%: %.go
	$(GO) build $(GOFLAG) -o $@ $<
