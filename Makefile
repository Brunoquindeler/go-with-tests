cover:
	@go test .\... -coverprofile coverage.out -covermode count
	@go tool cover -func coverage.out

test:
	@go test -race .\... 
	@go vet .\...
	@staticcheck .\...
	@go fmt .\...

gitall: 
	git add .
	git commit -m "$(m)"
	git push -u origin main