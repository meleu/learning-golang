# Learning Golang

Following the [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests/) gitbook.

Here I list some things I learned in each exercise.

## Hello, World

- <https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world>

### What I learned

#### Golang tooling

Start a new project as a module with:

```bash
mkdir hello
cd hello
go mod init hello
# check if the "go.mod" file was created
```

Awesome tool to rerun tests on file change: [watchexec](https://github.com/watchexec/watchexec). Example of usage:

```bash
# run 'go test' when a change happens on a file ending with .go
watchexec -e go -- go test -v
```

Local documentation:

```bash
# this is pretty handy!
go doc fmt
go doc fmt.Println

# Another useful one: pkgsite
# install it:
go install golang.org/x/pkgsite/cmd/pkgsite@latest

# use it to document your project:
pkgsite -open .
```

#### Golang basics

- a program have a `main` package defined with a `main` func inside.
- `func` defines a function with a name and a body (aka block)
- blocks are defined with `{`curly braces`}`
- `import "fmt"` is necessary to use `fmt.Println`
- `if` works like other programming languages, without `(`parenthesis`)`
- variables are ~~assigned~~ declared like this: `varName := value`
  - I [researched](https://stackoverflow.com/a/36513229/6354514) and realized that
    - `:=` for [short variable declarations](https://go.dev/ref/spec#Short_variable_declarations) (with type inference)
    - `=` for [variable declarations](https://go.dev/ref/spec#Variable_declarations) and [assignments](https://go.dev/ref/spec#Assignment_statements).
- constants are defined like `const myConst = "My String"`
- `PublicFunctions` start with a capital letter and `privateFunctions` start with a lowercase.
- `func greetingPrefix(language string) (prefix string)` creates a **named return value**
  - creates a variable called `prefix` in the function
  - it will be assigned the "zero" value. In this case (`string`): `""`
  - example (also showing a `switch` statement):

```go
func greetingPrefix(language string) (prefix string) {
 switch language {
 case spanish:
  prefix = spanishHelloPrefix
 case french:
  prefix = frenchHelloPrefix
 case portuguese:
  prefix = portugueseHelloPrefix
 default:
  prefix = englishHelloPrefix
 }
 return
}
```

- example of grouping constants:

```go
const (
 spanish    = "Spanish"
 french     = "French"
 portuguese = "Portuguese"

 englishHelloPrefix    = "Hello, "
 spanishHelloPrefix    = "Hola, "
 frenchHelloPrefix     = "Bonjour, "
 portugueseHelloPrefix = "Olá, "
)
```

#### Golang testing

- file name must be `${something}_test.go`
- `import "testing"`
- the test function must start with `Test`
- test function takes only one argument `t *testing.T` (it's your "hook" into the testing framework)
- `t.Errorf` prints a message when a test fails.
- `%q` means "string surrounded with double quotes", in the string format context
- subtests go in `t.Run("test name", testFunction)`. Example:

```go
func TestHello(t *testing.T) {
  // 👇 t.Run(testName, testFunction)
 t.Run("say hello to people", func(t *testing.T) {
  actual := Hello("Chris")
  expected := "Hello, Chris!"
  assertCorrectMessage(t, actual, expected)
 })

  // 👇 t.Run(testName, testFunction)
 t.Run("say 'Hello, World!' when passing empty string", func(t *testing.T) {
  actual := Hello("")
  expected := "Hello, World!"
  assertCorrectMessage(t, actual, expected)
 })
}

// comments about this helper function right after this codeblock
func assertCorrectMessage(t testing.TB, actual, expected string) {
 t.Helper() // <-- make it clear that this is a test helper function
 if actual != expected {
  t.Errorf("expected: %q; actual: %q", expected, actual)
 }
}
```

- For helper functions, accept `testing.TB` is a good idea.
- `t.Helper` is needed to report the caller line number when the test fails
  (not the line number in the helper function)

## Integers

- <https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/integers>

### Testable Examples

[Official article](https://go.dev/blog/examples).

Here's an example:

```go
func ExampleAdd() {
 sum := Add(1, 5)
 fmt.Println(sum)
 // Output: 6
}
```

The special comment `// Output: 6` makes the example to be executed.
**NOTE**: for this to work, the function MUST start with `Example`.

This example also goes to the documentation of your package. You can check by
running `pkgsite -open .`

## Iteration

- <https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration>

### Golang

In Go you iterate using `for`. There are **no** `while`, `do`, `until` keywords.

It's usually used like other C-like languages:

```go
for i := 0; i < 5; i++ {
  repeated += character
}
```

Other ways of using `for` are listed here: <https://gobyexample.com/for>

### Benchmarking

- Function must start with `Benchmark`
- Iterate over `b.Loop()`

Example:

```go
func BenchmarkRepeat(b *testing.B) {
  // ... setup...
  for b.Loop() {
    // ... code to measure...
  }
  // ... cleanup...
}
```

Run with:

```bash
go test -bench=.
```

When the benchmark code is executed, it measures how long it takes.
After `Loop()` returns false, `b.N` contains the amount of iterations.

The amount of times shouldn't matter, the framework determine what is a "good" value.

The results show how many times the code was executed and how many nanoseconds it took to run.

## Arrays

<https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices>

### Golang

Arrays can be initialized in two ways:

- `[N]type{value1, value2, ..., valueN}`
  - example: `numbers := [5]int{1, 2, 3, 4, 5}`
- `[...]type{value1, value2, ..., valueN}`
  - example: `numbers := [...]int{1, 2, 3, 4, 5}`

The `%v` placeholder print the variable in the "default" format (in this case
an array).

#### range

The `range` instruction interates over elements in a variety of data structures,
including arrays.

```go
func Sum(numbers [5]int) int {
  sum := 0
  for _, n := range numbers {
    sum += n
  }
  return sum
}
```

- `range` let you iterate over an array
- on each iteration it returns two values: the index and the value
- in the example we are choosing to ignore the index by using the
  `_` [blank identifier](https://go.dev/doc/effective_go#blank)

## Slices

### Golang

The [slice type](https://go.dev/doc/effective_go#slices) allows us to have
collections of any size. The syntax is very similar to arrays, just omit
the size.

Example: `mySlice := []int{1, 2, 3}`

Checking equality of slices:

```go
import "reflect"

reflect.DeepEqual(slice1, slice2)
```

Also:

> From Go 1.21, [slices](https://pkg.go.dev/slices#pkg-overview) standard
> package is available, which has [`slices.Equal`](https://pkg.go.dev/slices#Equal)
> function to do a simple shallow compare on slices, where you don't need to
> worry about the types (...). Note that this function expects the elements to
> be [comparable](https://pkg.go.dev/builtin#comparable).
> So, it can't be applied to slices with non-comparable elements like 2D slices.

Adding elements to a slice:

```go
// append() creates a new slice, therefore you need to assign the variable again
mySlice = append(mySlice, newElement)

// you can use append() to merge two slices:
mySlice = append(mySlice, anotherSlice...)
```

### Golang testing

Check the coverage with

```bash
go test -cover
```

TODO: I still didn't figure out how to have a visual indication of the lines
covered inside the editor (neovim).

## Structs, Methods and Interfaces

### Golang

```go
// declaring a struct
type Rectangle struct {
  Width  float64
  Height float64
}

// declaring a method for a struct
func (r Rectangle) Area() float64 {
  return r.Height * r.Width
}
```

An interesting thing about interfaces in Go (which makes me remember Duck Typing):
**interface resolution is implicit**.

Here's an example of an interface:

```go
type Shape interface {
  Area() float64
}
```

Once this 👆 is declared **any** struct with a method called `Area()` that returns
a `float64` is automatically considered a `Shape`.

We don't need to explicitly say "My type Foo implements interface Bar".

By the way, an interesting thing in Go is that any type with an Error() string
method fulfils the error interface.

### Golang testing

I learned about Table Driven Tests. It's useful but not that easy to read
(without some training).

Here's an example:

```go
func TestArea(t *testing.T) {
  // using Table Drive Tests <https://go.dev/wiki/TableDrivenTests>
  areaTests := []struct {
    name    string
    shape   Shape
    hasArea float64
  }{
    {
      name:    "Rectangle",
      shape:   Rectangle{Width: 12, Height: 6},
      hasArea: 72.0,
    },
    {
      name:    "Circle",
      shape:   Circle{Radius: 10},
      hasArea: 314.1592653589793,
    },
    {
      name:    "Triangle",
      shape:   Triangle{Base: 12, Height: 6},
      hasArea: 36.0,
    },
  }

  for _, tt := range areaTests {
    // using tt.name to use it as the `t.Run` test name
    t.Run(tt.name, func(t *testing.T) {
      got := tt.shape.Area()
      if got != tt.hasArea {
        // the `%#v` format string prints the struct with values in its fields
        t.Errorf("%#v got %g; want %g", tt.shape, got, tt.hasArea)
      }
    })
  }
}
```

## Pointers & errors

### Golang

#### Pointers

A word about pointers from [A Tour of Go](https://go.dev/tour/moretypes/1):

> A pointer holds the memory address of a value.
>
> The type `*T` is a pointer to a `T` value. Its zero value is `nil`.

If a method needs to change some property of its receiver, the method signature
must make it clear that the receiver is a pointer:

```go
// clearly states that the receiver is a pointer to a Wallet
func (w *Wallet) Deposit(amount int) {
  // here we don't need to use `(*w).balance` because struct pointers are
  // automatically dereferenced.
  w.balance += amount
}
```

#### "Rename" a type

This is not actually "renaming", but wrapping the original type in a name that
can be more meaningful to the domain.

```go
type NewName OriginalType

// example from the text:
type Bitcoin int
```

#### Errors

Creating a custom error is this simple:

```go
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")
```

> TODO: why we can't declare a custom error as a constant? Wouldn't it be safer?
>
> This article seems to have the answer: <https://dave.cheney.net/2016/04/07/constant-errors>

#### Unchecked errors

A tool to analyze your code and check if all possible errors are being handled.

```bash
# install it
go install github.com/kisielk/errcheck@latest

# use it with `errcheck $directory`
errcheck .
```

#### Golang testing

Use `t.Fatal` when you want to interrupt the current test and don't need to
go on with more assertions.

```go
func assertError(t testing.TB, got, want error) {
  t.Helper()
  if got == nil {
    // if there's no error, interrupt the test immediately
    t.Fatal("didn't get an error but wanted one")
  }

  if got != want {
    t.Errorf("got %q, want %q", got, want)
  }
}
```

## Maps

A map maps keys to values. It's like a Hash in Ruby.

The way to declare a map is `map[<key_type]<value_type>`.

Example:

```go
// a custom type where the key is a string and the value is also a string.
type Dictionary map[string]string
```

> An interesting property of maps is that you can modify them without passing
> as an address to it.

For example, when declaring a method we don't need to declare the receiver as
a pointer.

```go
func (d Dictionary) Search(word string) (string, error) {
  // ...
}
```

Interesting reading: [If a map isn't a reference variable, what is it?](https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it)

**Gotcha**: maps can be `nil`, and attempts to write toa `nil` map cause a
runtime panic. (recommended reading: [Go maps in action](https://go.dev/blog/maps))

#### Errors

An interesting thing in Go is that any type with an Error() string method
fulfils the error interface.

This is how we made the errors constant:

```go
type DictionaryErr string

func (e DictionaryErr) Error() string {
  return string(e)
}

const (
  ErrNotFound   = DictionaryErr("could not find the word you were looking for")
  ErrWordExists = DictionaryErr("cannot add word because it already exists")
)
```
