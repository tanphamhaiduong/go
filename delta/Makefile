start:
	@go run *.go

gen:
	@go run generator/*.go
	@goimports -w ./
	@gofmt -s -w ./