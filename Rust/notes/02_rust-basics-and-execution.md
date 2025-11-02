# Rust Learning Session 1 - Foundation & First Program

**Date:** November 2, 2025  
**Topics:** Memory Management, Ownership Review, Tooling Setup, Hello World

---

## Key Concepts Reviewed

### 1. Memory Management Fundamentals

**Stack vs Heap:**
- **Stack**: Fast, fixed-size data, LIFO, automatic cleanup
- **Heap**: Slower, dynamic-size data, requires management

**Three Approaches to Memory Management:**

| Language | Strategy | Pros | Cons |
|----------|----------|------|------|
| C/C++ | Manual (malloc/free) | Full control, fast | Memory leaks, crashes |
| Kotlin/Java | Garbage Collector | Safe, automatic | Runtime pauses, overhead |
| Rust | Ownership (compile-time) | Safe + Fast, no GC | Learning curve, strict rules |

### 2. Native Code vs Bytecode

**Kotlin (Bytecode):**
- Compiles to JVM bytecode
- Runs on JVM (runtime layer)
- Write once, run anywhere (portable)
- Slower execution due to JVM overhead

**Rust (Native Code):**
- Compiles to OS-specific machine code
- Runs directly on CPU (no runtime)
- Need separate builds per OS (.exe for Windows, different for Linux/Mac)
- Faster execution, instant startup

**Key Insight:** Rust's speed comes from producing native code, but this requires OS-specific tooling (like C++ build tools on Windows).

### 3. Ownership & Borrowing (Quick Recap)

**Ownership Rules:**
1. Each value has exactly one owner
2. Only one owner at a time
3. When owner goes out of scope, value is dropped (memory freed)

**The Dangling Reference Problem:**
```rust
fn process_data() -> &String {
    let data = String::from("important");
    &data  // ❌ ERROR: returning reference to data that will be freed!
}
```
Rust prevents this at **compile-time** - the memory would be freed when function ends, making the reference invalid.

**Why Rust Needs Lifetimes:**
- Ownership tracks "who cleans up"
- Lifetimes track "how long is this reference valid"
- Both are checked at compile-time for safety

---

## Rust Toolchain

### The Three Main Tools

**`rustup`** - Toolchain manager
- Installs and manages Rust versions
- Like `sdkman` for Kotlin, but integrated

**`rustc`** - The compiler
- Compiles `.rs` files to native code
- Like `kotlinc` but for native executables

**`cargo`** - Build system & package manager
- Like Gradle for Kotlin - does EVERYTHING
- Not just dependencies, but complete build orchestration

### Cargo Commands

```bash
cargo new <project>     # Create new project with directory
cargo init              # Initialize project in current directory
cargo build             # Compile in debug mode
cargo build --release   # Compile optimized for production
cargo run               # Compile + run
cargo test              # Run tests
cargo help              # Show all commands
```

### Cross-Compilation

**Can you build for Linux on Windows?**
Yes! Rust supports cross-compilation:
```bash
rustup target add x86_64-unknown-linux-gnu
cargo build --target x86_64-unknown-linux-gnu
```

Note: Some dependencies need platform-specific libraries, making it tricky. Most projects use CI/CD or Docker for multi-platform builds.

---

## Project Structure

### What `cargo new hello-rust` Creates:

```
hello-rust/
├── Cargo.toml          # Project config (like build.gradle)
├── src/
│   └── main.rs         # Source code (entry point)
└── .git/               # Git repository (auto-initialized)
```

### Cargo.toml Explained

```toml
[package]
name = "hello-rust"      # Project name
version = "0.1.0"        # Current version (manually updated)
edition = "2024"         # Rust language edition (like Java 8 vs Java 11)

[dependencies]           # Add external crates here (like Gradle dependencies)
```

**Edition:** Rust evolves the language over time. Editions allow:
- Old code to keep working (2018 rules)
- New code to use improvements (2021, 2024 rules)
- Both to coexist in the same project

---

## First Rust Program

### Hello World Code

```rust
fn main() {
    println!("Hello, world!");
}
```

**Syntax Comparison:**

| Feature | Kotlin | Rust |
|---------|--------|------|
| Function keyword | `fun` | `fn` |
| Entry point | `fun main()` | `fn main()` |
| Print | `println()` | `println!()` |
| Statement end | optional `;` | required `;` |

### The `!` Mystery - Macros

