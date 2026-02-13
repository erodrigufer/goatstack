set shell := ["/bin/sh", "-c"]

APP_NAME := "goatstack"
BUILD_TARGET := "./build/" +  APP_NAME

# List available just targets.
default:
  @just --list

# Remove build and tmp folders.
[group('build')]
clean:
  @rm -rf ./build ./tmp && echo "🆗 clean was successfully executed."

# go vet.
[group('go')]
vet:
  @go vet ./... && echo "🆗 vet was successful for {{ APP_NAME }}."

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

# Test the compiled binary by creating a project and running vet on it.
[group('build')]
binary: build
  @mkdir -p ./tmp
  @cd ./tmp && ../{{ BUILD_TARGET }} create --app testproject --module codeberg.org/testproject --daemon testprojectd
  @cd ./tmp && just vet
  @echo "✅ Binary test passed successfully."
