# User Input and Type Conversion

**📅 Created:** I/O and Conversion Series
**🏷️ Topics:** User Input, bufio, os.Stdin, Type Conversion, strconv, strings
**🔗 Related:** [01_variables-and-types.md](./01_variables-and-types.md), [Concepts/03userinput](../Concepts/03userinput/), [Concepts/04conversion](../Concepts/04conversion/)

---

## Overview

Reading user input and converting between types are fundamental operations in Go programs. This note covers the `bufio` package for efficient input reading, the `strconv` package for type conversion, and the `strings` package for string manipulation. Understanding these is essential for building interactive command-line applications.

**Why These Matter:**
- Build interactive CLI tools
- Parse user-provided data safely
- Convert between string and numeric types
- Handle input validation properly
- Process text efficiently

---

## Reading User Input

### The bufio Package

`bufio` provides buffered I/O, making reading input efficient and convenient.

**Basic Pattern:**

```go
import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Create a reader from standard input
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter your name: ")

    // Read until newline
    input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Hello,", input)
}
```

**Reasoning:** `bufio.NewReader` wraps `os.Stdin` for efficient buffered reading. `ReadString('\n')` reads until it encounters a newline character.

---

## Reading Different Input Types

### ReadString

Reads until a delimiter:

```go
reader := bufio.NewReader(os.Stdin)

// Read until newline
input, err := reader.ReadString('\n')

// Read until comma
data, err := reader.ReadString(',')
```

**Reasoning:** The delimiter is included in the returned string, so you'll need to trim it.

### ReadLine

Reads a line without the newline:

```go
line, isPrefix, err := reader.ReadLine()
// isPrefix is true if line was too long for buffer
```

**When to use:** When you need low-level control or are dealing with very long lines.

### Scan (Using fmt)

Quick way for simple input:

```go
var name string
var age int

fmt.Print("Enter name: ")
fmt.Scan(&name)

fmt.Print("Enter age: ")
fmt.Scan(&age)

fmt.Printf("Name: %s, Age: %d\n", name, age)
```

**Reasoning:** `fmt.Scan` is simpler but less flexible. It stops at whitespace and doesn't handle errors as gracefully.

---

## The Comma-OK (Comma-Err) Pattern

**Go's Error Handling Philosophy:**

Go doesn't use try-catch. Instead, functions return errors as values.

```go
// Comma-err syntax
value, err := someFunction()

if err != nil {
    // Handle error
    fmt.Println("Error:", err)
    return
}

// Use value safely
fmt.Println(value)
```

**Three Patterns:**

```go
// 1. Get both value and error
input, err := reader.ReadString('\n')

// 2. Ignore error (use with caution!)
input, _ := reader.ReadString('\n')

// 3. Ignore value, check error only
_, err := reader.ReadString('\n')
```

**Reasoning:** The `_` (blank identifier) explicitly discards values you don't need. This makes it clear you're intentionally ignoring something, not forgetting to handle it.

---

## String Manipulation

### The strings Package

Common operations on strings:

```go
import "strings"

input := "  Hello World  \n"

// Trim whitespace and newlines
clean := strings.TrimSpace(input)  // "Hello World"

// Case conversion
upper := strings.ToUpper(input)    // "  HELLO WORLD  \n"
lower := strings.ToLower(input)    // "  hello world  \n"

// Check contents
hasHello := strings.Contains(input, "Hello")  // true
starts := strings.HasPrefix(input, "  Hello") // true
ends := strings.HasSuffix(input, "World  \n") // true

// Split
parts := strings.Split("a,b,c", ",")  // ["a" "b" "c"]

// Join
joined := strings.Join(parts, "-")    // "a-b-c"

// Replace
replaced := strings.Replace(input, "World", "Go", 1)  // 1 = replace once
replaceAll := strings.ReplaceAll(input, " ", "_")
```

**Reasoning:** Always use `TrimSpace()` on user input to remove trailing newlines and whitespace. This prevents issues when converting to numbers or comparing strings.

---

