set shell := ["/bin/sh", "-c"]

APP_NAME := "goatstack"
BUILD_TARGET := "./build/" +  APP_NAME

# List available just targets.
default:
  @just --list

# Remove build and tmp folders.
[group('build')]
clean:
  rm -rf ./build ./tmp

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
  @go build -o {{ BUILD_TARGET }} ./cmd/{{ APP_NAME }} && echo "🆗 {{ APP_NAME }} was successfully built."

# Install on host.
[group('install')]
install : build
  @echo "Installing {{ APP_NAME }} on host..."
  @cp {{ BUILD_TARGET }} ~/bin && echo "✅ {{ APP_NAME }} was successfully installed on host."
