# Rust Learning Notes - Index

This directory contains comprehensive reference notes and session summaries for learning Rust. All notes are organized by topic and session.

---

## Quick Navigation

| Category | Topic | Description |
|----------|-------|-------------|
| **Fundamentals** | [Memory Management](#memory-management) | Stack vs Heap, Memory strategies |
| **Core Concepts** | [Ownership & Borrowing](#ownership--borrowing) | Rust's ownership system explained |
| **Language Features** | [Variables & Shadowing](#variables--shadowing) | Mutability and shadowing patterns |
| **Session Notes** | [Learning Sessions](#learning-sessions) | Session-by-session progress |

---

## Memory Management

### [00_memory-management-reference.md](./00_memory-management-reference.md)
**Core Topic:** Stack vs Heap, Memory Management Strategies

**Key Concepts Covered:**
- Stack memory characteristics (fast, LIFO, automatic)
- Heap memory characteristics (slower, dynamic, manual/GC)
- Memory layout visualization
- Comparison of memory strategies (C/C++, Java/Kotlin, Rust)
- Why Rust's ownership system is needed

**When to Read:** Start here for foundational understanding before ownership

**Quick Reference:**
```
Stack: Fast, fixed-size, automatic cleanup
Heap:  Slower, dynamic-size, needs management
Rust:  Ownership system = compile-time safety + no GC
```

---

## Ownership & Borrowing

### [01_rust-ownership-guide.md](./01_rust-ownership-guide.md)
**Core Topic:** Complete reference for Rust's ownership and borrowing rules

**Key Concepts Covered:**
- Three ownership rules
- Move semantics and why they exist
- Borrowing rules (immutable & mutable references)
- Copy trait and stack-only types
- Common gotchas and solutions
- Rust vs Kotlin comparison
- Reference lifetime basics

**When to Read:** Essential reading after memory management concepts

**Quick Reference:**
```
Ownership:  Each value has ONE owner
Move:       Ownership transfers, old owner invalid
Borrowing:  Multiple readers OR one writer (not both)
References: &T (read) or &mut T (write)
```

---

## Variables & Shadowing

### [03_shadowing_visualization.md](./03_shadowing_visualization.md)
**Core Topic:** How shadowing works at the stack/heap level

**Key Concepts Covered:**
- Stack behavior during shadowing (integers)
- Heap behavior during shadowing (Strings)
- When memory is freed during shadowing
- Compiler optimization vs guaranteed behavior
- Type transformation through shadowing

**When to Read:** After basic mutability, before complex transformations

**Quick Reference:**
```
Shadowing:     New binding, old scope ends
Stack values:  May reuse memory (optimization)
Heap values:   Dropped immediately when shadowed
Type changes:  Allowed with shadowing (compile-time safe)
```

---

### [04_variables_shadowing.md](./04_variables_shadowing.md)
**Core Topic:** Complete guide to mutability and shadowing patterns

**Key Concepts Covered:**
- Immutability by default philosophy
- `let` vs `let mut`
- Shadowing vs mutation (when to use each)
- Type changes through shadowing
- Shadowing in scopes
- Real-world patterns and use cases
- Rust vs Kotlin vs Python comparison

**When to Read:** Essential early topic, before ownership deep-dive

**Quick Reference:**
```
Immutable:  let x = 5;
Mutable:    let mut x = 5;
Shadow:     let x = x + 1;  (new binding)
Mutate:     x = x + 1;      (modify existing)
```

---

## Learning Sessions

### [02_rust-basics-and-execution.md](./02_rust-basics-and-execution.md)
**Session:** 1 - Foundation & First Program
**Date:** November 2, 2025

**Topics Covered:**
- Memory management fundamentals review
- Native code vs bytecode (Rust vs Kotlin)
- Ownership & borrowing quick recap
- Rust toolchain (rustup, rustc, cargo)
- Project structure and Cargo.toml
- First Rust program (Hello World)
- Macros vs functions (`println!`)
- Build process and modes (debug/release)
- Windows setup requirements
- Compile-time vs runtime philosophy

**Quick Commands:**
```bash
cargo new project-name    # Create new project
cargo build              # Compile (debug)
cargo run               # Build + run
cargo build --release   # Optimized build
```

---

## How to Use These Notes

### For New Learners
1. Start with [00_memory-management-reference.md](./00_memory-management-reference.md)
2. Move to [01_rust-ownership-guide.md](./01_rust-ownership-guide.md)
3. Read [04_variables_shadowing.md](./04_variables_shadowing.md)
4. Use [03_shadowing_visualization.md](./03_shadowing_visualization.md) for deep understanding
5. Reference [02_rust-basics-and-execution.md](./02_rust-basics-and-execution.md) for tooling

### For Quick Reference
- **Memory questions:** 00_memory-management-reference.md
- **Ownership errors:** 01_rust-ownership-guide.md (Common Gotchas section)
- **Shadowing questions:** 03_shadowing_visualization.md
- **Cargo commands:** 02_rust-basics-and-execution.md

### For Kotlin Developers
All notes include Kotlin comparisons to help you relate concepts:
- Memory management (GC vs Ownership)
- Variable mutability (`val`/`var` vs `let`/`let mut`)
- Type system differences
- Performance trade-offs

---

## Note Format Standards

All notes in this directory follow consistent formatting:

### Structure
- **Title and metadata** at the top
- **Table of contents** for long documents
- **Code examples** with syntax highlighting
- **Visualizations** for complex concepts (stack/heap diagrams)
- **Comparisons** with Kotlin/other languages
- **Quick reference** sections for easy scanning

### Conventions
- ✅ = Correct/allowed
- ❌ = Error/not allowed
- `code` = inline code
- ```rust = code blocks
- **Bold** = important terms
- *Italic* = emphasis

---

## Future Additions

When new notes are added to this directory, follow this format:

### Template for New Entries

```markdown
### [XX_topic-name.md](./XX_topic-name.md)
**Core Topic:** Brief description

**Key Concepts Covered:**
- Concept 1
- Concept 2
- Concept 3

**When to Read:** Context for when this note is most useful

**Quick Reference:**
```
Key syntax or rules
```
```

### Naming Convention
- `XX_descriptive-topic-name.md`
- XX = Sequential number (00, 01, 02...)
- Use hyphens for multi-word topics
- Keep names descriptive but concise

---

## Related Resources

**In This Repository:**
- `../code/` - Executable Rust examples
- `../PROJECT_INSTRUCTIONS.md` - Format for creating new learning files
- `../comprehensive-rust.pdf` - Google's Rust course
- `../The Rust Programming Language (2nd Edition).pdf` - Official Rust book

**External Resources:**
- [Official Rust Book](https://doc.rust-lang.org/book/)
- [Rust by Example](https://doc.rust-lang.org/rust-by-example/)
- [Rust Playground](https://play.rust-lang.org/) - Test code online

---

**Last Updated:** November 2, 2025
**Total Notes:** 5 files
**Status:** Active learning in progress

---

*This index is maintained alongside the notes. When adding new notes, update this file following the established format.*
