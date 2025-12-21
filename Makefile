PREFIX ?= /usr/local
BINDIR ?= $(PREFIX)/bin

INSTALL ?= install
INSTALL_PROGRAM ?= $(INSTALL) -m 755
INSTALL_DIR ?= $(INSTALL) -d -m 755

# For staged/package installations
DESTDIR ?=

BINARY_NAME = janus

.PHONY: all build install uninstall clean

all: build

build:
	go build -o $(BINARY_NAME) cmd/main.go

install: build
	sudo $(INSTALL_DIR) $(DESTDIR)$(BINDIR)
	sudo $(INSTALL_PROGRAM) $(BINARY_NAME) $(DESTDIR)$(BINDIR)/$(BINARY_NAME)

uninstall:
	sudo rm -f $(DESTDIR)$(BINDIR)/$(BINARY_NAME)

clean:
	sudo rm -f $(BINARY_NAME)
