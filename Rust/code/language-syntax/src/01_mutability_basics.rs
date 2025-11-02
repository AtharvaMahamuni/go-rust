// 01_mutability_basics.rs
// Topic: Variables are IMMUTABLE by default in Rust
// Compare with Kotlin: val (immutable) vs var (mutable)

fn main() {
    println!("=== Rust Mutability Basics ===\n");

    // Example 1: Immutable by default (like Kotlin's `val`)
    println!("Example 1: Immutable binding");
    let x = 5;
    println!("x = {}", x);
    
    // ❌ UNCOMMENT THIS LINE TO SEE THE ERROR:
    // x = 6;
    
    // The Rust compiler will say:
    // "cannot assign twice to immutable variable `x`"
    // And it will SHOW you where x was first assigned!
    // And it will SUGGEST: "consider making this binding mutable: `mut x`"

    println!("\n---\n");

    // Example 2: Explicit mutability (like Kotlin's `var`)
    println!("Example 2: Mutable binding with `mut`");
    let mut y = 10;
    println!("y before = {}", y);
    y = 20;  // ✅ This works because of `mut`
    println!("y after = {}", y);

    println!("\n---\n");

    // Example 3: The difference from Kotlin
    println!("Example 3: Comparing with your Kotlin example");
    
    let test1 = 1;           // immutable (like `val`)
    let mut test2 = test1;   // mutable (like `var`), copies the VALUE
    
    println!("test1 = {}", test1);
    println!("test2 before = {}", test2);
    
    test2 = 2;  // ✅ Works - test2 is mutable
    
    println!("test2 after = {}", test2);
    println!("test1 still = {}", test1);  // test1 unchanged!

    println!("\n---\n");

    // 🤔 QUESTION FOR YOU:
    // What's happening in memory here?
    // - test1 stores: 1
    // - test2 initially copies: 1  
    // - test2 reassignment stores: 2
    // 
    // For primitive types (integers), Rust COPIES the value
    // This is different from reference types - we'll explore that next!

    println!("Example 4: Multiple mutations");
    let mut counter = 0;
    println!("counter starts at: {}", counter);
    
    counter += 1;
    println!("after +=1: {}", counter);
    
    counter = counter * 2;
    println!("after *2: {}", counter);
    
    counter -= 1;
    println!("after -=1: {}", counter);
}

// 🎯 KEY TAKEAWAYS:
// 1. Rust variables are IMMUTABLE by default (safer default than Kotlin)
// 2. Use `mut` keyword to explicitly allow mutation
// 3. Compiler errors are HELPFUL - they suggest fixes!
// 4. Primitive types (i32, f64, bool, char) are COPIED, not referenced
// 5. `let x = 5` is like Kotlin's `val x = 5`
// 6. `let mut x = 5` is like Kotlin's `var x = 5`

// 🔥 POWER INSIGHT:
// In Kotlin, mutability is opt-OUT (var is default, val is choice)
// In Rust, mutability is opt-IN (immutable is default, mut is choice)
// This makes you THINK before adding mutation → safer concurrent code!
