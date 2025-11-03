# Math and Random Numbers in Go

**📅 Created:** Standard Library Series
**🏷️ Topics:** math package, random numbers, crypto/rand, math/big, arithmetic operations
**🔗 Related:** [01_variables-and-types.md](./01_variables-and-types.md), [Concepts/05mymaths](../Concepts/05mymaths/)

---

## Overview

Go provides comprehensive math functionality through the `math` package and secure random number generation through `crypto/rand`. Understanding these tools is essential for numerical computations, simulations, cryptography, and generating unique identifiers.

---

## Basic Arithmetic

### Type Compatibility

```go
var num1 int = 34
var num2 float64 = 5.6

// ❌ Error - mismatched types
// result := num1 + num2

// ✅ Correct - explicit type conversion
result := float64(num1) + num2  // 39.6
// or
result := num1 + int(num2)      // 39 (truncates decimal)
```

**Reasoning:** Go doesn't allow implicit type conversion in arithmetic. You must explicitly convert one operand to match the other's type.

### Basic Operations

```go
a, b := 10, 3

fmt.Println(a + b)   // 13  Addition
fmt.Println(a - b)   // 7   Subtraction
fmt.Println(a * b)   // 30  Multiplication
fmt.Println(a / b)   // 3   Integer division (truncates)
fmt.Println(a % b)   // 1   Modulo (remainder)

// Float division
fmt.Println(float64(a) / float64(b))  // 3.333...
```

**Reasoning:** Integer division truncates the decimal. Convert to float64 for decimal results.

---

## Math Package

### Common Functions

```go
import "math"

// Absolute value
math.Abs(-5.5)      // 5.5

// Power and roots
math.Pow(2, 8)      // 256 (2^8)
math.Sqrt(16)       // 4
math.Cbrt(27)       // 3 (cube root)

// Rounding
math.Floor(4.7)     // 4
math.Ceil(4.1)      // 5
math.Round(4.5)     // 5

// Min/Max
math.Max(10, 20)    // 20
math.Min(10, 20)    // 10

// Trigonometry (radians)
math.Sin(math.Pi / 2)   // 1
math.Cos(0)             // 1
math.Tan(math.Pi / 4)   // 1
```

### Constants

```go
math.Pi            // 3.141592653589793
math.E             // 2.718281828459045
math.Phi           // 1.618033988749895 (golden ratio)

math.MaxInt64      // 9223372036854775807
math.MinInt64      // -9223372036854775808
math.MaxFloat64    // 1.7976931348623157e+308
```

**Reasoning:** Math package provides standard mathematical functions. All trig functions use radians, not degrees.

---

## Random Numbers - math/rand

**⚠️ Warning:** `math/rand` is NOT cryptographically secure. Use for simulations, games, testing only.

### Basic Random Numbers

```go
import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Seed with current time (for different results each run)
    rand.Seed(time.Now().UnixNano())

    // Random int [0, n)
    fmt.Println(rand.Intn(100))      // [0, 100)

    // Range [min, max]
    min, max := 10, 20
    num := rand.Intn(max-min+1) + min  // [10, 20]

    // Random float [0.0, 1.0)
    fmt.Println(rand.Float64())

    // Random from range
    nums := []int{10, 20, 30, 40, 50}
    randomIndex := rand.Intn(len(nums))
    fmt.Println(nums[randomIndex])
}
```

**Reasoning:** Without `Seed()`, rand produces the same sequence every run. Seeding with current time gives different results.

### Common Patterns

```go
// Dice roll (1-6)
dice := rand.Intn(6) + 1

// Coin flip
coinFlip := rand.Intn(2)  // 0 or 1
if coinFlip == 0 {
    fmt.Println("Heads")
} else {
    fmt.Println("Tails")
}

// Random boolean
randomBool := rand.Intn(2) == 1

// Random element from slice
fruits := []string{"apple", "banana", "orange"}
randomFruit := fruits[rand.Intn(len(fruits))]

// Shuffle slice
rand.Shuffle(len(fruits), func(i, j int) {
    fruits[i], fruits[j] = fruits[j], fruits[i]
})
```

---

## Cryptographically Secure Random - crypto/rand

For security-sensitive operations (passwords, tokens, keys), use `crypto/rand`:

### Basic Usage

```go
import (
    "crypto/rand"
    "fmt"
    "math/big"
)

func main() {
    // Generate random int [0, n)
    max := big.NewInt(100)
    n, err := rand.Int(rand.Reader, max)
    if err != nil {
        panic(err)
    }

    fmt.Println(n)  // Random number [0, 100)
}
```

### Random Number in Range

```go
// Generate number in range [1, 6] (dice)
func secureDiceRoll() int {
    max := big.NewInt(6)
    n, _ := rand.Int(rand.Reader, max)

    // Add 1 to shift from [0, 6) to [1, 6]
    one := big.NewInt(1)
    result := n.Add(n, one)

    return int(result.Int64())
}
```

**Reasoning:** `crypto/rand` uses OS's random number generator (cryptographically secure). Returns `*big.Int`, not regular int.

### Generate Random Bytes

