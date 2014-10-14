test:
	VAULTDIR=$(PWD) \
	go test -v ./...

.PHONY: test
	
