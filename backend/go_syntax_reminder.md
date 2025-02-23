# Go Syntax Quick Reference

This is a quick guide to common Go syntax, including operators like `:=`, for when you need a refresher.

## Short Variable Declaration (`:=`)

- **Purpose**: Declares and initializes a variable in one line with type inference.
- **Scope**: Only usable inside functions (not at package level).
- **Key Features**:
  - Automatically infers the variable type from the value.
  - Cannot redeclare an existing variable in the same scope with `:=`.
  - Can declare multiple variables at once if at least one is new.

### Examples
```go
x := 42          // x is an int
name := "Alice"  // name is a string
a, b := 10, 20   // a and b are ints

x = 100          // Reassignment with = (works)
x := 200         // Error: no new variables

y := 5
y, z := 10, 15   // Works because z is new
```

---

## Other Common Go Syntax

### Variable Declaration with `var`
- Used for explicit type declaration or package-level variables.
- No type inference with `:=` here.
```go
var x int = 42
var name string = "Bob"
var a, b int = 1, 2  // Multiple variables
```

### Assignment (`=`)
- Reassigns values to already declared variables.
```go
x := 10
x = 20  // Updates x to 20
```

### Constants (`const`)
- Defines immutable values.
```go
const Pi = 3.14
const Name = "GoLang"
```

### Basic Data Types
- `int`, `float64`, `string`, `bool`, etc.
```go
x := 42          // int
y := 3.14        // float64
s := "hello"     // string
b := true        // bool
```

### Conditionals (`if`)
- No parentheses needed around conditions.
```go
x := 10
if x > 5 {
    fmt.Println("x is greater than 5")
}
```

### Loops (`for`)
- Go has only `for` loops (no `while`).
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)  // Prints 0 to 4
}
```

### Functions
- Define with `func`, specify return type after parameters.
```go
func add(a int, b int) int {
    return a + b
}
```

---

## Tips
- Use `:=` for quick, local variable setup.
- Use `var` for package-level variables or when type clarity matters.
- Keep this file nearby to jog your memory!

*Last Updated: February 23, 2025*