```go
// Generate 32 random bytes (for tokens, keys)
func generateToken() ([]byte, error) {
    token := make([]byte, 32)
    _, err := rand.Read(token)
    if err != nil {
        return nil, err
    }
    return token, nil
}

// Usage
token, _ := generateToken()
fmt.Printf("%x\n", token)  // Hex representation
```

### Random String

```go
import (
    "crypto/rand"
    "encoding/base64"
)

func generateRandomString(length int) string {
    bytes := make([]byte, length)
    rand.Read(bytes)
    return base64.URLEncoding.EncodeToString(bytes)[:length]
}

// Generate random ID
id := generateRandomString(16)
```

**Reasoning:** For tokens, passwords, session IDs, always use `crypto/rand`. It's slower but secure.

---

## math/big Package

For arbitrary-precision arithmetic:

### Big Integers

```go
import "math/big"

func main() {
    // Create big ints
    a := big.NewInt(12345)
    b := big.NewInt(67890)

    // Arithmetic
    sum := new(big.Int)
    sum.Add(a, b)  // 80235

    product := new(big.Int)
    product.Mul(a, b)  // 838102050

    // From string
    bigNum, _ := new(big.Int).SetString("123456789012345678901234567890", 10)
    fmt.Println(bigNum)
}
```

### Big Floats

```go
// Precise floating-point arithmetic
a := big.NewFloat(0.1)
b := big.NewFloat(0.2)

sum := new(big.Float)
sum.Add(a, b)

fmt.Println(sum)  // 0.3 (exactly, no floating point errors)
```

**Reasoning:** Use `math/big` when you need precision beyond int64/float64 limits, or for financial calculations requiring exact decimal arithmetic.

---

## Common Patterns

### Pattern 1: Random ID Generation

```go
func generateID() string {
    rand.Seed(time.Now().UnixNano())
    return fmt.Sprintf("%d", rand.Int63())
}
```

### Pattern 2: Random String with Charset

```go
func randomString(length int) string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    rand.Seed(time.Now().UnixNano())

    result := make([]byte, length)
    for i := range result {
        result[i] = charset[rand.Intn(len(charset))]
    }
    return string(result)
}

// Generate password
password := randomString(12)
```

### Pattern 3: Percentage Chance

```go
func hasChance(percentage int) bool {
    return rand.Intn(100) < percentage
}

// 30% chance
if hasChance(30) {
    fmt.Println("Rare event occurred!")
}
```

### Pattern 4: Weighted Random Selection

```go
func weightedRandom() string {
    weights := map[string]int{
        "common":   70,
        "uncommon": 20,
        "rare":     8,
        "legendary": 2,
    }

    total := 0
    for _, w := range weights {
        total += w
    }

    r := rand.Intn(total)
    current := 0

    for item, weight := range weights {
        current += weight
        if r < current {
            return item
        }
    }

    return "common"
}
```

---

## Performance Considerations

### math/rand vs crypto/rand

```go
import "testing"

func BenchmarkMathRand(b *testing.B) {
    for i := 0; i < b.N; i++ {
        rand.Intn(100)
    }
}

func BenchmarkCryptoRand(b *testing.B) {
    for i := 0; i < b.N; i++ {
        rand.Int(rand.Reader, big.NewInt(100))
    }
}

// math/rand: ~2 ns/op
// crypto/rand: ~300 ns/op
```

**Reasoning:** `math/rand` is ~150x faster but not secure. Choose based on your needs.

---

## Common Pitfalls

### 1. Forgetting to Seed

```go
// ❌ Same numbers every run
rand.Intn(100)

// ✅ Different numbers each run
rand.Seed(time.Now().UnixNano())
rand.Intn(100)
```

### 2. Using math/rand for Security

```go
// ❌ NEVER for passwords, tokens, keys
password := rand.Intn(1000000)

// ✅ Use crypto/rand
token, _ := rand.Int(rand.Reader, big.NewInt(1000000))
```

### 3. Integer Division Truncation

```go
a, b := 7, 2

result := a / b              // 3 (truncated)
result := float64(a) / float64(b)  // 3.5 (correct)
```

---

## Summary

**Key Takeaways:**
- Use `math` package for standard math functions
- `math/rand` for non-secure random numbers (games, simulations)
- `crypto/rand` for security (passwords, tokens, keys)
- Always seed `math/rand` with `time.Now().UnixNano()`
- Use `math/big` for arbitrary precision
- Explicit type conversion required for mixed-type arithmetic

**Quick Reference:**
```go
// Math
import "math"
math.Sqrt(16)         // 4
math.Pow(2, 8)        // 256
math.Round(4.5)       // 5

// Random (non-secure)
import "math/rand"
rand.Seed(time.Now().UnixNano())
rand.Intn(100)        // [0, 100)

// Random (secure)
import "crypto/rand"
import "math/big"
n, _ := rand.Int(rand.Reader, big.NewInt(100))

// Random bytes
bytes := make([]byte, 32)
rand.Read(bytes)
```

**When to Use This Note:**
- Numerical computations
- Random number generation
- Game development
- Simulations
- Security tokens/passwords
- Lottery/raffle systems

---

**📝 Last Updated:** Standard Library Series
**➡️ Related:** See [01_variables-and-types.md](./01_variables-and-types.md) for type conversion
**🔗 Example Code:** [Concepts/05mymaths](../Concepts/05mymaths/)
