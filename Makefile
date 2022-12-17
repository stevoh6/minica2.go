#!/usr/bin/env -S gmake -f

ifeq ($(PREFIX),)
    PREFIX := /opt
endif

.DEFAULT:
.PHONY:
all: minica2

minica2.go:
	[ -f $@ ] && rm $@ || true;\
	go build

minica2: minica2.go
	mv $< $@

.PHONY:
install: minica2
	install -d $(DESTDIR)$(PREFIX)/bin/
	install -m 777 $< $(DESTDIR)$(PREFIX)/bin/
