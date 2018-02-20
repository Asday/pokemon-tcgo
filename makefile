all: clean install

clean:
	./clean.sh

install:
	go install ./...
