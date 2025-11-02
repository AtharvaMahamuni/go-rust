# Shadowing: Stack and Heap Visualization

**📅 Created:** Session 2
**🏷️ Topics:** Shadowing, Stack Memory, Heap Memory, Memory Cleanup
**🔗 Related:** [04_variables_shadowing.md](./04_variables_shadowing.md), [01_rust-ownership-guide.md](./01_rust-ownership-guide.md)

---

## Overview

This note provides detailed visualizations of how shadowing works at the memory level, showing what happens on both the stack and heap when you shadow variables. It clarifies when memory is freed and what the compiler guarantees vs optimizes. Essential for understanding the memory implications of shadowing.

---

## Scenario 1: Stack Values (Integers)

```rust
{
    let x = 5;
    let x = x + 1;
    let x = x * 2;
}
```

### Timeline Visualization:

```
Step 1: let x = 5;
┌─────────────────┐
│   STACK         │
├─────────────────┤
│ x = 5          │  ← x binding created
└─────────────────┘

Step 2: let x = x + 1;
         ↓ Read x (gets 5)
         ↓ Calculate 5 + 1 = 6
         ↓ Create NEW x binding
         ↓ Old x binding scope ENDS
         
┌─────────────────┐
│   STACK         │
├─────────────────┤
│ x = 6          │  ← NEW x binding (old one's scope ended)
└─────────────────┘

Step 3: let x = x * 2;
         ↓ Read x (gets 6)
         ↓ Calculate 6 * 2 = 12
         ↓ Create NEW x binding
         ↓ Previous x binding scope ENDS
         
┌─────────────────┐
│   STACK         │
├─────────────────┤
│ x = 12         │  ← NEW x binding
└─────────────────┘

Step 4: Closing brace }
         ↓ x goes out of scope
         
┌─────────────────┐
│   STACK         │
├─────────────────┤
│ (empty/reused) │
└─────────────────┘
```

**Key Point:** Whether the compiler reuses the same 4-byte stack slot or uses new ones is an OPTIMIZATION DETAIL. From Rust's perspective, each `let x` creates a new binding and the old one's scope ends immediately.

---

## Scenario 2: Heap Values (String)

```rust
{
    let x = String::from("hello");
    let x = String::from("world");
}
```

### Timeline Visualization:

```
Step 1: let x = String::from("hello");

STACK                           HEAP
┌─────────────────┐            ┌──────────────┐
│ x:              │            │              │
│  ptr ───────────┼───────────→│ "hello"      │
│  len: 5         │            │              │
│  cap: 5         │            └──────────────┘
└─────────────────┘
     ↑ x binding owns this heap allocation


Step 2: let x = String::from("world");
         ↓ Create NEW String on heap
         ↓ Create NEW x binding
         ↓ OLD x binding scope ENDS ← IMPORTANT!
         ↓ Old x's ownership ends
         ↓ "hello" has no owner → DROP IT NOW!
         
STACK                           HEAP
┌─────────────────┐            ┌──────────────┐
│ x:              │            │ [freed]      │ ← "hello" cleaned up!
│  ptr ───────────┼───────┐    └──────────────┘
│  len: 5         │       │    
│  cap: 5         │       │    ┌──────────────┐
└─────────────────┘       └───→│ "world"      │
     ↑ NEW x binding            └──────────────┘


Step 3: Closing brace }
         ↓ x goes out of scope
         ↓ "world" has no owner → DROP IT NOW!
         
STACK                           HEAP
┌─────────────────┐            ┌──────────────┐
│ (empty)         │            │ [freed]      │
└─────────────────┘            └──────────────┘
```

---

## The Key Insight

### When does cleanup happen?

**Rust's Rule (one rule for everything):**
> "When a binding goes out of scope, its value is dropped"

**For shadowing specifically:**
- Old binding's scope ends when new `let` with same name executes
- If the value needs cleanup (heap allocation), it happens IMMEDIATELY
- If the value doesn't need cleanup (stack integer), the memory just sits there until the function returns (no harm, compiler might optimize)

---

## Memory Comparison

### What the compiler GUARANTEES (Ownership Rules):

```
let x = String::from("hello");  // x₁ owns "hello"
let x = String::from("world");  // x₁'s scope ends, "hello" dropped
                                 // x₂ owns "world"
```

### What the compiler MIGHT optimize (Implementation Detail):

For integers:
```rust
let x = 5;
let x = 6;
```

The compiler could:
- Option A: Use different stack slots → [5][6]
- Option B: Reuse same stack slot → [6] (5 overwritten)
- You don't care! Both are correct as long as you can't access the old value.

---

## Why This Matters

**Kotlin comparison:**
```kotlin
val x = "hello"
// val x = "world"  // ❌ Can't do this - "x already declared"
```

**Rust shadowing allows:**
```rust
let x = String::from("hello");
let x = String::from("world");  // ✅ Works! Old x's scope ended
```

**Even more powerful - TYPE CHANGES:**
```rust
let x = "5";           // &str (string)
let x = x.parse::<i32>().unwrap();  // i32 (integer)
// Same name, DIFFERENT TYPE! Impossible in Kotlin!
```

This is safe because:
1. Old binding's scope ends (no confusion about which x)
2. New binding can be completely different type
3. All checked at compile-time (type safe)

---

## Summary

**Key Takeaways:**
- Shadowing creates a NEW binding; old binding's scope ends immediately
- Stack values: Compiler may optimize and reuse memory slots
- Heap values: Memory is freed IMMEDIATELY when shadowed (no leak)
- Rust's ownership rule: "When a binding goes out of scope, its value is dropped"
- Type changes are safe with shadowing because each binding is independent
- All behavior is checked at compile-time for safety

**Visual Memory Model:**
```
Stack values:   [x₁=5] → [x₂=6]  (may reuse slot)
Heap values:    x₁→"hello" → [freed] when x₂→"world" created
```

**When to Use This Note:**
- When confused about shadowing's memory behavior
- To understand when heap memory is freed
- To differentiate between compiler guarantees and optimizations
- When learning about ownership and Drop trait

---

**📝 Last Updated:** Session 2
**➡️ Next:** See [04_variables_shadowing.md](./04_variables_shadowing.md) for complete shadowing guide and practical examples
