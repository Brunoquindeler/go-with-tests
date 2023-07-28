test:
	@go test .\... -coverprofile coverage.out -covermode count
	@go tool cover -func coverage.out

gitall: 
	git add .
	git commit -m "$(m)"
	git push -u origin main