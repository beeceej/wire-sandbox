wire-sandbox: generate
	go build -o wire-sandbox ./cmd

generate: clean
	go generate ./...
	go list ./... | xargs wire

test: generate
	go test ./...

clean:
	rm -f ./wire-sandbox
	rm -f ./cmd/wire_gen.go
	rm -rf ./pkg/models


.PHONY: clean generate test
