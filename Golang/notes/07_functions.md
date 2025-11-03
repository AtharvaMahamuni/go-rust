# Functions in Go

**📅 Created:** Language Fundamentals
**🏷️ Topics:** Functions, Parameters, Return Values, Variadic Functions, Anonymous Functions
**🔗 Related:** [08_methods.md](./08_methods.md), [09_defer.md](./09_defer.md), [Concepts/15functions](../Concepts/15functions/)

---

## Overview

Functions are the building blocks of Go programs. Go functions are first-class values - they can be assigned to variables, passed as arguments, and returned from other functions. Understanding functions is essential for writing clean, reusable code.

**Why Functions in Go Are Special:**
- Multiple return values (especially for errors)
- Named return values
- Variadic parameters
- First-class functions (can be assigned/passed)
- Deferred execution (defer keyword)

---

## Basic Function Syntax

```go
func functionName(param1 type1, param2 type2) returnType {
    // Function body
    return value
}
```

### Simple Example

```go
func sayHi() {
    fmt.Println("Hey, Welcome!")
}

func main() {
    sayHi()  // Call function
}
```

**Reasoning:** Functions without parameters use empty parentheses. Functions without return values don't specify a return type.

---

## Parameters

### Single Parameters

```go
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

greet("Alice")  // Hello, Alice!
```

### Multiple Parameters (Same Type)

```go
// When consecutive parameters have the same type
func adder(valOne int, valTwo int) int {
    return valOne + valTwo
}

// Shorthand (type once at end)
func adder(valOne, valTwo int) int {
    return valOne + valTwo
}

result := adder(3, 8)  // 11
```

**Reasoning:** Go allows you to specify the type once for consecutive parameters of the same type, reducing verbosity.

### Multiple Parameters (Different Types)

```go
func createUser(name string, age int, active bool) {
    fmt.Printf("User: %s, Age: %d, Active: %t\n", name, age, active)
}

createUser("Alice", 25, true)
```

---

## Return Values

### Single Return Value

```go
func add(a, b int) int {
    return a + b
}

sum := add(3, 5)  // 8
```

### Multiple Return Values

**This is one of Go's most powerful features:**

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Usage
result, err := divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)  // 5
}
```

**Reasoning:** Multiple returns eliminate the need for out parameters or error objects. The convention of returning `(value, error)` makes error handling explicit and impossible to ignore.

### Named Return Values

```go
func divide(a, b float64) (result float64, err error) {
    if b == 0 {
        err = fmt.Errorf("division by zero")
        return  // Returns result=0, err=error
    }
    result = a / b
    return  // Returns result and err
}
```

**Reasoning:** Named returns are automatically declared and zero-initialized. A bare `return` returns them. Useful for complex functions, but can reduce clarity if overused.

**When to use named returns:**
- Documentation purposes (clear what's returned)
- Functions with multiple exit points
- defer functions that modify return values

**When not to use:**
- Short, simple functions (overkill)
- Can make code less clear

---

## Variadic Functions

Accept any number of arguments of the same type:

```go
func proAdder(values ...int) (int, string) {
    sum := 0
    for _, value := range values {
        sum += value
    }
    return sum, "Pro function"
}

// Usage
result, msg := proAdder(4, 6, 8)        // sum: 18
result2, _ := proAdder(1, 2, 3, 4, 5)   // sum: 15
```

**Reasoning:** The `...type` syntax creates a slice inside the function. You can pass any number of arguments. This is how `fmt.Println` accepts multiple values.

### Passing a Slice to Variadic Function

```go
numbers := []int{1, 2, 3, 4, 5}

// Unpack slice with ...
sum, _ := proAdder(numbers...)
```

**Reasoning:** The `...` operator unpacks a slice into individual arguments. Without it, you'd pass a single slice argument (type mismatch).

---

## Anonymous Functions

Functions without names, defined inline:

```go
// Assign to variable
add := func(a, b int) int {
    return a + b
}

