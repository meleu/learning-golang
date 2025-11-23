# Notes about "The Power of Go Tools"

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

