test:
	@go test -v *.go

run:
	@go fmt
	@go fmt ./...
	@go test
	@go install
	@DEBUG=cipher cipher test.txt -r -k -i
	@DEBUG=cipher cipher test.txt -r -k -d -i

.PHONY: test run
