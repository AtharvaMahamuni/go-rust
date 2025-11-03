# Goroutines and Concurrency in Go

**📅 Created:** Advanced Topics Series
**🏷️ Topics:** Goroutines, Channels, Concurrency, sync.WaitGroup, sync.Mutex, CSP
**🔗 Related:** [02_pointers.md](./02_pointers.md), [Concepts/26goroutines](../Concepts/26goroutines/), [Concepts/27mutexAndAwaitGroups](../Concepts/27mutexAndAwaitGroups/), [Concepts/28channels](../Concepts/28channels/)

---

## Overview

This note covers Go's approach to concurrency: goroutines (lightweight threads), channels (communication between goroutines), and synchronization primitives. Go's concurrency model is one of its most powerful features, based on the principle "Don't communicate by sharing memory; instead, share memory by communicating."

---

## What are Goroutines?

**Simple Definition:**
A goroutine is a lightweight thread managed by the Go runtime.

```go
func sayHello() {
    fmt.Println("Hello")
}

func main() {
    go sayHello()  // Runs concurrently!
    fmt.Println("World")
}
```

**Key Characteristics:**
- **Lightweight** - thousands can run simultaneously (stack starts at 2KB)
- **Managed by Go runtime** - not OS threads
- **Cheap to create** - no significant overhead
- **Started with `go` keyword** - that's it!

---

## Creating Goroutines

### Method 1: Function Call

```go
func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }
}

func main() {
    go printNumbers()  // Runs concurrently
    time.Sleep(100 * time.Millisecond)  // Wait for goroutine
}
```

### Method 2: Anonymous Function

```go
func main() {
    go func() {
        fmt.Println("Running concurrently!")
    }()

    time.Sleep(100 * time.Millisecond)
}
```

### Method 3: With Parameters

```go
func greet(name string) {
    fmt.Println("Hello", name)
}

func main() {
    go greet("Alice")
    go greet("Bob")

    time.Sleep(100 * time.Millisecond)
}
```

---

## sync.WaitGroup - Proper Synchronization

**Problem:** Using `time.Sleep()` is unreliable.

**Solution:** Use `sync.WaitGroup` to wait for goroutines to finish.

```go
import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()  // Mark as done when function exits

    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)       // Increment counter
        go worker(i, &wg)
    }

    wg.Wait()  // Block until counter reaches 0
    fmt.Println("All workers done")
}
```

**How it works:**
1. `wg.Add(n)` - Increment counter by n
2. `wg.Done()` - Decrement counter by 1
3. `wg.Wait()` - Block until counter reaches 0

---

## Real-World Example: Concurrent HTTP Requests

```go
package main

import (
    "fmt"
    "net/http"
    "sync"
)

func getStatusCode(endpoint string, wg *sync.WaitGroup) {
    defer wg.Done()

    res, err := http.Get(endpoint)

    if err != nil {
        fmt.Printf("Error fetching %s\n", endpoint)
        return
    }

    fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}

func main() {
    var wg sync.WaitGroup

    endpoints := []string{
        "https://google.com",
        "https://github.com",
        "https://reddit.com",
    }

    for _, endpoint := range endpoints {
        wg.Add(1)
        go getStatusCode(endpoint, &wg)
    }

    wg.Wait()
    fmt.Println("All requests complete")
}
```

**Why this is powerful:**
- 3 requests run concurrently (not sequentially)
- Faster than making requests one by one
- Scales to hundreds or thousands of requests

---

## Shared Memory and Race Conditions

### The Problem

```go
var counter = 0
var wg sync.WaitGroup

func increment() {
    defer wg.Done()

    for i := 0; i < 1000; i++ {
        counter++  // ⚠️ RACE CONDITION!
    }
}

func main() {
    wg.Add(2)
    go increment()
    go increment()
    wg.Wait()

    fmt.Println(counter)  // Expected: 2000, Actual: varies!
}
```

**Why?** Multiple goroutines accessing `counter` simultaneously.

---

## sync.Mutex - Protecting Shared Data

**Mutex = Mutual Exclusion**

```go
var counter = 0
var wg sync.WaitGroup
var mut sync.Mutex  // Add mutex

func increment() {
    defer wg.Done()

    for i := 0; i < 1000; i++ {
        mut.Lock()    // Acquire lock
        counter++
        mut.Unlock()  // Release lock
    }
}

func main() {
    wg.Add(2)
    go increment()
    go increment()
    wg.Wait()

    fmt.Println(counter)  // Always 2000!
}
```

**How it works:**
- `mut.Lock()` - Only one goroutine can pass at a time
- `mut.Unlock()` - Release so others can enter
- Other goroutines wait their turn

### Real Example: Collecting Results Safely

```go
var signals []string
var wg sync.WaitGroup
var mut sync.Mutex

func getStatusCode(endpoint string) {
    defer wg.Done()

    res, err := http.Get(endpoint)
    if err != nil {
        return
    }

    mut.Lock()
    signals = append(signals, endpoint)  // Safely append
    mut.Unlock()

    fmt.Printf("%d for %s\n", res.StatusCode, endpoint)
}

func main() {
    endpoints := []string{
        "https://google.com",
        "https://github.com",
        "https://reddit.com",
    }

    for _, endpoint := range endpoints {
        wg.Add(1)
        go getStatusCode(endpoint)
    }

    wg.Wait()
    fmt.Println("Results:", signals)
}
```

