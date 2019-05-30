install:
	@go get -u github.com/pressly/goose/cmd/goose
	@GO111MODULE=on go mod vendor
start:
	@go run *.go

gen:
	@go run generator/*.go
	@goimports -w ./
	@gofmt -s -w ./
