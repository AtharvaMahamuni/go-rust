# Rust Learning Session 2 - Variables & Shadowing

**📅 Date:** November 2, 2025
**🏷️ Topics:** Mutability, Shadowing, Type Transformations, Immutability by Default
**🔗 Related:** [03_shadowing_visualization.md](./03_shadowing_visualization.md), [01_rust-ownership-guide.md](./01_rust-ownership-guide.md), [02_rust-basics-and-execution.md](./02_rust-basics-and-execution.md)

---

## Overview

This session explores Rust's unique approach to variables: immutability by default with the `let` keyword and explicit mutability with `let mut`. It covers the powerful shadowing feature that allows type transformations and data pipelines, including comparisons with Kotlin and real-world patterns. Essential for understanding Rust's safety-first design philosophy.

## Session Overview

This session covered Rust's approach to variable mutability and shadowing - two foundational concepts that shape how you write Rust code.

**Files Created:**
- `01_mutability_basics.rs` - Immutable by default, explicit mutation
- `02_shadowing.rs` - Variable shadowing and type transformations
- `shadowing_visualization.md` - Stack/heap behavior during shadowing
- `PROJECT_INSTRUCTIONS.md` - Learning format template

---

## Key Concepts Covered

### 1. Immutability by Default

**Rust's Philosophy:**
> Variables are IMMUTABLE by default - mutation is opt-IN, not opt-OUT

**Syntax Comparison:**

| Concept | Kotlin | Rust |
|---------|--------|------|
| Immutable | `val x = 5` | `let x = 5` |
| Mutable | `var x = 5` | `let mut x = 5` |
| Default | `var` (mutable) | `let` (immutable) |

**Key Insight:** Rust makes you explicitly choose mutability with `mut` keyword. This makes mutation **visible** in code and encourages immutable-first thinking.

**Example:**
```rust
let x = 5;     // Immutable
x = 6;         // ❌ Error: cannot assign twice to immutable variable

let mut y = 5; // Mutable  
y = 6;         // ✅ Works
```

---

### 2. Mutability and Concurrency

**The Threading Problem:**

In Kotlin:
```kotlin
var counter = 0
// Thread 1: counter++ 1000 times
// Thread 2: counter++ 1000 times
// Result: Inconsistent! (race condition)
```

**Rust's Approach:**
- Mutable variables (`mut`) compile without issue for single-threaded code
- Sharing mutable state across threads **won't compile** without explicit synchronization
- Error caught at **compile-time**, not runtime!

**Key Insight:** Rust's ownership system prevents data races at compile-time. You'll learn this in detail when covering ownership and borrowing.

---

### 3. Variable Shadowing

**What is Shadowing?**

Creating a NEW binding with the same name - the old binding's scope ends immediately.

**Syntax:**
```rust
let x = 5;
let x = x + 1;  // NEW x created, OLD x scope ended
let x = x * 2;  // Another NEW x, previous x scope ended
```

**How It Works:**

```
Step 1: let x = 5;
  → x₁ binding created with value 5

Step 2: let x = x + 1;
  → Read x₁ (value 5)
  → Calculate 5 + 1 = 6
  → Create NEW x₂ binding with value 6
  → x₁'s scope ENDS (dropped if heap-allocated)

Step 3: let x = x * 2;
  → Read x₂ (value 6)
  → Calculate 6 * 2 = 12
  → Create NEW x₃ binding with value 12
  → x₂'s scope ENDS
```

---

### 4. Shadowing vs Mutation

**When to use each:**

| Use Case | Use This | Why |
|----------|----------|-----|
| Counter, accumulator | `mut` | Same thing, changing over time |
| Type transformation | Shadowing | Converting to something new |
| Data pipeline | Shadowing | Multiple transformation steps |
| Loop variable | `mut` | Single value being updated |

**Example - Mutation:**
```rust
let mut count = 0;
count += 1;  // Modifying existing value
count += 1;
```

**Example - Shadowing:**
```rust
let value = "42";              // &str
let value = value.trim();      // &str (trimmed)
let value = value.parse::<i32>().unwrap();  // i32
// Same name, different types through transformation!
```

