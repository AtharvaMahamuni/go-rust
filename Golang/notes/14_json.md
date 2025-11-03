# JSON in Go

**📅 Created:** Data Formats Series
**🏷️ Topics:** JSON, encoding/json, Marshal, Unmarshal, Struct Tags
**🔗 Related:** [06_structs.md](./06_structs.md), [05_maps.md](./05_maps.md), [Concepts/22bitmorejson](../Concepts/22bitmorejson/)

---

## Overview

JSON (JavaScript Object Notation) is the most common data interchange format. Go's `encoding/json` package makes it easy to encode Go data structures to JSON and decode JSON to Go types. Understanding JSON is essential for APIs, configuration files, and data storage.

---

## Encoding (Go → JSON)

### Marshal - Basic

```go
import "encoding/json"

type Person struct {
    Name string
    Age  int
}

person := Person{Name: "Alice", Age: 25}

jsonData, err := json.Marshal(person)
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(jsonData))  // {"Name":"Alice","Age":25}
```

### MarshalIndent - Formatted

```go
jsonData, err := json.MarshalIndent(person, "", "  ")
// Parameters: data, prefix, indent

fmt.Println(string(jsonData))
// {
//   "Name": "Alice",
//   "Age": 25
// }
```

**Reasoning:** `MarshalIndent` adds newlines and indentation for readability. Use for debugging or human-readable output.

---

## Struct Tags

Control JSON field names and behavior:

```go
type Course struct {
    Name     string   `json:"coursename"`           // Rename field
    Price    int      `json:"price"`                // Lowercase
    Platform string   `json:"website"`              // Different name
    Password string   `json:"-"`                    // Exclude from JSON
    Tags     []string `json:"tags,omitempty"`       // Omit if empty
    Author   string   `json:"author,omitempty"`     // Omit if zero value
}

course := Course{
    Name:     "Go Programming",
    Price:    99,
    Platform: "example.com",
    Password: "secret123",  // Won't appear in JSON
    Tags:     []string{"programming", "go"},
}

jsonData, _ := json.MarshalIndent(course, "", "  ")
fmt.Println(string(jsonData))
```

**Output:**
```json
{
  "coursename": "Go Programming",
  "price": 99,
  "website": "example.com",
  "tags": ["programming", "go"]
}
```

**Tag Options:**
- `json:"fieldname"` - Custom field name
- `json:"-"` - Exclude field completely
- `json:"field,omitempty"` - Omit if zero value (0, "", nil, empty slice/map)
- No tag - Uses struct field name

**Reasoning:** Struct tags let you control JSON representation without changing Go code. `-` for passwords/secrets, `omitempty` to keep JSON clean.

---

## Decoding (JSON → Go)

### Unmarshal to Struct

```go
jsonStr := `{
    "coursename": "Go Programming",
    "price": 99,
    "website": "example.com",
    "tags": ["programming", "go"]
}`

var course Course

err := json.Unmarshal([]byte(jsonStr), &course)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("%+v\n", course)
// {Name:Go Programming Price:99 Platform:example.com Password: Tags:[programming go]}
```

**Key Points:**
- Pass pointer to `Unmarshal` (&course)
- JSON string must be []byte
- Unmatched JSON fields are ignored
- Missing struct fields remain zero values

**Reasoning:** Unmarshal needs a pointer to modify the struct. Converting string to []byte is required by the API.

### Unmarshal to map[string]interface{}

When structure is unknown:

```go
jsonStr := `{
    "name": "Alice",
    "age": 25,
    "active": true,
    "scores": [95, 87, 92]
}`

var data map[string]interface{}

json.Unmarshal([]byte(jsonStr), &data)

fmt.Println(data["name"])    // Alice
fmt.Println(data["age"])     // 25 (type: float64)
fmt.Println(data["active"])  // true (type: bool)

// Iterate
for key, value := range data {
    fmt.Printf("%s: %v (type: %T)\n", key, value, value)
}
```

**JSON to Go Type Mapping:**
- JSON number → `float64`
- JSON string → `string`
- JSON boolean → `bool`
- JSON null → `nil`
- JSON array → `[]interface{}`
- JSON object → `map[string]interface{}`

**Reasoning:** Use `map[string]interface{}` for dynamic JSON. You lose type safety but gain flexibility. Need type assertions to use values.

---

## Type Assertions for interface{}

```go
if name, ok := data["name"].(string); ok {
    fmt.Println("Name:", name)
}

if age, ok := data["age"].(float64); ok {
    fmt.Println("Age:", int(age))  // Convert to int
}

if scores, ok := data["scores"].([]interface{}); ok {
    for _, score := range scores {
        if s, ok := score.(float64); ok {
            fmt.Println("Score:", int(s))
        }
    }
}
```

