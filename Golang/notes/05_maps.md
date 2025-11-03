# Maps in Go - Key-Value Collections

**📅 Created:** Collections Series
**🏷️ Topics:** Maps, Hash Tables, Key-Value Pairs, make, delete, range
**🔗 Related:** [03_arrays-and-slices.md](./03_arrays-and-slices.md), [06_structs.md](./06_structs.md), [Concepts/10mymaps](../Concepts/10mymaps/)

---

## Overview

Maps are Go's built-in hash table implementation - collections that store key-value pairs. They provide fast lookups, additions, and deletions. Maps are reference types (like slices) and are one of the most commonly used data structures in Go for organizing and retrieving data efficiently.

**Why Maps Matter:**
- Fast O(1) average-case lookups
- Perfect for caching, counting, grouping
- Dynamic sizing (grow as needed)
- Type-safe with compile-time checking

---

## What is a Map?

**Simple Definition:**
A map is an unordered collection of key-value pairs where each key is unique.

```
Map Structure:
┌─────────┬──────────────┐
│  Key    │    Value     │
├─────────┼──────────────┤
│  "js"   │ "JavaScript" │
│  "py"   │ "Python"     │
│  "go"   │ "Golang"     │
└─────────┴──────────────┘
```

**Key Characteristics:**
- Keys must be unique
- Keys must be comparable (==, !=)
- Unordered (iteration order not guaranteed)
- Reference type (like slices)
- Zero value is `nil`

---

## Creating Maps

### Method 1: Using `make`

```go
// Create empty map
languages := make(map[string]string)

// Add key-value pairs
languages["js"] = "JavaScript"
languages["py"] = "Python"
languages["go"] = "Golang"

fmt.Println(languages)
// map[go:Golang js:JavaScript py:Python]
```

**Why use make:**
- Initialize empty map ready to use
- Can optionally specify initial capacity
- Most common way to create maps

### Method 2: Map Literal

```go
// Create and initialize in one step
languages := map[string]string{
    "js": "JavaScript",
    "py": "Python",
    "go": "Golang",
}

fmt.Println(languages)
```

**Why use literals:**
- Concise when you know initial values
- Clear and readable
- Common in configuration and setup code

### Method 3: With Capacity Hint

```go
// Pre-allocate space for better performance
languages := make(map[string]string, 10)
```

**Why specify capacity:**
- Reduces allocations when you know approximate size
- Improves performance for large maps
- Still grows automatically if needed

### Method 4: Nil Map (Avoid!)

```go
var languages map[string]string  // nil map

// ❌ PANIC: assignment to nil map
// languages["go"] = "Golang"

// ✅ Must initialize first
languages = make(map[string]string)
languages["go"] = "Golang"
```

**Why nil maps are tricky:**
- Can read from nil map (returns zero value)
- Cannot write to nil map (panics)
- Always use `make` to initialize

---

## Basic Operations

### Adding/Updating Elements

```go
languages := make(map[string]string)

// Add new key-value pair
languages["go"] = "Golang"

// Update existing value
languages["go"] = "Go Language"

fmt.Println(languages["go"])  // Go Language
```

**Reasoning:** Same syntax for add and update - Go checks if key exists and acts accordingly. This simplicity makes maps easy to use.

### Accessing Elements

```go
languages := map[string]string{
    "js": "JavaScript",
    "py": "Python",
}

// Access existing key
fmt.Println(languages["js"])  // JavaScript

// Access non-existent key
fmt.Println(languages["ruby"])  // "" (zero value)
```

**Reasoning:** Accessing non-existent keys returns the zero value (not an error). This is safe but can be misleading - use the comma-ok idiom to check existence.

### Checking Existence (Comma-Ok Idiom)

```go
languages := map[string]string{
    "go": "Golang",
}

// Check if key exists
value, exists := languages["go"]
if exists {
    fmt.Println("Found:", value)  // Found: Golang
} else {
    fmt.Println("Not found")
}

// Check non-existent key
value, exists = languages["ruby"]
fmt.Println(value, exists)  // "" false
```

