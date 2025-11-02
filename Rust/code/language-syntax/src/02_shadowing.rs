// 02_shadowing.rs
// Topic: Variable Shadowing - Reusing names for transformations
// Kotlin Comparison: Rust allows `let x` multiple times; Kotlin doesn't allow redeclaring `val x`
// Real-world use: Data transformation pipelines, parsing user input, config processing

// 🎯 REAL-WORLD SCENARIO:
// You're building a CLI tool that reads user input as a String,
// validates it, parses it to an integer, then applies business logic.
// Shadowing lets you reuse the name "value" through each transformation!

fn main() {
    println!("=== Shadowing in Rust ===\n");

    // ═══════════════════════════════════════════════════════════════════
    // Example 1: Basic Shadowing - Same name, new value
    // ═══════════════════════════════════════════════════════════════════
    println!("Example 1: Basic shadowing");
    
    let x = 5;
    println!("x first binding: {}", x);
    
    let x = x + 1;  // 🤔 Reading OLD x, creating NEW x
    println!("x second binding: {}", x);
    
    let x = x * 2;
    println!("x third binding: {}", x);
    
    // Question: How many times did x's scope end in the above code?
    // Answer: 2 times (when second let executed, and when third let executed)
    
    println!("\n---\n");

    // ═══════════════════════════════════════════════════════════════════
    // Example 2: The Power - TYPE CHANGES (impossible in Kotlin!)
    // ═══════════════════════════════════════════════════════════════════
    println!("Example 2: Shadowing with type changes");
    
    let spaces = "   ";              // Type: &str (string slice)
    println!("spaces as string: '{}'", spaces);
    println!("Length: {}", spaces.len());
    
    let spaces = spaces.len();       // Type: usize (integer)
    println!("spaces as number: {}", spaces);
    
    // 🎯 Why is this powerful?
    // In Kotlin you'd need: spacesStr and spacesNum (two names)
    // In Rust: Same logical concept = same variable name!
    
    println!("\n---\n");

    // ═══════════════════════════════════════════════════════════════════
    // Example 3: Real-world - Parsing user input
    // ═══════════════════════════════════════════════════════════════════
    println!("Example 3: Data transformation pipeline");
    
    // Simulating user input (normally from stdin)
    let user_input = "  42  ";           // Type: &str, has whitespace
    println!("Raw input: '{}'", user_input);
    
    let user_input = user_input.trim();  // Type: &str, trimmed
    println!("Trimmed: '{}'", user_input);
    
    let user_input: i32 = user_input.parse().unwrap();  // Type: i32
    println!("Parsed as integer: {}", user_input);
    
    let user_input = user_input * 2;     // Type: i32, transformed
    println!("After business logic: {}", user_input);
    
    // Beautiful! Same name through the entire pipeline!
    // Each transformation is clear and the type system keeps you safe.
    
    println!("\n---\n");

    // ═══════════════════════════════════════════════════════════════════
    // Example 4: Shadowing vs Mutation - When to use which?
    // ═══════════════════════════════════════════════════════════════════
    println!("Example 4: Shadowing vs mut");
    
    // Using mut (mutable variable):
    let mut count = 0;
    count += 1;
    count += 1;
    println!("Using mut: count = {}", count);
    
    // Using shadowing (new bindings):
    let count = 0;
    let count = count + 1;
    let count = count + 1;
    println!("Using shadowing: count = {}", count);
    
    // 🤔 When to use which?
    // mut: When you're modifying the SAME thing over time (counter, accumulator)
    // shadowing: When you're TRANSFORMING to something new (parsing, converting types)
    
    println!("\n---\n");

    // ═══════════════════════════════════════════════════════════════════
    // Example 5: Shadowing and Scopes
    // ═══════════════════════════════════════════════════════════════════
    println!("Example 5: Inner scopes and shadowing");
    
    let x = 5;
    println!("Outer x: {}", x);
    
    {
        let x = x * 2;  // Shadows outer x
        println!("Inner x: {}", x);
        
        let x = x + 1;  // Shadows the inner x!
        println!("Inner x again: {}", x);
    }  // Inner scopes end here
    
    println!("Back to outer x: {}", x);  // Original x is back!
    
    // Key insight: Inner scope shadows are temporary
    // Original binding comes back when inner scope ends
    
    println!("\n---\n");

    // ═══════════════════════════════════════════════════════════════════
    // Example 6: REAL-WORLD - Config file processing
    // ═══════════════════════════════════════════════════════════════════
    println!("Example 6: Real-world config processing");
    
    // Reading config from file (simulated)
    let config = "timeout=30";              // Raw string
    println!("1. Raw config: {}", config);
    
    let config = config.split('=').collect::<Vec<&str>>();  // Split to vec
    println!("2. Split config: {:?}", config);
    
    let config = config[1];                 // Extract value
    println!("3. Extracted value: {}", config);
    
    let config: u32 = config.parse().unwrap();  // Parse to number
    println!("4. Parsed as u32: {}", config);
    
    let config = config * 1000;             // Convert to milliseconds
    println!("5. Final timeout (ms): {}", config);
    
    // Same name "config" through 5 transformations!
    // Each step is a different type, all type-checked at compile-time
}