---

## Channels - The Go Way

**Philosophy:** "Don't communicate by sharing memory; share memory by communicating."

### What are Channels?

Channels are typed conduits for sending and receiving values between goroutines.

```go
ch := make(chan int)  // Channel of integers
```

### Basic Channel Operations

```go
// Create channel
ch := make(chan int)

// Send value (blocks until someone receives)
ch <- 42

// Receive value (blocks until someone sends)
value := <-ch
```

### Example: Simple Communication

```go
func main() {
    ch := make(chan string)

    go func() {
        ch <- "Hello from goroutine"  // Send
    }()

    msg := <-ch  // Receive
    fmt.Println(msg)
}
```

---

## Buffered Channels

**Unbuffered channel** (default):
- Blocks until both sender and receiver are ready
- Synchronous communication

**Buffered channel:**
- Has capacity
- Sender only blocks when buffer is full
- Receiver only blocks when buffer is empty

```go
ch := make(chan int, 2)  // Buffer of 2

ch <- 1  // Doesn't block
ch <- 2  // Doesn't block
// ch <- 3  // Would block (buffer full)

fmt.Println(<-ch)  // 1
fmt.Println(<-ch)  // 2
```

---

## Channel Directions

### Send-only Channel

```go
func sender(ch chan<- int) {
    ch <- 42  // ✅ Can send
    // val := <-ch  // ❌ Cannot receive
}
```

### Receive-only Channel

```go
func receiver(ch <-chan int) {
    val := <-ch  // ✅ Can receive
    // ch <- 42  // ❌ Cannot send
}
```

### Full Example

```go
func main() {
    ch := make(chan int, 2)
    wg := &sync.WaitGroup{}

    wg.Add(2)

    // Receiver goroutine (receive-only)
    go func(ch <-chan int, wg *sync.WaitGroup) {
        defer wg.Done()

        val, ok := <-ch
        if ok {
            fmt.Println("Received:", val)
        } else {
            fmt.Println("Channel closed")
        }
    }(ch, wg)

    // Sender goroutine (send-only)
    go func(ch chan<- int, wg *sync.WaitGroup) {
        defer wg.Done()

        ch <- 7
        close(ch)  // Important: close when done
    }(ch, wg)

    wg.Wait()
}
```

---

## Closing Channels

```go
ch := make(chan int)

go func() {
    for i := 1; i <= 5; i++ {
        ch <- i
    }
    close(ch)  // Signal no more values
}()

// Receive until channel is closed
for val := range ch {
    fmt.Println(val)
}
```

**Rules:**
- Only sender should close
- Receiving from closed channel returns zero value
- Sending to closed channel causes panic
- Closing is optional (useful for signaling)

---

## Select Statement

**Select** lets you wait on multiple channel operations:

```go
ch1 := make(chan string)
ch2 := make(chan string)

go func() {
    time.Sleep(1 * time.Second)
    ch1 <- "from ch1"
}()

go func() {
    time.Sleep(2 * time.Second)
    ch2 <- "from ch2"
}()

for i := 0; i < 2; i++ {
    select {
    case msg1 := <-ch1:
        fmt.Println(msg1)
    case msg2 := <-ch2:
        fmt.Println(msg2)
    }
}
```

---

## Common Patterns

### Pattern 1: Worker Pool

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        results <- job * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Start 3 workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send 5 jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect results
    for a := 1; a <= 5; a++ {
        <-results
    }
}
```

### Pattern 2: Timeout

```go
select {
case result := <-ch:
    fmt.Println("Got result:", result)
case <-time.After(1 * time.Second):
    fmt.Println("Timeout!")
}
```

### Pattern 3: Fan-out, Fan-in

```go
// Fan-out: One input, many workers
// Fan-in: Many workers, one output
```

---

## Best Practices

1. **Always close channels from sender**
2. **Use WaitGroup for coordination**
3. **Use Mutex for shared state**
4. **Prefer channels for communication**
5. **Don't start goroutines without knowing how they'll stop**
6. **Detect race conditions:** `go run -race main.go`

---

## Summary

**Key Takeaways:**
- Goroutines are lightweight concurrent functions
- Start with `go funcName()`
- Use `sync.WaitGroup` to wait for goroutines
- Use `sync.Mutex` to protect shared data
- Channels enable safe communication between goroutines
- Buffered channels have capacity
- Close channels to signal completion
- Select waits on multiple channel operations
- "Share memory by communicating"

**Quick Reference:**
```go
// Goroutines
go funcName()

// WaitGroup
wg.Add(1)
defer wg.Done()
wg.Wait()

// Mutex
mut.Lock()
mut.Unlock()

// Channels
ch := make(chan Type)
ch := make(chan Type, capacity)
ch <- value
value := <-ch
close(ch)

// Directions
chan<- Type  // send-only
<-chan Type  // receive-only
```

---

**📝 Last Updated:** Advanced Topics Series
**➡️ Related:** [sync package](https://pkg.go.dev/sync), [Go Concurrency Patterns](https://go.dev/blog/pipelines)
**🔗 Example Code:** [Concepts/26goroutines](../Concepts/26goroutines/), [Concepts/27mutexAndAwaitGroups](../Concepts/27mutexAndAwaitGroups/), [Concepts/28channels](../Concepts/28channels/)
