NAME := srhtctl
PREFIX ?= /usr/local
GO?=go
GOFLAGS?=

GOSRC!=find . -name '*.go'
GOSRC+=go.mod go.sum

all: $(NAME)

RM?=rm -f

clean:
	$(RM) srhtctl

$(NAME): $(GOSRC)
	$(GO) build $(GOFLAGS) \
		-o $@

install: $(NAME)
	install -dm755 '$(DESTDIR)$(PREFIX)/bin'
	install -m755 $(NAME) '$(DESTDIR)$(PREFIX)/bin/$(NAME)'

.PHONY: clean all
