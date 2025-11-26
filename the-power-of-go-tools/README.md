# Notes about "The Power of Go Tools"

My quick impressions about each chapter:

1. paperwork to be able to test program's output
2. paperwork to
    - allow/validate options
    - be able to test data input
3. how to
    - test CLI arguments
    - handling test data
    - [testscript](https://pkg.go.dev/github.com/rogpeppe/go-internal/testscript)

## 1. Packages

### modules and packages

> In Go, a _module_ is a collection of packages that share a common _version_.

### Zen mountaineering

> There's a Zen saying that applies here:
>
> > _if you want to climb a mountain, begin at the top._
>
> In other words, if we want to design a package, a great way to begin is by
> pretending it already exists, and writing code that uses it to solve our problem.

### structure of a test

About `t.Parallel`:

```go
func TestPrintPrintsHelloMessageToTerminal(t *testing.T) {
  t.Parallel()
  // ...
}
```

> The call to `t.Parallel` signals that the test should be run concurrently with
> other tests, and is a standard prelude to any test.

### `bytes.Buffer` is an off-the-shelf `io.Writer`

When we want to test something being printed:

```go
 buf := new(bytes.Buffer)
 hello.PrintTo(buf)
 got := buf.String()
```

For testability purposes, it's more convenient to use:

```go
// 👍 good
fmt.Fprint(os.Stdout, TEXT_TO_PRINT)
```

Rather than using the version that prints directly to stdout:

```go
// 👎 bad
fmt.Print(TEXT_TO_PRINT)
```

### gotestdox

Interesting tool to get a nice output when running the tests: gotestdox

```bash
$ # installation
$ go install github.com/bitfield/gotestdox/cmd/gotestdox@latest

$ # example of output
$ gotestdox
github.com/meleu/hello:
 ✔ PrintTo prints hello message to given writer (0.00s)
```

## 2. Paperwork

Currently we force the client of our package to pass the `os.Stdout` as an
argument in order to print the hello message.

```go
// we want to save the user from having to pass os.Stdout as argument
hello.PrintTo(os.Stdout)
```

In order to make things simpler for our users (and for our tests) we add some
complexity to our package's code.

```go
type Printer struct {
  Output io.Writer
}

// NewPrinter is the constructor of the default Printer (which prints to stdout).
func NewPrinter() *Printer {
  return &Printer{
    Output: os.Stdout
  }
}

func (p *Printer) Print() {
  fmt.Fprintln(p.Output, "Hello, world")
}
```

Test paperwork:

```go
// ...
buf := new(bytes.Buffer)
p := hello.NewPrinter()
p.Output = buf
p.Print()
got := buf.String()
// ...
```

### Simple line counter

Simple _ad-hoc_ solution:

```go
func main() {
  lines := 0
  input := bufio.NewScanner(os.Stdin)
  for input.Scan() {
    lines++
  }
  fmt.Println(lines)
}
```

Now, here comes all the paperwork needed to make such a simple solution testable
and "package-ready"...

```go
// counter is a struct to store the Input.
type counter struct {
  Input io.Reader
}

// NewCounter is a "constructor method" to define os.Stdin as the default Input.
func NewCounter() *counter {
  return &counter{
    Input: os.Stdin,
  }
}

// Lines returns the amount of lines as an int, rather than printing the number.
func (c *counter) Lines() int {
  lines := 0
  input := bufio.NewScanner(c.Input)
  for input.Scan() {
    lines++
  }
  return lines
}

// Main actually prints the amount of lines to os.Stdout
func Main() {
  fmt.Println(NewCounter().Lines())
}
```

### test data coming from/to stdin/stdout

It's actually about testing `fmt.Fprint(os.Stdout)` and
`bufio.NewScanner(os.Stdin).Scan()`.

The thing is that:

- `os.Stdout` and `bytes.Buffer`, both implement `io.Writer`
  - just create `bytes.Buffer` and write to it
- `os.Stdin` and `bytes.Buffer` both implement `io.Reader`
  - create variable via `bytes.NewBufferString(stringToBeRead)` and read from it

### Options as functions

One of the benefits of providing functions to set options is that it allows
validations.

Example (from page 47):

```go
// ...

type counter struct {
  input io.Reader
  output io.Writer
}

type option func(*counter) error

// WithInput validates the given input before using it.
func WitInput(input io.Reader) option {
  return func(c *counter) error {
    if input == nil {
      return errors.New("nil input reader")
    }
    c.input = input
    return nil
  }
}

// ...

// NewCounter is the constructor that applies all the given options (passed
// as functions).
func NewCounter(opts ...option) (*counter, error) {
  c := &counter{
    input: os.Stdin,
    output: os.Stdout,
  }
  // applying all given option functions
  for _, opt := range opts {
    err := opt(c)
    if err != nil {
      return nil, err
    }
  }
  return c, nil
}

// Main now needs to check if the constructor returns with no errors.
func Main() {
  c, err := NewCounter()
  if err != nil {
    panic(err)
  }
  fmt.Println(c.Lines())
}
```

And the test:

```go
func TestLinesCountsLinesInInput(t *testing.T) {
  t.Parallel()
  inputBuf := bytes.NewBufferString("1\n2\n3")
  c, err := count.NewCounter(
    count.WithInput(inputBuf),
  )
  if err != nil {
    t.Fatal(err)
  }
  want := 3
  got := c.Lines()
  if want != got {
    t.Errorf("want %d, got %d", want, got)
  }
}
```

### Methodical options

As mentioned in page 50, methodical options is an alternative that makes sense
when you want to allow **changing** the options after the object is created.

## 3. Arguments