// ═══════════════════════════════════════════════════════════════════════
// 🔬 EXPERIMENTS TO TRY:
// ═══════════════════════════════════════════════════════════════════════
// 1. In Example 1, add a println! between the second and third `let x`
//    What value prints? Why?

// 2. Try shadowing with a String instead of an integer:
//    let x = String::from("hello");
//    let x = x.len();
//    Run it. Did "hello" get cleaned up? When?

// 3. Try mixing mut and shadowing:
//    let mut x = 5;
//    x = 10;
//    let x = x * 2;
//    What's the final value? Can you still mutate x after shadowing?

// 4. BREAK IT: Try to access the old value after shadowing:
//    let x = 5;
//    let y = x;  // Capture old value
//    let x = 10;
//    println!("{} {}", x, y);  // What prints? Why?

// 5. Type confusion experiment:
//    let x = "5";
//    let x = x.parse::<i32>().unwrap();
//    let x = x + "0";  // What error do you get? Why?

// ═══════════════════════════════════════════════════════════════════════
// 🎯 KEY TAKEAWAYS:
// ═══════════════════════════════════════════════════════════════════════
// 1. Shadowing creates a NEW binding; the old binding's scope ends immediately
// 2. Each `let` can have a DIFFERENT TYPE - impossible in Kotlin!
// 3. Great for data transformation pipelines (parse, validate, transform)
// 4. Inner scopes can shadow outer scopes temporarily
// 5. Shadowing is immutable by default (need `mut` on new binding if you want mutation)
// 6. Old binding is dropped when shadowed (heap cleanup happens immediately)

// ═══════════════════════════════════════════════════════════════════════
// 🔥 POWER INSIGHT:
// ═══════════════════════════════════════════════════════════════════════
// Shadowing is Rust's answer to: "How do we allow type changes while staying
// type-safe?"
// 
// Python: Same name, any type, but runtime errors
// Kotlin: Different names required, compile-time safe but verbose
// Rust: Same name with shadowing, compile-time safe AND clean!
//
// This is why Rust can feel like a scripting language (concise) while
// being as safe as Java (compile-time checks).

// ═══════════════════════════════════════════════════════════════════════
// 🔄 KOTLIN COMPARISON:
// ═══════════════════════════════════════════════════════════════════════
// Kotlin (not allowed):
//   val x = "5"
//   val x = x.toInt()  // ❌ Error: x already declared!
//   
//   // You must do:
//   val xStr = "5"
//   val xInt = xStr.toInt()

// Rust (allowed):
//   let x = "5";
//   let x = x.parse::<i32>().unwrap();  // ✅ Works! Different type!

// Why different?
// - Kotlin: val means "this name-value binding is permanent"
// - Rust: let creates a new binding; old one's scope ends

// ═══════════════════════════════════════════════════════════════════════
// 📝 POST-RUN CHALLENGES:
// ═══════════════════════════════════════════════════════════════════════
// After running this code:

// 1. Explain in your own words: What is the difference between shadowing 
//    and mutation? Give a scenario where each is better.

// 2. Modify Example 3 to handle invalid input. What if user_input can't 
//    be parsed as an integer? (Hint: parse() returns a Result)

// 3. Write equivalent Kotlin code for Example 2 (the spaces example).
//    How many variables do you need? Compare the verbosity.

// 4. Predict: If you shadow a String, when does the old String's heap 
//    memory get freed? Test it by adding println! statements.

// 5. Challenge: Create a real transformation pipeline:
//    Start with "  100°F  " (string with spaces and °F)
//    End with celsius value as f32
//    Use shadowing for each step: trim → remove °F → parse → convert formula

// ═══════════════════════════════════════════════════════════════════════
// ➡️  NEXT TOPIC: 03_constants_and_statics.rs
//    Learn about `const` and `static` - the truly immutable values that
//    live for the entire program duration (unlike let bindings)
// ═══════════════════════════════════════════════════════════════════════