**Reasoning:** The comma-ok idiom (value, exists) distinguishes between "key exists with zero value" and "key doesn't exist". Critical for correct logic.

**Common Pattern:**
```go
if value, ok := myMap[key]; ok {
    // Key exists, use value
    fmt.Println(value)
} else {
    // Key doesn't exist
    fmt.Println("Key not found")
}
```

### Deleting Elements

```go
languages := map[string]string{
    "js": "JavaScript",
    "py": "Python",
    "go": "Golang",
}

// Delete a key
delete(languages, "py")

fmt.Println(languages)
// map[go:Golang js:JavaScript]

// Deleting non-existent key is safe (no-op)
delete(languages, "ruby")  // No error
```

**Reasoning:** `delete` is a built-in function, not a method. Deleting non-existent keys is safe (no panic), making it forgiving and easy to use.

---

## Iterating Over Maps

### Using `range`

```go
languages := map[string]string{
    "js": "JavaScript",
    "py": "Python",
    "go": "Golang",
}

// Iterate over key-value pairs
for key, value := range languages {
    fmt.Printf("Key: %s, Value: %s\n", key, value)
}

// Output (order not guaranteed):
// Key: js, Value: JavaScript
// Key: py, Value: Python
// Key: go, Value: Golang
```

**Reasoning:** `range` on maps gives you both key and value. Order is intentionally randomized by Go to prevent reliance on iteration order.

### Iterate Over Keys Only

```go
for key := range languages {
    fmt.Println(key)
}
```

### Iterate Over Values Only

```go
for _, value := range languages {
    fmt.Println(value)
}
```

**Reasoning:** Use `_` to ignore unwanted parts (key or value). This is idiomatic Go for expressing "I don't need this".

---

## Map Properties

### Maps Are Unordered

```go
m := map[string]int{
    "a": 1,
    "b": 2,
    "c": 3,
}

// Order is NOT guaranteed
for k, v := range m {
    fmt.Println(k, v)
}
```

**Reasoning:** Go intentionally randomizes map iteration order to prevent code from depending on it. Hash tables are inherently unordered - Go makes this explicit.

**If you need order:**
- Store keys in a slice and sort them
- Use a package like `orderedmap`

### Maps Are Reference Types

```go
original := map[string]int{"a": 1}
copy := original

copy["a"] = 999

fmt.Println(original["a"])  // 999 (modified!)
fmt.Println(copy["a"])      // 999
```

**Reasoning:** Maps are reference types (like slices). Assignment creates a reference, not a copy. Both variables point to the same underlying data.

**To make a real copy:**
```go
original := map[string]int{"a": 1, "b": 2}
copy := make(map[string]int)

for k, v := range original {
    copy[k] = v
}

copy["a"] = 999
fmt.Println(original["a"])  // 1 (unchanged)
```

### Nil Map vs Empty Map

```go
var nilMap map[string]int        // nil map
emptyMap := make(map[string]int) // empty map

fmt.Println(nilMap == nil)    // true
fmt.Println(emptyMap == nil)  // false

// Both have length 0
fmt.Println(len(nilMap))      // 0
fmt.Println(len(emptyMap))    // 0

// Can read from both
fmt.Println(nilMap["key"])    // 0 (no panic)
fmt.Println(emptyMap["key"])  // 0

// ❌ Cannot write to nil map
// nilMap["key"] = 1  // PANIC

// ✅ Can write to empty map
emptyMap["key"] = 1  // OK
```

**Reasoning:** Nil maps allow reads (return zero value) but not writes (panic). Always initialize with `make` before writing.

---

## Common Patterns

### Pattern 1: Counting Occurrences

