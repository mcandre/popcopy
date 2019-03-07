# popcopy - can i help you?

# ABOUT

popcopy recursively copies child file(s) to the desired destination path, excluding patterns you don't want copied.

# EXAMPLE

```go
import (
	"regexp"
	"github.com/mcandre/popcopy"
)

func main() {
	if err := popcopy.Copy(
		"business-presentations",
		"/media/usbstick",
		[]regexp.Regexp{regexp.MustCompile("Thumbs.db")},
	); err != nil {
		panic(err)
	}
}
```

# DOCUMENTATION

https://godoc.org/github.com/mcandre/popcopy

# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.11+

## Recommended

* [Mage](https://magefile.org/) (e.g., `go get github.com/magefile/mage`)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [nakedret](https://github.com/alexkohler/nakedret) (e.g. `go get github.com/alexkohler/nakedret`)
* [shadow](golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow) (e.g. `go get -u golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow`)
* [karp](https://github.com/mcandre/karp) (e.g., `go get github.com/mcandre/karp/...`)

# TEST REMOTELY

```console
$ go test github.com/mcandre/popcopy/...
```

# TEST LOCALLY

```console
$ go test
```

# COVERAGE

```console
$ mage coverageHTML
$ karp cover.html
```

# LINT

Keep the code tidy:

```console
$ mage lint
```
