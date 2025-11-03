# Go Learning Repository

A comprehensive, hands-on learning repository for mastering Go (Golang), with 28+ practical examples and detailed reference notes.

---

## Repository Structure

```
Golang/
├── README.md                    # This file - start here
├── notes.txt                    # Original learning notes
├── notes/                       # **Reference notes and guides**
│   ├── INDEX.md                 # **START HERE** - Complete index
│   ├── 00_golang-basics.md
│   ├── 01_variables-and-types.md
│   ├── 02_pointers.md
│   ├── 03_arrays-and-slices.md
│   └── 04_goroutines-and-concurrency.md
├── Concepts/                    # 28+ practical examples
│   ├── 01hello/
│   ├── 02variables/
│   ├── ... (see full list below)
│   └── 28channels/
├── DesignPatterns/              # Design pattern implementations
├── DSA/                         # Data Structures & Algorithms
├── Interview/                   # Interview preparation topics
└── Projects/                    # Larger Go projects
```

---

## Quick Start

### 1. New to Go?
Start with the **[notes directory](./notes/INDEX.md)** for comprehensive learning:

1. Read [notes/00_golang-basics.md](./notes/00_golang-basics.md) - Setup and introduction
2. Read [notes/01_variables-and-types.md](./notes/01_variables-and-types.md) - Type system
3. Read [notes/02_pointers.md](./notes/02_pointers.md) - Memory concepts
4. Read [notes/03_arrays-and-slices.md](./notes/03_arrays-and-slices.md) - Collections
5. Read [notes/04_goroutines-and-concurrency.md](./notes/04_goroutines-and-concurrency.md) - Concurrency
6. Practice with examples in `Concepts/`

### 2. Coming from Another Language?
- **Java/Kotlin:** No inheritance, explicit pointers, no exceptions
- **Python:** Static typing, compiled, explicit error handling
- **JavaScript:** No async/await, goroutines instead
- **C/C++:** GC instead of manual memory, no pointer arithmetic

### 3. Want Hands-On Practice?
Explore the `Concepts/` directory (28+ examples) in order:
- Start with basics (hello, variables, types)
- Progress through collections (arrays, slices, maps, structs)
- Master control flow (if/else, loops, switch)
- Learn concurrency (goroutines, channels, mutexes)
- Build real apps (web requests, APIs, databases)

---

## Notes Directory

The **[notes/](./notes/)** directory contains comprehensive reference materials.

### Available Notes

| File | Topic | Description |
|------|-------|-------------|
| **[INDEX.md](./notes/INDEX.md)** | **Navigation** | **Complete index - start here** |
| [00_golang-basics.md](./notes/00_golang-basics.md) | Basics | Go intro, setup, compilation, modules |
| [01_variables-and-types.md](./notes/01_variables-and-types.md) | Types | Type system, variables, zero values |
| [02_pointers.md](./notes/02_pointers.md) | Pointers | Memory addresses, pass-by-reference |
| [03_arrays-and-slices.md](./notes/03_arrays-and-slices.md) | Collections | Arrays vs slices, slice operations |
| [04_goroutines-and-concurrency.md](./notes/04_goroutines-and-concurrency.md) | Concurrency | Goroutines, channels, sync |
| [05_maps.md](./notes/05_maps.md) | Maps | Key-value collections, hash tables |
| [06_structs.md](./notes/06_structs.md) | Structs | Custom types, composition, embedding |
| [07_functions.md](./notes/07_functions.md) | Functions | Parameters, returns, closures |
| [08_methods.md](./notes/08_methods.md) | Methods | Receivers, value vs pointer |
| [09_defer.md](./notes/09_defer.md) | Defer | Cleanup, panic, recover |
| [10_control-flow.md](./notes/10_control-flow.md) | Control Flow | if/else, switch, loops |
| [11_user-input-and-conversion.md](./notes/11_user-input-and-conversion.md) | User Input | Reading input, type conversion |
| [12_time-handling.md](./notes/12_time-handling.md) | Time | Date/time operations, formatting |
| [13_file-operations.md](./notes/13_file-operations.md) | File I/O | Reading, writing, file operations |
| [14_json.md](./notes/14_json.md) | JSON | Encoding, decoding, struct tags |
| [15_web-and-http.md](./notes/15_web-and-http.md) | Web & HTTP | HTTP client, GET/POST, URLs |
| [16_building-rest-apis.md](./notes/16_building-rest-apis.md) | REST APIs | HTTP server, routing, handlers |
| [17_math-and-random.md](./notes/17_math-and-random.md) | Math & Random | Math operations, random numbers |
| [18_mongodb-integration.md](./notes/18_mongodb-integration.md) | MongoDB | Database integration, CRUD ops |

