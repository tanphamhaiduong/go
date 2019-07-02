install:
	@go get -u github.com/pressly/goose/cmd/goose
	@go get -u github.com/vektra/mockery/.../
	@GO111MODULE=on go mod vendor
start:
	@go run cmd/*.go

gen:
	@protoc -I ./pb ./pb/message.proto --go_out=plugins=grpc:./pb
	@go run pkg/generator/cmd/*.go
	@goimports -w ./
	@gofmt -s -w ./
