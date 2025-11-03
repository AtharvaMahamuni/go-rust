# Arrays and Slices in Go

**📅 Created:** Fundamentals Series
**🏷️ Topics:** Arrays, Slices, Dynamic Arrays, append, make, Slice Operations
**🔗 Related:** [02_pointers.md](./02_pointers.md), [04_maps-and-structs.md](./04_maps-and-structs.md), [Concepts/08myarray](../Concepts/08myarray/), [Concepts/09myslices](../Concepts/09myslices/)

---

## Overview

This note covers two fundamental collection types in Go: arrays (fixed-size) and slices (dynamic-size). Understanding the difference between them and when to use each is crucial for effective Go programming. Slices are used far more commonly than arrays in real-world Go code.

---

## Arrays: Fixed-Size Collections

### What is an Array?

An array is a **fixed-size** collection of elements of the same type.

```go
var fruitList [4]string    // Array of 4 strings

fruitList[0] = "Apple"
fruitList[1] = "Mango"
fruitList[3] = "Peach"

fmt.Println(fruitList)     // [Apple Mango  Peach]
```

**Key Characteristics:**
- **Fixed size** - declared at creation, cannot change
- **Size is part of the type** - `[3]int` is different from `[4]int`
- **Zero-indexed** - first element at index 0
- **Value type** - copying creates a new array
- **Zero values** - unassigned elements get zero value

### Array Declaration

```go
// Method 1: Declare with size
var numbers [5]int
fmt.Println(numbers)  // [0 0 0 0 0]

// Method 2: Declare and initialize
var fruits = [3]string{"Apple", "Mango", "Peach"}

// Method 3: Let Go count the size
fruits := [...]string{"Apple", "Mango", "Peach"}  // Size is 3

// Method 4: Specific indexes
nums := [5]int{1: 10, 3: 30}  // [0 10 0 30 0]
```

### Array Limitations

```go
var fruitList [4]string
fruitList[0] = "Apple"

// ❌ Cannot append (size is fixed)
// fruitList = append(fruitList, "Banana")

// ❌ Cannot change size
// fruitList = fruitList[:2]

// ✅ Can create slice from array
fruitSlice := fruitList[:2]  // Now a slice!
```

**Why Arrays Are Rarely Used:**
- Fixed size is too restrictive
- Can't grow or shrink
- Slices are more flexible and just as efficient

---

## Slices: Dynamic Collections

### What is a Slice?

A slice is a **dynamic, flexible** view into an array.

```go
var fruitList = []string{"Apple", "Tomato", "Mango"}
fruitList = append(fruitList, "Banana")  // ✅ Can grow!
```

**Key Characteristics:**
- **Dynamic size** - can grow and shrink
- **Reference type** - like a pointer to an underlying array
- **Three components:** pointer, length, capacity
- **Most common** collection type in Go

### Slice Structure

```
Slice internals:
┌─────────┬─────────┬──────────┐
│ pointer │ length  │ capacity │
│   ptr   │  len=3  │  cap=5   │
└────┬────┴─────────┴──────────┘
     │
     └──→ [Apple][Mango][Peach][ ][ ]
          └──── len ────┘
          └─────── cap ────────────┘
```

- **Length:** Number of elements in the slice
- **Capacity:** Number of elements in underlying array (from slice start)

### Creating Slices

#### Method 1: Literal Declaration

```go
fruits := []string{"Apple", "Mango", "Peach"}
fmt.Println(fruits)  // [Apple Mango Peach]
```

#### Method 2: Using `make`

```go
// make([]Type, length, capacity)
highScores := make([]int, 4)       // len=4, cap=4
highScores := make([]int, 4, 10)   // len=4, cap=10

highScores[0] = 123
highScores[1] = 213
highScores[2] = 312
highScores[3] = 321

fmt.Println(highScores)  // [123 213 312 321]
```

#### Method 3: From Arrays

```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4]  // [2 3 4]
```

#### Method 4: Nil Slice

```go
var slice []int    // nil slice (len=0, cap=0)
fmt.Println(slice == nil)  // true
```

---

## Slice Operations

### Appending Elements

```go
fruits := []string{"Apple", "Mango"}
fruits = append(fruits, "Banana")
fruits = append(fruits, "Peach", "Orange")

fmt.Println(fruits)  // [Apple Mango Banana Peach Orange]
```

**Important:** `append` returns a new slice (may reallocate if capacity exceeded).

### Slicing (Creating Sub-slices)

```go
fruits := []string{"Apple", "Mango", "Peach", "Banana", "Orange"}

// [start:end] - excludes end
fmt.Println(fruits[1:3])   // [Mango Peach]

// [:end] - from start
fmt.Println(fruits[:3])    // [Apple Mango Peach]

// [start:] - to end
fmt.Println(fruits[2:])    // [Peach Banana Orange]

// [:] - entire slice
fmt.Println(fruits[:])     // [Apple Mango Peach Banana Orange]
```

### Removing Elements

Go doesn't have a built-in remove function, but you can use slicing:

```go
highScores := []int{111, 123, 213, 222, 312, 321, 333}

// Remove element at index 2
index := 2
highScores = append(highScores[:index], highScores[index+1:]...)

fmt.Println(highScores)  // [111 123 222 312 321 333]
```

