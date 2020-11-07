test-all:
	@go test -p 1 ./...
test-api:
	@cd test/apiserver && GIN_MODE=test go test -parallel 1