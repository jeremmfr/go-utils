# go-utils

[![Go Reference](https://pkg.go.dev/badge/github.com/jeremmfr/go-utils.svg)](https://pkg.go.dev/github.com/jeremmfr/go-utils)
[![Release](https://img.shields.io/github/v/release/jeremmfr/go-utils)](https://github.com/jeremmfr/go-utils)
[![Report Card](https://goreportcard.com/badge/github.com/jeremmfr/go-utils)](https://goreportcard.com/report/github.com/jeremmfr/go-utils)
[![Go Status](https://github.com/jeremmfr/go-utils/workflows/Go%20Tests/badge.svg)](https://github.com/jeremmfr/go-utils/actions)
[![Lint Status](https://github.com/jeremmfr/go-utils/workflows/GolangCI-Lint/badge.svg)](https://github.com/jeremmfr/go-utils/actions)
[![codecov](https://codecov.io/gh/jeremmfr/go-utils/branch/main/graph/badge.svg)](https://codecov.io/gh/jeremmfr/go-utils)

Golang utility functions

## Usage

```go
package main

import (
     "fmt"

     "github.com/jeremmfr/go-utils/basiccheck"
)

func main() {
     s1 := []string{"foo", "bar"}
     s2 := []string{"bar", "foo"}

     if basiccheck.SimilarSlice(s1, s2) {
          fmt.Printf("%v =~ %v", s1, s2)
          // Output: [foo bar] =~ [bar foo]
     } else {
          fmt.Printf("%v != %v", s1, s2)
     }
}
```
