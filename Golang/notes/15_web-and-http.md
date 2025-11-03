# Web and HTTP in Go

**📅 Created:** Web Development Series
**🏷️ Topics:** HTTP Client, GET/POST Requests, net/http, URL Handling
**🔗 Related:** [14_json.md](./14_json.md), [04_goroutines-and-concurrency.md](./04_goroutines-and-concurrency.md), [Concepts/19webrequests](../Concepts/19webrequests/), [Concepts/20urls](../Concepts/20urls/)

---

## Overview

Go's `net/http` package provides a complete HTTP client and server. It's production-ready, powers many high-traffic services, and is remarkably simple to use. Understanding HTTP operations is essential for building web services, consuming APIs, and web scraping.

---

## HTTP GET Request - Basic

```go
import (
    "fmt"
    "io"
    "log"
    "net/http"
)

func main() {
    url := "https://api.github.com/users/golang"

    response, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()  // ALWAYS close response body

    fmt.Println("Status:", response.Status)        // 200 OK
    fmt.Println("Status Code:", response.StatusCode)  // 200

    // Read response body
    body, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(body))
}
```

**Key Points:**
1. **Always `defer response.Body.Close()`** - Prevents resource leaks
2. Response body is an `io.Reader` - Use `io.ReadAll` to read fully
3. Check errors at every step

**Reasoning:** HTTP connections are valuable resources. Closing the body ensures the connection can be reused. The defer ensures it happens even if errors occur.

---

## Response Structure

```go
response, _ := http.Get(url)

// Status information
response.Status       // "200 OK"
response.StatusCode   // 200 (int)

// Headers
response.Header       // map[string][]string
contentType := response.Header.Get("Content-Type")

// Body
response.Body         // io.ReadCloser

// Request info
response.Request      // *http.Request
response.ContentLength  // -1 if unknown
```

**Reasoning:** The Response struct contains everything about the HTTP response. Headers are a map where values are slices (headers can repeat).

---

## Custom Request with Headers

```go
// Create request
req, err := http.NewRequest("GET", url, nil)
if err != nil {
    log.Fatal(err)
}

// Add headers
req.Header.Add("Authorization", "Bearer token123")
req.Header.Add("User-Agent", "MyApp/1.0")
req.Header.Set("Accept", "application/json")

// Send request
client := &http.Client{}
response, err := client.Do(req)
if err != nil {
    log.Fatal(err)
}
defer response.Body.Close()
```

**Add vs Set:**
- `Add()` - Appends to existing header
- `Set()` - Replaces existing header

**Reasoning:** Custom requests give you control over headers, method, and body. The Client.Do pattern is the foundation of all HTTP operations in Go.

---

## HTTP POST Request

### POST with JSON

```go
import (
    "bytes"
    "encoding/json"
)

type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

user := User{Name: "Alice", Email: "alice@example.com"}

// Convert to JSON
jsonData, _ := json.Marshal(user)

// Create POST request
response, err := http.Post(
    "https://api.example.com/users",
    "application/json",
    bytes.NewBuffer(jsonData),
)
if err != nil {
    log.Fatal(err)
}
defer response.Body.Close()

// Read response
body, _ := io.ReadAll(response.Body)
fmt.Println(string(body))
```

**Reasoning:** POST typically sends data. JSON is the most common format. `bytes.NewBuffer` converts []byte to io.Reader required by Post.

### POST Form Data

```go
import "net/url"

// Create form data
formData := url.Values{}
formData.Set("username", "alice")
formData.Set("password", "secret123")

response, err := http.PostForm(
    "https://example.com/login",
    formData,
)
defer response.Body.Close()
```

**Reasoning:** `PostForm` is a convenience function for `application/x-www-form-urlencoded` data, common in HTML forms.

---

## Handling Different Response Types

### JSON Response

```go
type GitHubUser struct {
    Login     string `json:"login"`
    Name      string `json:"name"`
    PublicRepos int  `json:"public_repos"`
}

response, _ := http.Get("https://api.github.com/users/golang")
defer response.Body.Close()

var user GitHubUser
json.NewDecoder(response.Body).Decode(&user)

fmt.Println("User:", user.Name)
fmt.Println("Repos:", user.PublicRepos)
```

**Reasoning:** `json.NewDecoder` reads directly from response.Body without loading entire response into memory. Efficient for large JSON.

### Check Status Code

```go
response, _ := http.Get(url)
defer response.Body.Close()

switch response.StatusCode {
case 200:
    fmt.Println("Success")
case 404:
    fmt.Println("Not Found")
case 500:
    fmt.Println("Server Error")
default:
    fmt.Printf("Status: %d\n", response.StatusCode)
}

// Or check ranges
if response.StatusCode >= 200 && response.StatusCode < 300 {
    fmt.Println("Success")
} else if response.StatusCode >= 400 {
    fmt.Println("Error")
}
```

---

## Timeouts and Custom Client

```go
client := &http.Client{
    Timeout: 10 * time.Second,  // Request timeout
}

response, err := client.Get(url)
if err != nil {
    if err, ok := err.(net.Error); ok && err.Timeout() {
        fmt.Println("Request timed out")
    } else {
        fmt.Println("Error:", err)
    }
    return
}
defer response.Body.Close()
```

**Reasoning:** Default client has no timeout - can hang forever. Always set timeouts in production code.

### Custom Transport

```go
client := &http.Client{
    Timeout: 10 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 10,
        IdleConnTimeout:     90 * time.Second,
    },
}
```

**Reasoning:** Transport controls connection pooling and reuse. Custom transport improves performance for high-volume applications.

---

## URL Parsing and Building

