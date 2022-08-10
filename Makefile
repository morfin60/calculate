BINDIR			:= $(CURDIR)/bin
LDFLAGS			:= -w -s
GOFLAGS			:=
BINNAME			?= calculate

.PHONY: all
all: build

# ------------------------------------------------------------------------------
#  build
.PHONY: build
build: cmd/calculate/*.go
	GO111MODULE=on go build -o '$(BINDIR)'/$(BINNAME) ./cmd/calculate/*.go
