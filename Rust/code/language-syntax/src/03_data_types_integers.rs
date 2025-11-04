// 03_data_types_integers.rs
// Topic: Integer data types in Rust - choosing the right size and signedness
// Kotlin Comparison: Like Kotlin's Byte/Short/Int/Long but with unsigned variants
// Real-world use: Memory optimization, hardware interfacing, protocol implementations

// 🎯 REAL-WORLD SCENARIO:
// You're building a sensor data logger for an IoT device with only 4KB RAM.
// You need to store 1000 temperature readings (0-100°C) and 1000 timestamps.
// Wrong type choice: Uses 8KB (won't fit!). Right choice: Uses 2KB (fits perfectly!).
// In systems programming, type choice = memory efficiency = whether your code runs at all.

fn main() {
    println!("=== INTEGER TYPES IN RUST ===\n");

    // Example 1: The Integer Type Landscape
    // QUESTION: Look at these declarations. What pattern do you see in the naming?
    let tiny_signed: i8 = -128;      // 8-bit signed: -128 to 127
    let tiny_unsigned: u8 = 255;     // 8-bit unsigned: 0 to 255
    let small_signed: i16 = -32_768; // 16-bit signed
    let small_unsigned: u16 = 65_535; // 16-bit unsigned
    let medium_signed: i32 = -2_147_483_648; // 32-bit signed (DEFAULT!)
    let medium_unsigned: u32 = 4_294_967_295; // 32-bit unsigned
    let large_signed: i64 = -9_223_372_036_854_775_808; // 64-bit signed
    let large_unsigned: u64 = 18_446_744_073_709_551_615; // 64-bit unsigned
    
    println!("i8 range: {} to 127", tiny_signed);
    println!("u8 range: 0 to {}", tiny_unsigned);
    
    // QUESTION: Why does 'i' mean signed and 'u' mean unsigned?
    // HINT: Think "integer" vs "unsigned"
    
    // QUESTION: Can you calculate the max value formula?
    // For u8:  0 to 2^8 - 1  = 0 to 255
    // For i8: -2^7 to 2^7 - 1 = -128 to 127
    // Why does signed lose one value on the positive side?

    println!("\n=== TYPE INFERENCE VS EXPLICIT ANNOTATION ===\n");

    // Example 2: Rust's Type Inference is Smart
    let inferred = 42;  // What type is this? Run and see!
    println!("Inferred type size: {} bytes", std::mem::size_of_val(&inferred));
    // ANSWER: i32 (4 bytes) - Rust's default for integers
    
    // KOTLIN COMPARISON:
    // Kotlin: val x = 42  // Type: Int (always 32-bit)
    // Rust:   let x = 42  // Type: i32 (also 32-bit by default)
    // But Rust lets you be MORE specific when needed!
    
    let explicit_small: u8 = 42;
    let explicit_large: i64 = 42;
    println!("u8 size: {} byte", std::mem::size_of_val(&explicit_small));
    println!("i64 size: {} bytes", std::mem::size_of_val(&explicit_large));
    
    // QUESTION: When should you use explicit types vs let Rust infer?
    // Think about: memory efficiency, API requirements, code clarity

    println!("\n=== SIGNED VS UNSIGNED: THE CRITICAL CHOICE ===\n");

    // Example 3: When Signedness Matters
    let age: u8 = 25;           // Age is never negative - use unsigned!
    let temperature: i8 = -15;   // Temperature can be negative - use signed!
    let http_status: u16 = 404;  // HTTP codes are 0-599 - unsigned is perfect!
    
    println!("Age: {} (unsigned makes sense - no negative age)", age);
    println!("Temperature: {}°C (signed needed for freezing temps)", temperature);
    println!("HTTP Status: {} (unsigned - status codes start at 0)", http_status);
    
    // REAL-WORLD DECISION TREE:
    // Can your value be negative? → YES: use i8/i16/i32/i64
    //                             → NO: use u8/u16/u32/u64
    // 
    // How large can it get?
    // 0-255          → u8
    // 0-65,535       → u16  
    // 0-4 billion    → u32
    // Larger         → u64
    //
    // Can it be negative AND how large?
    // -128 to 127    → i8
    // -32K to 32K    → i16
    // -2B to 2B      → i32
    // Larger range   → i64

    println!("\n=== MEMORY EFFICIENCY IN ACTION ===\n");

    // Example 4: The IoT Sensor Problem (from scenario above)
    // Storing 1000 readings - which type to use?
    
    // BAD: Using i32 for temperature (0-100°C)
    let bad_choice_size = std::mem::size_of::<i32>() * 1000;
    println!("Using i32 for 1000 readings: {} bytes", bad_choice_size); // 4000 bytes!
    
    // GOOD: Using u8 for temperature (0-100°C fits in 0-255)
    let good_choice_size = std::mem::size_of::<u8>() * 1000;
    println!("Using u8 for 1000 readings: {} bytes", good_choice_size);  // 1000 bytes!
    
    println!("Memory saved: {} bytes ({}% reduction)", 
             bad_choice_size - good_choice_size,
             ((bad_choice_size - good_choice_size) * 100) / bad_choice_size);
    
    // QUESTION: On a 4KB (4096 bytes) device, which version leaves room for other data?

    println!("\n=== INTEGER LITERALS: DIFFERENT WAYS TO WRITE NUMBERS ===\n");

    // Example 5: Number Literal Formats
    let decimal = 98_222;           // Decimal (normal)
    let hex = 0xff;                 // Hexadecimal (base 16)
    let octal = 0o77;               // Octal (base 8)
    let binary = 0b1111_0000;       // Binary (base 2)
    let byte_literal = b'A';        // Byte literal (u8 only)
    
    println!("Decimal: {}", decimal);
    println!("Hex 0xff: {}", hex);              // QUESTION: What is 0xff in decimal?
    println!("Octal 0o77: {}", octal);          // QUESTION: What is 0o77 in decimal?
    println!("Binary 0b1111_0000: {}", binary); // QUESTION: What is this in decimal?
    println!("Byte literal b'A': {}", byte_literal); // ASCII code for 'A'
    
    // NOTICE: The underscore '_' is just for readability! (like 1_000_000)
    let million = 1_000_000;
    let same_million = 1000000;
    println!("\nAre they equal? {}", million == same_million); // true!
    
    // REAL-WORLD USE:
    // Hex: Color codes (0xFF0000 = red), memory addresses
    // Binary: Bit flags, hardware registers (0b1010 = specific bits set)
    // Byte: ASCII/raw byte data (b'A' = 65)

    println!("\n=== TYPE SUFFIX: TELLING RUST EXACTLY WHAT YOU WANT ===\n");

    // Example 6: Type Suffixes for Precision
    let small = 42u8;      // Force u8 type
    let medium = 42u32;    // Force u32 type
    let large = 42i64;     // Force i64 type
    
    println!("42u8 size: {} byte", std::mem::size_of_val(&small));
    println!("42u32 size: {} bytes", std::mem::size_of_val(&medium));
    println!("42i64 size: {} bytes", std::mem::size_of_val(&large));
    
    // WHEN YOU NEED THIS: Calling functions with specific parameter types
    // fn process_byte(value: u8) { ... }
    // process_byte(42u8);  // Clear intent!
    
    // KOTLIN COMPARISON:
    // Kotlin: val x: Byte = 42    // Must explicitly declare type on left
    // Rust:   let x = 42u8        // Can use suffix on right (more concise!)

    println!("\n=== REAL-WORLD APPLICATION: PROTOCOL PARSER ===\n");

    // Example 7: Parsing a Network Packet Header
    // HTTP status line has: version (1 byte), status code (2 bytes), length (4 bytes)
    
    let http_version: u8 = 1;           // HTTP/1.1 → just store "1"
    let status_code: u16 = 200;         // Status codes go up to 599
    let content_length: u32 = 1_048_576; // Could be up to 4GB
    
    let total_header_size = std::mem::size_of::<u8>() 
                          + std::mem::size_of::<u16>() 
                          + std::mem::size_of::<u32>();
    
    println!("Packet header:");
    println!("  HTTP Version: {} ({} byte)", http_version, std::mem::size_of::<u8>());
    println!("  Status Code: {} ({} bytes)", status_code, std::mem::size_of::<u16>());
    println!("  Content Length: {} bytes ({} bytes to store)", content_length, std::mem::size_of::<u32>());
    println!("  Total header size: {} bytes", total_header_size);
    
    // QUESTION: Why not use i32 for everything? What's the cost?
    // If we used i32 for all three: 4 + 4 + 4 = 12 bytes
    // Our optimized version:        1 + 2 + 4 = 7 bytes (41% smaller!)
    
    // In network protocols, this matters! Smaller headers = faster transmission.

    println!("\n=== SHADOWING WITH TYPE CHANGES (From Session 2!) ===\n");

    // Example 8: Combining Shadowing + Type Transformation
    let value = "42";                    // Type: &str
    println!("Original: '{}' (type: &str)", value);
    
    let value = value.parse::<u8>().unwrap();  // Type: u8
    println!("Parsed: {} (type: u8)", value);
    
    let value = value * 2;               // Type: u8 (still)
    println!("Calculated: {} (type: u8)", value);
    
    let value = value as i16;            // Type: i16 (explicit cast)
    println!("Cast: {} (type: i16)", value);
    
    // QUESTION: Why might we cast u8 to i16 in the last step?
    // HINT: What if the next operation could produce negative numbers?
}

