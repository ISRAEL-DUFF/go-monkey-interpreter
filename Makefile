build:
	@cd cmd/interpreter; go build -o ../../bin/interpreter

run: build
	@./bin/interpreter

test: 
	@go test ./... -v