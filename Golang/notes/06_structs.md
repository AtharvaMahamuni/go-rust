# Structs in Go - Custom Types and Composition

**📅 Created:** Data Types Series
**🏷️ Topics:** Structs, Custom Types, Composition, Embedding, struct tags
**🔗 Related:** [05_maps.md](./05_maps.md), [08_methods.md](./08_methods.md), [02_pointers.md](./02_pointers.md), [Concepts/11mystructs](../Concepts/11mystructs/)

---

## Overview

Structs are Go's way of creating custom data types by grouping fields together. They're the foundation of object-oriented-like programming in Go, but without classes or inheritance. Instead, Go uses composition and interfaces. Structs are value types, making them efficient and predictable.

**Why Structs Matter:**
- Create custom types that model your domain
- Group related data together
- Foundation for methods (functions on types)
- Composition over inheritance
- Zero allocation for stack-based usage

**Reasoning:** Go deliberately avoids inheritance complexity. Structs + composition + interfaces provide all the benefits of OOP without the pitfalls of deep inheritance hierarchies.

---

## Defining Structs

### Basic Struct Definition

```go
type User struct {
    Name   string
    Email  string
    Status bool
    Age    int
}
```

**Key Points:**
- `type` keyword creates a new type
- `struct` keyword defines it as a struct
- Fields have names and types
- Capital first letter = exported (public)
- Small first letter = unexported (private to package)

**Reasoning:** The capitalization visibility rule is consistent across all Go identifiers. It's simple and works at the package level, not file level like many languages.

### Creating Struct Instances

#### Method 1: Struct Literal (All Fields)

```go
atharva := User{
    Name:   "Atharva",
    Email:  "atharva@example.com",
    Status: true,
    Age:    23,
}
```

**Why this way:**
- Explicit field names
- Order doesn't matter
- Self-documenting
- Safe if struct changes
- **Most common and recommended**

#### Method 2: Positional (Not Recommended)

```go
// Works but fragile
atharva := User{"Atharva", "atharva@example.com", true, 23}
```

**Why avoid:**
- Breaks if struct field order changes
- Not self-documenting
- Easy to mix up similar types

#### Method 3: Partial Initialization

```go
user := User{
    Name: "Alice",
    Age:  25,
}
// Email = "" (zero value)
// Status = false (zero value)
```

**Reasoning:** Unspecified fields get zero values. This is safe because Go guarantees initialized memory, unlike C/C++.

#### Method 4: Zero Value

```go
var user User
// All fields have zero values:
// Name: ""
// Email: ""
// Status: false
// Age: 0
```

**Reasoning:** Every type has a useful zero value. For structs, all fields are set to their zero values, creating a valid (if empty) instance.

### Using `new`

```go
userPtr := new(User)
// Returns *User pointing to zero-valued User

userPtr.Name = "Bob"
fmt.Println(userPtr.Name)  // Bob
```

**Reasoning:** `new(T)` allocates a zero-valued T and returns a pointer to it. Useful when you need a pointer from the start, but struct literals with `&` are more common.

---

## Accessing Fields

### Dot Notation

```go
user := User{
    Name: "Atharva",
    Age:  23,
}

// Access fields
fmt.Println(user.Name)  // Atharva
fmt.Println(user.Age)   // 23

// Modify fields
user.Age = 24
fmt.Println(user.Age)   // 24
```

### With Pointers (Automatic Dereferencing)

```go
user := &User{Name: "Alice"}

// Go automatically dereferences
user.Name = "Bob"  // Shorthand

// Equivalent to:
(*user).Name = "Bob"
```

**Reasoning:** Go's syntactic sugar makes working with struct pointers clean. You don't need to explicitly dereference with `*` - Go does it automatically for field access.

---

## Anonymous Structs

**Use for one-off, local data structures:**

```go
// Declare and use immediately
person := struct {
    name string
    age  int
}{
    name: "Alice",
    age:  25,
}

fmt.Println(person.name)  // Alice
```

**When to use:**
- Temporary data structures
- JSON unmarshaling for simple cases
- Test data
- Configuration

**Reasoning:** Anonymous structs avoid polluting the namespace with one-time-use types. They're perfect for local scope where the structure is obvious from context.

**Common Pattern - Table-Driven Tests:**