---

### 5. The Power: Type Changes

**Impossible in Kotlin:**
```kotlin
val x = "5"
val x = x.toInt()  // ❌ Error: x already declared
// Must use: xStr and xInt (two names)
```

**Possible in Rust:**
```rust
let x = "5";                    // Type: &str
let x = x.parse::<i32>().unwrap();  // Type: i32  ✅
// Same name, different type!
```

**Why is this safe?**
- Each `let` creates a NEW binding (new variable)
- Old binding's scope ends (no name collision)
- Each binding has its own type
- All checked at compile-time (type safe)

**Real-World Use:**
```rust
let config = "timeout=30";                    // Raw string
let config = config.split('=').collect::<Vec<&str>>();  // Vector
let config = config[1];                       // String slice
let config: u32 = config.parse().unwrap();    // Integer
let config = config * 1000;                   // Transformed int
// Same logical concept = same variable name!
```

---

### 6. Memory Behavior

**Stack Values (integers):**
- Old bindings stay on stack until function returns
- No harm (compiler might optimize and reuse)
- No explicit cleanup needed

**Heap Values (String, Vec, etc.):**
- When shadowed, old binding's scope ends
- Ownership ends → Value DROPPED immediately
- Heap memory freed right away (no memory leak)

**Example:**
```rust
let x = String::from("hello");  // Allocates heap memory
let x = String::from("world");  // OLD x dropped HERE! "hello" freed
// "world" still alive
// Function ends: "world" dropped
```

---

### 7. Shadowing and Scopes

**Inner scopes can shadow temporarily:**

```rust
let x = 5;
{
    let x = 10;  // Shadows outer x
    println!("{}", x);  // Prints 10
}
println!("{}", x);  // Prints 5 (outer x back!)
```

**Key Insight:** Shadowing in inner scopes is temporary. Original binding becomes accessible again when inner scope ends.

---

## Rust vs Kotlin vs Python Comparison

| Feature | Python | Kotlin | Rust |
|---------|--------|--------|------|
| Variable mutability | All mutable | `var` (mutable), `val` (immutable ref) | `let` (immutable), `let mut` (mutable) |
| Default | Mutable | Mutable (`var`) | Immutable (`let`) |
| Type changes | ✅ Anytime (dynamic) | ❌ Must use new name | ✅ With shadowing (compile-time safe) |
| Name reuse | ✅ Always | ❌ Never | ✅ With `let` (shadowing) |
| Safety | Runtime | Compile-time (for types) | Compile-time (types + memory) |

---

## Key Takeaways

1. **Rust is immutable by default** - safer concurrent code, clearer mutation intent
2. **`mut` keyword** makes mutation explicit and visible
3. **Shadowing creates new bindings** - old binding's scope ends immediately
4. **Type changes through shadowing** - impossible in Kotlin, safe in Rust
5. **Shadowing ≠ mutation** - shadowing transforms, mutation modifies
6. **Heap cleanup is immediate** when shadowing heap-allocated values
7. **Inner scope shadows are temporary** - original binding returns

---

## Design Philosophy Insights

### Why Immutable by Default?

**Rust's Goal:** Catch bugs at compile-time, not runtime

