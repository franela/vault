test:
	PROJECTDIR=$(PWD) \
	go test ./...

debug:
	PROJECTDIR=$(PWD) \
	go test -v ./...

fmt: 
	go fmt ./...

.PHONY: test debug
	