```go
words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}

// Count frequency
wordCount := make(map[string]int)

for _, word := range words {
    wordCount[word]++  // Zero value (0) + 1 on first occurrence
}

fmt.Println(wordCount)
// map[apple:3 banana:2 cherry:1]
```

**Reasoning:** Leveraging zero values makes counting elegant - no need to check if key exists first.

### Pattern 2: Grouping Data

```go
type Person struct {
    Name string
    City string
}

people := []Person{
    {"Alice", "NYC"},
    {"Bob", "LA"},
    {"Charlie", "NYC"},
}

// Group by city
byCity := make(map[string][]Person)

for _, person := range people {
    byCity[person.City] = append(byCity[person.City], person)
}

fmt.Println(byCity["NYC"])
// [{Alice NYC} {Charlie NYC}]
```

**Reasoning:** Maps with slice values are perfect for grouping. The zero value of a slice is `nil`, which `append` handles correctly.

### Pattern 3: Set (Unique Elements)

```go
// Use map[Type]bool or map[Type]struct{}
visited := make(map[string]bool)

// Add to set
visited["page1"] = true
visited["page2"] = true
visited["page1"] = true  // Duplicate, no effect

// Check membership
if visited["page1"] {
    fmt.Println("Already visited")
}

// Iterate over set
for page := range visited {
    fmt.Println(page)
}
```

**Reasoning:** Maps provide O(1) lookups for checking membership. Using `bool` is simple; using `struct{}` saves memory (empty struct uses 0 bytes).

### Pattern 4: Caching/Memoization

```go
cache := make(map[int]int)

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }

    // Check cache
    if val, exists := cache[n]; exists {
        return val
    }

    // Compute and cache
    result := fibonacci(n-1) + fibonacci(n-2)
    cache[n] = result
    return result
}
```

**Reasoning:** Maps excel at caching - O(1) lookups dramatically improve performance for expensive computations.

### Pattern 5: Default Values

```go
config := map[string]int{
    "timeout": 30,
    "retries": 3,
}

// Get with default if not found
timeout := config["timeout"]
if timeout == 0 {
    timeout = 60  // default
}

// Or more elegantly
func getOrDefault(m map[string]int, key string, defaultVal int) int {
    if val, ok := m[key]; ok {
        return val
    }
    return defaultVal
}

timeout = getOrDefault(config, "timeout", 60)
```

**Reasoning:** Since zero values are returned for missing keys, you need explicit checking for true "not found" vs "stored zero value" cases.

---

## Map Key Restrictions

### Valid Key Types

Keys must support `==` and `!=` operators:

```go
// ✅ Valid key types
map[string]int          // strings
map[int]string          // integers
map[float64]bool        // floats
map[bool]int            // bools
map[rune]string         // runes
map[[3]int]string       // arrays (fixed size)
map[struct{x,y int}]bool // structs (if all fields comparable)
```

### Invalid Key Types

```go
// ❌ Invalid key types (not comparable)
map[[]int]string        // slices
map[map[string]int]bool // maps
map[func()]int          // functions
```

**Reasoning:** Hash tables need to compute a hash and check equality. Only comparable types support this. Slices, maps, and functions don't have defined equality.

**Workaround for slice keys:**
```go
// Convert slice to string representation
key := fmt.Sprintf("%v", mySlice)
myMap[key] = value
```

---

## Length and Capacity

```go
m := map[string]int{
    "a": 1,
    "b": 2,
}

// Get number of key-value pairs
fmt.Println(len(m))  // 2

// ❌ No cap() for maps
// fmt.Println(cap(m))  // Compile error
```

**Reasoning:** Maps grow dynamically as needed. Unlike slices, there's no exposed capacity concept - the internal implementation handles resizing automatically.

---

## Thread Safety

**Maps are NOT thread-safe!**

```go
var m = make(map[string]int)

// ❌ Race condition with multiple goroutines
go func() { m["key"] = 1 }()
go func() { m["key"] = 2 }()
```

**Solutions:**

### Option 1: Use sync.Mutex

