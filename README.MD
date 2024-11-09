
# Godx

**Godx** is a Go module providing various utilities to enhance developer experience. The package includes structured approaches to manage events, commands, and functional operations, leveraging the options pattern and middleware.

## Features

- **Events and Event Bus**: Create events, manage event handlers, and utilize a central event bus for dispatching.
- **Commands and Command Bus**: Define commands with an HTTP-like request structure, apply the options pattern, and use middleware for command handlers.
- **Coalesce Operations**: Utility functions for coalescing values (e.g., returning the first non-nil or non-zero element in a list).
- **Functional Slice Utilities**: Generic, functional utilities for slice operations such as `Map`, `Filter`, and `Reduce`.
- **Assertions**: Provides a set of assertion functions for validation, particularly useful in testing.

## Modules

### 1. `event`
A package for managing events with an HTTP-like structure. It supports:
- **Event Creation**: Events can carry headers, payloads, context, and a timestamp.
- **Event Bus**: Register and dispatch events to multiple handlers.
- **Options Pattern**: Use options to set headers, payload, context, and other event properties.
- **Example Usage**:
  ```go
  e := event.NewEvent("event1", "payload", event.WithHeader("key", "value"))
  eventBus.Publish("eventType", e)
  ```

### 2. `command`
Defines commands with an HTTP-like structure, including:
- **Command Creation**: Commands carry headers, payload, context, and timestamp.
- **Command Handler and Middleware**: Use middleware to extend command handlers, enabling logging, validation, etc.
- **Example Usage**:
  ```go
  cmd := command.NewCommand("cmd1", "payload", command.WithHeader("Authorization", "token"))
  commandHandler.Handle(cmd)
  ```

### 3. `coalesce`
Provides coalescing utilities for finding the first non-zero or non-nil value in a list.
- **Example Functions**:
  - `FirstNonNil`: Returns the first non-nil element in a list of pointers.
  - `FirstNonZero`: Returns the first non-zero element in a list.
  - `Coalesce`: General-purpose coalescing function for non-zero or non-nil values.

### 4. `xslice`
Functional utilities for slices.
- **Example Functions**:
  - `Map`: Applies a transformation to each element in a slice.
  - `Filter`: Filters elements based on a predicate.
  - `Reduce`: Reduces a slice to a single value by accumulating results.
  - `Unique`: Returns a slice with unique elements.

### 5. `assert`
Assertion functions for validation, commonly used in testing scenarios.
- **Example Functions**:
  - `Equal`: Asserts equality between two values.
  - `NotNil`: Asserts that a value is not nil.
  - `True`: Asserts that a condition is true.

## Installation

To install **godx**, use `go get`:
```sh
go get github.com/kindalus/godx
```

## Usage

Import the required packages from `godx` and use the modules as needed.

```go
import (
    "github.com/kindalus/godx/event"
    "github.com/kindalus/godx/command"
    "github.com/kindalus/godx/xslice"
    "github.com/kindalus/godx/coalesce"
    "github.com/kindalus/godx/assert"
)
```

### License
MIT License

---
Enjoy enhanced Go development with **godx**!
