# Building REST APIs in Go

**📅 Created:** Web Development Series
**🏷️ Topics:** HTTP Server, REST APIs, Routing, Handlers, gorilla/mux, JSON APIs
**🔗 Related:** [14_json.md](./14_json.md), [15_web-and-http.md](./15_web-and-http.md), [Concepts/24buildapi](../Concepts/24buildapi/)

---

## Overview

Go's `net/http` package provides everything needed to build production-ready HTTP servers and REST APIs. Combined with the `gorilla/mux` router, you can create powerful, performant APIs with minimal code. Understanding API development is essential for building web services, microservices, and backend applications.

**Why Build APIs in Go:**
- Built-in HTTP server (no external dependencies required)
- Excellent performance and concurrency
- Simple, clear request/response handling
- Easy JSON integration
- Production-ready standard library

---

## Basic HTTP Server

### Hello World Server

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    http.ListenAndServe(":8080", nil)
}
```

**Access:** `http://localhost:8080/`

**Reasoning:** `http.ListenAndServe` starts the server. The first parameter is the port, the second is the router (nil uses default). The server blocks and handles requests concurrently.

### Basic Handler Function

```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Welcome to my API</h1>"))
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.ListenAndServe(":8080", nil)
}
```

**Handler Signature:**
- `w http.ResponseWriter` - Write response to client
- `r *http.Request` - Read request from client

**Reasoning:** Handlers are just functions with this signature. `http.HandleFunc` registers the handler for a route.

---

## HTTP Methods

### Handling Different Methods

```go
func userHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        fmt.Fprintf(w, "Getting users")
    case "POST":
        fmt.Fprintf(w, "Creating user")
    case "PUT":
        fmt.Fprintf(w, "Updating user")
    case "DELETE":
        fmt.Fprintf(w, "Deleting user")
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprintf(w, "Method not allowed")
    }
}
```

**Reasoning:** `r.Method` contains the HTTP method as a string. You can route based on method manually, but routers make this easier.

---

## Gorilla Mux Router

The standard library router is basic. `gorilla/mux` provides powerful routing.

### Installation

```bash
go get -u github.com/gorilla/mux
```

