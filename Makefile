cover:
	@go test .\... -coverprofile coverage.out -covermode count
	@go tool cover -func coverage.out

test:
	@go test -race .\... 
	@go vet .\...
	@staticcheck .\...

gitall:
	@echo Commit Message: "$(m)"
	@git add .
	@git commit -m "$(m)"
	@git push -u origin main

gensvg:
	@echo Generating SVG %time%
	@go build -o .\cmd\clockface\clockface.exe .\cmd\clockface
	@.\cmd\clockface\clockface.exe > .\cmd\clockface\clock.svg