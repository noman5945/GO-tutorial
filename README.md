# Go Language Core Concepts

Welcome to the Go programming language tutorial! This document covers some of the core concepts of Go to help you get started.

## Table of Contents

1. [Variables](#variables)
2. [Data Types](#data-types)
3. [Control Structures](#control-structures)
4. [Functions](#functions)
5. [Structs](#structs)
6. [Interfaces](#interfaces)
7. [Concurrency](#concurrency)
8. [Error Handling](#error-handling)
9. [Packages](#packages)
10. [Resources](#resources)

---

## <a name="variables"></a>Variables

In Go, variables are explicitly declared and used by the compiler.

```go
var x int = 10
y := 20 // Short declaration
```

## <a name="data-types"></a>Data Types

Go has several built-in data types:

    Basic Types: int, float64, bool, string
    Composite Types: array, slice, map, struct

```go
    var name string = "Go"
    var isAwesome bool = true
```

## <a name="control-structures"></a>Control Structures

Go supports standard control structures:

    If-Else

    For Loop

    Switch

```go
if x > 10 {
    fmt.Println("x is greater than 10")
}

for i := 0; i < 5; i++ {
    fmt.Println(i)
}

switch day {
case "Monday":
    fmt.Println("Start of the week")
default:
    fmt.Println("Other day")
}
```

## <a name="functions"></a>Functions

Functions in Go are defined using the func keyword.

```go
   func add(a int, b int) int {
    return a + b
}
```

## <a name="structs"></a>Structs

Structs are used to define custom data types.

```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Alice", Age: 30}
```

## <a name="interfaces"></a>Interfaces

Interfaces define a set of method signatures.

```go
type Speaker interface {
    Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}
```

## <a name="concurrency"></a>Concurrency

Go has built-in support for concurrency using goroutines and channels.

```go
result, err := someFunction()
if err != nil {
    log.Fatal(err)
}
```

## <a name="packages"></a>Packages

Go code is organized into packages.

```go
import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Sqrt(16))
}
```

## <a name="resources"></a>Resources

<a href="https://go.dev/doc/"> Official Go Documentation </a>

<a href="https://go.dev/tour/welcome/1"> A Tour of Go </a>

<a href="https://go.dev/doc/effective_go"> Effective Go </a>
