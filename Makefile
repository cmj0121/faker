BIN    := $(subst .go,,$(wildcard cmd/*.go) $(wildcard examples/*.go))
PREFIX := /usr/local/bin

.PHONY: all clean help install

all: $(BIN)	# build all binary

clean:		# clean-up the environment
	rm -f $(BIN)

help:		# show this message
	@printf "Usage: make [OPTION]\n"
	@printf "\n"
	@perl -nle 'print $$& if m{^[\w-]+:.*?#.*$$}' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?#"} {printf "    %-18s %s\n", $$1, $$2}'

install: $(BIN)
	install -m755 cmd/faker $(PREFIX)/faker

GO      := go
GOFMT   := $(GO)fmt -w -s
GOFLAG  := -ldflags="-s -w"
GOTEST  := $(GO) test -race -cover -failfast -timeout 2s
GOBENCH := $(GO) test -bench=. -cover -failfast -benchmem

linter: .benchmark
	$(GOFMT) $(shell find . -name '*.go')

.benchmark:  $(wildcard *.go)
	$(GOTEST)
	$(GOBENCH)
	touch $@

$(BIN): linter

%: %.go
	$(GO) build $(GOFLAG) -o $@ $<
