# Pointers in Go

**📅 Created:** Fundamentals Series
**🏷️ Topics:** Pointers, Memory Addresses, Dereferencing, Pass by Reference
**🔗 Related:** [01_variables-and-types.md](./01_variables-and-types.md), [04_maps-and-structs.md](./04_maps-and-structs.md), [Concepts/07pointer](../Concepts/07pointer/)

---

## Overview

This note covers Go's pointer system: what pointers are, how to use them, when to use them, and how they differ from languages like C/C++ and Java/Kotlin. Pointers are essential for understanding how Go handles memory and enables efficient data manipulation.

---

## What is a Pointer?

**Simple Definition:**
A pointer is a variable that stores the **memory address** of another variable.

```
Variable:  [Value]
Pointer:   [Address] ──→ [Value]
```

**Why Pointers Matter:**
- Pass large data without copying
- Modify values in functions
- Work with dynamic data structures
- Understand how Go manages memory

---

## The Two Operators

### `&` (Address-of Operator)

Gets the memory address of a variable:

```go
number := 42
ptr := &number    // ptr now holds the address of number

fmt.Println(number)   // 42
fmt.Println(&number)  // 0xc0000140a0 (memory address)
fmt.Println(ptr)      // 0xc0000140a0 (same address)
```

### `*` (Dereference Operator)

Has TWO different uses depending on context:

**1. In type declaration - declares a pointer type:**
```go
var ptr *int    // ptr is a pointer to an int
```

**2. In expression - accesses the value at the address:**
```go
value := *ptr   // Get the value that ptr points to
*ptr = 50       // Set the value that ptr points to
```

---

## Basic Pointer Usage

### Example 1: Creating and Using Pointers

```go
package main

import "fmt"

func main() {
    number := 22
    ptr := &number     // ptr points to number

    fmt.Println("number is:", number)      // 22
    fmt.Println("&number is:", &number)    // 0xc0000140a0
    fmt.Println("ptr is:", ptr)            // 0xc0000140a0
    fmt.Println("*ptr is:", *ptr)          // 22
}
```

**Memory Visualization:**

```
Stack:
┌────────────────┐
│ number = 22    │  at address 0xc0000140a0
└────────────────┘
         ↑
         │
┌────────────────┐
│ ptr = 0xc...a0 │  (stores the address)
└────────────────┘
```

### Example 2: Modifying Through Pointers

```go
package main

import "fmt"

func main() {
    number := 22
    ptr := &number

    fmt.Println("Before:", number)   // 22

    *ptr = *ptr + 3    // Modify the value through pointer

    fmt.Println("After:", number)    // 25
}
```

**What happens:**
1. `*ptr` gets the current value (22)
2. Add 3 to get 25
3. `*ptr = 25` sets the value at that address
4. `number` now equals 25 (same memory location)

---

## Pointer Declarations

### Method 1: Declare and Initialize

```go
number := 42
ptr := &number    // ptr is *int, points to number
```

### Method 2: Declare Without Initialization

```go
var ptr *int      // ptr is nil (doesn't point to anything yet)

fmt.Println(ptr)  // <nil>

// Later, assign it
number := 42
ptr = &number     // Now ptr points to number
```

### Method 3: Using `new` Keyword

```go
ptr := new(int)   // Allocates memory for an int, returns pointer
*ptr = 42         // Set the value

fmt.Println(*ptr) // 42
```

**What `new` does:**
- Allocates memory for the type
- Initializes with zero value
- Returns a pointer to that memory

---

## Nil Pointers

**Zero value of a pointer is `nil`:**

```go
var ptr *int
fmt.Println(ptr)  // <nil>

if ptr == nil {
    fmt.Println("Pointer is nil")
}
```

**WARNING: Dereferencing nil causes panic:**

```go
var ptr *int
fmt.Println(*ptr)  // PANIC: runtime error
```

**Always check before dereferencing:**

```go
var ptr *int
if ptr != nil {
    fmt.Println(*ptr)  // Safe
} else {
    fmt.Println("Pointer is nil")
}
```

---

## Pointers in Functions

### Pass by Value (Default in Go)

