# Methods in Go

**📅 Created:** Object-Oriented Patterns
**🏷️ Topics:** Methods, Receivers, Value vs Pointer Receivers, Method Sets
**🔗 Related:** [06_structs.md](./06_structs.md), [07_functions.md](./07_functions.md), [02_pointers.md](./02_pointers.md), [Concepts/16methods](../Concepts/16methods/)

---

## Overview

Methods are functions with a special receiver argument. They allow you to attach behavior to types, similar to methods in classes but without the complexity of inheritance. Methods are how Go achieves object-oriented-like programming through composition and interfaces.

**Key Insight:** When functions go inside structs (or are associated with types), we call them methods!

---

## Defining Methods

**Syntax:**
```go
func (receiver ReceiverType) methodName(parameters) returnType {
    // Method body
}
```

### Basic Example

```go
type User struct {
    Name   string
    Email  string
    Status bool
    Age    int
}

// Method with value receiver
func (u User) GetStatus() {
    fmt.Println("User status is:", u.Status)
}

// Usage
atharva := User{"Atharva", "atharva@example.com", true, 23}
atharva.GetStatus()  // User status is: true
```

**Reasoning:** The `(u User)` before the function name is the receiver. It's like the implicit `this` or `self` in other languages, but explicit in Go.

---

## Value Receivers vs Pointer Receivers

### Value Receiver

```go
func (u User) NewMail() {
    u.Email = "test@go.dev"
    fmt.Println("New user email is:", u.Email)
}

// Usage
atharva := User{"Atharva", "atharva@example.com", true, 2}
atharva.NewMail()  // Prints: New user email is: test@go.dev
fmt.Println(atharva.Email)  // Still: atharva@example.com
```

**What happens:**
- The method receives a COPY of the struct
- Changes to the copy don't affect the original
- Original struct remains unchanged

**Reasoning:** Value receivers work on a copy. This is safe (no accidental mutations) but inefficient for large structs.

### Pointer Receiver

```go
func (u *User) UpdateEmail(newEmail string) {
    u.Email = newEmail
}

// Usage
atharva := User{"Atharva", "atharva@example.com", true, 2}
atharva.UpdateEmail("new@example.com")
fmt.Println(atharva.Email)  // new@example.com (changed!)
```

**What happens:**
- The method receives a pointer to the struct
- Changes affect the original struct
- No copying of the struct

**Reasoning:** Pointer receivers allow mutations and avoid copying. Go automatically takes the address for you when calling methods.

---

## Choosing Between Value and Pointer Receivers

### Use Pointer Receiver When:

1. **Method needs to modify the receiver:**
```go
func (u *User) Activate() {
    u.Status = true
}
```

2. **Struct is large (avoid copying):**
```go
type LargeStruct struct {
    data [1000000]int
}

func (ls *LargeStruct) Process() {
    // Avoid copying 8MB
}
```

3. **Consistency - if ANY method uses pointer, ALL should:**
```go
type Counter struct {
    count int
}

// All pointer receivers for consistency
func (c *Counter) Increment() { c.count++ }
func (c *Counter) Get() int { return c.count }
```

### Use Value Receiver When:

1. **Struct is small:**
```go
type Point struct {
    X, Y int
}

func (p Point) Distance() float64 {
    return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}
```

2. **Method doesn't modify receiver:**
```go
func (u User) DisplayName() string {
    return u.Name
}
```

3. **Type is not a struct (built-in types, type aliases):**
```go
type MyInt int

func (m MyInt) Double() MyInt {
    return m * 2
}
```

**Rule of Thumb:** When in doubt, use pointer receivers. They're more flexible and avoid accidental copying.

---

## Automatic Dereferencing

Go provides convenient syntax:

```go
user := User{Name: "Alice"}
userPtr := &User{Name: "Bob"}

// Both work the same way:
user.GetStatus()     // Value
userPtr.GetStatus()  // Pointer (Go auto-dereferences)

// Go handles these conversions:
// (&user).Method()  - Takes address automatically
// (*userPtr).Method() - Dereferences automatically
```

**Reasoning:** Go's syntactic sugar makes working with methods clean. You don't have to think about whether you have a value or pointer - Go does the right thing.

---

## Methods on Non-Struct Types

You can define methods on ANY type you define:

### Custom Integer Type

```go
type Celsius float64

func (c Celsius) ToFahrenheit() Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}

type Fahrenheit float64

func (f Fahrenheit) ToCelsius() Celsius {
    return Celsius((f - 32) * 5 / 9)
}

// Usage
temp := Celsius(100)
fmt.Println(temp.ToFahrenheit())  // 212
```

**Reasoning:** Defining methods on custom types makes code expressive and type-safe. You can't accidentally mix Celsius and Fahrenheit.

### Slice Type

```go
type IntSlice []int

func (is IntSlice) Sum() int {
    total := 0
    for _, v := range is {
        total += v
    }
    return total
}

// Usage
numbers := IntSlice{1, 2, 3, 4, 5}
fmt.Println(numbers.Sum())  // 15
```

---

## Method Sets and Interfaces

**Important Rule:**
- A type with value receiver implements the interface
- A type with pointer receiver does NOT automatically implement for values

```go
type Speaker interface {
    Speak() string
}

type Dog struct {
    Name string
}

// Pointer receiver
func (d *Dog) Speak() string {
    return "Woof! I'm " + d.Name
}

// Usage
var s Speaker

// ✅ OK: Pointer implements interface
dog := &Dog{Name: "Buddy"}
s = dog

// ❌ ERROR: Value doesn't implement (method has pointer receiver)
// dog2 := Dog{Name: "Max"}
// s = dog2  // Compile error
```

**Reasoning:** If a method modifies the receiver, only pointers should satisfy the interface (to ensure consistent behavior). This prevents subtle bugs.

---

## Method vs Function

### When to Use Methods:

```go
// Method - belongs to User
func (u *User) Activate() {
    u.Status = true
}
```

### When to Use Functions:

```go
// Function - operates on any user
func ValidateEmail(email string) bool {
    return strings.Contains(email, "@")
}
```

**Guidelines:**
- Use methods when behavior belongs to a type
- Use functions for general utilities
- Use methods for interface implementation

---

## Chaining Methods

```go
type StringBuilder struct {
    str string
}

func (sb *StringBuilder) Append(s string) *StringBuilder {
    sb.str += s
    return sb
}

func (sb *StringBuilder) String() string {
    return sb.str
}

// Usage
result := (&StringBuilder{}).
    Append("Hello").
    Append(" ").
    Append("World").
    String()

fmt.Println(result)  // Hello World
```

**Reasoning:** Returning the receiver enables method chaining (fluent interface). Common in builders and configuration APIs.

---

## Summary

**Key Takeaways:**
- Methods are functions with a receiver
- Value receiver = copy (safe, but can't modify)
- Pointer receiver = reference (can modify, efficient)
- Go auto-converts between value and pointer
- Use pointer receivers by default
- Methods enable interface implementation
- Can define methods on any custom type
- Explicit receiver (no implicit this/self)

**Quick Reference:**
```go
// Value receiver
func (t Type) Method() { }

// Pointer receiver
func (t *Type) Method() { }

// Usage
value.Method()   // Works
pointer.Method() // Also works (Go auto-handles)
```

**When to Use:**
- **Pointer receiver:** Method modifies receiver, large struct, or for consistency
- **Value receiver:** Small struct, immutable operation

---

**📝 Last Updated:** Object-Oriented Patterns
**➡️ Next Topic:** [Defer, Panic, Recover](./09_defer.md)
**🔗 Example Code:** [Concepts/16methods](../Concepts/16methods/)
