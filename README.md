# Learn Go with Tests

A practice repository for learning Go by following the [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/) book.

## Overview

This repository contains exercises and implementations from each chapter of the book, organized into topical modules. Each module focuses on a specific Go concept with accompanying tests.

## Modules

| Module | Topics Covered |
|--------|----------------|
| `helloworld` | Basic Go syntax, functions, returning strings |
| `integers` | Integer operations, addition |
| `iteration` | `for` loops, `sum` functions |
| `structs_methods_interfaces` | Structs, methods, interfaces |
| `pointers` | Pointers, values, errors |
| `maps` | Map operations, CRUD |
| `arrays` | Array handling, slice operations |
| `dependency-injection` | Dependency injection, io.Writer |
| `mocking` | Mocking, interfaces, spies |
| `concurrency` | Goroutines, channels, concurrent operations |
| `select` | `select` statements, racing URLs, timeouts |
| `sync` | `sync` package, WaitGroups, mutexes |
| `context` | Context package, cancellation, deadlines |
| `arraysWithGenerics` | Arrays, generics, Reduce, Find, Sum operations |
| `generics` | Generic types, constraints, type parameters |
| `templating` | HTML templating, rendering, file-based templates |

## Running Tests

Run all tests:
```bash
go test ./...
```

Run tests for a specific module:
```bash
go test ./<module-name>
```

For example:
```bash
go test ./helloworld
go test ./concurrency
```

## Structure

Each module contains:
- `<module>.go` - Implementation code
- `<module>_test.go` - Tests

## Goals

- Learn Go through test-driven development
- Build solid foundation in Go fundamentals
- Practice writing idiomatic Go code