```go
func increment(x int) {
    x = x + 1    // Only changes local copy
}

func main() {
    num := 10
    increment(num)
    fmt.Println(num)  // Still 10 (unchanged)
}
```

### Pass by Pointer (Reference-like Behavior)

```go
func increment(x *int) {
    *x = *x + 1  // Changes the original value
}

func main() {
    num := 10
    increment(&num)   // Pass address
    fmt.Println(num)  // 11 (changed!)
}
```

**Memory View:**

```
Before increment(&num):
main():           increment():
num = 10          x = &num ──→ num = 10

After *x = *x + 1:
main():           increment():
num = 11          x = &num ──→ num = 11
```

---

## When to Use Pointers

### Use Pointers When:

1. **You need to modify the original value:**
```go
func reset(x *int) {
    *x = 0
}
```

2. **Working with large structs (avoid copying):**
```go
type BigStruct struct {
    data [1000000]int
}

// Bad: Copies entire struct
func processBad(b BigStruct) { }

// Good: Only copies pointer (8 bytes)
func processGood(b *BigStruct) { }
```

3. **Need to represent "no value" (nil):**
```go
var user *User = nil  // No user
if user == nil {
    // Handle no user case
}
```

4. **Working with methods that modify receivers:**
```go
type Counter struct {
    count int
}

func (c *Counter) Increment() {
    c.count++  // Modifies original
}
```

### Don't Use Pointers When:

1. **Small values (int, bool, float):**
```go
// Overkill - int is cheap to copy
func add(x *int, y *int) int {
    return *x + *y
}

// Better
func add(x int, y int) int {
    return x + y
}
```

2. **Immutable operations:**
```go
// If you don't need to modify, don't use pointers
func display(name string) {
    fmt.Println(name)
}
```

---

## Pointers vs Other Languages

### C/C++ vs Go

**C/C++:**
```c
int x = 10;
int* ptr = &x;
*ptr = 20;

// Pointer arithmetic (dangerous!)
ptr++;
*(ptr + 5) = 30;

// Manual memory management
int* p = malloc(sizeof(int));
free(p);
```

**Go:**
```go
x := 10
ptr := &x
*ptr = 20

// ❌ NO pointer arithmetic (safer!)
// ptr++  // Compile error

// ✅ Automatic memory management
ptr := new(int)  // GC handles cleanup
```

**Key Differences:**
- Go: No pointer arithmetic
- Go: Automatic garbage collection
- Go: Safer, but still efficient

### Java/Kotlin vs Go

**Java/Kotlin:**
```kotlin
// Everything (except primitives) is a reference
val user = User("Alice")  // user is a reference
modify(user)              // Automatically by reference

fun modify(u: User) {
    u.name = "Bob"       // Modifies original
}
```

**Go:**
```go
// Explicit choice
user := User{name: "Alice"}
modify(&user)             // Must explicitly pass pointer

func modify(u *User) {
    u.name = "Bob"        // Modifies original
}
```

**Key Differences:**
- Go: Explicit about pointers (& and *)
- Java/Kotlin: Implicit references (automatic)
- Go: More control, clearer intent

---

## Common Patterns

### Pattern 1: Returning Pointers

```go
func createUser(name string) *User {
    user := User{name: name}
    return &user  // Safe! Go moves to heap if needed
}

func main() {
    user := createUser("Alice")
    fmt.Println(user.name)
}
```

**Note:** Go's escape analysis determines if `user` should be on heap or stack.

### Pattern 2: Pointer to Struct Field

```go
type Person struct {
    name string
    age  int
}

func main() {
    p := Person{name: "Alice", age: 25}

    // Pointer to entire struct
    ptr := &p
    ptr.age = 26  // Go automatically dereferences

    // Equivalent to:
    (*ptr).age = 26
}
```

**Syntactic Sugar:** Go automatically dereferences struct pointers.

### Pattern 3: Optional Values with Pointers

```go
type Config struct {
    timeout *int  // nil means "use default"
}

func process(cfg Config) {
    timeout := 30  // default
    if cfg.timeout != nil {
        timeout = *cfg.timeout  // use provided value
    }
    fmt.Println("Timeout:", timeout)
}

func main() {
    // Use default
    process(Config{})

    // Provide value
    t := 60
    process(Config{timeout: &t})
}
```

