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
# installing local documentation
go install golang.org/x/tools/cmd/godoc@latest

# NOTE: sometimes godoc won't be in your path.
# in my case it went to `~/.asdf/` structure and I created an alias:
alias godoc=$(find $HOME/.asdf/installs/golang -type f -path '*packages/bin/godoc')

# now you can launch the local documentation server
godoc -http :8000
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
	t.Helper() // <-- pra que isso?
	if actual != expected {
		t.Errorf("expected: %q; actual: %q", expected, actual)
	}
}
```

- For helper functions, accept `testing.TB` is a good idea.
- `t.Helper` is needed to report the caller line number when the test fails
  (not the line number in the helper function)

## Integers

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

This example also goes to the documentation of your package. You can check by
running `godoc -http :8000` and looking for the `Integers` package.

## Iteration

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

Example:

```go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
```

- `testing.B` gives you access to the (cryptic) `b.N`.
- the benchmark code is executed `b.N` times and measures how long it takes.
  - the amount of times shouldn't matter, the framework determine what is a "good" value.
- run the benchmark with `go test -bench=.`
- the results show how many times the code was executed and how many nanoseconds it took to run.

## Arrays

### Golang

Arrays can be initialized in two ways:

- `[N]type{value1, value2, ..., valueN}`
  - example: `numbers := [5]int{1, 2, 3, 4, 5}`
- `[...]type{value1, value2, ..., valueN}`
  - example: `numbers := [...]int{1, 2, 3, 4, 5}`

There's also the [slice type](https://go.dev/doc/effective_go#slices) which
allows us to have collections of any size. The syntax is very similar to arrays,
just omit the size.

Example: `mySlice := []int{1, 2, 3}`

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
