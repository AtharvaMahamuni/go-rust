# Time Handling in Go

**📅 Created:** Standard Library Series
**🏷️ Topics:** time package, Date/Time Manipulation, Formatting, Parsing, Duration
**🔗 Related:** [11_user-input-and-conversion.md](./11_user-input-and-conversion.md), [Concepts/06mytime](../Concepts/06mytime/)

---

## Overview

The `time` package provides functionality for measuring and displaying time. Understanding time handling is essential for logging, scheduling, timeouts, and timestamps. Go's time package is powerful but has some quirks, especially around formatting.

**Why Time Matters:**
- Timestamps for logs and records
- Scheduling and delays
- Measuring performance
- Timeouts in network operations
- Date calculations

---

## Getting Current Time

```go
import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Println(now)  // 2025-11-02 14:30:45.123456789 +0000 UTC
}
```

**What You Get:**
- Date and time
- Nanosecond precision
- Timezone information

**Reasoning:** `time.Now()` returns the current time with nanosecond precision. The actual precision depends on the operating system, but Go provides the API for nanoseconds.

---

## Time Structure

```go
now := time.Now()

// Access components
year := now.Year()          // 2025
month := now.Month()        // November (type: time.Month)
day := now.Day()            // 2
hour := now.Hour()          // 14
minute := now.Minute()      // 30
second := now.Second()      // 45
nano := now.Nanosecond()    // 123456789

// Weekday
weekday := now.Weekday()    // Saturday (type: time.Weekday)

// Unix timestamp
unix := now.Unix()          // Seconds since Jan 1, 1970
unixNano := now.UnixNano()  // Nanoseconds since Jan 1, 1970
```

**Reasoning:** Each component has its own method. Some return custom types (Month, Weekday) for type safety.

---

## Creating Specific Times

### Using time.Date

```go
// time.Date(year, month, day, hour, min, sec, nsec, location)
specificTime := time.Date(2000, time.April, 5, 4, 30, 0, 0, time.Local)

fmt.Println(specificTime)
// 2000-04-05 04:30:00 +0000 Local
```

**Parameters:**
- `year`: int (e.g., 2025)
- `month`: time.Month (January through December constants)
- `day`: int (1-31)
- `hour`: int (0-23)
- `minute`: int (0-59)
- `second`: int (0-59)
- `nanosecond`: int (0-999,999,999)
- `location`: *time.Location (time.UTC, time.Local, or custom)

**Reasoning:** `time.Date` normalizes values - you can pass day=32 and it'll roll over to the next month. This is useful for date arithmetic.

---

## Time Formatting - The Magic Reference

**Go's Unique Approach:**

Go uses a **reference time** for formatting: `Mon Jan 2 15:04:05 MST 2006`

This specific date/time represents: `01/02 03:04:05PM '06 -0700`

**Why this weird date?** It's easy to remember: `01/02 03:04:05 06 07`

```go
now := time.Now()

// Common formats
fmt.Println(now.Format("2006-01-02"))              // 2025-11-02
fmt.Println(now.Format("02/01/2006"))              // 02/11/2025
fmt.Println(now.Format("02-01-2006 Monday"))       // 02-11-2025 Saturday
fmt.Println(now.Format("15:04:05"))                // 14:30:45
fmt.Println(now.Format("03:04:05 PM"))             // 02:30:45 PM
fmt.Println(now.Format("Jan 02, 2006"))            // Nov 02, 2025
fmt.Println(now.Format("Monday, January 2, 2006")) // Saturday, November 2, 2025
```

**Format Components:**

| Component | Format | Example |
|-----------|--------|---------|
| **Year** | `2006` | 2025 |
| | `06` | 25 |
| **Month** | `01` | 11 (numeric) |
| | `1` | 11 (no leading zero) |
| | `Jan` | Nov (short) |
| | `January` | November (full) |
| **Day** | `02` | 02 |
| | `2` | 2 (no leading zero) |
| | `_2` | 2 (space padded) |
| **Day of Week** | `Mon` | Sat (short) |
| | `Monday` | Saturday (full) |
| **Hour** | `15` | 14 (24-hour) |
| | `03` | 02 (12-hour) |
| | `3` | 2 (no leading zero) |
| **Minute** | `04` | 30 |
| | `4` | 30 (no leading zero) |
| **Second** | `05` | 45 |
| | `5` | 45 (no leading zero) |
| **AM/PM** | `PM` | PM |
| | `pm` | pm |
| **Timezone** | `MST` | PST (name) |
| | `-0700` | -0800 (offset) |
| | `-07:00` | -08:00 (with colon) |