**Reasoning:** Type assertions extract concrete types from `interface{}`. Always use comma-ok pattern to avoid panics.

---

## Validating JSON

```go
jsonStr := `{"name": "Alice", "age": 25}`

isValid := json.Valid([]byte(jsonStr))
if isValid {
    fmt.Println("JSON is valid")
} else {
    fmt.Println("Invalid JSON")
}
```

**Reasoning:** Check validity before unmarshaling to avoid errors. Useful for validating user input or external data.

---

## Encoding Arrays/Slices

```go
courses := []Course{
    {Name: "Go", Price: 99},
    {Name: "Rust", Price: 149},
}

jsonData, _ := json.MarshalIndent(courses, "", "  ")
fmt.Println(string(jsonData))
```

**Output:**
```json
[
  {
    "coursename": "Go",
    "price": 99
  },
  {
    "coursename": "Rust",
    "price": 149
  }
]
```

---

## Common Patterns

### Pattern 1: API Response

```go
type APIResponse struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

response := APIResponse{
    Success: true,
    Message: "Data retrieved",
    Data:    someData,
}

json.NewEncoder(w).Encode(response)  // Write to http.ResponseWriter
```

### Pattern 2: Config File

```go
type Config struct {
    Port     int    `json:"port"`
    Host     string `json:"host"`
    Database struct {
        Name string `json:"name"`
        User string `json:"user"`
    } `json:"database"`
}

// Read config
data, _ := os.ReadFile("config.json")
var config Config
json.Unmarshal(data, &config)

// Write config
data, _ := json.MarshalIndent(config, "", "  ")
os.WriteFile("config.json", data, 0644)
```

### Pattern 3: Pretty Print

```go
func prettyPrint(data interface{}) {
    jsonData, _ := json.MarshalIndent(data, "", "  ")
    fmt.Println(string(jsonData))
}

prettyPrint(myStruct)
```

---

## Streaming JSON

For large JSON data:

### Encoder (Write)

```go
file, _ := os.Create("output.json")
defer file.Close()

encoder := json.NewEncoder(file)
encoder.SetIndent("", "  ")

for _, item := range items {
    encoder.Encode(item)  // Write each item
}
```

### Decoder (Read)

```go
file, _ := os.Open("input.json")
defer file.Close()

decoder := json.NewDecoder(file)

var data MyStruct
for {
    err := decoder.Decode(&data)
    if err == io.EOF {
        break
    }
    if err != nil {
        log.Fatal(err)
    }
    // Process data
}
```

**Reasoning:** Streaming is memory-efficient for large files. Don't load entire JSON into memory.

---

## Common Pitfalls

### 1. Forgetting & in Unmarshal

```go
var data MyStruct
json.Unmarshal(jsonBytes, data)  // ❌ Wrong!
json.Unmarshal(jsonBytes, &data) // ✅ Correct
```

### 2. Numbers Become float64

```go
var data map[string]interface{}
json.Unmarshal([]byte(`{"age": 25}`), &data)
age := data["age"]  // float64, not int!

// Must convert
ageInt := int(age.(float64))
```

### 3. Unexported Fields

```go
type Person struct {
    name string  // ❌ Won't be in JSON (lowercase)
    Name string  // ✅ Will be in JSON
}
```

**Reasoning:** Only exported (capitalized) fields are marshaled. Use struct tags to control JSON names.

---

## Summary

**Key Takeaways:**
- `json.Marshal` → Go to JSON
- `json.Unmarshal` → JSON to Go
- Use struct tags to control JSON field names
- `-` to exclude fields, `omitempty` to skip zero values
- Pass pointer to `Unmarshal`
- `map[string]interface{}` for dynamic JSON
- JSON numbers unmarshal as `float64`
- Use `json.Valid()` to validate
- Streaming with Encoder/Decoder for large data

**Quick Reference:**
```go
// Encode
jsonData, _ := json.Marshal(data)
jsonData, _ := json.MarshalIndent(data, "", "  ")

// Decode
json.Unmarshal(jsonBytes, &struct)

// Struct tags
`json:"name"`
`json:"-"`
`json:"name,omitempty"`

// Validate
json.Valid(jsonBytes)
```

---

**📝 Last Updated:** Data Formats Series
**➡️ Next Topic:** [Web & HTTP](./15_web-and-http.md)
**🔗 Example Code:** [Concepts/22bitmorejson](../Concepts/22bitmorejson/)
