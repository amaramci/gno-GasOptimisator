.PHONY: install test

install:
	@cd ../gno && make install

test:
	cd example/hello && gno test
