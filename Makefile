run:
	@go run cmd/main.go
templ:
	@templ generate
download:
	@go get github.com/labstack/echo/v4
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get github.com/a-h/templ
	@go get github.com/joho/godotenv
tidy:
	@go mod tidy
build:
	@templ generate
	@go build cmd/main.go
	@./main.exe

