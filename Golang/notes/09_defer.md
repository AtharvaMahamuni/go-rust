# Defer, Panic, and Recover

**📅 Created:** Control Flow Series
**🏷️ Topics:** defer, panic, recover, cleanup, error handling
**🔗 Related:** [07_functions.md](./07_functions.md), [08_methods.md](./08_methods.md), [Concepts/17defer](../Concepts/17defer/)

---

## Overview

`defer`, `panic`, and `recover` are Go's mechanisms for handling cleanup, errors, and exceptional situations. Defer ensures code runs at function exit, panic stops normal execution, and recover catches panics. Together they provide robust error handling without traditional try-catch blocks.

---

## Defer

### What is Defer?

`defer` schedules a function call to run AFTER the surrounding function returns:

```go
func main() {
    defer fmt.Println("World!")
    fmt.Println("Hello")
}

// Output:
// Hello
// World!
```

**Reasoning:** Defer executes at the END of the function, not where it's written. This is perfect for cleanup operations that must happen regardless of how the function exits.

---

## Common Use Cases

### 1. File Cleanup

```go
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()  // Guaranteed to run

    // Work with file...
    // Even if error occurs, file.Close() will run
    return nil
}
```

**Reasoning:** Without defer, you'd need Close() before every return. Defer ensures cleanup happens once, at the right time.

### 2. Mutex Unlock

```go
var mu sync.Mutex

func criticalSection() {
    mu.Lock()
    defer mu.Unlock()  // Always unlocks

    // Critical code...
    // Even if panic, mutex gets unlocked
}
```

### 3. Database Connections

```go
func query() error {
    db, err := sql.Open("mysql", "...")
    if err != nil {
        return err
    }
    defer db.Close()

    // Run queries...
    return nil
}
```

**Pattern:** Acquire resource → immediately defer cleanup → use resource safely.

---

## Defer Execution Order

Defers execute in **LIFO (Last In, First Out)** order:

```go
func main() {
    defer fmt.Println("One")
    defer fmt.Println("Two")
    defer fmt.Println("Three")
    fmt.Println("Hello")
}

// Output:
// Hello
// Three
// Two
// One
```

**Reasoning:** Think of defers as a stack. Last deferred call executes first. This matches cleanup order - reverse of acquisition.

### Multiple Defers Example

```go
func deferCounter() {
    defer fmt.Println()  // Prints newline last
    for i := 0; i < 5; i++ {
        defer fmt.Print(i)  // Each iteration adds to stack
    }
}

// Output: 43210
// (prints in reverse: 4, 3, 2, 1, 0, then newline)
```

**Reasoning:** Each `defer` in the loop adds to the stack. When function exits, they execute in reverse order.

---

## Defer Arguments Are Evaluated Immediately

```go
func main() {
    x := 10
    defer fmt.Println(x)  // x evaluated NOW (10)

    x = 20
    fmt.Println(x)  // 20
}

// Output:
// 20
// 10
```

**Reasoning:** Defer evaluates arguments immediately, but delays the function call. This can be surprising - the value is captured, not the reference.

### To Capture Current Value

```go
func main() {
    x := 10

    // Wrap in anonymous function
    defer func() {
        fmt.Println(x)  // Captures x by reference
    }()

    x = 20
    fmt.Println(x)
}

// Output:
// 20
// 20  (captures updated value)
```

---

## Defer With Named Returns

Defer can modify named return values:

```go
func increment() (result int) {
    defer func() {
        result++  // Modifies return value!
    }()

    return 5  // Returns 5, then defer increments to 6
}

fmt.Println(increment())  // 6
```

**Reasoning:** Named returns are variables. Defer runs after `return` assigns the value but before function exits, allowing modification.

**Practical Use - Timing:**
```go
func slowFunction() (duration time.Duration) {
    start := time.Now()
    defer func() {
        duration = time.Since(start)
    }()

    // Do work...
    time.Sleep(1 * time.Second)
    return
}
```