```go
import "net/url"

// Parse URL
rawURL := "https://example.com/path?key=value&foo=bar#section"
parsedURL, err := url.Parse(rawURL)
if err != nil {
    log.Fatal(err)
}

fmt.Println("Scheme:", parsedURL.Scheme)      // https
fmt.Println("Host:", parsedURL.Host)          // example.com
fmt.Println("Path:", parsedURL.Path)          // /path
fmt.Println("RawQuery:", parsedURL.RawQuery)  // key=value&foo=bar
fmt.Println("Fragment:", parsedURL.Fragment)  // section

// Parse query parameters
params := parsedURL.Query()  // url.Values (map[string][]string)
fmt.Println("key:", params.Get("key"))  // value
fmt.Println("foo:", params.Get("foo"))  // bar

// Build URL with parameters
baseURL := "https://api.example.com/search"
params := url.Values{}
params.Add("q", "golang")
params.Add("limit", "10")

fullURL := baseURL + "?" + params.Encode()
// https://api.example.com/search?limit=10&q=golang
```

**Reasoning:** URL parsing handles encoding, special characters, and structure. Always use `url.Values` to build query strings - it handles encoding correctly.

---

## Common Patterns

### Pattern 1: API Client

```go
type APIClient struct {
    BaseURL    string
    HTTPClient *http.Client
    Token      string
}

func NewAPIClient(baseURL, token string) *APIClient {
    return &APIClient{
        BaseURL: baseURL,
        Token:   token,
        HTTPClient: &http.Client{
            Timeout: 30 * time.Second,
        },
    }
}

func (c *APIClient) Get(endpoint string) ([]byte, error) {
    url := c.BaseURL + endpoint

    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "Bearer "+c.Token)

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return nil, fmt.Errorf("status: %d", resp.StatusCode)
    }

    return io.ReadAll(resp.Body)
}

// Usage
client := NewAPIClient("https://api.example.com", "token123")
data, err := client.Get("/users")
```

### Pattern 2: Retry Logic

```go
func httpGetWithRetry(url string, maxRetries int) (*http.Response, error) {
    var resp *http.Response
    var err error

    for i := 0; i < maxRetries; i++ {
        resp, err = http.Get(url)
        if err == nil && resp.StatusCode < 500 {
            return resp, nil
        }

        // Close body if request succeeded but status was 500+
        if resp != nil {
            resp.Body.Close()
        }

        time.Sleep(time.Second * time.Duration(i+1))  // Backoff
    }

    return nil, fmt.Errorf("max retries exceeded: %w", err)
}
```

### Pattern 3: Concurrent Requests

```go
func fetchURLs(urls []string) []string {
    results := make([]string, len(urls))
    var wg sync.WaitGroup

    for i, url := range urls {
        wg.Add(1)

        go func(index int, u string) {
            defer wg.Done()

            resp, err := http.Get(u)
            if err != nil {
                results[index] = "Error: " + err.Error()
                return
            }
            defer resp.Body.Close()

            results[index] = resp.Status
        }(i, url)
    }

    wg.Wait()
    return results
}
```

---

## Download File

```go
func downloadFile(url, filepath string) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, resp.Body)
    return err
}

// Usage
downloadFile("https://example.com/file.zip", "downloaded.zip")
```

**Reasoning:** `io.Copy` efficiently streams from response to file without loading entire content into memory.

---

## Context for Cancellation

```go
import "context"

func fetchWithContext() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        // Could be timeout or cancellation
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    // Process response...
}
```

**Reasoning:** Context allows cancellation and timeouts at the request level. Essential for long-running operations or user cancellations.

---

## Common Pitfalls

### 1. Not Closing Response Body

```go
// ❌ Leak - body never closed
resp, _ := http.Get(url)

// ✅ Always close
resp, _ := http.Get(url)
defer resp.Body.Close()
```

### 2. Not Reading Body on Error Status

```go
resp, _ := http.Get(url)
defer resp.Body.Close()

// ❌ Wrong - may leak connection
if resp.StatusCode != 200 {
    return fmt.Errorf("bad status")
}

// ✅ Read and discard body
if resp.StatusCode != 200 {
    io.Copy(io.Discard, resp.Body)
    return fmt.Errorf("bad status: %d", resp.StatusCode)
}
```

### 3. No Timeout

```go
// ❌ Can hang forever
resp, _ := http.Get(url)

// ✅ Set timeout
client := &http.Client{Timeout: 10 * time.Second}
resp, _ := client.Get(url)
```

---

## Summary

**Key Takeaways:**
- Always `defer response.Body.Close()`
- Use `io.ReadAll` to read response body
- Set timeouts (default client has none)
- Check `StatusCode` before processing
- Use `NewRequest` + `Client.Do` for control
- JSON: `json.NewDecoder(resp.Body).Decode()`
- URL parsing: `url.Parse()` and `url.Values`
- Read body even on error to reuse connections
- Use Context for cancellation

**Quick Reference:**
```go
// GET
resp, _ := http.Get(url)
defer resp.Body.Close()
body, _ := io.ReadAll(resp.Body)

// POST JSON
jsonData, _ := json.Marshal(data)
resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonData))

// Custom request
req, _ := http.NewRequest("GET", url, nil)
req.Header.Set("Authorization", "Bearer token")
client := &http.Client{Timeout: 10 * time.Second}
resp, _ := client.Do(req)

// Parse URL
u, _ := url.Parse(rawURL)
params := u.Query()
```

**When to Use This Note:**
- Building API clients
- Consuming REST APIs
- Web scraping
- Downloading files
- Understanding HTTP in Go

---

**📝 Last Updated:** Web Development Series
**➡️ Related:** See [04_goroutines-and-concurrency.md](./04_goroutines-and-concurrency.md) for concurrent requests
**🔗 Example Code:** [Concepts/19webrequests](../Concepts/19webrequests/), [Concepts/20urls](../Concepts/20urls/)