```go
var m = make(map[string]int)
var mu sync.Mutex

// Safe write
mu.Lock()
m["key"] = 1
mu.Unlock()

// Safe read
mu.Lock()
val := m["key"]
mu.Unlock()
```

### Option 2: Use sync.Map (for specific use cases)

```go
var m sync.Map

// Store
m.Store("key", 1)

// Load
val, ok := m.Load("key")

// Delete
m.Delete("key")
```

**Reasoning:** Regular maps optimize for single-threaded performance. For concurrent access, use explicit synchronization (mutex) or the specialized `sync.Map` for concurrent read-heavy workloads.

---

## Performance Considerations

### Pre-allocate When Possible

```go
// Bad: Grows multiple times
m := make(map[string]int)
for i := 0; i < 10000; i++ {
    m[fmt.Sprintf("key%d", i)] = i
}

// Better: Pre-allocate
m := make(map[string]int, 10000)
for i := 0; i < 10000; i++ {
    m[fmt.Sprintf("key%d", i)] = i
}
```

**Reasoning:** Pre-allocating reduces the number of internal resizing operations, improving performance for large maps.

### Choose Good Key Types

```go
// Good: Simple, fast hash
map[string]int
map[int]string

// Less efficient: Complex struct
type Key struct {
    field1 string
    field2 string
    field3 int
}
map[Key]int
```

**Reasoning:** Simpler keys hash faster. Complex keys (large structs) slow down operations.

---

## Common Pitfalls

### Pitfall 1: Writing to Nil Map

```go
var m map[string]int
m["key"] = 1  // ❌ PANIC: assignment to nil map
```

**Fix:**
```go
m := make(map[string]int)
m["key"] = 1  // ✅ OK
```

### Pitfall 2: Assuming Order

```go
m := map[int]string{1: "one", 2: "two", 3: "three"}

for k, v := range m {
    fmt.Println(k, v)
}
// ❌ Order is NOT 1, 2, 3
```

**Fix:** Store keys in slice and sort if order matters.

### Pitfall 3: Comparing Maps

```go
m1 := map[string]int{"a": 1}
m2 := map[string]int{"a": 1}

// ❌ Cannot compare maps with ==
// fmt.Println(m1 == m2)  // Compile error
```

**Fix:** Compare manually or use reflect.DeepEqual.

### Pitfall 4: Concurrent Access

```go
m := make(map[string]int)

go func() { m["key"] = 1 }()
go func() { m["key"] = 2 }()
// ❌ Race condition
```

**Fix:** Use mutex or sync.Map.

---

## Summary

**Key Takeaways:**
- Maps are built-in hash tables (key-value pairs)
- Create with `make(map[KeyType]ValueType)` or literals
- Fast O(1) average-case operations
- Keys must be comparable (support ==)
- Unordered (iteration order not guaranteed)
- Reference types (assignment doesn't copy)
- Use comma-ok idiom to check existence
- Zero value is nil (can't write to nil map)
- NOT thread-safe (use mutex or sync.Map)
- Perfect for: counting, grouping, caching, sets

**Quick Reference:**
```go
// Create
m := make(map[string]int)
m := map[string]int{"a": 1}

// Operations
m[key] = value           // Add/update
value := m[key]          // Get
value, ok := m[key]      // Check existence
delete(m, key)           // Delete
len(m)                   // Size

// Iterate
for k, v := range m { }
```

**When to Use Maps:**
- Lookups by key
- Counting occurrences
- Grouping data
- Caching results
- Implementing sets
- Configuration storage

**When to Use This Note:**
- Learning about hash tables in Go
- Choosing between map, slice, array
- Understanding reference types
- Debugging map-related issues
- Optimizing map performance

---

**📝 Last Updated:** Collections Series
**➡️ Next Topic:** [Structs and Methods](./06_structs.md)
**🔗 Example Code:** [Concepts/10mymaps](../Concepts/10mymaps/)
