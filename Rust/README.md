# Rust Learning Repository

A comprehensive, hands-on learning repository for mastering Rust, with focus on transitioning from Kotlin/JVM background.

---

## Repository Structure

```
Rust/
├── README.md                    # This file - start here
├── PROJECT_INSTRUCTIONS.md      # Template for creating new learning files
├── code/                        # Executable Rust examples
│   └── language-syntax/         # Syntax learning exercises
├── notes/                       # Reference notes and session summaries
│   ├── INDEX.md                 # **START HERE** - Complete notes index
│   ├── 00_memory-management-reference.md
│   ├── 01_rust-ownership-guide.md
│   ├── 02_rust-basics-and-execution.md
│   ├── 03_shadowing_visualization.md
│   └── 04_variables_shadowing.md
├── comprehensive-rust.pdf       # Google's Rust course
└── The Rust Programming Language (2nd Edition).pdf
```

---

## Quick Start

### 1. New to Rust?
Start with the **[notes directory](./notes/INDEX.md)** for comprehensive learning materials:

1. Read [notes/00_memory-management-reference.md](./notes/00_memory-management-reference.md) - Understand stack/heap
2. Read [notes/01_rust-ownership-guide.md](./notes/01_rust-ownership-guide.md) - Master ownership
3. Read [notes/04_variables_shadowing.md](./notes/04_variables_shadowing.md) - Learn mutability
4. Run examples in `code/language-syntax/` to practice

### 2. Coming from Kotlin?
All notes include **Kotlin comparisons** to help you relate concepts:
- Memory management: GC vs Ownership
- Variable mutability: `val`/`var` vs `let`/`let mut`
- Type system differences
- Performance trade-offs

### 3. Want Hands-On Practice?
Explore the `code/language-syntax/` directory for executable examples following the format in [PROJECT_INSTRUCTIONS.md](./PROJECT_INSTRUCTIONS.md).

---

## Learning Path

### Phase 1: Foundations (Current)
- [x] Memory management basics (stack vs heap)
- [x] Ownership and borrowing rules
- [x] Variables and shadowing
- [x] First Rust program (Hello World)
- [ ] Data types and functions
- [ ] Control flow and pattern matching

### Phase 2: Core Language Features
- [ ] Error handling (Result, Option)
- [ ] Structs and enums
- [ ] Traits and generics
- [ ] Lifetimes (advanced)
- [ ] Smart pointers

### Phase 3: Real-World Rust
- [ ] Modules and packages
- [ ] Testing
- [ ] Concurrency
- [ ] Building CLI tools
- [ ] Building web services

---

## Key Resources

### In This Repository
- **[notes/INDEX.md](./notes/INDEX.md)** - Complete index of all learning notes
- **[PROJECT_INSTRUCTIONS.md](./PROJECT_INSTRUCTIONS.md)** - Template for creating new examples
- **code/** - Hands-on exercises and examples
- **PDFs** - Comprehensive Rust and official Rust book

### External Resources
- [Official Rust Book](https://doc.rust-lang.org/book/) - The definitive guide
- [Rust by Example](https://doc.rust-lang.org/rust-by-example/) - Learn by doing
- [Rust Playground](https://play.rust-lang.org/) - Test code online
- [Rustlings](https://github.com/rust-lang/rustlings) - Interactive exercises

### Toolchain
```bash
# Setup
rustup update              # Update Rust
rustup doc                # Open offline docs

# Development
cargo new my-project      # Create new project
cargo build              # Compile
cargo run               # Build + run
cargo test              # Run tests

# Verification
rustc --version
cargo --version
```

---

## Notes Directory

The **[notes/](./notes/)** directory contains comprehensive reference materials and session summaries.

### Available Notes

| File | Topic | Description |
|------|-------|-------------|
| **[INDEX.md](./notes/INDEX.md)** | **Navigation** | **Complete index with all notes** |
| [00_memory-management-reference.md](./notes/00_memory-management-reference.md) | Memory | Stack vs Heap fundamentals |
| [01_rust-ownership-guide.md](./notes/01_rust-ownership-guide.md) | Ownership | Complete ownership & borrowing guide |
| [02_rust-basics-and-execution.md](./notes/02_rust-basics-and-execution.md) | Session 1 | Tooling, Hello World, first program |
| [03_shadowing_visualization.md](./notes/03_shadowing_visualization.md) | Shadowing | Stack/heap behavior during shadowing |
| [04_variables_shadowing.md](./notes/04_variables_shadowing.md) | Session 2 | Mutability and shadowing patterns |

**See [notes/INDEX.md](./notes/INDEX.md) for detailed descriptions and reading order.**

---

## Code Examples

### Language Syntax (code/language-syntax/)

Executable examples following the structured learning format:

- `01_mutability_basics.rs` - Immutability by default, explicit mutation
- *(More examples to be added as learning progresses)*

Each file follows the format defined in [PROJECT_INSTRUCTIONS.md](./PROJECT_INSTRUCTIONS.md):
- Real-world scenarios
- Progressive examples
- Kotlin comparisons
- Experiments to try
- Key takeaways

---

## Learning Approach

This repository follows a **learn-by-doing** philosophy:

1. **Read** the concept in notes
2. **Run** the example code
3. **Observe** the output
4. **Modify** the code
5. **Break** the code intentionally
6. **Fix** it and understand errors
7. **Compare** with Kotlin
8. **Explain** the concept in your own words

---

## Progress Tracking

**Current Focus:** Variables, shadowing, and basic types
**Last Updated:** November 2, 2025
**Sessions Completed:** 2
**Next Up:** Data types and functions

---

## How to Add New Content

### Adding New Notes

1. Create file in `notes/` following naming convention: `XX_topic-name.md`
2. Add entry to [notes/INDEX.md](./notes/INDEX.md) following the established format
3. Update this README's notes table
4. Commit with descriptive message

### Adding New Code Examples

1. Create file in `code/language-syntax/` following naming convention: `XX_topic_name.rs`
2. Follow the format template in [PROJECT_INSTRUCTIONS.md](./PROJECT_INSTRUCTIONS.md)
3. Include real-world scenarios, Kotlin comparisons, and experiments
4. Update this README's code examples section
5. Test the code runs successfully

---

## Contributing to Your Own Learning

Keep this repository as your **living documentation**:
- Add notes after each learning session
- Document "aha!" moments
- Record common mistakes and solutions
- Build a personal Rust reference guide
- Compare with your Kotlin experience

---

## Philosophy

**Rust vs Kotlin Trade-offs:**

| Aspect | Kotlin | Rust |
|--------|--------|------|
| Memory Safety | GC (runtime) | Ownership (compile-time) |
| Performance | Good (JVM overhead) | Excellent (native code) |
| Learning Curve | Gentler | Steeper (strict rules) |
| Concurrency | Possible data races | Compile-time prevention |
| Null Safety | Nullable types | Option type |
| Startup Time | Slower (JVM) | Instant (native) |

**The Rust Promise:**
> Memory safety + Performance + Thread safety with zero-cost abstractions

The learning curve is steep, but the payoff is code that's both safe and fast.

---

**Happy Learning!** 🦀

*Remember: The Rust compiler is your teacher. Read error messages carefully - they're designed to help you learn.*