```go
tests := []struct {
    input    int
    expected int
}{
    {2, 4},
    {3, 9},
    {4, 16},
}

for _, test := range tests {
    result := square(test.input)
    if result != test.expected {
        t.Errorf("Expected %d, got %d", test.expected, result)
    }
}
```

---

## Structs Are Value Types

```go
user1 := User{Name: "Alice", Age: 25}
user2 := user1  // Copy, not reference

user2.Name = "Bob"

fmt.Println(user1.Name)  // Alice (unchanged)
fmt.Println(user2.Name)  // Bob
```

**Reasoning:** Structs are value types - assignment copies all fields. This is different from maps and slices (which are reference types). Value semantics make code easier to reason about - no hidden sharing.

**When copying is expensive:**
```go
// For large structs, use pointers
func processUser(u *User) {
    // Work with pointer to avoid copying
}
```

---

## Composition (No Inheritance!)

### Embedding Structs

Go uses composition instead of inheritance:

```go
type User struct {
    Name  string
    Email string
}

type Admin struct {
    User        // Embedded struct (anonymous field)
    AccessLevel int
}

// Usage
admin := Admin{
    User: User{
        Name:  "Alice",
        Email: "alice@example.com",
    },
    AccessLevel: 10,
}

// Can access embedded fields directly
fmt.Println(admin.Name)  // Alice (promoted field)
fmt.Println(admin.Email) // alice@example.com

// Or explicitly
fmt.Println(admin.User.Name)  // Also Alice
```

**Reasoning:** Embedding promotes the embedded struct's fields to the outer struct. This provides composition - the Admin "has-a" User, not "is-a" User. No inheritance hierarchy, no fragile base class problems.

**Why composition over inheritance:**
1. **Clearer relationships** - explicit about what you're using
2. **No diamond problem** - can't happen with composition
3. **Flexible** - can compose multiple types
4. **Simple** - no super/parent/override keywords

### Multiple Embedding

```go
type Timestamp struct {
    CreatedAt time.Time
    UpdatedAt time.Time
}

type Address struct {
    Street string
    City   string
}

type User struct {
    Name      string
    Timestamp // Embedded
    Address   // Embedded
}

user := User{
    Name: "Alice",
    Timestamp: Timestamp{
        CreatedAt: time.Now(),
    },
    Address: Address{
        City: "NYC",
    },
}

// Access promoted fields
fmt.Println(user.CreatedAt)  // Direct access
fmt.Println(user.City)       // Direct access
```

**Reasoning:** You can embed multiple structs to compose behavior. Fields are promoted unless there's a name conflict, in which case you must use the explicit path.

---

## Comparing Structs

### Structs Can Be Compared

```go
type Point struct {
    X, Y int
}

p1 := Point{1, 2}
p2 := Point{1, 2}
p3 := Point{2, 3}

fmt.Println(p1 == p2)  // true
fmt.Println(p1 == p3)  // false
```

**Reasoning:** Structs are comparable if all their fields are comparable. Go compares field by field.

### Structs With Non-Comparable Fields

```go
type Container struct {
    Data []int  // Slices are not comparable
}

c1 := Container{Data: []int{1, 2, 3}}
c2 := Container{Data: []int{1, 2, 3}}

// fmt.Println(c1 == c2)  // ❌ Compile error
```

**Reasoning:** If any field is non-comparable (slice, map, function), the struct becomes non-comparable. Use `reflect.DeepEqual` or write custom comparison logic.

---

## Struct Tags

**Metadata for reflection and libraries:**

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age,omitempty"`
}
```

**Common uses:**
- JSON encoding/decoding
- Database ORM mapping
- Validation rules
- Form parsing

**Reasoning:** Struct tags don't affect the struct's behavior directly - they're metadata that libraries can read via reflection. This keeps behavior separate from declaration.

### JSON Example

```go
import "encoding/json"

type User struct {
    Name    string `json:"name"`
    Email   string `json:"email"`
    Age     int    `json:"age,omitempty"`
    private string // Not exported, won't be in JSON
}

user := User{Name: "Alice", Email: "alice@example.com", Age: 25}

// Marshal to JSON
jsonData, _ := json.Marshal(user)
fmt.Println(string(jsonData))
// {"name":"Alice","email":"alice@example.com","age":25}
```

---

## Empty Struct

```go
type Empty struct{}

