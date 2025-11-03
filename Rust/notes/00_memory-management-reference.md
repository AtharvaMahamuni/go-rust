# Memory Management - Quick Reference Notes

## The Two Memory Regions

### Stack Memory
```
┌─────────────────────────────────────────────────┐
│  STACK                                          │
├─────────────────────────────────────────────────┤
│  • Fast (1ns access)                            │
│  • Small (~1-8 MB limit)                        │
│  • LIFO (Last In, First Out)                    │
│  • Automatic cleanup when scope ends            │
│  • Stores:                                      │
│    - Local variables with known size            │
│    - Function call info (return addresses)      │
│    - References/pointers to heap                │
│                                                 │
│  Example: val x = 42                            │
└─────────────────────────────────────────────────┘
```

### Heap Memory
```
┌─────────────────────────────────────────────────┐
│  HEAP                                           │
├─────────────────────────────────────────────────┤
│  • Slower (~100ns access)                       │
│  • Large (GBs available)                        │
│  • Random access                                │
│  • Needs manual or automatic cleanup            │
│  • Stores:                                      │
│    - Large data                                 │
│    - Dynamic size data                          │
│    - Data that outlives function scope          │
│                                                 │
│  Example: val arr = IntArray(1000000)           │
└─────────────────────────────────────────────────┘
```

## Stack vs Heap Decision Tree

```
Is size known at compile time?
├─ YES ─→ Is it small (< few KB)?
│         ├─ YES ─→ STACK
│         └─ NO ──→ HEAP (too big for stack)
└─ NO ──→ HEAP (dynamic size)
```

## Memory Layout Visualization

```
MEMORY LAYOUT when running a function:

┌─────────────────────────────────────┐
│           STACK                      │  (Small, Fast, Auto-managed)
│                                      │
│  ┌──────────────────────┐           │
│  │ fn example()         │           │
│  │  x = 42    ───────┐  │           │  <- Value IS here
│  │  numRef = 0x1A2B ─┼──┼─┐         │  <- Just an address
│  │  userRef = 0x3C4D─┼──┼─┼─┐       │  <- Just an address
│  └──────────────────────┘ │ │ │     │
└────────────────────────────┼─┼─┼─────┘
                             │ │ │
┌────────────────────────────┼─┼─┼─────┐
│           HEAP              │ │ │     │  (Large, Slower, Manual/GC)
│                             ↓ │ │     │
│  [1,2,3,4,5,6...]  <- Array  │ │     │
│                               ↓ │     │
│  User {                         │     │
│    name: "Bob"  <───────────────┘     │
│    age: 25                            │
│  }                                    │
└───────────────────────────────────────┘
```

## Memory Management Strategies

| Language | Strategy | Pros | Cons |
|----------|----------|------|------|
| **C/C++** | Manual (malloc/free, new/delete) | Fast, full control | Easy to leak/crash |
| **Java/Kotlin/Python** | Garbage Collector | Safe, automatic | Runtime cost, pauses |
| **Rust** | Ownership (compile-time) | Safe + Fast + No GC | Learning curve |

## Key Concepts

### Stack Overflow
Happens when you exceed stack limit due to:
- Too many function calls (deep recursion)
- Large local variables on stack

### Memory Leak
Heap memory that's allocated but never freed. 
- In GC languages: Prevented automatically
- In manual languages: Developer must remember to free

### Garbage Collection
Runtime process that finds and frees unused heap objects by:
1. **Scanning** for unreachable objects
2. **Marking** what's in use (reachable from stack)
3. **Sweeping** (freeing) what's not referenced

**Cost:** Periodic pauses, runtime overhead

## Kotlin Example

```kotlin
fun example() {
    // Stack:
    val x = 42                    // Value on stack
    
    // Heap (with stack reference):
    val user = User("Alice")      // userRef on stack → User object on heap
    val numbers = IntArray(1000)  // numbersRef on stack → Array on heap
    
} // Stack clears immediately, heap objects cleaned by GC later
```

## The Rust Innovation

**Question:** What if the compiler could figure out at **compile time**:
- When to allocate heap memory
- When to free it (without GC!)
- And guarantee no memory leaks or use-after-free bugs?

**Answer:** This is Rust's **Ownership System** - the topic we're learning next!

---

## Stack Growth Example

```
When main() starts:
STACK: [main's variables, main's return address]

When main() calls helper():
STACK: [main's variables, main's return address]
       [helper's variables, helper's return address] <- Added on top!

When helper() returns:
STACK: [main's variables, main's return address]  <- Top removed!
```

## Speed Comparison

| Operation | Stack | Heap |
|-----------|-------|------|
| Allocation | ~1 ns (just move pointer) | ~100 ns (search for space) |
| Access | Very fast (contiguous memory) | Slower (scattered memory) |
| Deallocation | Instant (move pointer back) | Complex (GC scan or manual free) |

## Why Stack is Faster

**Stack:** 
- Always add to top
- Always remove from top
- No searching, no bookkeeping
- Simple pointer arithmetic

**Heap:**
- Find free space big enough
- Mark as used
- Track allocation metadata
- Later: find and free it

---

**Last Updated:** Learning Rust - Session 1
**Next Topic:** Rust Ownership System