**Why `println!` has `!`?**

The `!` marks a **macro** - not a regular function.

**Regular Function (runtime):**
```kotlin
// Kotlin checks format at runtime
println("Value: %d".format(42))
```

**Macro (compile-time):**
```rust
// Rust checks format at compile-time
println!("Value: {}", 42);  // Compiler verifies placeholder count matches arguments
```

**Benefits of Macros:**
- Errors caught during compilation, not when program runs
- Variable number of arguments without varargs
- Code generation at compile-time for better performance

**Example Error Caught at Compile-Time:**
```rust
println!("Two: {} {}", 42);  // ❌ Compiler error: not enough arguments
```

---

## Build Process

### What Happens During `cargo run`

```
Step 1: Compilation
  main.rs → rustc → object files → link.exe → hello-rust.exe

Step 2: Execution
  Running target\debug\hello-rust.exe
  Hello, world!
```

### Build Modes

**Debug Mode (default):**
- Fast compilation
- Slower runtime
- Includes debugging symbols
- Output: `target/debug/hello-rust.exe`

**Release Mode:**
- Slower compilation
- Fast runtime (optimized)
- No debugging symbols
- Output: `target/release/hello-rust.exe`
- Use: `cargo build --release`

---

## Windows Setup Notes

### Why C++ Build Tools Required?

Rust compiles to native Windows code, which requires:
- **link.exe** - Windows linker (connects your code with system libraries)
- **Windows SDK** - Access to Windows APIs
- **MSVC toolchain** - Microsoft's C++ compiler infrastructure

Installation:
1. Download `rustup-init.exe` from rust-lang.org
2. Install Visual Studio Build Tools
3. Check "Desktop development with C++" workload
4. Verify: `rustc --version` and `cargo --version`

### Common Error:
```
error: linker `link.exe` not found
```
**Solution:** Install/modify Visual Studio with "Desktop development with C++" checked.

---

## Compile-Time vs Runtime Philosophy

**Rust's Core Philosophy:**
> Catch errors at **compile-time**, not **runtime**

**Examples:**

| Error Type | Runtime (Kotlin/Java) | Compile-Time (Rust) |
|------------|----------------------|---------------------|
| Type mismatch | RuntimeException | Won't compile |
| Use after free | Crashes/corruption | Won't compile |
| Data races | Silent bugs | Won't compile |
| Format string errors | Runtime crash | Won't compile |
| Null pointer | NullPointerException | Won't compile (no nulls!) |

**Trade-off:** Stricter compiler = longer learning curve, but safer and faster code.

---

## Key Takeaways

1. **Rust = Safety + Speed**: Combines memory safety of GC languages with performance of C/C++
2. **Ownership is Rust's superpower**: Prevents memory bugs at compile-time
3. **Cargo is your friend**: One tool for everything (like Gradle but better integrated)
4. **Macros (`!`) = Compile-time magic**: More powerful than functions, caught early
5. **Native code = OS-specific**: Need different builds for Windows/Linux/Mac
6. **Rust tooling is helpful**: Error messages guide you to solutions

---

## Next Steps

1. **Practice ownership in code**: Write programs that move and borrow values
2. **Compare Rust & Kotlin syntax**: Learn through familiar patterns
3. **Basic Rust types**: Understand primitives, strings, vectors
4. **Control flow**: if, match, loops with Rust syntax
5. **Error handling**: Result and Option types (Rust's approach to no null)

---

## Quick Reference Commands

```bash
# Setup
rustup update                    # Update Rust
rustup doc                       # Open offline docs

# Project
cargo new my-project            # New project
cargo init                      # Init in current dir
cd my-project

# Development
cargo check                     # Quick syntax check (no build)
cargo build                     # Debug build
cargo run                       # Build + run
cargo build --release          # Optimized build
cargo test                     # Run tests

# Verification
rustc --version
cargo --version
```

---

## Useful Resources

- Official Rust Book: https://doc.rust-lang.org/book/
- Rust by Example: https://doc.rust-lang.org/rust-by-example/
- Project Knowledge Files:
  - `rust-ownership-guide.md` - Detailed ownership reference
  - `memory-management-reference.md` - Stack/heap fundamentals
  - "The Rust Programming Language" book (PDF in project)

---

**Session Complete!** ✅  
First Rust program compiled and running. Ready to explore ownership through actual code in the next session.