// 🔬 EXPERIMENTS TO TRY:
// 1. Uncomment below and predict the error:
// let overflow: u8 = 256;  // What happens? Why?

// 2. Try changing line with `let inferred = 42` to `let inferred = 42u8`
//    How does the output change? Why?

// 3. Calculate memory for your own use case:
//    If storing 10,000 user IDs (max value: 100,000), what type should you use?
//    Calculate: std::mem::size_of::<TYPE>() * 10_000

// 4. Uncomment and run to see type inference failure:
// let ambiguous = "42".parse().unwrap();  // Compiler error!
// Fix it by adding type annotation: let ambiguous: u32 = ...

// 5. Try mixing signed and unsigned in calculations:
// let unsigned: u8 = 10;
// let signed: i8 = -5;
// let result = unsigned + signed;  // What's the error? How to fix?

// 🎯 KEY TAKEAWAYS:
// - Rust has 8 main integer types: i8/u8, i16/u16, i32/u32, i64/u64
// - 'i' = signed (can be negative), 'u' = unsigned (only positive)
// - Number = bits used (8, 16, 32, 64)
// - Default type: i32 (like Kotlin's Int)
// - Right type choice = memory efficiency = faster code
// - Type suffixes (42u8) let you be explicit
// - Underscores in numbers (1_000_000) are just for readability