---

## Panic

`panic` stops normal execution and begins unwinding the stack:

```go
func main() {
    fmt.Println("Start")
    panic("Something went wrong!")
    fmt.Println("Never printed")
}

// Output:
// Start
// panic: Something went wrong!
// ... stack trace ...
```

**When to Use Panic:**
- Unrecoverable errors (should crash the program)
- Programming errors (invariants violated)
- Initialization failures

**When NOT to Use:**
- Expected errors (use error returns)
- Validation failures (return errors)
- User input errors (return errors)

**Reasoning:** Panic is for exceptional situations. Go prefers explicit error handling with return values for normal error cases.

### Panic Unwinding

```go
func main() {
    defer fmt.Println("Defer 1")
    defer fmt.Println("Defer 2")

    panic("Oops!")

    fmt.Println("Never runs")
}

// Output:
// Defer 2
// Defer 1
// panic: Oops!
```

**Reasoning:** Even during panic, deferred functions execute (in LIFO order). This ensures cleanup happens before program exits.

---

## Recover

`recover` stops a panic and returns the panic value:

```go
func safeDivide(a, b float64) (result float64, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic: %v", r)
        }
    }()

    if b == 0 {
        panic("division by zero")
    }

    return a / b, nil
}

func main() {
    result, err := safeDivide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)  // Error: panic: division by zero
    } else {
        fmt.Println("Result:", result)
    }
}
```

**Reasoning:** Recover only works inside deferred functions. It converts a panic into an error, allowing graceful handling.

### Recover Rules

1. **Must be called in deferred function:**
```go
// ❌ Won't work
func bad() {
    recover()  // Does nothing
    panic("oops")
}

// ✅ Works
func good() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()
    panic("oops")
}
```

2. **Only stops current goroutine's panic:**
```go
go func() {
    panic("goroutine panic")
}()
// Can't recover from another goroutine's panic
```

---

## Real-World Pattern

### HTTP Handler Protection

```go
func handler(w http.ResponseWriter, r *http.Request) {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Panic: %v", r)
            http.Error(w, "Internal Server Error", 500)
        }
    }()

    // Handle request...
    // If panic occurs, recovered and returns 500
}
```

**Reasoning:** Prevents one handler's panic from crashing the entire server.

---

## Common Patterns

### Pattern 1: Cleanup Resources

```go
func processFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()

    // Process file
    return nil
}
```

### Pattern 2: Timing Functions

```go
func timeTrack(start time.Time, name string) {
    fmt.Printf("%s took %v\n", name, time.Since(start))
}

func slowFunc() {
    defer timeTrack(time.Now(), "slowFunc")
    // Do work...
}
```

### Pattern 3: Recover Pattern

```go
func safeCall(fn func()) (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic: %v", r)
        }
    }()

    fn()
    return nil
}
```

---

## Summary

**Key Takeaways:**
- `defer` schedules function to run at function exit
- Executes in LIFO order (last defer runs first)
- Arguments evaluated immediately, call delayed
- Perfect for cleanup (files, locks, connections)
- `panic` stops normal execution (use sparingly)
- `recover` catches panics (only in defer)
- Go prefers error returns over panic/recover

**Quick Reference:**
```go
// Defer
defer file.Close()
defer mu.Unlock()

// Multiple defers (LIFO)
defer fmt.Println("Last")   // Runs second
defer fmt.Println("First")  // Runs first

// Panic
panic("error message")

// Recover
defer func() {
    if r := recover(); r != nil {
        // Handle panic
    }
}()
```

**When to Use:**
- **Defer:** Cleanup, unlocking, closing
- **Panic:** Unrecoverable programmer errors
- **Recover:** Preventing program crashes (servers, libraries)

---

**📝 Last Updated:** Control Flow Series
**➡️ Next Topic:** [Control Flow](./10_control-flow.md)
**🔗 Example Code:** [Concepts/17defer](../Concepts/17defer/)