**Reasoning:** Instead of format codes like `%Y %m %d`, Go uses the reference date. Whatever format you want to see, you write using those exact numbers. Weird but consistent.

---

## Parsing Time Strings

```go
// Parse string to time
layout := "02-01-2006 15:04:05"
str := "05-04-2000 04:30:00"

t, err := time.Parse(layout, str)
if err != nil {
    fmt.Println("Parse error:", err)
} else {
    fmt.Println("Parsed:", t)
}
```

**Common Layouts:**

```go
// Predefined constants
time.ANSIC       // "Mon Jan _2 15:04:05 2006"
time.UnixDate    // "Mon Jan _2 15:04:05 MST 2006"
time.RubyDate    // "Mon Jan 02 15:04:05 -0700 2006"
time.RFC822      // "02 Jan 06 15:04 MST"
time.RFC1123     // "Mon, 02 Jan 2006 15:04:05 MST"
time.RFC3339     // "2006-01-02T15:04:05Z07:00"
time.Kitchen     // "3:04PM"

// Usage
t, _ := time.Parse(time.RFC3339, "2025-11-02T14:30:00Z")
```

**Reasoning:** Layouts must match exactly. If your string has `02/01/2006`, your layout must be `02/01/2006`. Wrong layout = parse error.

---

## Duration Type

Duration represents elapsed time in nanoseconds:

```go
// Creating durations
d1 := time.Second             // 1 second
d2 := 5 * time.Minute         // 5 minutes
d3 := time.Hour + 30*time.Minute  // 1.5 hours

// Constants available
time.Nanosecond   // 1 ns
time.Microsecond  // 1000 ns
time.Millisecond  // 1000000 ns
time.Second       // 1000000000 ns
time.Minute       // 60 seconds
time.Hour         // 60 minutes

// Get duration components
hours := d3.Hours()           // 1.5
minutes := d3.Minutes()       // 90
seconds := d3.Seconds()       // 5400

// As integers
h := int(d3.Hours())          // 1
m := int(d3.Minutes()) % 60   // 30
```

**Reasoning:** Duration is just an int64 representing nanoseconds. The constants and methods make it human-readable.

---

## Time Arithmetic

### Adding/Subtracting Time

```go
now := time.Now()

// Add duration
future := now.Add(24 * time.Hour)         // Tomorrow
past := now.Add(-1 * time.Hour)           // 1 hour ago

// Add specific components
nextMonth := now.AddDate(0, 1, 0)         // AddDate(years, months, days)
nextYear := now.AddDate(1, 0, 0)
tomorrow := now.AddDate(0, 0, 1)
```

**Reasoning:** `Add()` works with Duration. `AddDate()` handles calendar-aware operations (months have different days, leap years, etc.).

### Time Difference

```go
start := time.Now()

// Do some work...
time.Sleep(2 * time.Second)

end := time.Now()
duration := end.Sub(start)

fmt.Println("Elapsed:", duration)         // 2.001234s
fmt.Println("Seconds:", duration.Seconds())  // 2.001234
```

**Reasoning:** `Sub()` returns a Duration (the difference). Positive if end is after start, negative if before.

---

## Comparing Times

```go
t1 := time.Now()
time.Sleep(1 * time.Second)
t2 := time.Now()

// Comparison
if t2.After(t1) {
    fmt.Println("t2 is after t1")
}

if t1.Before(t2) {
    fmt.Println("t1 is before t2")
}

if t1.Equal(t2) {
    fmt.Println("Same time")
}
```

**Important:** Use `Equal()`, not `==`:

```go
// ❌ Don't use ==
if t1 == t2 { }  // May fail due to location differences

// ✅ Use Equal()
if t1.Equal(t2) { }  // Compares instant in time
```

**Reasoning:** `Equal()` compares the actual instant, ignoring timezone. `==` compares the entire struct including location.

---

## Sleep and Timers

### Sleep

```go
fmt.Println("Starting...")
time.Sleep(2 * time.Second)
fmt.Println("Done!")  // After 2 seconds
```

