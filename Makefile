# test-all:
	# @go test -p 1 ./...
test-model:
	cd $(PWD)/internal/model && go test
test-unit-api: test-model
	cd $(PWD)/internal/apiserver && go test ./... -tags mock
test-integration-api: test-unit-api
	@cd $(PWD)/test/apiserver && GIN_MODE=test go test -parallel 1
gen-pbs:
	cd pb && protoc --go_out=paths=source_relative:../internal/pb \
		--go-grpc_out=paths=source_relative:../internal/pb \
		*.proto 