### Basic Routing

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    // Routes
    r.HandleFunc("/", homeHandler).Methods("GET")
    r.HandleFunc("/users", getAllUsers).Methods("GET")
    r.HandleFunc("/user", createUser).Methods("POST")
    r.HandleFunc("/user/{id}", getUser).Methods("GET")
    r.HandleFunc("/user/{id}", updateUser).Methods("PUT")
    r.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")

    // Start server
    log.Fatal(http.ListenAndServe(":8080", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>API Home</h1>"))
}
```

**Key Features:**
- `.Methods()` - Specify allowed HTTP methods
- `{id}` - Path parameters (dynamic routes)
- `log.Fatal()` - Log and exit if server fails to start

**Reasoning:** Mux provides method-based routing, URL parameters, and better organization than the standard router.

---

## Path Parameters

```go
func getUser(w http.ResponseWriter, r *http.Request) {
    // Extract path parameters
    params := mux.Vars(r)
    userID := params["id"]

    fmt.Fprintf(w, "Getting user with ID: %s", userID)
}

// Route: /user/{id}
// Access: /user/123 → userID = "123"
```

**Multiple Parameters:**

```go
r.HandleFunc("/user/{id}/posts/{postID}", getPost)

func getPost(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    userID := params["id"]
    postID := params["postID"]

    fmt.Fprintf(w, "User: %s, Post: %s", userID, postID)
}
```

**Reasoning:** `mux.Vars(r)` returns a map of path parameters. Clean URLs without query strings.

---

## JSON APIs

### Sending JSON Response

```go
type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func getUser(w http.ResponseWriter, r *http.Request) {
    // Set content type
    w.Header().Set("Content-Type", "application/json")

    user := User{
        ID:    "1",
        Name:  "Alice",
        Email: "alice@example.com",
    }

    // Encode and send
    json.NewEncoder(w).Encode(user)
}
```

**Output:**
```json
{
  "id": "1",
  "name": "Alice",
  "email": "alice@example.com"
}
```

**Reasoning:** `json.NewEncoder(w).Encode()` writes JSON directly to the response. Always set Content-Type header.

### Receiving JSON Request

```go
func createUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Check if body is empty
    if r.Body == nil {
        http.Error(w, "Request body is empty", http.StatusBadRequest)
        return
    }

    // Decode JSON
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Process user...
    fmt.Printf("Created user: %+v\n", user)

    // Send response
    json.NewEncoder(w).Encode(user)
}
```

**Reasoning:** `json.NewDecoder(r.Body).Decode()` reads and parses JSON from request body. Always validate input.

---

## Complete CRUD API Example

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "strconv"
    "time"

    "github.com/gorilla/mux"
)

type Course struct {
    CourseID    string  `json:"courseid"`
    CourseName  string  `json:"coursename"`
    CoursePrice int     `json:"price"`
    Author      *Author `json:"author"`
}

type Author struct {
    Fullname  string `json:"fullname"`
    TwitterID string `json:"twitterid"`
}

// Fake database
var courses []Course

// Helper method
func (c *Course) IsEmpty() bool {
    return c.CourseName == ""
}

func main() {
    fmt.Println("Building REST API")

    r := mux.NewRouter()

    // Seed data
    courses = append(courses, Course{
        CourseID:    "1",
        CourseName:  "Go Programming",
        CoursePrice: 299,
        Author:      &Author{Fullname: "Alice", TwitterID: "@alice"},
    })

    // Routes
    r.HandleFunc("/", serveHome).Methods("GET")
    r.HandleFunc("/courses", getAllCourses).Methods("GET")
    r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
    r.HandleFunc("/course", createOneCourse).Methods("POST")
    r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
    r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

    // Start server
    log.Fatal(http.ListenAndServe(":4000", r))
}

// Handlers

func serveHome(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Welcome to Course API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Get all courses")
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Get one course")
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)

    for _, course := range courses {
        if course.CourseID == params["id"] {
            json.NewEncoder(w).Encode(course)
            return
        }
    }

    json.NewEncoder(w).Encode("No course found with ID: " + params["id"])
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Create one course")
    w.Header().Set("Content-Type", "application/json")

    // Check if body is empty
    if r.Body == nil {
        json.NewEncoder(w).Encode("Please send course data")
        return
    }

    // Decode request
    var course Course
    _ = json.NewDecoder(r.Body).Decode(&course)

    if course.IsEmpty() {
        json.NewEncoder(w).Encode("No data inside JSON")
        return
    }

    // Generate ID
    rand.Seed(time.Now().UnixNano())
    course.CourseID = strconv.Itoa(rand.Intn(100))

    // Add to database
    courses = append(courses, course)

    json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Update one course")
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)

    // Find and remove old entry
    for index, course := range courses {
        if course.CourseID == params["id"] {
            // Remove from slice
            courses = append(courses[:index], courses[index+1:]...)

            // Decode new data
            var updatedCourse Course
            _ = json.NewDecoder(r.Body).Decode(&updatedCourse)

            // Keep same ID
            updatedCourse.CourseID = params["id"]

            // Add back
            courses = append(courses, updatedCourse)

            json.NewEncoder(w).Encode(updatedCourse)
            return
        }
    }

    json.NewEncoder(w).Encode("Course not found")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Delete one course")
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)

    for index, course := range courses {
        if course.CourseID == params["id"] {
            courses = append(courses[:index], courses[index+1:]...)
            break
        }
    }

    json.NewEncoder(w).Encode("Deleted course with ID: " + params["id"])
}
```

**Reasoning:** This is a complete CRUD API with in-memory storage. Real apps would use a database, but the pattern remains the same.

---

## Status Codes

### Setting Status Codes

```go
// Success
w.WriteHeader(http.StatusOK)           // 200
w.WriteHeader(http.StatusCreated)      // 201

// Client errors
w.WriteHeader(http.StatusBadRequest)   // 400
w.WriteHeader(http.StatusNotFound)     // 404

// Server errors
w.WriteHeader(http.StatusInternalServerError)  // 500
```

### Using http.Error

```go
if err != nil {
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
}

if user == nil {
    http.Error(w, "User not found", http.StatusNotFound)
    return
}
```

**Reasoning:** `http.Error` sets status code and sends error message in one call. Must call before writing other data.

---

## Headers

### Setting Response Headers

```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("X-Custom-Header", "MyValue")

    // Must set headers BEFORE WriteHeader or Write
    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(data)
}
```

