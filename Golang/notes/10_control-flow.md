# Control Flow in Go

**📅 Created:** Language Fundamentals
**🏷️ Topics:** if/else, switch, for loops, range, break, continue
**🔗 Related:** [07_functions.md](./07_functions.md), [Concepts/12ifelse](../Concepts/12ifelse/), [Concepts/13switchcase](../Concepts/13switchcase/), [Concepts/14loops](../Concepts/14loops/)

---

## Overview

Go's control flow is deliberately simple - only `if`, `switch`, and `for`. No `while`, no `do-while`, no ternary operator. This simplicity makes code easier to read and maintain.

---

## If/Else Statements

### Basic If

```go
if condition {
    // Code
}
```

### If-Else

```go
if x > 10 {
    fmt.Println("Greater than 10")
} else {
    fmt.Println("10 or less")
}
```

### If-Else If-Else

```go
if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")
} else if score >= 70 {
    fmt.Println("C")
} else {
    fmt.Println("F")
}
```

### If With Short Statement

```go
// Initialize variable in if scope
if num := getValue(); num > 0 {
    fmt.Println("Positive:", num)
} else {
    fmt.Println("Non-positive:", num)
}
// num not accessible here
```

**Reasoning:** Variables declared in the if statement are scoped to the if/else block. This reduces variable pollution and makes intent clear.

---

## Switch Statements

### Basic Switch

```go
day := "Monday"

switch day {
case "Monday":
    fmt.Println("Start of week")
case "Friday":
    fmt.Println("End of week")
case "Saturday", "Sunday":
    fmt.Println("Weekend!")
default:
    fmt.Println("Midweek")
}
```

**Key Difference from C/Java:** No `break` needed! Cases automatically break.

**Reasoning:** Automatic break prevents the common bug of forgetting `break`. To fall through, use `fallthrough` keyword explicitly.

### Switch With Multiple Values

```go
switch day {
case "Saturday", "Sunday":  // Multiple values in one case
    fmt.Println("Weekend")
}
```

### Switch With Expressions

```go
score := 85

switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
case score >= 70:
    fmt.Println("C")
default:
    fmt.Println("F")
}
```

**Reasoning:** Omitting the switch expression makes it like an if-else chain, but often clearer for multiple conditions.

### Switch With Short Statement

```go
switch num := getValue(); {
case num < 0:
    fmt.Println("Negative")
case num == 0:
    fmt.Println("Zero")
default:
    fmt.Println("Positive")
}
```

### Type Switch

```go
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    default:
        fmt.Printf("Unknown type\n")
    }
}
```

**Reasoning:** Type switches are perfect for handling interface types. The type assertion syntax `.(type)` is only valid in switch statements.

---

## Loops

**Go has only ONE loop keyword: `for`**

### Traditional For Loop

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### While-Style Loop

```go
count := 0
for count < 5 {
    fmt.Println(count)
    count++
}
```

**Reasoning:** No separate `while` keyword - `for` with just a condition works the same way.

### Infinite Loop

```go
for {
    // Runs forever
    // Use break to exit
}
```

### Range Over Slice/Array

```go
numbers := []int{1, 2, 3, 4, 5}

// Get index and value
for i, num := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", i, num)
}

// Just values
for _, num := range numbers {
    fmt.Println(num)
}

// Just indexes
for i := range numbers {
    fmt.Println(i)
}
```

### Range Over Map

```go
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
}

for name, age := range ages {
    fmt.Printf("%s is %d years old\n", name, age)
}
```

### Range Over String

```go
for i, char := range "Hello" {
    fmt.Printf("%d: %c\n", i, char)
}
```

**Reasoning:** Range iterates over Unicode code points (runes), not bytes.

---

## Break and Continue

### Break

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break  // Exit loop
    }
    fmt.Println(i)
}
```

### Continue

```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue  // Skip even numbers
    }
    fmt.Println(i)  // Only prints odd
}
```

### Labeled Break

```go
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i*j > 4 {
            break outer  // Break outer loop
        }
        fmt.Printf("(%d, %d) ", i, j)
    }
}
```

**Reasoning:** Labels allow breaking/continuing outer loops from nested loops. More explicit than goto, safer than exceptions.

---

## Summary

**Key Takeaways:**
- Only one loop: `for` (no while/do-while)
- Switch cases don't fall through (no break needed)
- Range for convenient iteration
- If with short statement for scoped variables
- Labels for breaking nested loops

**Quick Reference:**
```go
// If
if condition { }
if x := getValue(); x > 0 { }

// Switch (auto-break)
switch value {
case 1:
case 2, 3:  // Multiple values
default:
}

// Loops
for i := 0; i < 10; i++ { }           // Traditional
for condition { }                      // While-style
for { }                                // Infinite
for i, v := range slice { }            // Range
```

---

**📝 Last Updated:** Language Fundamentals
**➡️ Next Topics:** See [INDEX.md](./INDEX.md) for complete topic list
**🔗 Example Code:** [Concepts/12ifelse](../Concepts/12ifelse/), [Concepts/13switchcase](../Concepts/13switchcase/), [Concepts/14loops](../Concepts/14loops/)
