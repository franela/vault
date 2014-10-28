test:
	chmod 700 testdata/bob
	go get -v -d -t ./...
	PROJECTDIR=$(PWD) \
	go test ./...

debug:
	go get -v -d -t ./...
	PROJECTDIR=$(PWD) \
	go test -v ./...

fmt: 
	go fmt ./...

install:
	go get github.com/mitchellh/gox
	sudo gox -build-toolchain
build:
	rm -fr build
	gox --output "build/{{.OS}}/{{.Arch}}/vault" ./... 

.PHONY: test debug build install fmt	