**See [notes/INDEX.md](./notes/INDEX.md) for detailed descriptions and learning paths.**

---

## Concepts Directory (28+ Examples)

Practical, runnable examples covering all Go fundamentals:

### Basics (1-6)
- **[01hello](./Concepts/01hello/)** - Hello World and project setup
- **[02variables](./Concepts/02variables/)** - Variable declarations and types
- **[03userinput](./Concepts/03userinput/)** - Reading user input (bufio, os)
- **[04conversion](./Concepts/04conversion/)** - Type conversions (strconv)
- **[05mymaths](./Concepts/05mymaths/)** - Math operations and random numbers
- **[06mytime](./Concepts/06mytime/)** - Time handling and formatting

### Memory & Collections (7-11)
- **[07pointer](./Concepts/07pointer/)** - Pointers and memory addresses
- **[08myarray](./Concepts/08myarray/)** - Arrays (fixed-size)
- **[09myslices](./Concepts/09myslices/)** - Slices (dynamic arrays)
- **[10mymaps](./Concepts/10mymaps/)** - Maps (key-value pairs)
- **[11mystructs](./Concepts/11mystructs/)** - Structs and composition

### Control Flow (12-14)
- **[12ifelse](./Concepts/12ifelse/)** - If/else statements
- **[13switchcase](./Concepts/13switchcase/)** - Switch statements (no break!)
- **[14loops](./Concepts/14loops/)** - For loops and range

### Functions (15-17)
- **[15functions](./Concepts/15functions/)** - Function declarations, multiple returns
- **[16methods](./Concepts/16methods/)** - Methods on types
- **[17defer](./Concepts/17defer/)** - Defer keyword and execution order

### File & Web (18-22)
- **[18files](./Concepts/18files/)** - File I/O operations
- **[19webrequests](./Concepts/19webrequests/)** - HTTP GET/POST requests
- **[20urls](./Concepts/20urls/)** - URL parsing and manipulation
- **[21gofrontend](./Concepts/21gofrontend/)** - Frontend integration
- **[22bitmorejson](./Concepts/22bitmorejson/)** - JSON encoding/decoding

### Modules & APIs (23-25)
- **[23mymodules](./Concepts/23mymodules/)** - Creating and using modules
- **[24buildapi](./Concepts/24buildapi/)** - Building REST APIs
- **[25mongoapi](./Concepts/25mongoapi/)** - MongoDB integration

### Concurrency (26-28)
- **[26goroutines](./Concepts/26goroutines/)** - Goroutines and WaitGroups
- **[27mutexAndAwaitGroups](./Concepts/27mutexAndAwaitGroups/)** - Mutex for synchronization
- **[28channels](./Concepts/28channels/)** - Channels for communication

---

## Additional Directories