## Type Conversion with strconv

### String to Number

```go
import "strconv"

// String to int
str := "42"
num, err := strconv.Atoi(str)
if err != nil {
    fmt.Println("Not a valid integer")
} else {
    fmt.Println(num)  // 42 (int)
}

// String to int64 (with base)
num64, err := strconv.ParseInt("42", 10, 64)
// Parameters: string, base (2-36), bitSize

// String to float
pi, err := strconv.ParseFloat("3.14", 64)
// Parameters: string, precision (32 or 64)

// String to bool
flag, err := strconv.ParseBool("true")  // Also accepts: 1, t, T, TRUE, True
```

**Reasoning:** Go requires explicit conversion - no implicit casting. This prevents subtle bugs where strings accidentally get treated as numbers.

### Number to String

```go
// Int to string
num := 42
str := strconv.Itoa(num)  // "42"

// Int64 to string
num64 := int64(100)
str = strconv.FormatInt(num64, 10)  // base 10

// Float to string
pi := 3.14159
str = strconv.FormatFloat(pi, 'f', 2, 64)
// Parameters: value, format, precision, bitSize
// Formats: 'f' (123.45), 'e' (1.23e+02), 'g' (auto)

// Or use fmt.Sprintf (simpler but allocates more)
str = fmt.Sprintf("%d", num)        // int
str = fmt.Sprintf("%.2f", pi)       // float with 2 decimals
```

**Reasoning:** `strconv` functions are more efficient than `fmt.Sprintf`, but `Sprintf` is more convenient and readable for complex formatting.

---

## Complete Example: Rating Input

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Rate our pizza (1-10): ")

    // Read input
    input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    // Clean input (remove newline and whitespace)
    rating := strings.TrimSpace(input)

    // Convert string to float
    floatRating, err := strconv.ParseFloat(rating, 64)
    if err != nil {
        fmt.Println("Invalid number:", err)
        return
    }

    // Convert to int and validate
    intRating := int(floatRating)
    if intRating < 1 || intRating > 10 {
        fmt.Println("Rating must be between 1 and 10")
        return
    }

    fmt.Printf("Thanks for rating: %d/10\n", intRating)
    fmt.Printf("That's %.0f stars!\n", floatRating)
}
```

**Step-by-Step:**
1. Create buffered reader from stdin
2. Read string until newline
3. Trim whitespace (crucial!)
4. Parse string to float
5. Convert float to int for comparison
6. Validate range
7. Use the validated input

**Reasoning:** This pattern handles errors at each step, validates input, and provides clear error messages. It's the foundation for robust CLI applications.

---

## Numeric Conversion Between Types

### Explicit Casting

```go
// Int to float
intNum := 42
floatNum := float64(intNum)  // 42.0

// Float to int (truncates decimal)
pi := 3.14
intPi := int(pi)  // 3 (not rounded!)

// For rounding, use math package
import "math"
rounded := int(math.Round(pi))  // 3

// Int types
var small int8 = 100
var large int64 = int64(small)

// Different int sizes
var x int32 = 42
var y int64 = int64(x)  // Must cast explicitly
```

**Reasoning:** Go never does implicit numeric conversion. This prevents accidental loss of precision or overflow. Every conversion must be explicit with `Type(value)` syntax.

---

## Common Patterns

### Pattern 1: Validated Integer Input

```go
func getIntInput(prompt string, min, max int) (int, error) {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print(prompt)

    input, err := reader.ReadString('\n')
    if err != nil {
        return 0, err
    }

    input = strings.TrimSpace(input)
    num, err := strconv.Atoi(input)
    if err != nil {
        return 0, fmt.Errorf("invalid integer: %w", err)
    }

    if num < min || num > max {
        return 0, fmt.Errorf("number must be between %d and %d", min, max)
    }

    return num, nil
}

