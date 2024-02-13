NAME=canon
PREFIX=/usr

all:
	go build

clean:
	rm -f ${NAME}

install:
	mkdir -p ${DESTDIR}${PREFIX}/bin
	cp -f ${NAME} ${DESTDIR}${PREFIX}/bin
	chmod 755 ${DESTDIR}${PREFIX}/bin/${NAME}

.PHONY: all clean install

