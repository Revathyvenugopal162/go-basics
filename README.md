# Go Language Study Notes

## Extensions and Commands in Go

- **Extension:** In Go, extensions are not typically used as in some other languages. Go files typically have a `.go` extension.
- **Command Palette and Go Install:** The Command Palette is a tool used in integrated development environments (IDEs) like Visual Studio Code. You can run various Go-related commands using the Command Palette, including `go install all` which installs all the Go packages.
- **`go mod init`:** Initializes a new Go module in the current directory.
- **`go run`:** Compiles and runs Go programs.
- **`fmt` Package:** The `fmt` package in Go is used for formatting strings and printing to the standard output.

## Simple Data Types in Go

- **Strings:**
  - Interpreted strings are enclosed in double quotes `"`
  - Raw strings are enclosed in backticks ``

- **Numbers:**
  - Integers
  - Unsigned integers
  - Floating point numbers
  - Complex numbers

- **Booleans:**
  - `true` and `false` are the two boolean values in Go.

- **Errors:**
  - Errors in Go are represented by a type that implements the error interface.
  - The `error` interface in Go is defined as follows:
    ```go
    type error interface {
        Error() string
    }
    ```

  - More information about built-in types in Go can be found at [pkg.go.dev/builtin](https://pkg.go.dev/builtin).

## Variables in Go

- **Variable Declaration:**
  - `var MyName string`: Declares a variable named `MyName` of type string.

- **Variable Initialization:**
  - `var myName string = "revathy"`: Declares and initializes a string variable with the value "revathy".
  - `var myName = "rev"`: Initializes the variable `myName` with the inferred type string.
  - `myName := "mike"`: Short declaration syntax, declares and initializes a variable with inferred type string.

## Type Conversions in Go

Type conversions in Go are performed using explicit syntax. For example:
```go
var myInt int = 42
var myFloat float64 = float64(myInt)
```
 
# Arithmetic Operations in Go

- **Addition (`+`), Subtraction (`-`), Multiplication (`*`), Division (`/`), Modulus (`%`), etc.**

# Logical and Comparison Operations in Go

- Logical AND (`&&`), OR (`||`), NOT (`!`) are logical operators.
- Comparison operators include equal (`==`), not equal (`!=`), greater than (`>`), less than (`<`), greater than or equal to (`>=`), and less than or equal to (`<=`).
- For detailed information, refer to the [official Go Language Specification](https://go.dev/ref/spec).

# Constants in Go

- **`iota`**: A special identifier in Go which is used in const declarations to simplify definitions of incrementing numbers.
- Constants in Go are created using the `const` keyword and cannot be reassigned after their initial value is set.
- For more detailed information on Go constants and their usage, refer to [Effective Go](https://go.dev/doc/effective_go#constants).

# Pointers and Values

- Pointers in Go are used to share memory addresses.
- Example: `a := 42`, `&a` represents the memory address of `a`, `*b` holds the value in `a`.
- Use pointers to share memory and use copies whenever possible to avoid unintended side effects.

Create a cli application