# recotem-cli

[![CI](https://github.com/codelibs/recotem-cli/workflows/CI/badge.svg)](https://github.com/codelibs/recotem-cli/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/codelibs/recotem-cli)](https://goreportcard.com/report/github.com/codelibs/recotem-cli)

A command-line interface for the [Recotem](https://github.com/codelibs/recotem) recommendation system server (v2 API).

## Features

- **Authentication** -- JWT-based login/logout with automatic token refresh, API key support
- **Project Management** -- Create, list, delete projects and view project summaries
- **Training Data** -- Upload, list, delete, download, and preview training datasets
- **Item Metadata** -- Upload, list, delete, and download item metadata
- **Trained Models** -- Create, list, delete, download models and run recommendations
- **Model Configuration** -- Create, list, update, and delete model configurations
- **Evaluation Config** -- Configure evaluation metrics (NDCG, MAP, Recall, Hit)
- **Split Config** -- Configure train/test data splitting strategies
- **Parameter Tuning** -- Run hyperparameter optimization jobs
- **Deployment Slots** -- Manage model deployment slots for production serving
- **A/B Testing** -- Create, manage, and analyze A/B tests with conversion tracking
- **Retraining Schedules** -- Schedule automatic model retraining with cron expressions
- **API Key Management** -- Create, list, revoke, and delete API keys
- **User Management** -- Create, list, update, activate/deactivate users
- **Output Formats** -- Text, JSON, and YAML output for all commands

## Installation

### Pre-built Binaries

Download the latest release from the [releases page](https://github.com/codelibs/recotem-cli/releases).

### From Source

```bash
git clone https://github.com/codelibs/recotem-cli.git
cd recotem-cli
make build
# Binary is at ./bin/recotem
```

Or install directly to your `$GOPATH/bin`:

```bash
make install
```

**Requires:** Go 1.25 or later.

## Quick Start

### 1. Configure the server URL

Create `~/.recotem/config.yaml`:

```yaml
url: http://localhost:8000
```

### 2. Authenticate

```bash
# Interactive login (prompts for credentials)
recotem login

# Or provide credentials directly
recotem login -u admin -p password

# Or use an API key instead
recotem --api-key "your-api-key" project list
```

### 3. Use the CLI

```bash
# List projects
recotem project list

# Create a project
recotem project create --name "my-project" --user-column user_id --item-column item_id

# Upload training data
recotem training-data upload --project-id 1 --file ./interactions.csv

# Get JSON output
recotem project list -o json

# Check server health
recotem ping
```

## Commands

| Command | Alias | Description |
|---------|-------|-------------|
| `login` | | Authenticate with username/password (JWT) |
| `logout` | | Clear stored tokens |
| `ping` | | Check server connectivity |
| `version` | | Print version information |
| `completion` | | Generate shell completion (bash/zsh/fish/powershell) |
| `project` | `p` | Project management (list, create, delete, summary) |
| `training-data` | `td` | Training data (list, upload, delete, download, preview) |
| `item-meta-data` | `imd` | Item metadata (list, upload, delete, download) |
| `trained-model` | `tm` | Trained models (list, create, delete, download, recommend, sample-recommend, recommend-profile) |
| `model-configuration` | `mc` | Model config (list, create, update, delete) |
| `evaluation-config` | `ec` | Evaluation config (list, create, update, delete) |
| `split-config` | `sc` | Split config (list, create, update, delete) |
| `parameter-tuning-job` | `ptj` | Tuning jobs (list, create, delete) |
| `api-key` | `ak` | API keys (list, create, get, revoke, delete) |
| `deployment-slot` | `ds` | Deployment slots (list, create, get, update, delete) |
| `ab-test` | `ab` | A/B tests (list, create, get, update, delete, start, stop, results, promote-winner) |
| `conversion-event` | `ce` | Conversion events (list, create, batch-create, get) |
| `retraining-schedule` | `rs` | Retraining schedules (list, create, get, update, delete, trigger) |
| `retraining-run` | `rr` | Retraining runs (list, get) |
| `task-log` | `tl` | Task logs (list) |
| `user` | `u` | User management (list, create, get, update, deactivate, activate, reset-password) |

### Global Flags

```
-o, --output string   Output format: text, json, yaml (default "text")
    --api-key string  API key for authentication (overrides stored tokens)
```

### Shell Completion

```bash
# Bash
source <(recotem completion bash)

# Zsh
source <(recotem completion zsh)

# Fish
recotem completion fish | source
```

## Authentication

The CLI supports three authentication methods (in priority order):

1. **API Key** (`--api-key` flag or `api_key` in config) -- Sent as `X-API-Key` header
2. **JWT Tokens** (via `recotem login`) -- Access token with automatic refresh
3. **Legacy Token** (`token` in config) -- Sent as `Token` header (backward compatible)

Tokens are stored in `~/.recotem/config.yaml`:

```yaml
url: http://localhost:8000
access_token: eyJ...
refresh_token: eyJ...
expires_at: 2025-01-01T00:00:00Z
```

The CLI automatically refreshes expired JWT tokens using the stored refresh token.

## Configuration

The configuration file is located at `~/.recotem/config.yaml`.

| Field | Description |
|-------|-------------|
| `url` | Recotem server URL (required) |
| `access_token` | JWT access token (set by `login`) |
| `refresh_token` | JWT refresh token (set by `login`) |
| `expires_at` | Token expiration time (set by `login`) |
| `api_key` | API key for authentication (optional) |
| `token` | Legacy token (backward compatible) |

## Development

### Requirements

- Go 1.25 or later
- [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) v2 (for code generation)
- [golangci-lint](https://golangci-lint.run/) (for linting)

### Building and Testing

```bash
make build          # Build for current platform
make build-all      # Cross-compile for linux/darwin/windows (amd64/arm64)
make test           # Run tests with race detector
make test-coverage  # Run tests with coverage report
make lint           # Run golangci-lint
make check          # Run all checks (fmt, vet, lint, test)
```

### OpenAPI Code Generation

The API client is generated from the OpenAPI schema using [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) v2:

```bash
make install-tools  # Install oapi-codegen
make generate       # Regenerate client from pkg/openapi/recotem.yaml
```

### Project Structure

```
recotem-cli/
├── main.go                 # Entry point
├── Makefile                # Build automation
├── pkg/
│   ├── api/                # API client (one file per resource)
│   ├── cfg/                # Configuration management (JWT, load/save)
│   ├── cmd/                # CLI commands (cobra)
│   ├── openapi/            # OpenAPI schema and generated client
│   └── utils/              # Output formatting, string helpers
├── .github/workflows/      # CI/CD (test on Go 1.25, lint, build)
└── .golangci.yml           # Linter configuration
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/your-feature`)
3. Make your changes
4. Run checks (`make check`)
5. Commit and push
6. Open a Pull Request

## License

Apache License 2.0