---

## Pointer Pitfalls

### Pitfall 1: Forgetting to Dereference

```go
func double(x *int) {
    x = x * 2  // ❌ Wrong! Modifying pointer itself
}

// ✅ Correct
func double(x *int) {
    *x = *x * 2  // Dereference to modify value
}
```

### Pitfall 2: Returning Pointer to Local Variable (in C)

```go
// In Go, this is SAFE (escape analysis)
func getNumber() *int {
    x := 42
    return &x  // ✅ Go moves x to heap
}

// In C, this would be dangerous (dangling pointer)
```

### Pitfall 3: Nil Pointer Dereference

```go
var ptr *int
*ptr = 10  // ❌ PANIC: nil pointer dereference

// ✅ Check first
if ptr != nil {
    *ptr = 10
}
```

### Pitfall 4: Pointer to Pointer Confusion

```go
var x int = 10
var ptr *int = &x
var pptr **int = &ptr  // Pointer to pointer

fmt.Println(**pptr)  // 10 (double dereference)
```

---

## Practical Examples

### Example 1: Swapping Values

```go
func swap(a *int, b *int) {
    temp := *a
    *a = *b
    *b = temp
}

func main() {
    x, y := 10, 20
    fmt.Printf("Before: x=%d, y=%d\n", x, y)

    swap(&x, &y)

    fmt.Printf("After: x=%d, y=%d\n", x, y)
}
// Output:
// Before: x=10, y=20
// After: x=20, y=10
```

### Example 2: Modifying Struct

```go
type BankAccount struct {
    balance int
}

func (acc *BankAccount) Deposit(amount int) {
    acc.balance += amount  // Modifies original
}

func main() {
    account := BankAccount{balance: 100}
    account.Deposit(50)
    fmt.Println(account.balance)  // 150
}
```

### Example 3: Large Data Efficiency

```go
type LargeData struct {
    data [1000000]int
}

// Bad: Copies 8MB every call
func processBad(d LargeData) {
    // Process data...
}

// Good: Copies 8 bytes (pointer size)
func processGood(d *LargeData) {
    // Process data...
}

func main() {
    large := LargeData{}
    processGood(&large)  // Efficient!
}
```

---

## Memory Management Notes

### Stack vs Heap Allocation

**Stack Allocation (faster):**
```go
func example() {
    x := 10  // On stack (if doesn't escape)
}
```

**Heap Allocation (when escapes):**
```go
func example() *int {
    x := 10
    return &x  // Escapes to heap (returned)
}
```

**Go's Escape Analysis:**
- Compiler decides stack vs heap automatically
- You write code normally
- Optimization happens behind the scenes

### Garbage Collection

```go
ptr := new(int)
*ptr = 42
// Use ptr...
// No need to free! GC handles it
```

**Key Point:** Go has automatic garbage collection - you never manually `free` memory.

---

## Summary

**Key Takeaways:**
- Pointers store memory addresses
- `&` gets address, `*` dereferences (in expressions)
- `*Type` declares pointer type
- Zero value is `nil` - always check before dereferencing
- Use pointers for: large structs, need to modify, optional values
- No pointer arithmetic (safer than C/C++)
- Automatic garbage collection (no manual free)
- Go's escape analysis handles stack vs heap

**Pointer Syntax:**
```go
var ptr *int          // Declare pointer
ptr = &variable       // Get address
value = *ptr          // Dereference
*ptr = newValue       // Modify through pointer
ptr = new(int)        // Allocate with new
```

**Quick Reference:**
```
&variable  → Get address
*ptr       → Get value at address
*Type      → Pointer type
nil        → Null pointer
new(Type)  → Allocate and return pointer
```

**When to Use This Note:**
- Understanding memory addresses
- Passing by reference in functions
- Optimizing large struct handling
- Implementing methods that modify receivers
- Comparing Go pointers with C/C++

---

**📝 Last Updated:** Fundamentals Series
**➡️ Next Topic:** [Arrays and Slices](./03_arrays-and-slices.md)
**🔗 Example Code:** [Concepts/07pointer](../Concepts/07pointer/)