**Benefits:**
- Easier to reason about code (value doesn't change unexpectedly)
- Safer concurrent code (fewer race conditions)
- Compiler optimizations (immutable = cacheable)
- Forces you to think before adding mutation

**Trade-off:** More explicit code (must write `mut`) but clearer intent

---

### Why Allow Shadowing?

**The Problem:** 
- Python: Type changes are unsafe (runtime errors)
- Kotlin: Must use different names (verbose: `xStr`, `xInt`)

**Rust's Solution:**
- Each `let` creates new binding (compile-time checked)
- Can have different types (transformation safe)
- Same name (code stays clean and readable)

**Result:** Safety of static typing + conciseness of scripting languages

---

## Common Patterns

### Pattern 1: Data Transformation Pipeline
```rust
let data = get_raw_input();       // String
let data = data.trim();           // Trimmed string
let data = data.parse().unwrap(); // Parsed type
let data = transform(data);       // Business logic
```

### Pattern 2: Configuration Processing
```rust
let config = read_file("config.txt");
let config = parse_config(config);
let config = validate(config);
let config = apply_defaults(config);
```

### Pattern 3: Type Narrowing
```rust
let value: &str = "42";           // Wide type
let value: i32 = value.parse()?;  // Narrow to specific type
```

---

## Practice Exercises Completed

1. ✅ Basic mutability with `let` and `let mut`
2. ✅ Understanding race conditions in concurrent code
3. ✅ Variable shadowing with same types
4. ✅ Variable shadowing with type changes
5. ✅ Data transformation pipelines
6. ✅ Shadowing vs mutation decision making
7. ✅ Inner scope shadowing behavior
8. ✅ Real-world config processing example

---

## Common Gotchas Avoided

**Gotcha 1:** Thinking shadowing is mutation
- **Wrong:** "Shadowing changes the variable"
- **Right:** "Shadowing creates a new variable; old one's scope ends"

**Gotcha 2:** Forgetting `mut` for mutation
- Compiler error is helpful: "consider making this binding mutable: `mut x`"

**Gotcha 3:** Confusing Kotlin's `val` with Rust's `let`
- Kotlin `val`: Immutable reference (can't rebind)
- Rust `let`: Immutable binding (can shadow with new `let`)

---

## Mental Models

### Reference vs Value Immutability

**Kotlin:**
```kotlin
val list = mutableListOf(1, 2, 3)
list.add(4)  // ✅ Works - reference is immutable, data is mutable
list = mutableListOf(5)  // ❌ Error - can't reassign reference
```

**Rust:**
```rust
let mut list = vec![1, 2, 3];
list.push(4);  // ✅ Works - mutable binding
list = vec![5];  // ✅ Works - can reassign mutable binding

let list = vec![1, 2, 3];
list.push(4);  // ❌ Error - can't mutate immutable binding
```

**Key Difference:** 
- Kotlin's `val` protects the REFERENCE
- Rust's `let` protects the BINDING and VALUE

---

## Next Steps

In Session 3, we'll explore:
1. **Data Types** - Scalar types (integers, floats, booleans, chars)
2. **Functions** - Basic syntax, parameters, return values, expressions
3. **Pattern Matching** - Destructuring tuples and compound types

**Choose your path based on what you want to tackle next!**

---

## Quick Reference

### Mutability
```rust
let x = 5;        // Immutable
let mut y = 5;    // Mutable
```

### Shadowing
```rust
let x = 5;
let x = x + 1;    // New binding, old scope ended
let x = "hello";  // Different type - still works!
```

### When to Use
- **`let`**: Default for everything
- **`let mut`**: When value changes over time (counters, accumulators)
- **Shadowing**: When transforming to new value/type

---

## Resources Used

- The Rust Programming Language (Chapter 3)
- Project knowledge base
- Hands-on code examples and experiments

---

## Summary

**Key Takeaways:**
- Rust is immutable by default (`let`) - mutation requires explicit `mut` keyword
- Shadowing creates new bindings - allows type changes safely
- `let mut` for values that change over time (counters, accumulators)
- Shadowing for transformations and type conversions (data pipelines)
- Heap-allocated values are dropped immediately when shadowed
- Inner scope shadowing is temporary - original binding returns
- Type changes through shadowing are impossible in Kotlin, safe in Rust

**Quick Reference:**
```rust
let x = 5;           // Immutable
let mut y = 5;       // Mutable
let x = x + 1;       // Shadow (new binding)
let x = "hello";     // Shadow with type change
```

**Common Patterns:**
```rust
// Data transformation pipeline
let data = get_input();
let data = data.trim();
let data = data.parse().unwrap();

// Configuration processing
let config = read_file();
let config = parse(config);
let config = validate(config);
```

**When to Use This Note:**
- Understanding Rust's immutability-first philosophy
- Deciding between mutation and shadowing
- Learning type transformation patterns
- Comparing Rust with Kotlin variable handling
- Reference for common shadowing patterns

---

**Session Complete!** ✅
**📝 Date:** November 2, 2025
**➡️ Next Session:** Data Types and Functions (Session 3)
