test:
	PROJECTDIR=$(PWD) \
	go test ./...

debug:
	PROJECTDIR=$(PWD) \
	go test -v ./...

.PHONY: test debug
	
