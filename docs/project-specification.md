# Railgun - Project Specification

---

## Overview

Railgun is a high-performance Go library for streaming data into PostgreSQL at 10-50x the speed of traditional row-by-row operations.
Starting with CSV support, Railgun provides a clean, extensible architecture for adding additional formats while maintaining consistent performance and API design.

## Problem Statement

Engineers loading data files into PostgreSQL face an impossible choice:

- **PostgreSQL COPY**: Blazing fast (100k+ rows/sec) but inflexible. No transformations, all-or-nothing error handling
- **Row-by-row INSERT**: Flexible but painfully slow (1-5k rows/sec)

This results in:

- 10GB CSV files taking 2+ hours to import
- Custom COPY implementations for every project
- Memory crashes from naive implementations
- Repeated work for each new format

## Solution

Railgun provides a streaming pipeline that combines the speed of COPY with the flexibility of row-by-row operations:

```
[Source] → [Transform] → [Buffer] → [PostgreSQL COPY]
```

- **Stream processing**: Constant memory usage regardless of file size
- **Transformation pipeline**: Modify data in-flight
- **Extensible architecture**: Add new formats without changing core
- **Production ready**: Comprehensive error handling and monitoring

## Goals

### v0.1

- Stream CSV files without memory issues
- Achieve 20-50k rows/second throughput
- Support simple transformations
- Provide progress monitoring
- Establish extensible architecture

### Future Versions

- JSON Lines and TSV support
- Parallel processing and remote sources
- Plugin ecosystem for custom formats

## Architecture

### Design Principles

1. **Format-Agnostic Core**: Pipeline works identically regardless of input format
2. **Interface-Driven**: Clean contracts enable format additions without core changes
3. **Streaming First**: Process data in chunks, never load entire files
4. **Zero-Copy Operations**: Minimize data copying between stages

### Component Architecture

```plaintext
User API
    ├── Fluent Builder API
    ├── Execution Controller
    └── Progress Monitoring

Core Pipeline
    ├── Source Interface (format abstraction)
    ├── Transformation Engine
    ├── Buffer Manager
    └── PostgreSQL Writer

Format Implementations
    ├── CSV
    ├── JSON Lines (planned)
    └── TSV (planned)
```

### Data Flow

1. **Source** reads input file and emits row stream
2. **Transformer** applies user-defined modifications
3. **Buffer** accumulates rows for batch processing
4. **Writer** executes PostgreSQL COPY protocol

Each stage operates independently, connected through Go channels for natural backpressure handling.

## Technical Specifications

### Source Interface

The Source interface abstracts format-specific parsing:

- **Read**: Returns channel of rows for streaming
- **Schema**: Provides column information
- **Close**: Ensures resource cleanup

New formats implement this interface without modifying the core pipeline.

### Row Model

Rows are represented as maps for flexibility:

- Key-value pairs (column name to value)
- Type-agnostic storage
- Metadata support (row numbers, errors)

### Transformation Pipeline

Transformations are composable functions:

- Type conversions (string to int, date parsing)
- String operations (trim, case conversion)
- Field operations (rename, reorder)
- Custom functions

Transformations work identically across all formats.

### Buffer Management

Intelligent batching for optimal performance:

- Size-based flushing (row count)
- Memory-based flushing (byte limit)
- Adaptive sizing based on throughput

Implementation uses ring buffers for efficiency.

### PostgreSQL Writer

High-performance insertion using:

- COPY protocol for maximum speed
- Transaction management
- Connection pooling
- Automatic error recovery

## API Design

### Basic Usage

```
railgun.New().
    Source(railgun.CSV("data.csv")).
    Target(db, "users").
    Execute()
```

### With Transformations

```
railgun.New().
    Source(railgun.CSV("data.csv")).
    Transform(transformer).
    Target(db, "users").
    Execute()
```

### With Progress Monitoring

```
railgun.New().
    Source(railgun.CSV("large.csv")).
    WithProgress(callback).
    Target(db, "users").
    Execute()
```

## Performance Targets

### Throughput

- Minimum: 20,000 rows/second
- Target: 50,000 rows/second
- Benchmark: 10x faster than row-by-row

### Resource Usage

- Memory: O(1) constant usage
- Target: <100MB for any file size
- CPU: Efficient single-core, future parallel support

### Benchmarks

Comprehensive benchmark suite comparing:

- Railgun vs row-by-row
- Railgun vs raw COPY
- Different file sizes (10K to 10M rows)

## Format Support

### v0.1: CSV

- RFC 4180 compliant
- Configurable delimiters
- Header support
- UTF-8 encoding

## Error Handling

### Error Modes

1. **Skip**: Log errors and continue
2. **Fail**: Stop on first error
3. **Collect**: Gather all errors, report at end

### Error Context

- Row number
- Column name
- Error type
- Original value
- Suggested fixes

## Testing Strategy

### Test Coverage

- Unit tests: Component isolation
- Integration tests: End-to-end pipeline
- Performance tests: Benchmark suite
- Format tests: Format-specific edge cases

### Coverage Goals

- Core pipeline: >90%
- Format implementations: >85%
- Overall: >80%

## Contributing

### Adding a Format

1. Implement the Source interface
2. Add configuration options
3. Write comprehensive tests
4. Document usage
5. Submit pull request

### Development Setup

- Go 1.24+ required
- PostgreSQL 12+ for testing
- Standard Go toolchain
- No special dependencies

## Dependencies

### Core

- `database/sql` Database interface
- `pgx` or `lib/pq` PostgreSQL driver
- Standard library for everything else

### Development

- Testing framework
- Benchmark tools
- Docker for integration tests

## Success Metrics

### Technical

- 50k rows/second achieved
- Memory usage constant
- Zero memory leaks
- Clean architecture for extensions

### Community

- 100+ GitHub stars (month 1)
- 5+ contributors (quarter 1)
- Production usage reports
- Format contributions

## Comparison

| Solution | Pros | Cons |
|----------|------|------|
| pgx CopyFrom | Fast, built-in | No transformations, low-level API |
| Custom scripts | Flexible | Not reusable, maintenance burden |
| Railgun | Fast, transformable, extensible | New project, CSV only initially |

## FAQ

**Q: Why not support all formats immediately?**
A: Starting with CSV allows us to perfect the core pipeline and API before expanding.

**Q: How does this compare to ETL tools?**
A: Railgun is a library, not a framework. It solves one problem well: fast data loading.

**Q: Can I use this in production?**
A: v0.1 will be production-ready for CSV imports with comprehensive testing.

**Q: Will you support other databases?**
A: PostgreSQL is the initial focus. The architecture could support others in the future.

## License

MIT License. Use freely in personal and commercial projects.
