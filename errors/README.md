# `errors` Package

The `errors` package in `go-utils` provides standardized error types for consistent error handling across Go repositories. The package introduces a unified `Error` interface and several custom error types to represent common error cases, making it easy to categorize and manage errors across multiple services or applications.

## Purpose

This package is designed to help developers handle errors in a consistent way across various repositories. While the error types are structured with a naming convention similar to HTTP response errors, they are primarily intended for internal use. Developers can use these error types to check error categories, log issues uniformly, and convert them to appropriate HTTP errors as needed in the calling code.

## Installation

```bash
go get github.com/yourusername/go-utils/errors
```

## Usage

### Defining and Creating Errors

The `errors` package defines an `Error` interface, implemented by various error types like `InternalServerError`, `InputValidationError`, `NotFoundError`, and others:

```go
// Error is an interface for representing execution errors.
type Error interface {
    error
}
```

Each error type implements the `Error` interface and includes a message field. Here’s an example of defining and creating an `InternalServerError`:

```go
package main

import (
    "github.com/kanvstnrn/go-utils/errors"
)

func main() {
    err := errors.NewInternalServerError("an unexpected error occurred")
    if errors.IsInternalServerError(err) {
        // Handle or convert to HTTP 500 if needed
    }
}
```

### Checking Error Types

The package provides helper functions to check for specific error types. These allow the caller to assert and convert errors to HTTP responses or other formats as needed:

```go
if errors.IsConflictError(err) {
    // Handle as HTTP 500 or log accordingly
}
```

### Available Error Types

- **`InternalServerError`**: Represents unexpected or server-side errors.
- **`InputValidationError`**: Used when input validation fails.
- **`NotFoundError`**: Used when a requested resource is not found.
- **`ConflictError`**: Indicates a conflict in the current state of the resource (e.g., duplicate entries).
- **`MethodNotAllowedError`**: Used when an HTTP method is not supported for a requested resource.
- **`NotImplementedError`**: Represents functionality that is not yet implemented.
- **`ServiceUnavailableError`**: Used when a service is temporarily unavailable.
- **`TimeoutError`**: Indicates an operation timed out before completion.

These error types cover a wide range of common error cases, allowing you to handle them appropriately across different applications and services.

Additional error types can be added over time based on project requirements.

### Example

```go
package main

import (
    "fmt"
	
    "github.com/yourusername/go-utils/errors"
)

func main() {
    err := // your error here
	
    if errors.IsNotFoundError(err) {
        fmt.Println("Handle 404: Resource not found")
    } else if errors.IsInternalServerError(err) {
        fmt.Println("Handle 500: Internal server error")
    }
}
```

## Extensibility

Each error type currently includes only a message field, but the types can be extended with additional fields like codes or metadata in the future. As a result, the package can evolve to meet new error handling requirements.