### DesignPatterns/
Go implementations of design patterns:
- **cmd/** - Command pattern examples
- **models/** - Model structures
- **pets/** - Practical examples
- **StandAloneExamples/** - Isolated pattern demonstrations

### DSA/
Data Structures and Algorithms in Go:
- **BinaryTrees/** - Binary tree implementations
- **LinkedList/** - Linked list operations
- **Recursion/** - Recursive algorithms
- **Stack/** - Stack data structure

### Interview/
Interview preparation topics:
- **Closures/** - Closure examples and patterns
- **Composition/** - Composition over inheritance
- **SOLID/** - SOLID principles in Go

### Projects/
Larger, real-world Go projects demonstrating full applications.

---

## Learning Path

### Phase 1: Foundations (Weeks 1-2)
- [x] Setup Go environment
- [x] Hello World and project structure
- [x] Variables, types, and type system
- [x] Pointers and memory
- [x] Arrays, slices, and maps
- [x] Structs and composition
- [x] Functions and methods
- [x] Defer, panic, recover

### Phase 2: Core Language (Weeks 3-4)
- [x] Control flow (if/else, loops, switch)
- [x] User input and conversion
- [x] Time handling and formatting
- [x] File I/O operations
- [ ] Error handling patterns (custom errors, wrapping)
- [ ] Interfaces and polymorphism

### Phase 3: Concurrency (Week 5)
- [x] Goroutines
- [x] Channels
- [x] sync.WaitGroup
- [x] sync.Mutex
- [x] Concurrent patterns

### Phase 4: Web & APIs (Week 6+)
- [x] HTTP client and web requests
- [x] URL parsing and building
- [x] JSON encoding/decoding
- [x] Math operations and random numbers
- [x] Building REST APIs and HTTP servers
- [x] Database integration (MongoDB)
- [ ] Middleware patterns
- [ ] Testing and benchmarking

### Phase 5: Advanced Topics
- [ ] Interfaces and generics
- [ ] Context package
- [ ] Testing (unit, integration)
- [ ] Profiling and optimization
- [ ] Design patterns

---

## Key Go Concepts

### What Makes Go Special?

1. **Fast Compilation** - Compiles to native code, instant startup
2. **Simple Syntax** - Easy to learn, read, and maintain
3. **Built-in Concurrency** - Goroutines and channels are first-class
4. **Strong Standard Library** - Most things included
5. **Static Typing** - Catch errors at compile time
6. **Garbage Collection** - Automatic memory management
7. **Cross-Platform** - Build for any OS from any OS

### Go Philosophy

- **Simplicity over complexity** - Features must earn their keep
- **Composition over inheritance** - No classes, use structs + interfaces
- **Explicit over implicit** - No hidden magic, clear intentions
- **Fast compilation** - Dependency management built for speed
- **Concurrency as a core feature** - Not an afterthought

---

## Quick Reference

### Essential Commands

```bash
# Project setup
go mod init github.com/username/project
go mod tidy

# Development
go run main.go              # Run without building
go build                    # Build executable
go build -o appname         # Build with custom name
go fmt ./...                # Format all files
go vet ./...                # Lint code

# Testing
go test ./...               # Run all tests
go test -v                  # Verbose output
go test -race               # Detect race conditions
go test -cover              # Coverage report

# Dependencies
go get package              # Add dependency
go get -u package           # Update dependency
go mod download             # Download dependencies

# Documentation
go doc package              # Show package docs
go doc package.Function     # Show function docs
```

### Basic Syntax

```go
// Package declaration
package main

// Imports
import "fmt"
import (
    "fmt"
    "net/http"
)

// Variables
var x int = 10
var x = 10
x := 10  // Functions only

// Functions
func add(a int, b int) int {
    return a + b
}

// Multiple returns
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Structs
type Person struct {
    Name string
    Age  int
}

// Methods
func (p *Person) Greet() {
    fmt.Println("Hello, I'm", p.Name)
}

// Goroutines
go myFunction()

// Channels
ch := make(chan int)
ch <- 42
value := <-ch
```

---

## Resources

### Official Documentation
- **Go Website:** https://go.dev/
- **Package Docs:** https://pkg.go.dev/
- **Go Spec:** https://go.dev/ref/spec
- **Effective Go:** https://go.dev/doc/effective_go
- **Go Blog:** https://go.dev/blog/

### Learning Resources
- **Tour of Go:** https://go.dev/tour/
- **Go by Example:** https://gobyexample.com/
- **Go Playground:** https://go.dev/play/
- **Exercism Go Track:** https://exercism.org/tracks/go

### Books
- The Go Programming Language (Donovan & Kernighan)
- Learning Go (Jon Bodner)
- Concurrency in Go (Katherine Cox-Buday)

### Video Course
- LCO YouTube Course: https://www.youtube.com/playlist?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa

---

## Common Go Patterns

### Error Handling
```go
result, err := someFunction()
if err != nil {
    return err  // Handle error
}
// Use result
```

### Defer for Cleanup
```go
file, err := os.Open("file.txt")
if err != nil {
    return err
}
defer file.Close()  // Guaranteed to run
```

### Interface Implementation
```go
type Speaker interface {
    Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}
// Dog automatically implements Speaker
```

---

## Progress Tracking

**Current Focus:** Advanced concurrency patterns and web APIs
**Last Updated:** November 2, 2025
**Concepts Completed:** 28/28
**Next Up:** Testing, middleware, production deployment

---

## Contributing to Your Learning

Keep this repository as your **personal Go knowledge base**:
- Add notes after each learning session
- Document "aha!" moments and gotchas
- Build and commit practice projects
- Maintain a running list of questions and answers
- Compare Go patterns with other languages you know

---

## Why Go?

**Use Go when you need:**
- High performance (close to C/C++)
- Simple concurrency (goroutines beat threads)
- Fast compilation (faster than Java/C++)
- Single binary deployment (no runtime dependencies)
- Cloud services / microservices
- CLI tools
- Network servers

**Go powers:**
- Docker (containerization)
- Kubernetes (orchestration)
- Terraform (infrastructure)
- Prometheus (monitoring)
- Ethereum (blockchain)

---

**Happy Learning!** 🚀

*Remember: Go is designed for simplicity. If something feels complicated, there's probably a simpler way. Trust the language and its idioms.*
