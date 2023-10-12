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

```go
    fmt.Println("hello")
	var a string
	a = "foo"
	fmt.Println(a)

	var b int = 99
	fmt.Println(b)

	d := 3.14
	fmt.Println(d)

	// var e int = d cannot use d (variable of type float64) as int value in variable declaration
	var c int = int(d)
	fmt.Println(c)
	var e int8 = 54
	fmt.Println(e)

	x, y := 10, 5
	z := x == y
	fmt.Println(z)
	fmt.Println(x - y)
	fmt.Println(x + y)
	fmt.Println(x / y)
	fmt.Println(x % y)
	fmt.Println(x * y)

	// constants
	const g = 42
	var h int = g
	var i float64 = g
	fmt.Println(i, h)

	const j = iota

	fmt.Println(j)

	const (
		k = 2 * 5
		l
		m = iota
	)

	fmt.Println(k, l, m)

	s := "hello"
	p := &s
	fmt.Println(p)
	fmt.Println(*p)

	// dereferncing
	*p = "hello revathy"
	fmt.Println(s)
```

# Create a cli application

fmt - For formatted I/O.
bufio - For buffered I/O operations.
os - For operating system functions.
strings - For string manipulation functions.
net/http - For creating HTTP server and client.

- os
    - stdin
    - stdout
    - strerr
- fmt
    - scan
    - print
    - bufio pkg.go.dev/bufio

 ```go
    fmt.Println("what would you like to me screem")
	in := bufio.NewReader(os.Stdin) //decorator wrapig around the input function
	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	fmt.Println(s + "!")
 ```

 ## web service

 2 clients interacting
 all these by single package 

https://pkg.go.dev/net/http
 ```
 package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe("localhost:7700", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./menu.txt")
	io.Copy(w, f)

}
 ```

# Datatype - aggrigate
## Array

examples arr := [3]string{"foo", "bar", "qux"}
## slice

refrence datatypes

s = append (s, 5, 10, 15 ) // 5,10,15
s = slices.Delete(s, 1, 2) //remve indices 1 and 2 5

## map types

value and keys

var m map[string] int
fmt.println(m)
m=map[string]int{ "foo" :1, "bar":2}
fmt.println(m)

```go
    var m map[string][]string
	fmt.Println(m)

	m = map[string][]string{
		"cofee": {"a", "b", "c"},
	}
	fmt.Println((m))
```

## struct datatype

similar to c struct
structs re copied by values as arrays
```
	fmt.Println("what would you like to me screem")
	fmt.Println(" print1 ")
	in := bufio.NewReader(os.Stdin) //decorator wrapig around the input function
	choice, _ := in.ReadString('\n')
	choice = strings.TrimSpace(choice)
	type mystr struct {
		name  string
		price map[string]float64
	}
	menu := []mystr{
		{name: "coffee", price: map[string]float64{"small": 10, "medium": 20, "large": 30}},
		{name: "tea", price: map[string]float64{"small": 10, "medium": 20, "large": 30}},
	}
	fmt.Println(menu)
```

# looping

infinite loop - for { .. }
loop till condition - for condition { }
counter-based loop - for initialisear;test; post cause{}
```go
    for i := 1; i < 5; i++ {
		fmt.Println(i)
	}
```
# branching

## if statement
```
age := 20

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}
```
## switch statement
logical operators like && (and) and || (or) for complex conditions in if statements and switch cases.
```
day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("It's the start of the week.")
	case "Friday", "Saturday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a regular weekday.")
	}
```
## logical switches
```
age := 25
	isStudent := false

	if age >= 18 && !isStudent {
		fmt.Println("You are eligible to vote.")
	} else {
		fmt.Println("You are not eligible to vote.")
	}
```

```go
type mystr struct {
		name  string
		price map[string]float64
	}
	menu := []mystr{
		{name: "coffee", price: map[string]float64{"small": 10, "medium": 20, "large": 30}},
		{name: "tea", price: map[string]float64{"small": 10, "medium": 20, "large": 30}},
	}
loop:
	for {
		in := bufio.NewReader(os.Stdin)
		fmt.Println(("pelsea select an item in the menu"))
		fmt.Println("1) print menu")
		fmt.Println("2) add item")
		fmt.Println("3) q")
		choice, _ := in.ReadString(('\n'))

		switch strings.TrimSpace(choice) {
		case "1":
			for _, item := range menu {
				fmt.Println(item.name)
				fmt.Println(strings.Repeat("-", 10))
				for size, cost := range item.price {
					fmt.Printf("\t%10s%10.2f\n", size, cost)
				}
			}
		case "2":
			{
				fmt.Println("enter the item")
				name, _ := in.ReadString('\n')
				menu = append(menu, mystr{name: name, price: make(map[string]float64)})
				fmt.Println()
			}
		case "3":
			{
				break loop
			}

		default:
			fmt.Println(("invalid"))

		}
	}
```
##  Deffered function

defer statement that schedules a function call to be run after the function completes. It's often used for cleanup actions.
![Alt text](image.png)

## panic and recovery 

To handle abnormal situation.

## go to statement

