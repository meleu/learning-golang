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

Recommended linter: [GolangCI-Lint](https://golangci-lint.run/). It can be installed with Homebrew or `asdf` (I used `asdf`).

Useful config to be used in VSCode:

```json
{
  "go.lintTool": "golangci-lint",
  "go.lintFlags": ["--fast"],
  "go.coverOnSave": true,
  "go.coverageDecorator": {
    "type": "gutter",
    "coveredHighlightColor": "rgba(64,128,128,0.5)",
    "uncoveredHighlightColor": "rgba(128,64,64,0.25)",
    "coveredGutterStyle": "blockgreen",
    "uncoveredGutterStyle": "slashred"
  }
}
```

Awesome tool to rerun tests on file change: [watchexec](https://github.com/watchexec/watchexec). Example of usage:

```bash
# run 'go test' when a change happens on a file ending with .go
watchexec -e go -- go test -v
```

Local documentation:

```bash
# There are two methods:

# 1. just run `go doc <PKGNAME>`
# example: check `fmt` documentation directly on the terminal
go doc fmt
# the output is massive...

# 2. using `pkgsite`
# a) install the pkgsite command, with the official package viewing website
go install golang.org/x/pkgsite/cmd/pkgsite@latest

# b) run it on your pkg's dir
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
- `t.Errorf` mark a test as failed and prints a formatted message.
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
 t.Helper() // <-- pra que isso?
 if actual != expected {
  t.Errorf("expected: %q; actual: %q", expected, actual)
 }
}
```

- For testing helper functions, use `testing.TB` so you can use `t.Helper`
- `t.Helper` is useful to report the caller line number when the test fails
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

## Iteration

- <https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration>

### Golang

#### for Loops

In Go you iterate using `for`. There are **no** `while`, `do`, `until` keywords.

It's usually used like other C-like languages:

```go
for i := 0; i < 5; i++ {
 repeated += character
}
```

Other ways of using `for` are listed here: <https://gobyexample.com/for>

#### string and strings.Builder

Strings are immutable, therefore each concatenation involves copying memory to accommodate the new string (which impacts performance).

The standard library provides the `strings.Builder` type. It implements a `WriteString` method that can be used to concatenate strings. Like this:

```go
const repeatCount = 5

func Repeat(character string) string {
 var repeated strings.Builder
 for range repeatCount {
  repeated.WriteString(character)
 }
 return repeated.String()
}
```

### Benchmarking

Typical structure of a benchmark:

```go
func Benchmark(b *testing.B) {
  // ... setup...
 for b.Loop() {
  // ... code to measure...
 }
  // ... cleanup...
}
```

- official documentation: <https://golang.org/pkg/testing/#hdr-Benchmarks>
- `b.Loop()` returns true as long as the benchmark should continue running.
- `testing.B` gives you access to the (cryptic) `b.N`.
- the benchmark code is executed `b.N` times and measures how long it takes.
  - the amount of times shouldn't matter, the framework determine what is a "good" value.
- run the benchmark with `go test -bench=.`
- the results show how many times the code was executed and how many nanoseconds it took to run.

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

Let's check the `range` instruction:

```go
func Sum(numbers [5]int) int {
 sum := 0
 // numbers is the array given as argument
 for _, number := range numbers {
  sum += number
 }
 return sum
}
```

- `range` let you iterate over an array
- on each iteration it returns two values: the index and the value
- in the example we are choosing to ignore the index by using the
  `_` [blank identifier](https://go.dev/doc/effective_go#blank)

## Slices

- <https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices>

### Golang

#### Slices

The [slice type](https://go.dev/doc/effective_go#slices) allows us to have
collections of any size. The syntax is very similar to arrays, just omit
the size.

Example: `mySlice := []int{1, 2, 3}`

#### Adding elements to a slice

```go
// append() creates a new slice, therefore you need to assign the variable again
mySlice = append(mySlice, newElement)

// you can use append() to merge two slices:
mySlice = append(mySlice, anotherSlice...)
```

#### Slicing slices

Slices can be sliced with `slice[low:high]`. If you omit the value on one of
the sides of the `:` it captures everything to that side of it.

Example: `number[1:]` means "take from 1 to the end".

#### Checking equality of slices

Starting from Go 1.21, the slices standard package has [the `slices.Equal`](https://pkg.go.dev/slices#Equal)
function to do simple shallow compare on slices.

```go
func TestSumAll(t *testing.T) {
 got := SumAll([]int{1, 2}, []int{0, 9})
 want := []int{3, 9}

 if !slices.Equal(got, want) {
  t.Errorf("got %v want %v", got, want)
 }
}
```

Before Go 1.21, the text was using [the `reflect.DeepEqual`](https://pkg.go.dev/reflect#DeepEqual) function.

```go
import "reflect"

reflect.DeepEqual(slice1, slice2)
```

### Golang testing

Check the coverage with

```bash
go test -cover
```

## Structs, Methods and Interfaces

- <https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces>

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

Once this 👆 is declared **any** struct with a method called `Area()` returning
a `float64` is automatically considered a `Shape`.

We don't need to explicitly say "My type Foo implements interface Bar".

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

## Pointers & Errors

- <https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors>

### Golang

#### Pointer receivers

Go passes arguments by value. This includes the receiver of a method. So, if
your method needs to mutate state, you need to use pointers.

Example:

```go
// example where the wallet's balance needs to be updated
func (w *Wallet) Deposit(amount Bitcoin) { w.balance += amount}
```

- It's usual to always use pointers for receivers, for the sake of consistency.
- Auto-derefenrencing:
  - you write `w.balance`, not `(*w).balance` (although both are valid)
  - similarly, you can call `wallet.Deposit(...)`, not `&wallet.Deposit(...)`
- other reasons to use pointers:
  - large structs (avoid copying)
  - types that must be shared (e.g.: DB pools, mutexes)

#### Named types from primitives

Example:

```go
type Bitcoin int
```

It's not a typedef or alias, it's a distinct type. The compiler refuses to mix
`Bitcoin` and `int` without explicit conversion.

Example:

- `var b Bitcoin = 5` compiles
- `var b Bitcoin = someInt` ERROR! The value must be converted with `Bitcoin(someInt)`

It's useful to get domain meaning and a type to hang methods. This way we can
make "primitive-like" types to satisfy interfaces.

**Note**: as `Bitcoin` has the `int` type, it can use the underlying `int`
operators, like `+`, `-`, `<`, `>`. That's why `w.balance += amount` works.

#### The Stringer interface

The go-way to declare a `toString()` method is by defining a `String()` method
that returns a `string`.

```go
func (b Bitcoin) String() string {
  return fmt.Sprintf("%d BTC", b)
}
```

This is enough to satisfy the `Stringer` interface, and  now we can do this:

```go
bitcoint := Bitcoin(10)
fmt.Printf("%s", bitcoin)
// => 10 BTC
```

#### Errors

Go has no "Exceptions". Functions signal failure by returning an `error` as the
last value, and the caller expcitly checks it.

The `error` "type" is actually an interface: `interface{ Error() string}`.
Therefore any time with an `Error() string` method can be considered an `error`.

**Errors are values**. It's usual to declare them once as a package var, using
the `Err` prefix (e.g.: `ErrInsufficientFunds`). Callers can compare with
`err == ErrSomething` or `errors.Is(err, ErrSomething`

Although `err.Error()` returns a string, your tests should not check this string,
use the `ErrXxx` instead.

If error is `nil`, it means success. Return `nil` on the happy path. Callers
check success with `if err != nil { erroHandling... }`. You'll see this a lot
in Go code.

Creating errors:

- `errors.New`: for static messages (useful for custom errors definitions)
- `fmt.Errorf("...: %w", err)`: useful for when you need to wrap/format

Interesting reading about constant errors: <https://dave.cheney.net/2016/04/07/constant-errors>

## Maps

- <https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/maps>

### Constant Errors

Useful technique to make errors immutable:

```go
// making the errors immutable and more reusable.
// See: https://dave.cheney.net/2016/04/07/constant-errors
const (
  ErrNotFound         = DictionaryErr("could not find the word you were looking for")
  ErrWordExists       = DictionaryErr("this word already has a definition")
  ErrWordDoesNotExist = DictionaryErr("this word does not exist in the dictionary")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
  return string(e)
}
```

Another useful reading: <https://go.dev/blog/go1.13-errors>