var e Empty
fmt.Println(unsafe.Sizeof(e))  // 0 bytes
```

**Use cases:**

### 1. Signal-Only Channels

```go
done := make(chan struct{})

go func() {
    // Do work...
    done <- struct{}{}  // Signal completion
}()

<-done  // Wait for signal
```

**Reasoning:** `struct{}` uses zero bytes - perfect for signaling where you don't need to send data, just notification.

### 2. Set Implementation

```go
type Set map[string]struct{}

set := make(Set)
set["item"] = struct{}{}

if _, exists := set["item"]; exists {
    fmt.Println("Item in set")
}
```

**Reasoning:** `map[string]bool` would use 1 byte per value. `map[string]struct{}` uses 0 bytes - more efficient for large sets.

---

## Common Patterns

### Pattern 1: Constructor Function

```go
func NewUser(name, email string) *User {
    return &User{
        Name:   name,
        Email:  email,
        Status: true,  // Default value
    }
}

// Usage
user := NewUser("Alice", "alice@example.com")
```

**Reasoning:** Constructor functions (conventionally named `New` or `NewType`) encapsulate creation logic, set defaults, and perform validation. Returning a pointer is common for larger structs.

### Pattern 2: Builder Pattern

```go
type UserBuilder struct {
    user User
}

func (b *UserBuilder) WithName(name string) *UserBuilder {
    b.user.Name = name
    return b
}

func (b *UserBuilder) WithEmail(email string) *UserBuilder {
    b.user.Email = email
    return b
}

func (b *UserBuilder) Build() User {
    return b.user
}

// Usage
user := (&UserBuilder{}).
    WithName("Alice").
    WithEmail("alice@example.com").
    Build()
```

**Reasoning:** Builder pattern provides a fluent API for complex construction. Useful when you have many optional fields.

### Pattern 3: Option Functions

```go
type User struct {
    Name  string
    Email string
    Age   int
}

type UserOption func(*User)

func WithAge(age int) UserOption {
    return func(u *User) {
        u.Age = age
    }
}

func NewUser(name, email string, opts ...UserOption) *User {
    u := &User{Name: name, Email: email}
    for _, opt := range opts {
        opt(u)
    }
    return u
}

// Usage
user := NewUser("Alice", "alice@example.com", WithAge(25))
```

**Reasoning:** Functional options pattern provides flexibility for optional parameters without overloading constructors. Extensible and readable.

---

## Structs vs Other Languages

### Go vs Java/Kotlin

**Java/Kotlin:**
```kotlin
class User(
    val name: String,
    val email: String
) {
    var status: Boolean = true
}
```

**Go:**
```go
type User struct {
    Name   string
    Email  string
    Status bool
}
```

**Key Differences:**
- Go: No classes, just structs + methods
- Go: No inheritance, use composition
- Go: No constructors, use factory functions
- Go: No this/self keyword in methods
- Go: Simpler, more explicit

---

## Summary

**Key Takeaways:**
- Structs group related data into custom types
- Value types (copying creates independent copies)
- Use composition via embedding (not inheritance)
- Fields with capital letter are exported (public)
- Zero value is a struct with all fields zero-valued
- Comparable if all fields are comparable
- Struct tags provide metadata for libraries
- Empty struct uses zero bytes
- Constructor functions for encapsulation
- Pointers avoid copying large structs

**Quick Reference:**
```go
// Define
type User struct {
    Name string
    Age  int
}

// Create
user := User{Name: "Alice", Age: 25}
user := User{"Alice", 25}  // Positional
userPtr := &User{Name: "Alice"}
userPtr := new(User)

// Access
user.Name = "Bob"
userPtr.Name = "Bob"  // Auto-dereference

// Embedding
type Admin struct {
    User        // Embed
    AccessLevel int
}
```

**When to Use This Note:**
- Defining custom types
- Understanding composition vs inheritance
- Working with struct pointers
- Struct tags for JSON/DB mapping
- Memory efficiency considerations

---

**📝 Last Updated:** Data Types Series
**➡️ Next Topic:** [Control Flow](./07_control-flow.md)
**🔗 Example Code:** [Concepts/11mystructs](../Concepts/11mystructs/)
