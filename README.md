# Railgun

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.24-blue)](https://go.dev/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-12--16-blue)](https://www.postgresql.org/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![Status](https://img.shields.io/badge/status-in%20development-orange)](https://github.com/benidevo/railgun)

High-performance streaming data pipeline for PostgreSQL. Load any format, transform on-the-fly, maintain constant memory usage.

**⚠️ Early development. Nothing here yet.**

## What I'm Building

See [Project Specification](docs/project-specification.md) for detailed requirements and design decisions.

This will be a Go library for high-performance data loading into PostgreSQL - targeting 50k+ rows/second with constant memory usage regardless of file size.

## Getting Started

```bash
# Clone the repository
git clone https://github.com/benidevo/railgun.git
cd railgun

# Install dependencies
go mod download

# Run tests (once implemented)
go test ./...
```

## License

MIT License. See [LICENSE](LICENSE) for details.
