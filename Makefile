install:
	go mod tidy
	which templ || go install github.com/a-h/templ/cmd/templ@latest
