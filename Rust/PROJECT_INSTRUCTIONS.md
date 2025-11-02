# Project Instructions: Rust Syntax Learning Files

## Format for All Rust Syntax Code Files

### File Naming Convention
`XX_topic_name.rs` where XX is sequential number (01, 02, 03...)

### File Structure Template

```rust
// XX_topic_name.rs
// Topic: [Clear one-line description]
// Kotlin Comparison: [How this differs/relates to Kotlin]
// Real-world use: [Where you'll see this in actual codebases]

// 🎯 REAL-WORLD SCENARIO:
// [Brief 2-3 line scenario that motivates WHY this syntax matters]
// Example: "You're building a web server that handles user sessions..."

fn main() {
    // Example 1: [Descriptive name]
    // [Brief explanation + question to think about]
    
    // Example 2: [Next concept]
    
    // ... 3-5 examples per file
    
    // Example N: REAL-WORLD APPLICATION
    // [Larger example showing how all concepts work together]
}

// 🔬 EXPERIMENTS TO TRY:
// 1. [Uncomment line X and predict the error]
// 2. [Modify value Y to Z - what changes?]
// 3. [Try breaking this code by ___. What happens?]

// 🎯 KEY TAKEAWAYS:
// - [Core concept 1]
// - [Core concept 2]
// - [Core concept 3]

// 🔥 POWER INSIGHT:
// [One deep insight connecting to bigger Rust philosophy]

// 🔄 KOTLIN COMPARISON:
// Kotlin: [code example]
// Rust:   [equivalent code]
// Why different: [reason]

// 📝 POST-RUN CHALLENGES:
// After running this:
// 1. Explain in your own words: [question]
// 2. Modify the code to: [task]
// 3. Write equivalent Kotlin code for Example X
// 4. Predict what happens if: [scenario]

// ➡️  NEXT TOPIC: [Preview of next file]
```

## Teaching Principles

### For Each Code File:

1. **Start with real-world motivation** - Why does this syntax exist?
2. **Progress from simple to complex** - 3-5 examples building on each other
3. **Include experiments:**
   - Commented code to uncomment
   - Values to modify and predict outcomes
   - Ways to intentionally break code
4. **Ask questions in comments** rather than just explaining
5. **Compare with Kotlin** when relevant
6. **End with comprehensive review:**
   - Key takeaways
   - Power insights (deeper "why")
   - Post-run challenges
   - Next topic preview

### Learning Flow:

```
Read code → Run code → Observe output → Answer questions → 
Modify code → Break code → Fix code → Compare with Kotlin → 
Explain understanding → Ready for next topic
```

### Real-World Examples Should Include:

- Web servers / API handling
- Data processing pipelines
- Configuration management
- CLI tools
- File operations
- Concurrent operations (later)
- Error handling in production code

### Question Types to Include:

- **Prediction**: "What will this print?"
- **Analysis**: "Why does this work but that doesn't?"
- **Comparison**: "How would Kotlin handle this?"
- **Application**: "Modify this to do X"
- **Breaking**: "Try making this fail - what's the error?"
- **Synthesis**: "Explain this concept to someone else"

## Learning Pace

- **One file = One focused concept**
- Each file should take 15-20 minutes to work through
- Include 3-5 experiments per file
- Balance between "show" and "discover"

## File Organization in Project

```
outputs/
├── 01_mutability_basics.rs          ✅ Done
├── 02_shadowing.rs                   ← Next
├── 03_data_types_integers.rs
├── 04_compound_types_tuples.rs
├── 05_compound_types_arrays.rs
├── 06_functions_basics.rs
├── 07_expressions_vs_statements.rs
├── 08_control_flow_if.rs
├── 09_control_flow_loops.rs
├── 10_pattern_matching_basics.rs
└── ... (continue based on chapter 3 topics)
```

## Success Criteria

Student should be able to:
1. Read the code and predict behavior
2. Run it and verify predictions
3. Modify it successfully
4. Break it intentionally and understand errors
5. Explain the concept in their own words
6. Compare/contrast with Kotlin
7. Recognize the pattern in real codebases

---

**Remember:** 
- More questions, fewer direct explanations
- Real-world context always
- Learn by doing, breaking, and fixing
- Connect everything back to Kotlin for familiarity
- Make Rust compiler errors your teaching assistant!
