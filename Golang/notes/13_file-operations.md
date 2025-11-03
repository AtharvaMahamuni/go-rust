# File Operations in Go

**📅 Created:** I/O Series
**🏷️ Topics:** File I/O, Reading Files, Writing Files, os package, io package
**🔗 Related:** [09_defer.md](./09_defer.md), [11_user-input-and-conversion.md](./11_user-input-and-conversion.md), [Concepts/18files](../Concepts/18files/)

---

## Overview

File operations are fundamental to most applications. Go's `os` and `io` packages provide simple yet powerful APIs for reading, writing, and manipulating files. Understanding these operations is essential for data persistence, configuration management, and logging.

---

## Writing Files

### Simple Write - os.WriteFile

```go
content := []byte("Hello, World!")
err := os.WriteFile("output.txt", content, 0644)
if err != nil {
    log.Fatal(err)
}
```

**Parameters:**
- Filename
- Data ([]byte)
- Permissions (0644 = rw-r--r--)

**Reasoning:** Simplest way for small files. Creates or truncates file. One function call does everything.

### Create and Write

```go
file, err := os.Create("myfile.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()  // Always close!

content := "This is the file content."
length, err := io.WriteString(file, content)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Wrote %d bytes\n", length)
```

**Reasoning:** `defer file.Close()` ensures file closes even if error occurs. `io.WriteString` writes string directly without converting to []byte.

---

## Reading Files

### Read Entire File - os.ReadFile

```go
data, err := os.ReadFile("myfile.txt")
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(data))  // Convert []byte to string
```

**Reasoning:** Best for small files that fit in memory. Returns []byte - convert to string if needed.

### Open and Read

```go
file, err := os.Open("myfile.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

data := make([]byte, 100)  // Buffer
n, err := file.Read(data)
if err != nil && err != io.EOF {
    log.Fatal(err)
}

fmt.Printf("Read %d bytes: %s\n", n, data[:n])
```

**Reasoning:** For large files or when you need control. Read into buffer, may need multiple reads for large files.

---

## File Operations

### Check if File Exists

```go
func fileExists(filename string) bool {
    _, err := os.Stat(filename)
    return !os.IsNotExist(err)
}

if fileExists("myfile.txt") {
    fmt.Println("File exists")
}
```

### Delete File

```go
err := os.Remove("myfile.txt")
if err != nil {
    log.Fatal(err)
}
```

### Rename/Move File

```go
err := os.Rename("old.txt", "new.txt")
if err != nil {
    log.Fatal(err)
}
```

### Copy File

```go
func copyFile(src, dst string) error {
    source, err := os.Open(src)
    if err != nil {
        return err
    }
    defer source.Close()

    destination, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destination.Close()

    _, err = io.Copy(destination, source)
    return err
}
```

**Reasoning:** Go doesn't have a built-in file copy function. Use `io.Copy` to efficiently copy between files.

---

## Directory Operations

### Create Directory

```go
// Create single directory
err := os.Mkdir("mydir", 0755)

// Create nested directories
err := os.MkdirAll("path/to/mydir", 0755)
```

### List Directory

```go
entries, err := os.ReadDir(".")
if err != nil {
    log.Fatal(err)
}

for _, entry := range entries {
    fmt.Println(entry.Name(), entry.IsDir())
}
```

### Remove Directory

```go
// Remove empty directory
err := os.Remove("mydir")

// Remove directory and contents
err := os.RemoveAll("mydir")
```

---

## File Permissions

Unix-style permissions (3 octal digits):

```go
0644  // rw-r--r--  (Owner: read/write, Others: read)
0755  // rwxr-xr-x  (Owner: all, Others: read/execute)
0600  // rw-------  (Owner: read/write only)
0777  // rwxrwxrwx  (All permissions for everyone)
```

**Breakdown:**
- First digit: Owner permissions
- Second digit: Group permissions
- Third digit: Other permissions

**Values:**
- 4 = Read
- 2 = Write
- 1 = Execute
- Sum them: 7 = 4+2+1 = rwx

---

## Buffered I/O

For efficient reading/writing:

```go
import "bufio"

// Buffered writing
file, _ := os.Create("output.txt")
defer file.Close()

writer := bufio.NewWriter(file)
writer.WriteString("Line 1\n")
writer.WriteString("Line 2\n")
writer.Flush()  // Important! Flush buffer to disk

// Buffered reading
file, _ := os.Open("input.txt")
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
    fmt.Println(line)
}
```

**Reasoning:** Buffering reduces system calls, improving performance for many small operations.

---

## Summary

**Key Takeaways:**
- `os.WriteFile` / `os.ReadFile` for simple operations
- Always `defer file.Close()` after opening
- Check errors at every step
- Use buffered I/O for performance
- File permissions in octal (0644, 0755)
- `os` package for basic operations, `io` for streaming

**Quick Reference:**
```go
// Write
os.WriteFile("file.txt", []byte("data"), 0644)

// Read
data, _ := os.ReadFile("file.txt")

// Create/Write
file, _ := os.Create("file.txt")
defer file.Close()
io.WriteString(file, "content")

// Exist/Delete/Rename
os.Stat(path)
os.Remove(path)
os.Rename(old, new)
```

---

**📝 Last Updated:** I/O Series
**➡️ Next Topic:** [Web & HTTP](./14_web-and-http.md)
**🔗 Example Code:** [Concepts/18files](../Concepts/18files/)
