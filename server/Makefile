.PHONY: all test

all: test

main: main.go
	go build $<

test: main
	python3 test/api_test.py
