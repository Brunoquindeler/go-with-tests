test:
	@go test .\... -coverprofile coverage.out -covermode count
	@go tool cover -func coverage.out