### Timer (One-shot)

```go
timer := time.NewTimer(2 * time.Second)
<-timer.C  // Block until timer expires
fmt.Println("Timer expired")

// Cancel timer if needed
timer.Stop()
```

### Ticker (Repeating)

```go
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop()

for i := 0; i < 5; i++ {
    <-ticker.C
    fmt.Println("Tick", i)
}
```

**Reasoning:** Ticker sends on its channel at regular intervals. Always `Stop()` tickers when done to free resources.

---

## Timezones

### Working with Locations

```go
// UTC
utc := time.Now().UTC()

// Local (system timezone)
local := time.Now().Local()

// Specific timezone
loc, err := time.LoadLocation("America/New_York")
if err != nil {
    fmt.Println("Error loading location:", err)
} else {
    ny := time.Now().In(loc)
    fmt.Println("NYC time:", ny)
}

// Convert between timezones
localTime := time.Now()
utcTime := localTime.UTC()
nyTime := localTime.In(loc)
```

**Common Locations:**
- `"UTC"`
- `"America/New_York"`
- `"Europe/London"`
- `"Asia/Tokyo"`
- `"Local"` (system timezone)

**Reasoning:** Timezones are tricky. Always be explicit about which timezone you're working in, especially when storing/transmitting times.

---

## Common Patterns

### Pattern 1: Timestamp Logging

```go
func log(message string) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    fmt.Printf("[%s] %s\n", timestamp, message)
}

log("Application started")
// [2025-11-02 14:30:45] Application started
```

### Pattern 2: Measuring Execution Time

```go
func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    fmt.Printf("%s took %s\n", name, elapsed)
}

func slowFunction() {
    defer timeTrack(time.Now(), "slowFunction")
    // Do work...
    time.Sleep(2 * time.Second)
}
```

### Pattern 3: Timeout Context

```go
func doWork(timeout time.Duration) error {
    done := make(chan bool)

    go func() {
        // Do work...
        time.Sleep(3 * time.Second)
        done <- true
    }()

    select {
    case <-done:
        return nil
    case <-time.After(timeout):
        return fmt.Errorf("operation timed out")
    }
}

err := doWork(2 * time.Second)  // Will timeout
```

### Pattern 4: Parse and Validate Date

```go
func validateDate(dateStr string) (time.Time, error) {
    t, err := time.Parse("2006-01-02", dateStr)
    if err != nil {
        return time.Time{}, fmt.Errorf("invalid date format: %w", err)
    }

    // Check if date is in the future
    if t.After(time.Now()) {
        return time.Time{}, fmt.Errorf("date cannot be in the future")
    }

    return t, nil
}
```

---

## Zero Value

```go
var t time.Time

fmt.Println(t.IsZero())  // true
fmt.Println(t)           // 0001-01-01 00:00:00 +0000 UTC
```

**Reasoning:** Zero time is `January 1, year 1, 00:00:00 UTC`. Use `IsZero()` to check if a time has been initialized.

---

## Summary

**Key Takeaways:**
- Use `time.Now()` for current time
- Format with reference date: `Mon Jan 2 15:04:05 MST 2006`
- Duration represents elapsed time (nanoseconds)
- `Add()` for duration, `AddDate()` for calendar operations
- Use `Equal()` not `==` for comparison
- `Sleep()` for delays, `Ticker` for intervals
- Always be explicit about timezones
- `time.Time` zero value is Jan 1, year 1

**Quick Reference:**
```go
// Current time
now := time.Now()

// Format
now.Format("2006-01-02 15:04:05")

// Parse
time.Parse("2006-01-02", "2025-11-02")

// Duration
2 * time.Hour + 30 * time.Minute

// Arithmetic
now.Add(24 * time.Hour)
now.AddDate(0, 1, 0)  // Add 1 month

// Compare
t1.Before(t2)
t1.After(t2)
t1.Equal(t2)
```

**When to Use This Note:**
- Adding timestamps to logs
- Scheduling and timers
- Parsing date strings
- Time calculations
- Performance measurement
- Working with timezones

---

**📝 Last Updated:** Standard Library Series
**➡️ Next Topic:** [File Operations](./13_file-operations.md)
**🔗 Example Code:** [Concepts/06mytime](../Concepts/06mytime/)
