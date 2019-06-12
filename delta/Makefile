install:
	@go get -u github.com/pressly/goose/cmd/goose
	@GO111MODULE=on go mod vendor
start:
	@go run cmd/*.go

gen:
	@go run pkg/generator/cmd/*.go
	@goimports -w ./
	@gofmt -s -w ./
