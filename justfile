set shell := ["/bin/sh", "-c"]

# List available just targets.
default:
  @just --list

# Remove build and tmp folders.
[group('build')]
clean:
  rm -rf ./build ./tmp

# Start the server with air.
[group('go')]
dev: vet test
  air

# go vet.
[group('go')]
vet:
  go vet ./...

# go test.
[group('go')]
test:
  go tool gotest ./...

# go build.
[group('build')]
build: vet test clean
  go build -o ./build/goatstack ./cmd/goatstack
