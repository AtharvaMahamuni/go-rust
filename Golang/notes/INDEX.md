# Go Learning Notes - Index

This directory contains comprehensive reference notes for learning Go. All notes are organized by topic with links to example code in the Concepts folder.

---

## Quick Navigation

| Category | Topic | Description |
|----------|-------|-------------|
| **Fundamentals** | [Go Basics](#go-basics) | Introduction, compilation, project setup |
| **Core Concepts** | [Variables & Types](#variables-and-types) | Type system, declarations, zero values |
| **Memory** | [Pointers](#pointers) | Memory addresses, references, pass-by-pointer |
| **Collections** | [Arrays & Slices](#arrays-and-slices) | Fixed and dynamic collections |
| **Concurrency** | [Goroutines](#goroutines-and-concurrency) | Concurrent programming, channels, sync |

---

## Go Basics

### [00_golang-basics.md](./00_golang-basics.md)
**Core Topic:** Go Introduction, Compilation, Modules, Project Setup

**Key Concepts Covered:**
- What makes Go unique (compiled, fast, simple)
- Go's philosophy (no inheritance, explicit errors, simplicity)
- Go toolchain: rustup, rustc, cargo equivalents
- Project structure with go mod
- Native compilation vs bytecode
- Cross-compilation for different OS
- go commands (run, build, test, fmt)

**When to Read:** Starting point for all Go learners

**Quick Reference:**
```bash
go mod init project    # Initialize project
go run main.go        # Run quickly
go build              # Create executable
go fmt ./...          # Format code
```

**🔗 Example Code:** [Concepts/01hello](../Concepts/01hello/)

---

## Variables and Types

### [01_variables-and-types.md](./01_variables-and-types.md)
**Core Topic:** Type System, Variable Declarations, Zero Values

**Key Concepts Covered:**
- Static typing with type safety
- Three declaration styles (`var`, type inference, `:=`)
- Basic types (string, int, float, bool, complex)
- Type aliases (byte, rune)
- Zero values (no uninitialized variables)
- No implicit type conversions
- Constants and type conversion
- fmt.Printf format verbs

**When to Read:** After basics, essential foundation

**Quick Reference:**
```go
var x int = 10        // Explicit
var x = 10            // Inferred
x := 10               // Short (functions only)
```

**🔗 Example Code:** [Concepts/02variables](../Concepts/02variables/), [Concepts/04conversion](../Concepts/04conversion/)

---

## Pointers

### [02_pointers.md](./02_pointers.md)
**Core Topic:** Memory Addresses, Dereferencing, Pass-by-Reference

**Key Concepts Covered:**
- What pointers are and why they matter
- `&` (address-of) and `*` (dereference) operators
- Pointer declarations and nil pointers
- Pass by value vs pass by pointer
- When to use pointers (large structs, modification)
- Pointers vs C/C++ (no arithmetic, safer)
- Pointers vs Java/Kotlin (explicit vs implicit)
- Common patterns and pitfalls

**When to Read:** After types, before structs

**Quick Reference:**
```go
ptr := &variable      // Get address
value := *ptr         // Dereference
*ptr = newValue       // Modify through pointer
```

**🔗 Example Code:** [Concepts/07pointer](../Concepts/07pointer/)

---

## Arrays and Slices

### [03_arrays-and-slices.md](./03_arrays-and-slices.md)
**Core Topic:** Fixed vs Dynamic Collections, Slice Internals

**Key Concepts Covered:**
- Arrays: fixed-size, rarely used
- Slices: dynamic, most common collection
- Slice internals (pointer, length, capacity)
- Creating slices (`make`, literals, from arrays)
- append, copy, and slice operations
- Sorting with sort package
- make() vs new()
- Common gotchas (shared underlying array)
- Performance patterns (pre-allocate capacity)

**When to Read:** Essential for working with collections

**Quick Reference:**
```go
// Arrays (fixed)
arr := [5]int{1, 2, 3, 4, 5}

// Slices (dynamic)
slice := []int{1, 2, 3}
slice = append(slice, 4)
slice = slice[1:3]
```

**🔗 Example Code:** [Concepts/08myarray](../Concepts/08myarray/), [Concepts/09myslices](../Concepts/09myslices/)

---

## Maps

### [05_maps.md](./05_maps.md)
**Core Topic:** Key-Value Collections, Hash Tables

**Key Concepts Covered:**
- Creating maps with make and literals
- Adding, accessing, deleting elements
- Comma-ok idiom for checking existence
- Iterating with range (unordered)
- Maps are reference types
- Nil vs empty maps
- Thread safety (not thread-safe)
- Common patterns (counting, grouping, sets, caching)
- Valid vs invalid key types

**When to Read:** Essential for working with key-value data

**Quick Reference:**
```go
m := make(map[string]int)
m["key"] = value
value, ok := m["key"]
delete(m, "key")
for k, v := range m { }
```

**🔗 Example Code:** [Concepts/10mymaps](../Concepts/10mymaps/)

---

## Structs

### [06_structs.md](./06_structs.md)
**Core Topic:** Custom Types, Composition, Embedding

**Key Concepts Covered:**
- Defining and creating structs
- Field visibility (capital = public)
- Struct literals vs new()
- Anonymous structs
- Structs are value types
- Composition via embedding (no inheritance!)
- Comparing structs
- Struct tags (JSON, DB mapping)
- Empty struct (zero bytes)
- Constructor patterns

**When to Read:** Foundation for object-oriented patterns

**Quick Reference:**
```go
type User struct {
    Name string
    Age  int
}

user := User{Name: "Alice", Age: 25}
userPtr := &User{Name: "Bob"}
```

**🔗 Example Code:** [Concepts/11mystructs](../Concepts/11mystructs/)

---

## Functions

### [07_functions.md](./07_functions.md)
**Core Topic:** Functions, Parameters, Return Values, Closures

**Key Concepts Covered:**
- Function syntax and declarations
- Multiple return values (value, error pattern)
- Named return values
- Variadic functions (...args)
- Anonymous functions and closures
- Functions as first-class values
- Function types
- Error handling patterns
- Option functions pattern

**When to Read:** Core language feature

**Quick Reference:**
```go
func add(a, b int) int { return a + b }
func divide(a, b float64) (float64, error) { }
func sum(nums ...int) int { }
```

**🔗 Example Code:** [Concepts/15functions](../Concepts/15functions/)

---

## Methods

### [08_methods.md](./08_methods.md)
**Core Topic:** Methods, Receivers, Value vs Pointer

**Key Concepts Covered:**
- Defining methods with receivers
- Value receivers (copy, immutable)
- Pointer receivers (reference, mutable)
- Choosing between value and pointer
- Automatic dereferencing
- Methods on non-struct types
- Method sets and interfaces
- Method chaining

**When to Read:** After structs and functions

**Quick Reference:**
```go
func (u User) GetName() string { }      // Value receiver
func (u *User) SetName(n string) { }    // Pointer receiver
```

**🔗 Example Code:** [Concepts/16methods](../Concepts/16methods/)

---

## Defer, Panic, Recover

### [09_defer.md](./09_defer.md)
**Core Topic:** Cleanup, Error Recovery, Control Flow

**Key Concepts Covered:**
- Defer for cleanup (LIFO execution)
- Defer with named returns
- Arguments evaluated immediately
- Panic for unrecoverable errors
- Recover to catch panics (only in defer)
- Common patterns (file cleanup, timing, recovery)
- When to use vs error returns

**When to Read:** Essential for resource management

**Quick Reference:**
```go
defer file.Close()
defer mu.Unlock()

defer func() {
    if r := recover(); r != nil {
        // Handle panic
    }
}()
```

**🔗 Example Code:** [Concepts/17defer](../Concepts/17defer/)

---

## Control Flow

### [10_control-flow.md](./10_control-flow.md)
**Core Topic:** if/else, switch, loops, range

**Key Concepts Covered:**
- If with short statement
- Switch (auto-break, no fallthrough)
- Switch with expressions and type switching
- For loops (only loop in Go)
- While-style loops (for with condition)
- Range over slices, maps, strings
- Break, continue, labeled breaks

**When to Read:** Basic language constructs

**Quick Reference:**
```go
if x := getValue(); x > 0 { }
switch value { case 1: }
for i := 0; i < 10; i++ { }
for _, v := range slice { }
```

**🔗 Example Code:** [Concepts/12ifelse](../Concepts/12ifelse/), [Concepts/13switchcase](../Concepts/13switchcase/), [Concepts/14loops](../Concepts/14loops/)

---

## Goroutines and Concurrency

### [04_goroutines-and-concurrency.md](./04_goroutines-and-concurrency.md)
**Core Topic:** Concurrent Programming, Channels, Synchronization

**Key Concepts Covered:**
- What are goroutines (lightweight threads)
- Starting goroutines with `go` keyword
- sync.WaitGroup for coordination
- Race conditions and sync.Mutex
- Channels for communication
- Buffered vs unbuffered channels
- Channel directions (send-only, receive-only)
- Closing channels
- Select statement
- Worker pool pattern
- "Share memory by communicating"

**When to Read:** After basics, Go's most powerful feature

**Quick Reference:**
```go
// Goroutines
go funcName()

// WaitGroup
wg.Add(1)
defer wg.Done()
wg.Wait()

// Channels
ch := make(chan int)
ch <- value
value := <-ch
```

**🔗 Example Code:** [Concepts/26goroutines](../Concepts/26goroutines/), [Concepts/27mutexAndAwaitGroups](../Concepts/27mutexAndAwaitGroups/), [Concepts/28channels](../Concepts/28channels/)

---

## User Input and Conversion

### [11_user-input-and-conversion.md](./11_user-input-and-conversion.md)
**Core Topic:** Reading Input, Type Conversion, String Handling

**Key Concepts Covered:**
- Reading user input with bufio
- `os.Stdin` and `bufio.NewReader`
- Comma-err pattern for error handling
- String manipulation with strings package
- Type conversion with strconv (Atoi, ParseFloat)
- Number to string conversion
- Input validation patterns
- Explicit vs implicit conversion

**When to Read:** Essential for CLI applications

**Quick Reference:**
```go
reader := bufio.NewReader(os.Stdin)
input, _ := reader.ReadString('\n')
clean := strings.TrimSpace(input)
num, err := strconv.Atoi(clean)
```

**🔗 Example Code:** [Concepts/03userinput](../Concepts/03userinput/), [Concepts/04conversion](../Concepts/04conversion/)

---

## Time Handling

### [12_time-handling.md](./12_time-handling.md)
**Core Topic:** Date/Time Operations, Formatting, Parsing

**Key Concepts Covered:**
- Getting current time with `time.Now()`
- Creating specific times with `time.Date`
- Go's unique formatting with reference time
- Parsing time strings
- Duration type and arithmetic
- Time comparison methods
- Sleep, timers, and tickers
- Timezone handling
- Common patterns (timestamps, measuring execution)

**When to Read:** Working with dates, timestamps, scheduling

**Quick Reference:**
```go
now := time.Now()
now.Format("2006-01-02 15:04:05")
time.Parse("2006-01-02", "2025-11-02")
duration := 2 * time.Hour + 30 * time.Minute
```

**🔗 Example Code:** [Concepts/06mytime](../Concepts/06mytime/)

---

## File Operations

### [13_file-operations.md](./13_file-operations.md)
**Core Topic:** File I/O, Reading, Writing, Directories

**Key Concepts Covered:**
- Reading files with `os.ReadFile`
- Writing files with `os.WriteFile`
- Creating and writing with `os.Create`
- File operations (exist, delete, rename, copy)
- Directory operations (create, list, remove)
- File permissions (Unix-style)
- Buffered I/O for performance
- Always `defer file.Close()`

**When to Read:** File handling in applications

**Quick Reference:**
```go
// Write
os.WriteFile("file.txt", []byte("data"), 0644)

// Read
data, _ := os.ReadFile("file.txt")

// Create and write
file, _ := os.Create("file.txt")
defer file.Close()
io.WriteString(file, "content")
```

**🔗 Example Code:** [Concepts/18files](../Concepts/18files/)

---

## JSON

### [14_json.md](./14_json.md)
**Core Topic:** JSON Encoding/Decoding, Struct Tags

**Key Concepts Covered:**
- Marshal (Go → JSON)
- Unmarshal (JSON → Go)
- Struct tags for field control
- `json:"-"` to exclude fields
- `json:"name,omitempty"` for optional fields
- Decoding to `map[string]interface{}`
- Type assertions for interface{}
- Validating JSON with `json.Valid`
- Streaming with Encoder/Decoder

**When to Read:** Working with JSON APIs, config files

**Quick Reference:**
```go
// Encode
jsonData, _ := json.Marshal(data)
json.MarshalIndent(data, "", "  ")

// Decode
json.Unmarshal(jsonBytes, &struct)

// Tags
`json:"name"`
`json:"-"`
`json:"name,omitempty"`
```

**🔗 Example Code:** [Concepts/22bitmorejson](../Concepts/22bitmorejson/)

---

## Web and HTTP

### [15_web-and-http.md](./15_web-and-http.md)
**Core Topic:** HTTP Client, Requests, URL Handling

**Key Concepts Covered:**
- HTTP GET requests with `http.Get`
- POST requests (JSON and form data)
- Custom requests with headers
- Response handling and status codes
- Always close response body
- Timeouts and custom clients
- URL parsing and building
- Common patterns (API client, retry, concurrent)
- Context for cancellation

**When to Read:** Building API clients, web scraping

**Quick Reference:**
```go
// GET
resp, _ := http.Get(url)
defer resp.Body.Close()
body, _ := io.ReadAll(resp.Body)

// POST JSON
jsonData, _ := json.Marshal(data)
http.Post(url, "application/json", bytes.NewBuffer(jsonData))

// Custom client
client := &http.Client{Timeout: 10 * time.Second}
```

**🔗 Example Code:** [Concepts/19webrequests](../Concepts/19webrequests/), [Concepts/20urls](../Concepts/20urls/)

---

## Building REST APIs

### [16_building-rest-apis.md](./16_building-rest-apis.md)
**Core Topic:** HTTP Server, REST APIs, Routing, Handlers

**Key Concepts Covered:**
- Basic HTTP server with `http.ListenAndServe`
- Handler functions and signatures
- gorilla/mux router for advanced routing
- Path parameters and query strings
- JSON request/response handling
- Complete CRUD API implementation
- Status codes and headers
- Middleware patterns (logging, CORS)
- File uploads
- Testing APIs with curl

**When to Read:** Building web services and APIs

**Quick Reference:**
```go
// Basic server
http.HandleFunc("/", handler)
http.ListenAndServe(":8080", nil)

// Mux router
r := mux.NewRouter()
r.HandleFunc("/users", getUsers).Methods("GET")
r.HandleFunc("/user/{id}", getUser).Methods("GET")

// JSON response
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(data)
```

**🔗 Example Code:** [Concepts/24buildapi](../Concepts/24buildapi/)

---

## Math and Random Numbers

### [17_math-and-random.md](./17_math-and-random.md)
**Core Topic:** Math Operations, Random Number Generation

**Key Concepts Covered:**
- Basic arithmetic and type compatibility
- Math package functions (Sqrt, Pow, Round, etc.)
- `math/rand` for non-secure random numbers
- Seeding with `time.Now().UnixNano()`
- `crypto/rand` for cryptographically secure randoms
- `math/big` for arbitrary precision
- Random patterns (dice, coin flip, shuffling)
- Performance considerations

**When to Read:** Numerical computations, games, security

**Quick Reference:**
```go
// Math
math.Sqrt(16)         // 4
math.Pow(2, 8)        // 256

// Random (non-secure)
rand.Seed(time.Now().UnixNano())
rand.Intn(100)        // [0, 100)

// Random (secure)
n, _ := rand.Int(rand.Reader, big.NewInt(100))
```

**🔗 Example Code:** [Concepts/05mymaths](../Concepts/05mymaths/)

---

## MongoDB Integration

### [18_mongodb-integration.md](./18_mongodb-integration.md)
**Core Topic:** MongoDB, Database Operations, CRUD

**Key Concepts Covered:**
- Connecting to MongoDB with mongo-driver
- Context package for timeouts
- BSON tags and `primitive.ObjectID`
- Insert operations (InsertOne, InsertMany)
- Find operations with cursors
- Update operations with `$set` operator
- Delete operations
- Query filters and operators
- Complete MongoDB API example
- MVC pattern (Model, Controller, Router)

**When to Read:** Database integration, NoSQL applications

**Quick Reference:**
```go
// Connect
client, _ := mongo.Connect(ctx, options.Client().ApplyURI(uri))
collection := client.Database("db").Collection("coll")

// Insert
collection.InsertOne(ctx, document)

// Find
collection.FindOne(ctx, filter).Decode(&result)

// Update
collection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"field": value}})

// Delete
collection.DeleteOne(ctx, filter)
```

**🔗 Example Code:** [Concepts/25mongoapi](../Concepts/25mongoapi/)

---

## Additional Topics (Code Examples Only)

### Advanced Web
- **Frontend Integration:** [Concepts/21gofrontend](../Concepts/21gofrontend/) - Go backend with frontend
- **Modules:** [Concepts/23mymodules](../Concepts/23mymodules/) - Creating and using modules

**Note:** These topics have code examples. Refer to the code directly for hands-on learning.

---

## How to Use These Notes

### For New Go Learners

**Foundations (Start Here):**
1. [00_golang-basics.md](./00_golang-basics.md) - Setup and intro
2. [01_variables-and-types.md](./01_variables-and-types.md) - Type system
3. [02_pointers.md](./02_pointers.md) - Memory concepts
4. [03_arrays-and-slices.md](./03_arrays-and-slices.md) - Collections
5. [05_maps.md](./05_maps.md) - Key-value data
6. [06_structs.md](./06_structs.md) - Custom types

**Core Language:**
7. [07_functions.md](./07_functions.md) - Function fundamentals
8. [08_methods.md](./08_methods.md) - Methods on types
9. [09_defer.md](./09_defer.md) - Cleanup and recovery
10. [10_control-flow.md](./10_control-flow.md) - if/switch/loops

**Practical Skills:**
11. [11_user-input-and-conversion.md](./11_user-input-and-conversion.md) - I/O and conversion
12. [12_time-handling.md](./12_time-handling.md) - Dates and time
13. [13_file-operations.md](./13_file-operations.md) - File I/O
14. [17_math-and-random.md](./17_math-and-random.md) - Math operations and random numbers
15. [14_json.md](./14_json.md) - JSON handling
16. [15_web-and-http.md](./15_web-and-http.md) - HTTP client requests

**Web & Database:**
17. [16_building-rest-apis.md](./16_building-rest-apis.md) - Building HTTP servers and APIs
18. [18_mongodb-integration.md](./18_mongodb-integration.md) - MongoDB database operations

**Advanced:**
19. [04_goroutines-and-concurrency.md](./04_goroutines-and-concurrency.md) - Concurrency
20. Explore Concepts folder (28+ examples)

### For Quick Reference
- **Go commands:** 00_golang-basics.md
- **Type conversion:** 01_variables-and-types.md
- **Pointer syntax:** 02_pointers.md
- **Slice operations:** 03_arrays-and-slices.md
- **Channel patterns:** 04_goroutines-and-concurrency.md

### Coming from Other Languages
- **Java/Kotlin developers:** Focus on explicit pointers, no inheritance, error handling
- **Python developers:** Study static typing, compilation, performance benefits
- **JavaScript developers:** Learn about goroutines vs async/await
- **C/C++ developers:** Understand GC, no pointer arithmetic, safer memory

---

## Note Format Standards

All notes follow consistent formatting:

### Structure
- **Metadata** (date, topics, related links) at top
- **Overview** section explaining what and why
- **Code examples** with clear explanations
- **Comparisons** with other languages when relevant
- **Common patterns** and best practices
- **Pitfalls** section with what to avoid
- **Summary** with key takeaways and quick reference
- **Links** to example code

### Visual Aids
- Memory diagrams for pointers and slices
- Tables for comparisons
- Code blocks with comments
- Step-by-step walkthroughs

---

## Adding New Notes

When adding notes, follow this template:

```markdown
# Topic Name

**📅 Created:** Series Name
**🏷️ Topics:** Topic1, Topic2, Topic3
**🔗 Related:** [other-note.md](./other-note.md), [Concepts/folder](../Concepts/folder/)

---

## Overview

Brief description of what this note covers...

---

## Main Content

...

---

## Summary

**Key Takeaways:**
- Point 1
- Point 2

**Quick Reference:**
```go
// Code examples
```

**When to Use This Note:**
- Use case 1
- Use case 2

---

**📝 Last Updated:** Date
**➡️ Next Topic:** [Next Note](./next-note.md)
**🔗 Example Code:** [Concepts/folder](../Concepts/folder/)
```

---

## Concepts Folder Structure

The `../Concepts/` folder contains 28+ practical examples:

```
Concepts/
├── 01hello/              - Hello World
├── 02variables/          - Variables and types
├── 03userinput/          - User input
├── 04conversion/         - Type conversion
├── 05mymaths/            - Math operations
├── 06mytime/             - Time handling
├── 07pointer/            - Pointers
├── 08myarray/            - Arrays
├── 09myslices/           - Slices
├── 10mymaps/             - Maps
├── 11mystructs/          - Structs
├── 12ifelse/             - If/else
├── 13switchcase/         - Switch statements
├── 14loops/              - Loops
├── 15functions/          - Functions
├── 16methods/            - Methods
├── 17defer/              - Defer keyword
├── 18files/              - File operations
├── 19webrequests/        - HTTP requests
├── 20urls/               - URL handling
├── 21gofrontend/         - Frontend integration
├── 22bitmorejson/        - JSON handling
├── 23mymodules/          - Modules
├── 24buildapi/           - REST API
├── 25mongoapi/           - MongoDB
├── 26goroutines/         - Goroutines
├── 27mutexAndAwaitGroups/ - Synchronization
└── 28channels/           - Channels
```

---

## External Resources

### Official Documentation
- **Go Website:** https://go.dev/
- **Package Docs:** https://pkg.go.dev/
- **Go Spec:** https://go.dev/ref/spec
- **Effective Go:** https://go.dev/doc/effective_go

### Learning Resources
- **Tour of Go:** https://go.dev/tour/
- **Go by Example:** https://gobyexample.com/
- **Go Playground:** https://go.dev/play/

### Community
- **Go Blog:** https://go.dev/blog/
- **Go Forum:** https://forum.golangbridge.org/
- **r/golang:** https://reddit.com/r/golang

---

**Last Updated:** November 3, 2025
**Total Notes:** 19 comprehensive notes + 28+ practical code examples
**Coverage:** Complete - from basics to concurrency, I/O, web, APIs, and databases
**Status:** Comprehensive reference ready

---

*This index is maintained alongside the notes. All notes follow the standardized format established for consistency.*
