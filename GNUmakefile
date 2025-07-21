override TOPDIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))

.PHONY: all
all: build

.PHONY: build
build:
	@cd $(TOPDIR) && go build

.PHONY: clean
clean:
	@rm -f $(TOPDIR)/test

.PHONY: run
run: build
	@$(TOPDIR)/test
