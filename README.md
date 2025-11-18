# recotem-cli

[![CI](https://github.com/codelibs/recotem-cli/workflows/CI/badge.svg)](https://github.com/codelibs/recotem-cli/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/codelibs/recotem-cli)](https://goreportcard.com/report/github.com/codelibs/recotem-cli)

Command-line interface for the Recotem recommendation system API.

## Features

- **Project Management**: Create, list, and delete recommendation projects
- **Training Data**: Upload, list, and delete training datasets
- **Model Configuration**: Configure recommendation models with specific parameters
- **Trained Models**: Create, list, delete, and download trained models
- **Evaluation Config**: Configure model evaluation metrics (NDCG, MAP, Recall, Hit)
- **Split Config**: Configure train/test data splitting strategies
- **Parameter Tuning**: Run hyperparameter optimization jobs
- **Item Metadata**: Manage item metadata for recommendations

## Installation

### Using Go Install

```bash
go install recotem.org/cli/recotem@latest
```

### Using Make (from source)

```bash
git clone https://github.com/codelibs/recotem-cli.git
cd recotem-cli
make install
```

### Download Pre-built Binaries

Download the latest release from the [releases page](https://github.com/codelibs/recotem-cli/releases).

## Usage

### Basic Commands

```bash
# Get help
recotem --help

# Login to get an access token
recotem login

# List projects
recotem project list

# Create a project
recotem project create --name "my-project" --user-column "user_id" --item-column "item_id"

# Upload training data
recotem training-data upload --project-id 1 --file ./data.csv

# List trained models
recotem trained-model list
```

### Available Commands

- `login` - Get an access token
- `project` (alias: `p`) - Project management tasks
- `training-data` (alias: `td`) - Training data operations
- `item-meta-data` (alias: `imd`) - Item metadata management
- `trained-model` (alias: `tm`) - Trained model operations
- `model-configuration` (alias: `mc`) - Model configuration
- `evaluation-config` (alias: `ec`) - Evaluation configuration
- `split-config` (alias: `sc`) - Data splitting configuration
- `parameter-tuning-job` (alias: `ptj`) - Hyperparameter tuning

## Development

### Requirements

- Go 1.23 or later
- Make (optional, for using Makefile commands)

### Building

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Run tests
make test

# Run tests with coverage
make test-coverage

# Run linters
make lint

# Format code
make fmt
```

### Project Structure

```
recotem-cli/
├── main.go              # Application entry point
├── pkg/
│   ├── api/            # API client implementations
│   ├── cfg/            # Configuration management
│   ├── cmd/            # CLI command definitions
│   ├── utils/          # Utility functions
│   └── openapi/        # OpenAPI specification and generated code
├── .github/
│   └── workflows/      # GitHub Actions CI/CD
├── Makefile            # Build automation
└── .golangci.yml       # Linter configuration
```

### OpenAPI Code Generation

#### Setup

```bash
go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
```

#### Update Generated Code

Update `pkg/openapi/recotem.yaml` and then regenerate:

```bash
oapi-codegen -generate types,client pkg/openapi/recotem.yaml > pkg/openapi/recotem.gen.go
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Run benchmarks
go test -bench=. ./...
```

### Code Quality

This project uses:
- **gofmt** for code formatting
- **golangci-lint** for comprehensive linting
- **GitHub Actions** for CI/CD
- Automated testing on multiple Go versions

## Configuration

The CLI stores configuration in `~/.recotem/config.yaml`:

```yaml
url: http://localhost:8000
token: your-access-token
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests (`make test`)
5. Run linters (`make lint`)
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

## License

Apache License 2.0
