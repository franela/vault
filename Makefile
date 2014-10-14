test:
	VAULTDIR=$(PWD) \
	go test ./...

debug:
	VAULTDIR=$(PWD) \
	go test -v ./...

.PHONY: test debug
	