### Reading Request Headers

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Get header
    auth := r.Header.Get("Authorization")

    // Check if exists
    if auth == "" {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Process request...
}
```

**Reasoning:** Headers must be set before writing response body. `Get` returns empty string if header doesn't exist.

---

## Common Patterns

### Pattern 1: API Response Structure

```go
type APIResponse struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

func sendSuccess(w http.ResponseWriter, data interface{}, message string) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(APIResponse{
        Success: true,
        Message: message,
        Data:    data,
    })
}

func sendError(w http.ResponseWriter, message string, statusCode int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(APIResponse{
        Success: false,
        Error:   message,
    })
}
```

### Pattern 2: Middleware

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", homeHandler)

    // Apply middleware
    r.Use(loggingMiddleware)

    http.ListenAndServe(":8080", r)
}
```

### Pattern 3: CORS Headers

```go
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}
```

**Reasoning:** Middleware wraps handlers to add cross-cutting concerns like logging, auth, CORS without duplicating code.

---

## Query Parameters

```go
func searchHandler(w http.ResponseWriter, r *http.Request) {
    // Get query parameters
    query := r.URL.Query()

    search := query.Get("q")
    limit := query.Get("limit")

    // With defaults
    if limit == "" {
        limit = "10"
    }

    fmt.Fprintf(w, "Search: %s, Limit: %s", search, limit)
}

// Access: /search?q=golang&limit=20
```

**Reasoning:** `r.URL.Query()` returns url.Values (map[string][]string). Use `Get()` to retrieve single value.

---

## File Uploads

```go
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    // Limit size (10 MB)
    r.ParseMultipartForm(10 << 20)

    // Retrieve file
    file, handler, err := r.FormFile("uploadfile")
    if err != nil {
        http.Error(w, "Error retrieving file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    fmt.Printf("Uploaded: %s\n", handler.Filename)
    fmt.Printf("Size: %d bytes\n", handler.Size)

    // Save file
    dst, _ := os.Create("./uploads/" + handler.Filename)
    defer dst.Close()

    io.Copy(dst, file)

    fmt.Fprintf(w, "File uploaded successfully")
}
```

**Reasoning:** Multipart forms handle file uploads. Always limit size to prevent abuse.

---

## Testing APIs

### Manual Testing with curl

```bash
# GET
curl http://localhost:4000/courses

# POST
curl -X POST http://localhost:4000/course \
  -H "Content-Type: application/json" \
  -d '{"coursename":"Go Basics","price":199}'

# GET one
curl http://localhost:4000/course/1

# PUT
curl -X PUT http://localhost:4000/course/1 \
  -H "Content-Type: application/json" \
  -d '{"coursename":"Go Advanced","price":299}'

# DELETE
curl -X DELETE http://localhost:4000/course/1
```

---

## Summary

**Key Takeaways:**
- `http.ListenAndServe` starts the server
- Handlers: `func(w http.ResponseWriter, r *http.Request)`
- gorilla/mux for advanced routing
- `mux.Vars(r)` for path parameters
- `json.NewEncoder(w).Encode()` for JSON responses
- `json.NewDecoder(r.Body).Decode()` for JSON requests
- Set headers before writing body
- Use middleware for cross-cutting concerns
- Always validate input
- Return proper status codes

**Quick Reference:**
```go
// Basic server
http.HandleFunc("/", handler)
http.ListenAndServe(":8080", nil)

// Mux router
r := mux.NewRouter()
r.HandleFunc("/users", getUsers).Methods("GET")
r.HandleFunc("/user/{id}", getUser).Methods("GET")
http.ListenAndServe(":8080", r)

// Path params
params := mux.Vars(r)
id := params["id"]

// JSON response
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(data)

// JSON request
var data MyStruct
json.NewDecoder(r.Body).Decode(&data)

// Error response
http.Error(w, "Error message", http.StatusBadRequest)
```

**When to Use This Note:**
- Building REST APIs
- Creating web services
- Microservices architecture
- Backend development
- Understanding HTTP servers in Go

---

**📝 Last Updated:** Web Development Series
**➡️ Related:** See [14_json.md](./14_json.md) for JSON handling, [15_web-and-http.md](./15_web-and-http.md) for HTTP client
**🔗 Example Code:** [Concepts/24buildapi](../Concepts/24buildapi/)