result := add(3, 5)  // 8
```

### Immediately Invoked Function Expression (IIFE)

```go
func main() {
    result := func(a, b int) int {
        return a + b
    }(3, 5)  // Call immediately

    fmt.Println(result)  // 8
}
```

**Reasoning:** Anonymous functions are useful for closures, goroutines, and callbacks. IIFEs are common in initialization code or to create local scopes.

---

## Closures

Functions that reference variables from their surrounding scope:

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// Usage
c := counter()
fmt.Println(c())  // 1
fmt.Println(c())  // 2
fmt.Println(c())  // 3
```

**Reasoning:** The inner function "closes over" the `count` variable, maintaining state between calls. Each call to `counter()` creates a new independent counter.

### Practical Example - Custom Filters

```go
func filter(numbers []int, test func(int) bool) []int {
    result := []int{}
    for _, num := range numbers {
        if test(num) {
            result = append(result, num)
        }
    }
    return result
}

// Usage
numbers := []int{1, 2, 3, 4, 5, 6}

evens := filter(numbers, func(n int) bool {
    return n%2 == 0
})

fmt.Println(evens)  // [2 4 6]
```

**Reasoning:** Passing functions as arguments enables powerful abstractions. This is the foundation of functional programming patterns in Go.

---

## Function Types

Functions are first-class types:

```go
type MathFunc func(int, int) int

func apply(a, b int, f MathFunc) int {
    return f(a, b)
}

func add(a, b int) int { return a + b }
func multiply(a, b int) int { return a * b }

// Usage
result1 := apply(3, 5, add)       // 8
result2 := apply(3, 5, multiply)  // 15
```

**Reasoning:** Defining function types makes signatures clearer and enables type checking for function parameters.

---

## Common Patterns

### Pattern 1: Error Handling

```go
func readFile(filename string) ([]byte, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("failed to read %s: %w", filename, err)
    }
    return data, nil
}
```

**Reasoning:** Always return errors as the last value. Check errors immediately after function calls. Use `fmt.Errorf` with `%w` to wrap errors, preserving the error chain.

### Pattern 2: Options/Config Functions

```go
type Server struct {
    host string
    port int
}

type ServerOption func(*Server)

func WithPort(port int) ServerOption {
    return func(s *Server) {
        s.port = port
    }
}

func NewServer(host string, opts ...ServerOption) *Server {
    s := &Server{host: host, port: 8080}
    for _, opt := range opts {
        opt(s)
    }
    return s
}

// Usage
server := NewServer("localhost", WithPort(3000))
```

**Reasoning:** Functional options pattern provides extensible, readable configuration without constructor overloading.

---

## Defer, Panic, Recover (Brief)

### Defer

```go
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()  // Runs when function exits

    // Work with file...
    return nil
}
```

**Reasoning:** `defer` ensures cleanup happens even if function returns early or panics. See [09_defer.md](./09_defer.md) for details.

---

## Summary

**Key Takeaways:**
- Functions can return multiple values
- Error as last return value is idiomatic
- Variadic functions accept any number of arguments
- Named returns for documentation and defer
- First-class functions (can assign/pass/return)
- Closures capture surrounding variables
- Defer for cleanup (always executes)

**Quick Reference:**
```go
// Basic
func name(params) returnType { }

// Multiple params (same type)
func add(a, b int) int { }

// Multiple returns
func divide(a, b float64) (float64, error) { }

// Named returns
func calc() (result int, err error) { }

// Variadic
func sum(numbers ...int) int { }

// Anonymous
func(x int) int { return x * 2 }

// Type
type MathFunc func(int, int) int
```

---

**📝 Last Updated:** Language Fundamentals
**➡️ Next Topic:** [Methods](./08_methods.md)
**🔗 Example Code:** [Concepts/15functions](../Concepts/15functions/)