**The `...` operator:** Unpacks a slice into individual elements.

### Length and Capacity

```go
slice := make([]int, 3, 5)

fmt.Println(len(slice))  // 3 (current elements)
fmt.Println(cap(slice))  // 5 (underlying array size)
```

### Sorting Slices

```go
import "sort"

numbers := []int{5, 2, 8, 1, 9}
sort.Ints(numbers)
fmt.Println(numbers)  // [1 2 5 8 9]

fruits := []string{"Peach", "Apple", "Mango"}
sort.Strings(fruits)
fmt.Println(fruits)  // [Apple Mango Peach]

// Check if sorted
fmt.Println(sort.IntsAreSorted(numbers))  // true
```

---

## make() vs new()

### `make()` - For Slices, Maps, Channels

```go
slice := make([]int, 5)       // Creates slice with len=5
// Allocates memory AND initializes
// Non-zeroed storage (ready to use)
```

### `new()` - For Any Type

```go
ptr := new(int)               // Returns pointer to int
// Allocates memory but NO initialization
// Zeroed storage
```

**Rule of Thumb:**
- Slices, maps, channels → use `make()`
- Everything else → use literals or `new()`

---

## Arrays vs Slices: Key Differences

| Feature | Array | Slice |
|---------|-------|-------|
| **Size** | Fixed | Dynamic |
| **Type** | Size part of type | Size not part of type |
| **Declaration** | `[4]int` | `[]int` |
| **Append** | ❌ Not possible | ✅ `append()` |
| **Resize** | ❌ Not possible | ✅ Automatically |
| **Copy behavior** | Copies all elements | Copies pointer/len/cap |
| **Common usage** | Rare | Very common |

---

## Slice Internals and Gotchas

### Gotcha 1: Slices Share Underlying Array

```go
original := []int{1, 2, 3, 4, 5}
slice1 := original[1:4]  // [2 3 4]

slice1[0] = 999

fmt.Println(original)  // [1 999 3 4 5] - CHANGED!
fmt.Println(slice1)    // [999 3 4]
```

**Solution:** Copy to new slice:
```go
slice2 := make([]int, len(original))
copy(slice2, original)
```

### Gotcha 2: Append May Reallocate

```go
slice := make([]int, 2, 5)
slice[0] = 1
slice[1] = 2

// While cap > len, appends don't reallocate
slice = append(slice, 3, 4, 5)  // Still same array

// Exceeding capacity creates new array
slice = append(slice, 6, 7)  // NEW array allocated
```

### Gotcha 3: Nil vs Empty Slice

```go
var nilSlice []int           // nil slice
emptySlice := []int{}        // empty slice (not nil)

fmt.Println(nilSlice == nil)   // true
fmt.Println(emptySlice == nil) // false

// Both have len=0 and cap=0
// Use len() to check if empty (works for both)
```

---

## Common Patterns

### Pattern 1: Pre-allocate Capacity

```go
// Bad: Reallocates many times
result := []int{}
for i := 0; i < 10000; i++ {
    result = append(result, i)
}

// Good: Pre-allocate capacity
result := make([]int, 0, 10000)
for i := 0; i < 10000; i++ {
    result = append(result, i)
}
```

### Pattern 2: Filter Slice

```go
numbers := []int{1, 2, 3, 4, 5, 6}
var evens []int

for _, num := range numbers {
    if num%2 == 0 {
        evens = append(evens, num)
    }
}

fmt.Println(evens)  // [2 4 6]
```

### Pattern 3: Copy Slice Safely

```go
original := []int{1, 2, 3, 4, 5}

// Create new slice with same length
copied := make([]int, len(original))

// Copy elements
copy(copied, original)

// Now modifications are independent
copied[0] = 999
fmt.Println(original)  // [1 2 3 4 5]
fmt.Println(copied)    // [999 2 3 4 5]
```

---

## Summary

**Key Takeaways:**
- **Arrays**: Fixed size, rarely used, value type
- **Slices**: Dynamic size, most common, reference type
- Use `make()` to create slices with specific length/capacity
- `append()` adds elements (may reallocate)
- Slicing creates views into same underlying array
- Use `copy()` for independent copies
- Pre-allocate capacity for performance
- Check length with `len()`, capacity with `cap()`

**Quick Reference:**
```go
// Arrays (fixed)
var arr [5]int
arr := [...]int{1, 2, 3}

// Slices (dynamic)
slice := []int{1, 2, 3}
slice := make([]int, len, cap)
slice = append(slice, 4)
slice = slice[1:3]
copy(dest, src)
len(slice), cap(slice)
```

**When to Use:**
- **Arrays**: When size is truly fixed and known at compile time
- **Slices**: Almost always (99% of cases)

**When to Use This Note:**
- Choosing between array and slice
- Understanding slice internals
- Debugging slice-related issues
- Optimizing slice operations
- Learning about dynamic collections

---

**📝 Last Updated:** Fundamentals Series
**➡️ Next Topic:** [Maps and Structs](./04_maps-and-structs.md)
**🔗 Example Code:** [Concepts/08myarray](../Concepts/08myarray/), [Concepts/09myslices](../Concepts/09myslices/)
