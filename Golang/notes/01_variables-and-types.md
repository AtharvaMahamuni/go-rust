# Variables and Types in Go

**📅 Created:** Fundamentals Series
**🏷️ Topics:** Variables, Data Types, Type System, Zero Values, Type Inference
**🔗 Related:** [00_golang-basics.md](./00_golang-basics.md), [02_pointers.md](./02_pointers.md), [Concepts/02variables](../Concepts/02variables/)

---

## Overview

This note covers Go's type system, variable declarations, basic and advanced types, zero values, type inference, and the special walrus operator (`:=`). Essential for understanding how Go handles data and enforces type safety at compile time.

---

## Go's Type System Philosophy

**Key Principles:**
- **Everything is a type** - even functions
- **Statically typed** - types must be known at compile time
- **Type safety** - no implicit conversions (explicit only)
- **Zero values** - uninitialized variables get safe default values

**Capitalization Matters:**
- `PublicVariable` - Capital first letter = **Public** (exported)
- `privateVariable` - Small first letter = **Private** (package-only)

---

## Basic Types

### Strings

```go
var username string = "Atharva"
fmt.Println(username)                      // Atharva
fmt.Printf("Type: %T\n", username)         // Type: string
```

**Characteristics:**
- Immutable (can't change individual characters)
- UTF-8 encoded by default
- Zero value: `""` (empty string)

### Booleans

```go
var isLoggedIn bool = true
fmt.Println(isLoggedIn)                    // true
fmt.Printf("Type: %T\n", isLoggedIn)       // Type: bool
```

**Characteristics:**
- Only two values: `true` or `false`
- Zero value: `false`

### Integers

```go
var smallInt uint8 = 255      // 0 to 255
var regularInt int = 42       // -2^31 to 2^31-1 (on 32-bit systems)
var bigInt int64 = 9223372036854775807
```

**Integer Types:**

| Type | Size | Range |
|------|------|-------|
| `int8` | 1 byte | -128 to 127 |
| `uint8` | 1 byte | 0 to 255 |
| `int16` | 2 bytes | -32,768 to 32,767 |
| `uint16` | 2 bytes | 0 to 65,535 |
| `int32` | 4 bytes | -2^31 to 2^31-1 |
| `uint32` | 4 bytes | 0 to 2^32-1 |
| `int64` | 8 bytes | -2^63 to 2^63-1 |
| `uint64` | 8 bytes | 0 to 2^64-1 |
| `int` | Platform dependent | 32 or 64 bits |
| `uint` | Platform dependent | 32 or 64 bits |

**Zero value:** `0`

### Floating Point

```go
var smallFloat float32 = 25.36237864366435
fmt.Println(smallFloat)                    // 25.362379 (precision limited)

var bigFloat float64 = 3.141592653589793
fmt.Println(bigFloat)                      // Full precision
```

**Types:**
- `float32` - Single precision (~6-7 decimal digits)
- `float64` - Double precision (~15-16 decimal digits)

**Zero value:** `0.0`

**Note:** Always use `float64` unless you have specific memory constraints

### Complex Numbers

```go
var c complex64 = 1 + 2i
var d complex128 = complex(3, 4)  // 3 + 4i

fmt.Println(real(c))  // 1
fmt.Println(imag(c))  // 2
```

**Types:**
- `complex64` - float32 real and imaginary parts
- `complex128` - float64 real and imaginary parts

**Zero value:** `0 + 0i`

---

## Type Aliases

Go provides convenient aliases for common types:

```go
// Aliases in standard library
byte    // alias for uint8 (raw data)
rune    // alias for int32 (Unicode code points)
int     // alias for int32 or int64 (platform dependent)
uint    // alias for uint32 or uint64 (platform dependent)
```

**Example Usage:**

```go
var data byte = 255               // Raw byte data
var char rune = '😀'              // Unicode character
var count int = 42                // General purpose integer
```

---

## Zero Values

**Unlike many languages, Go NEVER has uninitialized variables!**

Every variable gets a "zero value" by default:

| Type | Zero Value |
|------|------------|
| `int`, `uint`, `float` | `0` |
| `string` | `""` (empty string) |
| `bool` | `false` |
| `pointer` | `nil` |
| `slice`, `map`, `channel` | `nil` |
| `struct` | All fields set to their zero values |

**Example:**

```go
var name string         // ""
var age int             // 0
var isActive bool       // false
var ptr *int            // nil
```

**Why this matters:**
- No "undefined" or "null" surprises
- Safe defaults prevent crashes
- Explicit about initialization

---

## Variable Declaration Styles

### 1. Explicit Type Declaration

```go
var username string = "Atharva"
var age int = 25
var isActive bool = true
```

**Use when:**
- You want to be explicit about the type
- Working with interfaces or specific numeric types

### 2. Type Inference (Implicit Type)

```go
var username = "Atharva"      // Go infers: string
var age = 25                  // Go infers: int
var price = 19.99             // Go infers: float64
```

**How it works:**
- Go looks at the value and determines the type
- Once set, the type is FIXED (can't change later)
- Clean syntax while maintaining type safety

### 3. Short Variable Declaration (Walrus Operator)

```go
numberOfUsers := 42           // int
website := "github.com"       // string
isValid := true               // bool
```

**The `:=` operator:**
- Declares AND assigns in one step
- Type is inferred from the value
- **Can ONLY be used inside functions** (not for global variables)
- Most common style in Go code

**Restrictions:**

```go
// ❌ ERROR: Cannot use := outside functions
package main
numberOfUsers := 42   // Compile error

// ✅ OK: Use var for package-level variables
package main
var numberOfUsers = 42

func main() {
    // ✅ OK: := works inside functions
    count := 10
}
```

---

## Multiple Variable Declarations

### Multiple Variables - Same Type

```go
var x, y, z int = 1, 2, 3

// Or with inference
var a, b, c = 1, 2, 3

// Or with walrus
x, y, z := 1, 2, 3
```

### Multiple Variables - Different Types

```go
var (
    name     string = "Atharva"
    age      int    = 25
    isActive bool   = true
)

// Or in one line
name, age, isActive := "Atharva", 25, true
```

---

## Type System Rules

### 1. No Implicit Conversions

```go
var x int = 10
var y float64 = x      // ❌ ERROR: Cannot use int as float64

// ✅ Must convert explicitly
var y float64 = float64(x)
```

### 2. Type Cannot Change

```go
x := 10          // x is int
x = "hello"      // ❌ ERROR: Cannot assign string to int

// Once declared, type is fixed
```

### 3. Unused Variables Are Errors

```go
func main() {
    x := 10      // ❌ ERROR: x declared and not used
}

// Go forces you to use what you declare
```

**Exception:** Use `_` to explicitly ignore:

```go
value, _ := someFunction()  // Ignore second return value
```

---

## Constants

Constants are immutable values known at compile time:

```go
const Pi = 3.14159
const Language = "Go"
const MaxUsers = 1000

// Multiple constants
const (
    StatusOK       = 200
    StatusNotFound = 404
    StatusError    = 500
)
```

**Rules:**
- Must be assigned at declaration
- Value must be known at compile time
- Cannot use := for constants
- Can be untyped (more flexible)

**Untyped Constants:**

```go
const x = 42         // Untyped constant
var a int = x        // ✅ OK: x becomes int
var b float64 = x    // ✅ OK: x becomes float64
```

---

## Type Conversion

Go requires **explicit type conversion** (no automatic conversions):

```go
// Integer conversions
var x int32 = 10
var y int64 = int64(x)     // Explicit conversion

// Float conversions
var a float64 = 3.14
var b int = int(a)         // b = 3 (truncates decimal)

// String conversions (see conversion package)
import "strconv"

// String to int
str := "42"
num, err := strconv.Atoi(str)
if err != nil {
    // Handle error
}

// Int to string
num := 42
str := strconv.Itoa(num)

// String to float
str := "3.14"
num, err := strconv.ParseFloat(str, 64)

// Float to string
num := 3.14
str := strconv.FormatFloat(num, 'f', 2, 64)
```

**More on type conversion:** See [04_conversion](../Concepts/04conversion/)

---

## The `fmt` Package - Printing and Formatting

### Basic Printing

```go
import "fmt"

fmt.Println("Hello")          // With newline
fmt.Print("Hello")            // Without newline
```

### Formatted Printing

```go
name := "Atharva"
age := 25

// %v = value in default format
fmt.Printf("Name: %v, Age: %v\n", name, age)

// %T = type
fmt.Printf("Type: %T\n", age)  // Type: int

// %d = decimal integer
fmt.Printf("Age: %d\n", age)

// %s = string
fmt.Printf("Name: %s\n", name)

// %f = floating point
price := 19.99
fmt.Printf("Price: %.2f\n", price)  // Price: 19.99
```

**Common Format Verbs:**

| Verb | Meaning | Example |
|------|---------|---------|
| `%v` | Value (default format) | `%v` |
| `%+v` | Value with field names (structs) | `%+v` |
| `%T` | Type of value | `int`, `string` |
| `%d` | Decimal integer | `42` |
| `%f` | Floating point | `3.14` |
| `%s` | String | `"hello"` |
| `%t` | Boolean | `true` |
| `%p` | Pointer address | `0xc0000140a0` |
| `%x` | Hexadecimal | `2a` |
| `%b` | Binary | `101010` |

---

## Practical Examples

### Example 1: All Declaration Styles

```go
package main

import "fmt"

// Package-level variables (must use var)
var globalCount = 100

func main() {
    // Explicit type
    var name string = "Atharva"

    // Type inference
    var age = 25

    // Short declaration (walrus)
    city := "Mumbai"

    fmt.Printf("Name: %s, Age: %d, City: %s\n", name, age, city)
    fmt.Printf("Global: %d\n", globalCount)
}
```

### Example 2: Working with Different Types

```go
package main

import "fmt"

func main() {
    // Integers
    var small uint8 = 255
    var big int64 = 9223372036854775807

    // Floats
    pi := 3.14159265359

    // Strings
    message := "Hello, Go!"

    // Booleans
    isActive := true

    fmt.Printf("Small: %d (type: %T)\n", small, small)
    fmt.Printf("Big: %d (type: %T)\n", big, big)
    fmt.Printf("Pi: %.2f (type: %T)\n", pi, pi)
    fmt.Printf("Message: %s (type: %T)\n", message, message)
    fmt.Printf("Active: %t (type: %T)\n", isActive, isActive)
}
```

### Example 3: Zero Values

```go
package main

import "fmt"

func main() {
    var s string
    var i int
    var f float64
    var b bool

    fmt.Printf("string: '%s'\n", s)    // string: ''
    fmt.Printf("int: %d\n", i)         // int: 0
    fmt.Printf("float: %f\n", f)       // float: 0.000000
    fmt.Printf("bool: %t\n", b)        // bool: false
}
```

---

## Common Pitfalls

### Pitfall 1: Using := Outside Functions

```go
// ❌ ERROR
package main
name := "Atharva"  // Cannot use := at package level

// ✅ CORRECT
package main
var name = "Atharva"
```

### Pitfall 2: Implicit Type Conversion

```go
var x int = 10
var y float64 = x  // ❌ ERROR

// ✅ CORRECT
var y float64 = float64(x)
```

### Pitfall 3: Unused Variables

```go
func main() {
    x := 10
    // ❌ ERROR: x declared and not used
}

// ✅ CORRECT: Use it or remove it
func main() {
    x := 10
    fmt.Println(x)
}
```

### Pitfall 4: Redeclaring with :=

```go
name := "Atharva"
name := "New Name"  // ❌ ERROR: no new variables on left side

// ✅ CORRECT: Use = for reassignment
name = "New Name"
```

---

## Summary

**Key Takeaways:**
- Go is statically typed with compile-time type checking
- Three declaration styles: `var name type = value`, `var name = value`, `name := value`
- `:=` (walrus) only works inside functions
- Zero values ensure no uninitialized variables
- No implicit type conversions (explicit only)
- Capitalization determines public/private visibility
- Unused variables cause compile errors
- Use `fmt.Printf` with format verbs for formatted output

**Type System:**
```
Basic Types:    string, bool, int, uint, float, complex
Aliases:        byte (uint8), rune (int32)
Zero Values:    "", 0, false, nil
Constants:      const Name = value
```

**Declaration Styles:**
```go
var x int = 10        // Explicit
var x = 10            // Inferred
x := 10               // Short (functions only)
```

**When to Use This Note:**
- Learning Go's type system
- Understanding variable declarations
- Reference for type conversions
- Debugging type errors
- Comparing with dynamically typed languages

---

**📝 Last Updated:** Fundamentals Series
**➡️ Next Topic:** [Pointers and Memory](./02_pointers.md)
**🔗 Example Code:** [Concepts/02variables](../Concepts/02variables/), [Concepts/04conversion](../Concepts/04conversion/)
