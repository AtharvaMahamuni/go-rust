# Go Basics - Getting Started with Go

**📅 Created:** Initial Setup
**🏷️ Topics:** Go Introduction, Compilation, Go Modules, Project Setup
**🔗 Related:** [01_variables-and-types.md](./01_variables-and-types.md), [Concepts/01hello](../Concepts/01hello/)

---

## Overview

This note covers the fundamental concepts of Go programming: what makes Go unique, how it compiles, setting up projects with Go modules, and running your first program. Essential starting point for understanding Go's philosophy and tooling.

---

## What is Go?

### Key Characteristics

**Compiled Language:**
- Go compiles directly to machine code (no VM like Java/Kotlin)
- Can generate executables for different operating systems
- Fast execution with instant startup time
- Single binary deployment (no dependencies needed)

**Purpose:**
- System apps to web apps
- Cloud applications (reducing server resource requirements)
- Already widely used in production (Docker, Kubernetes, etc.)

**Philosophy:**
- Simplicity over complexity
- "Don't bring baggage" - start fresh without assumptions
- Similarities with C, Java, Pascal (familiar yet different)

---

## Object-Oriented? Yes and No

**The Question:** Is Go object-oriented?

**Answer:** Yes and No
- Go has its own way of doing things
- No traditional classes, but has structs and methods
- No inheritance, but has composition and interfaces
- Focus on "what works" rather than strict OOP principles

**Key Insight:** Go asks "Do you really need that feature, or is there a better way?"

---

## What's Missing (Intentionally)

Go deliberately omits features common in other languages:

1. **No try-catch** - Uses explicit error handling instead
2. **No exceptions** - Functions return errors as values
3. **No classes/inheritance** - Uses composition over inheritance
4. **No generics** (until Go 1.18) - Initially focused on simplicity
5. **Automatic semicolon insertion** - Lexer adds them for you

**Philosophy:** "Is it really missing, or do we not need it?"

---

## The Lexer

**What is a Lexer?**
- Part of the compiler that checks grammar
- Handles syntax validation

**Go's Lexer Magic:**
- Automatically inserts semicolons (no need to write them explicitly)
- You CAN use semicolons in complex statements for clarity
- Makes code cleaner and more readable

```go
// No semicolons needed
x := 5
y := 10
z := x + y

// But you CAN use them for one-liners
if x > 0 { fmt.Println("Positive"); return }
```

