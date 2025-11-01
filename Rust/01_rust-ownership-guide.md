# Rust Ownership & Borrowing - Complete Reference Guide

## Table of Contents
1. [Memory Management Fundamentals](#memory-management-fundamentals)
2. [Ownership Rules](#ownership-rules)
3. [Move Semantics](#move-semantics)
4. [Borrowing Rules](#borrowing-rules)
5. [Rust vs Kotlin Comparison](#rust-vs-kotlin-comparison)
6. [Common Gotchas](#common-gotchas)
7. [Quick Reference](#quick-reference)

---

## Memory Management Fundamentals

### Stack vs Heap

**Stack:**
- Fast, automatic memory management
- Fixed-size data known at compile time
- LIFO (Last In, First Out) structure
- Examples: `i32`, `bool`, `char`, tuples of fixed types

**Heap:**
- Slower, requires allocation
- Dynamic-size data or data unknown at compile time
- Requires explicit management
- Examples: `String`, `Vec<T>`, `Box<T>`

### Memory Layout Visualization

```
STACK (Fast)              HEAP (Flexible)
┌──────────────┐         ┌─────────────────┐
│ x = 42       │         │                 │
│ ptr ─────────┼────────→│ "hello" data    │
│ len = 5      │         │ (String/Vec)    │
└──────────────┘         └─────────────────┘
```

### Three Approaches to Memory Management

| Language | Strategy | Pros | Cons |
|----------|----------|------|------|
| **C/C++** | Manual (malloc/free) | Full control, fast | Memory leaks, crashes, security bugs |
| **Kotlin/Java** | Garbage Collector | Safe, automatic | Runtime pauses, overhead |
| **Rust** | Ownership (compile-time) | Safe + Fast, no GC | Learning curve, stricter rules |

---

## Ownership Rules

Rust's ownership system is enforced at **compile time** with zero runtime cost.

### The Three Rules

1. **Each value in Rust has an owner**
2. **There can only be one owner at a time**
3. **When the owner goes out of scope, the value is dropped (freed)**

### Example

```rust
fn main() {
    let s = String::from("hello");  // s owns the String
    
    // ... use s ...
    
}  // s goes out of scope → heap memory automatically freed
```

### Key Insights

- **Automatic cleanup:** No need for `free()` or waiting for GC
- **Deterministic:** You know exactly when memory is freed
- **Safe:** Compiler prevents use-after-free and double-free bugs

---

## Move Semantics

### The Problem: Double Free

If Rust allowed copying pointers (like Kotlin does), both variables would try to free the same memory:

```rust
let s1 = String::from("hello");
let s2 = s1;  // If both had pointers to same data...
}  // Both s1 and s2 try to free → DOUBLE FREE BUG! 💥
```

### Rust's Solution: Move Ownership

```rust
let s1 = String::from("hello");
let s2 = s1;  // Ownership MOVES to s2, s1 becomes invalid

println!("{}", s1);  // ❌ COMPILE ERROR: value borrowed after move
println!("{}", s2);  // ✅ Works fine
```

**What happens:**
1. Heap data stays in place (not copied)
2. Ownership transfers from `s1` to `s2`
3. `s1` becomes invalid (compiler enforces this)
4. Only `s2` will free the memory at end of scope

### Memory After Move

```
BEFORE MOVE:
STACK:                    HEAP:
┌──────────┐             ┌─────────────┐
│ s1       │             │             │
│ ptr ─────┼────────────→│  "hello"    │
└──────────┘             └─────────────┘

AFTER MOVE (let s2 = s1):
STACK:                    HEAP:
┌──────────┐             ┌─────────────┐
│ s1       │ (INVALID)   │             │
├──────────┤             │  "hello"    │
│ s2       │             │             │
│ ptr ─────┼────────────→│  (same data)│
└──────────┘             └─────────────┘
```

### Copy Types (Stack-Only Data)

Small, fixed-size types implement the `Copy` trait and are copied instead of moved:

```rust
let x = 5;       // i32 implements Copy
let y = x;       // x is copied, not moved
println!("{}", x);  // ✅ Still works! Both x and y are valid
```

**Types that implement Copy:**
- All integers: `i32`, `u64`, etc.
- `bool`, `char`
- Floating points: `f32`, `f64`
- Tuples of Copy types: `(i32, i32)`

---

## Borrowing Rules

**Borrowing** allows multiple parts of code to access data without taking ownership.

### Two Types of References

1. **Immutable Reference (`&T`):** Read-only access
2. **Mutable Reference (`&mut T`):** Read-write access

### The Core Borrowing Rules

At any given time, you can have **EITHER:**
- **Any number of immutable references (`&`)** - multiple readers
- **OR exactly one mutable reference (`&mut`)** - single writer

**But NOT both at the same time!**

### Immutable Borrowing (Multiple Readers)

```rust
let s = String::from("hello");

let r1 = &s;  // OK
let r2 = &s;  // OK - multiple immutable borrows allowed
let r3 = &s;  // OK

println!("{}, {}, {}", r1, r2, r3);  // All can read simultaneously
```

### Mutable Borrowing (Single Writer)

```rust
let mut s = String::from("hello");

let r1 = &mut s;  // OK - one mutable borrow
r1.push_str(" world");

// let r2 = &mut s;  // ❌ ERROR: cannot borrow as mutable more than once
```

### Cannot Mix Readers and Writers

```rust
let mut s = String::from("hello");

let r1 = &s;        // Immutable borrow (reader)
let r2 = &s;        // Another reader - OK
let r3 = &mut s;    // ❌ ERROR: cannot borrow as mutable while immutable borrows exist

println!("{}, {}", r1, r2);
```

**Why?** If `r3` modifies the data while `r1` and `r2` are reading, they could see inconsistent state or freed memory.

### Reference Lifetimes (Last Use Matters)

The borrow checker is smart - references are valid until their **last use**, not just until the scope ends:

```rust
let mut s = String::from("hello");

let r1 = &s;
let r2 = &s;
println!("{}, {}", r1, r2);  // Last use of r1 and r2

// r1 and r2 are no longer in use here
let r3 = &mut s;  // ✅ OK now! No active immutable borrows
r3.push_str(" world");
```

### Who Frees the Memory?

```rust
fn main() {
    let s = String::from("hello");  // s OWNS the data
    
    let r1 = &s;   // r1 BORROWS (just a pointer)
    let r2 = &s;   // r2 BORROWS (just a pointer)
    
}  // s goes out of scope → ONLY s frees the heap memory
   // r1 and r2 just disappear (they don't own anything)
```

**Key point:** Borrowers don't free memory - only the owner does.

---

## Rust vs Kotlin Comparison

### Multiple References to Same Data

**Kotlin:**
```kotlin
val s1 = "hello"
val s2 = s1  // Both point to same data
val s3 = s1  // All valid, GC handles cleanup

println("$s1, $s2, $s3")  // ✅ All work
```

**Rust (Ownership):**
```rust
let s1 = String::from("hello");
let s2 = s1;  // Ownership moves to s2

println!("{}", s1);  // ❌ ERROR: s1 no longer valid
```

**Rust (Borrowing):**
```rust
let s1 = String::from("hello");
let s2 = &s1;  // Borrow, not move
let s3 = &s1;  // Another borrow

println!("{}, {}, {}", s1, s2, s3);  // ✅ All work
```

### Type System Differences

**Kotlin:**
```kotlin
val s1: String = "hello"
val s2: String = s1  // Same type, both "own" in GC sense
```

**Rust:**
```rust
let s1: String = String::from("hello");   // Owner type
let s2: &String = &s1;                     // Borrower type (different!)
```

The **type system** encodes ownership:
- `String` = owner (must free memory)
- `&String` = borrower (just a pointer)

### Memory Management Timing

| Aspect | Kotlin | Rust |
|--------|--------|------|
| When cleaned up | Unpredictable (GC decides) | Predictable (end of scope) |
| Who decides | Runtime (GC) | Compiler (at compile time) |
| Runtime cost | Yes (GC pauses) | No (zero cost) |
| Developer effort | Low | Higher (stricter rules) |

---

## Common Gotchas

### 1. Moving by Accident

```rust
let s1 = String::from("hello");
some_function(s1);  // s1 moved into function
println!("{}", s1);  // ❌ ERROR: s1 no longer valid
```

**Solution:** Borrow instead of moving:
```rust
let s1 = String::from("hello");
some_function(&s1);  // Borrow, don't move
println!("{}", s1);   // ✅ Still valid
```

### 2. Forgetting `mut` on Variable

```rust
let s = String::from("hello");
let r = &mut s;  // ❌ ERROR: cannot borrow as mutable
```

**Solution:** Declare the variable as mutable:
```rust
let mut s = String::from("hello");  // Add mut here
let r = &mut s;  // ✅ Now OK
```

### 3. Holding References Too Long

```rust
let mut v = vec![1, 2, 3];
let first = &v[0];   // Immutable borrow
v.push(4);           // ❌ ERROR: cannot borrow as mutable
println!("{}", first);
```

**Why?** `push()` might reallocate the vector, making `first` point to freed memory.

**Solution:** Limit the reference scope:
```rust
let mut v = vec![1, 2, 3];
{
    let first = &v[0];
    println!("{}", first);
}  // first goes out of scope
v.push(4);  // ✅ Now OK
```

### 4. Multiple Mutable Borrows

```rust
let mut s = String::from("hello");
let r1 = &mut s;
let r2 = &mut s;  // ❌ ERROR: cannot borrow more than once as mutable
```

**Why?** Two writers could cause data races and inconsistent state.

**Solution:** Use one mutable reference at a time:
```rust
let mut s = String::from("hello");
{
    let r1 = &mut s;
    r1.push_str(" world");
}  // r1 scope ends
let r2 = &mut s;  // ✅ Now OK
```

### 5. Mixing Readers and Writers

```rust
let mut s = String::from("hello");
let r1 = &s;        // Reader
let r2 = &mut s;    // Writer - ❌ ERROR
println!("{}", r1);
```

**Why?** Writer could invalidate what readers see.

**Solution:** Ensure no readers exist when you need a writer:
```rust
let mut s = String::from("hello");
let r1 = &s;
println!("{}", r1);  // Last use of r1

let r2 = &mut s;  // ✅ OK - r1 no longer active
```

### 6. Returning References Without Lifetimes

```rust
fn dangle() -> &String {  // ❌ ERROR: missing lifetime
    let s = String::from("hello");
    &s  // s is dropped, returning dangling reference!
}
```

**Solution:** Return owned data:
```rust
fn no_dangle() -> String {
    let s = String::from("hello");
    s  // Move ownership out
}
```

---

## Quick Reference

### Ownership Quick Facts

- ✅ Each value has exactly one owner
- ✅ Owner is responsible for cleanup
- ✅ When owner goes out of scope, value is dropped
- ✅ Ownership can be transferred (moved)
- ✅ Small fixed-size types are copied, not moved

### Borrowing Quick Facts

- ✅ `&T` = immutable reference (read-only)
- ✅ `&mut T` = mutable reference (read-write)
- ✅ Multiple readers OR one writer (not both)
- ✅ References valid until last use
- ✅ Only owner frees memory, not borrowers

### Syntax Cheat Sheet

```rust
// OWNERSHIP
let s = String::from("hello");      // s owns the String
let s2 = s;                         // Move: s invalid, s2 owns now

// IMMUTABLE BORROWING
let r1 = &s;                        // Borrow immutably
let r2 = &s;                        // Multiple OK

// MUTABLE BORROWING
let mut s = String::from("hello");  // Must be mut
let r = &mut s;                     // Borrow mutably (only one)

// FUNCTION SIGNATURES
fn read(s: &String)                 // Borrows (doesn't take ownership)
fn modify(s: &mut String)           // Borrows mutably
fn consume(s: String)               // Takes ownership
```

### Decision Tree

```
Need to modify data?
├─ NO  → Use immutable borrow (&)
└─ YES → Need exclusive access?
         ├─ YES → Use mutable borrow (&mut)
         └─ NO  → Need multiple owners?
                  ├─ YES → Use Rc<T> or Arc<T> (advanced)
                  └─ NO  → Transfer ownership (move)
```

---

## The Big Picture

**Rust's ownership system solves the memory management trilemma:**

1. **Memory safety** ✅ (like GC languages)
2. **No runtime overhead** ✅ (like manual management)
3. **Ease of use** ⚠️ (requires learning the rules)

The trade-off: You must learn and follow the ownership/borrowing rules, but in return you get:
- Zero-cost abstractions
- No garbage collector pauses
- Thread safety (data race freedom)
- Compile-time guarantees (not runtime crashes)

---

## Connection to Kotlin

| Concept | Kotlin | Rust |
|---------|--------|------|
| **Multiple refs** | Always allowed | Only via borrowing |
| **Type system** | Reference types uniform | Owner vs borrower types |
| **Memory cleanup** | GC (runtime) | Ownership (compile-time) |
| **Data races** | Possible | Prevented by borrow checker |
| **Performance** | GC overhead | Zero overhead |
| **Learning curve** | Gentler | Steeper (strict rules) |

**Key Insight:** Rust moves memory safety checks from **runtime to compile time**, trading developer convenience for performance and safety guarantees.

---

## Next Steps

After mastering ownership and borrowing, explore:
1. **Lifetimes** - Explicit lifetime annotations
2. **Smart Pointers** - `Box<T>`, `Rc<T>`, `Arc<T>`
3. **Interior Mutability** - `RefCell<T>`, `Mutex<T>`
4. **Traits and Generics** - The Copy and Clone traits
5. **Concurrency** - Fearless concurrency with ownership

---

*Last Updated: 2025-11-01*
*Created from first principles learning session*
