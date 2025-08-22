# vibecheck build and release automation

# Default recipe
default: build

# Build for current platform
build:
    go build -o vibecheck .

# Clean build artifacts
clean:
    rm -f vibecheck
    rm -rf dist/

# Test the application
test:
    go test ./...

# Run linting
lint:
    go vet ./...
    go fmt ./...

# Install locally
install:
    go install .

# Cross-compile for all platforms
build-all: clean
    mkdir -p dist
    # Linux
    GOOS=linux GOARCH=amd64 go build -o dist/vibecheck-linux-amd64 .
    GOOS=linux GOARCH=arm64 go build -o dist/vibecheck-linux-arm64 .
    # macOS
    GOOS=darwin GOARCH=amd64 go build -o dist/vibecheck-darwin-amd64 .
    GOOS=darwin GOARCH=arm64 go build -o dist/vibecheck-darwin-arm64 .
    # Windows
    GOOS=windows GOARCH=amd64 go build -o dist/vibecheck-windows-amd64.exe .
    GOOS=windows GOARCH=arm64 go build -o dist/vibecheck-windows-arm64.exe .

# Package binaries
package: build-all
    cd dist && \
    tar -czf vibecheck-linux-amd64.tar.gz vibecheck-linux-amd64 && \
    tar -czf vibecheck-linux-arm64.tar.gz vibecheck-linux-arm64 && \
    tar -czf vibecheck-darwin-amd64.tar.gz vibecheck-darwin-amd64 && \
    tar -czf vibecheck-darwin-arm64.tar.gz vibecheck-darwin-arm64 && \
    zip vibecheck-windows-amd64.zip vibecheck-windows-amd64.exe && \
    zip vibecheck-windows-arm64.zip vibecheck-windows-arm64.exe

# Create checksums
checksums: package
    cd dist && sha256sum *.tar.gz *.zip > checksums.txt

# Create a new release (requires version tag)
release version: checksums
    #!/usr/bin/env bash
    set -euo pipefail
    
    # Check if version starts with 'v'
    if [[ ! "{{version}}" =~ ^v ]]; then
        echo "Version must start with 'v' (e.g., v1.0.0)"
        exit 1
    fi
    
    # Check if tag already exists
    if git tag -l | grep -q "^{{version}}$"; then
        echo "Tag {{version}} already exists"
        exit 1
    fi
    
    # Create and push tag
    git tag {{version}}
    git push origin {{version}}
    
    # Create GitHub release
    gh release create {{version}} \
        --title "vibecheck {{version}}" \
        --generate-notes \
        dist/*.tar.gz \
        dist/*.zip \
        dist/checksums.txt

# Create a draft release (for testing)
draft-release version: checksums
    #!/usr/bin/env bash
    set -euo pipefail
    
    if [[ ! "{{version}}" =~ ^v ]]; then
        echo "Version must start with 'v' (e.g., v1.0.0)"
        exit 1
    fi
    
    if git tag -l | grep -q "^{{version}}$"; then
        echo "Tag {{version}} already exists"
        exit 1
    fi
    
    git tag {{version}}
    git push origin {{version}}
    
    gh release create {{version}} \
        --title "vibecheck {{version}}" \
        --generate-notes \
        --draft \
        dist/*.tar.gz \
        dist/*.zip \
        dist/checksums.txt

# Quick local development test
dev-test:
    go run . --vibes 80 --ai 40 --agent claude --artifact test-feature
    go run . --vibes 80 --ai 40 --agent claude --artifact test-feature --output json
    go run . --vibes 80 --ai 40 --agent claude --artifact test-feature --output github-markdown

# Test AMP thread
test-amp:
    go run . --vibes 90 --ai 85 --agent amp --artifact T-007dd158-76eb-4daf-967f-8c6d683cde7b --output github-markdown

# Show help
help:
    @just --list