**Learn More:** [Go Language Specification](https://go.dev/ref/spec)

---

## Setting Up a Go Project

### Step 1: Create Project Structure

```bash
mkdir myproject
cd myproject
```

### Step 2: Initialize Go Module

```bash
go mod init github.com/username/myproject
```

This creates a `go.mod` file - Go's dependency management file (similar to `package.json` or `build.gradle`).

**What is go.mod?**
- Defines your module path (import path for others)
- Tracks dependencies and their versions
- Required for any Go project (even simple ones)

### Step 3: Create main.go

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Step 4: Run Your Program

```bash
go run main.go
```

This compiles and runs the program in one step (no executable created).

---

## Go Commands

### Essential Commands

```bash
# Get help
go help

# Run without creating executable
go run main.go

# Build executable
go build

# Build for specific OS (cross-compilation)
GOOS=linux GOARCH=amd64 go build
GOOS=windows GOARCH=amd64 go build

# Install dependencies
go get github.com/some/package

# Tidy dependencies (clean up go.mod)
go mod tidy

# Format code (automatically)
go fmt ./...

# Run tests
go test ./...
```

### Build vs Run

| Command | Purpose | Output |
|---------|---------|--------|
| `go run main.go` | Quick testing | No file created |
| `go build` | Create executable | Binary file created |
| `go install` | Build and install | Binary in $GOPATH/bin |

---

## Project Structure

### Minimal Project

```
myproject/
├── go.mod          # Module definition
├── go.sum          # Dependency checksums (auto-generated)
└── main.go         # Entry point
```

### Typical Project

```
myproject/
├── go.mod
├── go.sum
├── main.go         # Entry point
├── README.md
├── cmd/            # Command-line applications
│   └── app/
│       └── main.go
├── pkg/            # Reusable packages
│   └── mylib/
│       └── mylib.go
├── internal/       # Private packages (can't be imported)
│   └── utils/
│       └── helper.go
└── api/            # API definitions
```

---

## The main Package and main Function

**Special Package:** `package main`
- Creates an executable program
- Must have a `main()` function as entry point

**Non-main Packages:**
- Create reusable libraries
- Can be imported by other packages

```go
// Executable program
package main

func main() {
    // Entry point
}
```

```go
// Reusable library
package mylib

func MyFunction() {
    // Can be imported and used
}
```

---

## Compilation Model

### How Go Compiles

```
Source Code (.go files)
    ↓
Lexer (syntax check, add semicolons)
    ↓
Parser (build syntax tree)
    ↓
Type Checker (verify types)
    ↓
Compiler (generate machine code)
    ↓
Linker (create executable)
    ↓
Native Binary
```

### Comparison with Other Languages

| Language | Compilation | Runtime | Deployment |
|----------|-------------|---------|------------|
| **Go** | Native code | None | Single binary |
| **Java/Kotlin** | Bytecode | JVM required | JAR + JVM |
| **Python** | Interpreted | Python interpreter | .py files + interpreter |
| **JavaScript** | Interpreted/JIT | Node.js/Browser | .js files + runtime |

**Go Advantage:** No runtime dependencies - just copy the binary!

---

## Cross-Compilation

Go makes it easy to build for different operating systems:

```bash
# Build for Linux
GOOS=linux GOARCH=amd64 go build -o myapp-linux

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o myapp.exe

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o myapp-mac

# Build for ARM (Raspberry Pi, etc.)
GOOS=linux GOARCH=arm go build -o myapp-arm
```

**Supported Platforms:**
- Windows, Linux, macOS
- FreeBSD, OpenBSD, NetBSD
- ARM, ARM64, 386, amd64
- WebAssembly (WASM)

---

## Go Workspace (Pre-Modules Era)

**Old Way (Before Go 1.11):**
```
$GOPATH/
├── bin/     # Compiled executables
├── pkg/     # Compiled packages
└── src/     # Source code
    └── github.com/
        └── username/
            └── project/
```

**Modern Way (Go Modules):**
- Work anywhere on your filesystem
- No need for $GOPATH/src
- Dependencies managed in go.mod
- Much simpler!

---

## Hello World - Explained

```go
package main              // This is an executable package

import "fmt"              // Import standard library package

func main() {             // Entry point function
    fmt.Println("Hello")  // Print with newline
}
```

**Breaking it down:**

1. `package main` - Creates executable (not library)
2. `import "fmt"` - Import formatting package from standard library
3. `func main()` - Entry point (must exist in package main)
4. `fmt.Println()` - Function from fmt package

---

## Important Resources

### Official Documentation

- **Go Website:** https://go.dev/
- **Package Documentation:** https://pkg.go.dev/
- **Go Spec:** https://go.dev/ref/spec
- **Effective Go:** https://go.dev/doc/effective_go
- **Go Blog:** https://go.dev/blog/

### Learning Resources

- **Go by Example:** https://gobyexample.com/
- **Tour of Go:** https://go.dev/tour/
- **Go Playground:** https://go.dev/play/ (test code online)

### Tools

- **Go Modules:** https://go.dev/blog/using-go-modules
- **VS Code Extension:** Official Go extension
- **GoLand:** JetBrains IDE for Go

---

## Summary

**Key Takeaways:**
- Go compiles to native code (fast, no runtime needed)
- Use `go mod init` to start a project
- `go run` for quick testing, `go build` for executables
- No try-catch, no classes, no inheritance (by design)
- Lexer automatically adds semicolons
- Single binary deployment (no dependencies)
- Cross-compilation is built-in and easy
- package main + main() = executable program

**Quick Start Commands:**
```bash
go mod init myproject    # Initialize project
go run main.go          # Run quickly
go build                # Create executable
go fmt ./...            # Format code
go help                 # Get help
```

**When to Use This Note:**
- Starting a new Go project
- Understanding Go's philosophy
- Quick reference for go commands
- Learning about compilation model

---

**📝 Last Updated:** Initial Version
**➡️ Next Topic:** [Variables and Types](./01_variables-and-types.md)
**🔗 Example Code:** [Concepts/01hello](../Concepts/01hello/)
