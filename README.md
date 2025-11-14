# casbin-pgx-adapter

A PostgreSQL adapter for [Casbin](https://casbin.org/) using the [pgx](https://github.com/jackc/pgx) driver.

## Status

⚠️ **Early Development** - This project is in its initial stages. The adapter interfaces are defined, but implementation is still in progress. Contributions and feedback are welcome!

## Overview

Casbin is a powerful and efficient open-source access control library that supports various access control models (ACL, RBAC, ABAC, etc.). This adapter enables Casbin to use PostgreSQL as its policy storage backend via the pgx driver, providing:

- Standard Casbin adapter interface
- Context-aware operations for request-scoped transactions
- Batch operations for efficient bulk policy updates
- Filtered adapter support for policy filtering
- Updatable adapter support for policy modifications

## Installation

```bash
go get github.com/noho-digital/casbin-pgx-adapter
```

## Usage

*Note: Implementation examples will be added as development progresses.*

### Basic Adapter

```go
import (
    "github.com/casbin/casbin/v2"
    pgxadapter "github.com/noho-digital/casbin-pgx-adapter"
)

// TODO
```

## Supported Interfaces

This adapter implements the following Casbin adapter interfaces:

- `Adapter` - Standard Casbin adapter interface
- `ContextAdapter` - Context-aware operations
- `BatchAdapter` - Batch add/remove operations
- `FilteredAdapter` - Policy filtering support
- `UpdatableAdapter` - Policy update operations

## Requirements

- Go 1.24.4 or later
- PostgreSQL database
- Casbin v2

## Related Projects

- [Casbin](https://github.com/casbin/casbin) - The authorization library
- [pgx](https://github.com/jackc/pgx) - PostgreSQL driver and toolkit for Go