// Usage
age, err := getIntInput("Enter age: ", 0, 120)
if err != nil {
    log.Fatal(err)
}
```

### Pattern 2: Menu Selection

```go
func showMenu() int {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Menu:")
    fmt.Println("1. Add")
    fmt.Println("2. Remove")
    fmt.Println("3. Exit")
    fmt.Print("Choose option: ")

    input, _ := reader.ReadString('\n')
    choice, err := strconv.Atoi(strings.TrimSpace(input))

    if err != nil || choice < 1 || choice > 3 {
        fmt.Println("Invalid choice")
        return showMenu()  // Retry
    }

    return choice
}
```

### Pattern 3: Yes/No Confirmation

```go
func confirm(prompt string) bool {
    reader := bufio.NewReader(os.Stdin)

    fmt.Printf("%s (y/n): ", prompt)
    input, _ := reader.ReadString('\n')

    answer := strings.ToLower(strings.TrimSpace(input))
    return answer == "y" || answer == "yes"
}

// Usage
if confirm("Delete file?") {
    // Delete file
}
```

---

## Error Handling Best Practices

### Don't Panic on User Input

```go
// ❌ Bad: Panics on invalid input
rating, err := strconv.Atoi(input)
if err != nil {
    panic(err)  // Crashes program!
}

// ✅ Good: Return error or handle gracefully
rating, err := strconv.Atoi(input)
if err != nil {
    fmt.Println("Please enter a valid number")
    return
}
```

**Reasoning:** User input errors are expected and should be handled gracefully. `panic` should only be used for truly exceptional situations (programmer errors, initialization failures).

### Wrap Errors for Context

```go
rating, err := strconv.ParseFloat(input, 64)
if err != nil {
    return fmt.Errorf("failed to parse rating: %w", err)
}
```

**Reasoning:** Wrapping errors with context helps debugging. The `%w` verb preserves the original error for error checking.

---

## Comparison with Other Languages

### Go vs Python

**Python:**
```python
name = input("Enter name: ")  # Simple, one line
age = int(input("Enter age: "))  # Implicit conversion, can crash
```

**Go:**
```go
reader := bufio.NewReader(os.Stdin)
input, err := reader.ReadString('\n')  // Explicit error handling
age, err := strconv.Atoi(strings.TrimSpace(input))
if err != nil {
    // Handle error
}
```

**Reasoning:** Go is more verbose but safer - errors are impossible to ignore. Python's simplicity can hide errors until runtime.

---

## Quick Reference Commands

```go
// Reading input
reader := bufio.NewReader(os.Stdin)
input, err := reader.ReadString('\n')

// Cleaning
clean := strings.TrimSpace(input)

// String to number
num, err := strconv.Atoi(clean)           // string to int
f, err := strconv.ParseFloat(clean, 64)   // string to float64

// Number to string
str := strconv.Itoa(42)                   // int to string
str := strconv.FormatFloat(3.14, 'f', 2, 64)

// Number conversions
floatVal := float64(intVal)
intVal := int(floatVal)
intVal := int(math.Round(floatVal))       // Rounded
```

---

## Summary

**Key Takeaways:**
- Use `bufio.NewReader(os.Stdin)` for reading input
- Always `TrimSpace()` on user input (removes newlines)
- Use `strconv` for type conversion (Atoi, ParseFloat, etc.)
- Check errors explicitly (comma-err pattern)
- No implicit conversions - all casting is explicit
- Don't panic on user input errors
- Use `strings` package for string manipulation
- Validate input ranges after conversion

**Quick Pattern:**
```go
reader := bufio.NewReader(os.Stdin)
input, _ := reader.ReadString('\n')
clean := strings.TrimSpace(input)
num, err := strconv.Atoi(clean)
if err != nil {
    // Handle error
}
// Use num
```

**When to Use This Note:**
- Building interactive CLI applications
- Parsing user input
- Converting between types
- Understanding Go's error handling
- Input validation patterns

---

**📝 Last Updated:** I/O Series
**➡️ Next Topic:** [Time Handling](./12_time-handling.md)
**🔗 Example Code:** [Concepts/03userinput](../Concepts/03userinput/), [Concepts/04conversion](../Concepts/04conversion/)
