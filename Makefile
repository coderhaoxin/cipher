run:
	@go fmt
	@go fmt ./...
	@go test
	@go install
	@DEBUG=cipher cipher test.txt -r
	@DEBUG=cipher cipher test.txt -r -d

test:
	@go test -v *.go