// 🔥 POWER INSIGHT:
// In Kotlin/Java, you often use Int for everything and waste memory.
// In Rust, choosing precise types is EXPECTED - you're programming the hardware directly!
// u8 vs i32 isn't just "smaller" - it's 4x less RAM, 4x better CPU cache, 4x more data processed.
// This is why Rust is used for: embedded systems, game engines, OS kernels, databases.
// Every byte matters when you're writing the foundation that everything else runs on.

// 🔄 KOTLIN COMPARISON:
// Kotlin: Int, Long, Short, Byte (all signed, wider range)
// val x: Byte = 127  // -128 to 127
// val y: Int = 42    // Default type
// 
// Rust: i8/u8, i16/u16, i32/u32, i64/u64 (signed AND unsigned options)
// let x: u8 = 255    // 0 to 255 (unsigned gives more range for positives!)
// let y = 42         // Infers to i32 (default)
// let z = 42u8       // Type suffix for precision
// 
// Why different: 
// - Kotlin has GC → less focus on memory efficiency
// - Rust has no GC → you control every byte
// - Unsigned types give you 2x range for positive-only values

// 📝 POST-RUN CHALLENGES:
// After running this:
// 1. Explain in your own words: Why does u8 go to 255 but i8 only goes to 127?
// 2. Calculate: If you store 1 million IPv4 addresses (4 bytes each), 
//    would you use [u8; 4] or u32? Calculate memory difference.
// 3. Write equivalent Kotlin code for Example 3 (age, temperature, http_status)
// 4. Predict: What happens if you do `let x: u8 = 200 + 100;` in debug mode?
// 5. Design: You're storing game scores (0-999). Which type? Why?
// 6. Compare: Open Session 2's shadowing file. How does type-changing 
//    shadowing connect to what you learned here?

// ➡️ NEXT TOPIC: 04_compound_types_tuples.rs
// Now that you understand scalar types (single values), we'll explore
// compound types (multiple values together): tuples and arrays.
// Question to think about: How do you think Rust handles memory for 
// a tuple of (u8, u32, i16)? Same as Kotlin or different?
