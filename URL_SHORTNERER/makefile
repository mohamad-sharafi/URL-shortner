build:
	@go build -o bin/main CMD/main.go

test:
	@go test -v ./...

run: build
	@./bin/main

clean:
	@rm -rf bin