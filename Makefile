.PHONY: build install generate generate_schema generate_sdks clean test build_nodejs build_python build_dotnet

PROVIDER_NAME := danubedata
VERSION := 0.1.0
WORKING_DIR := $(shell pwd)
TFGEN := $(WORKING_DIR)/bin/pulumi-tfgen-$(PROVIDER_NAME)
PROVIDER := $(WORKING_DIR)/bin/pulumi-resource-$(PROVIDER_NAME)

OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH := $(shell uname -m)
ifeq ($(ARCH),x86_64)
	ARCH := amd64
endif
ifeq ($(ARCH),aarch64)
	ARCH := arm64
endif

# Build the provider binary
build: provider

provider:
	cd provider && go build -o $(PROVIDER) -ldflags "-X main.version=$(VERSION)" ./cmd/pulumi-resource-$(PROVIDER_NAME)

# Build the tfgen binary
tfgen:
	cd provider && go build -o $(TFGEN) ./cmd/pulumi-tfgen-$(PROVIDER_NAME)

# Generate the Pulumi schema
generate_schema: tfgen
	$(TFGEN) schema --out provider/cmd/pulumi-resource-$(PROVIDER_NAME)
	@echo "Schema generated successfully"

# Generate all SDKs
generate_sdks: generate_schema
	$(TFGEN) nodejs --out sdk/nodejs
	$(TFGEN) python --out sdk/python
	$(TFGEN) go --out sdk/go
	$(TFGEN) dotnet --out sdk/dotnet
	@echo "SDKs generated successfully"

# Generate everything (schema + SDKs)
generate: generate_sdks

# Install the provider locally
install: provider
	mkdir -p ~/.pulumi/plugins/resource-$(PROVIDER_NAME)-v$(VERSION)
	cp $(PROVIDER) ~/.pulumi/plugins/resource-$(PROVIDER_NAME)-v$(VERSION)/

# Install for development
install_dev: provider
	mkdir -p ~/.pulumi/plugins
	cp $(PROVIDER) ~/.pulumi/plugins/pulumi-resource-$(PROVIDER_NAME)

# Build and install NodeJS SDK
build_nodejs: generate_sdks
	cd sdk/nodejs && \
		sed -i.bak 's/$${VERSION}/$(VERSION)/g' package.json && \
		rm -f package.json.bak && \
		npm install && \
		npm run build

# Build and install Python SDK
build_python: generate_sdks
	cd sdk/python && \
		PULUMI_VERSION=$(VERSION) python3 -m venv venv && \
		. venv/bin/activate && \
		pip install build && \
		PULUMI_VERSION=$(VERSION) python -m build

# Build .NET SDK
build_dotnet: generate_sdks
	cd sdk/dotnet && \
		sed -i.bak 's/<Version>0.0.0<\/Version>/<Version>$(VERSION)<\/Version>/g' *.csproj && \
		rm -f *.csproj.bak && \
		echo "$(VERSION)" > version.txt && \
		dotnet build

# Run tests
test:
	cd provider && go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/
	rm -rf sdk/

# Tidy up Go modules
tidy:
	cd provider && go mod tidy

# Development helper - build and install everything
dev: build install_dev
	@echo "Provider installed for development"

# Show help
help:
	@echo "Available targets:"
	@echo "  build          - Build the provider binary"
	@echo "  tfgen          - Build the tfgen binary"
	@echo "  generate       - Generate schema and all SDKs"
	@echo "  generate_schema - Generate only the Pulumi schema"
	@echo "  generate_sdks  - Generate all language SDKs"
	@echo "  install        - Install provider to Pulumi plugins directory"
	@echo "  install_dev    - Install provider for local development"
	@echo "  build_nodejs   - Build NodeJS SDK"
	@echo "  build_python   - Build Python SDK"
	@echo "  build_dotnet   - Build .NET SDK"
	@echo "  test           - Run provider tests"
	@echo "  clean          - Remove build artifacts"
	@echo "  tidy           - Tidy Go modules"
	@echo "  dev            - Build and install for development"
