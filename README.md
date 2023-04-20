# go-assert-slices
[![Go](https://github.com/kampanosg/go-assert-slices/actions/workflows/go.yml/badge.svg)](https://github.com/kampanosg/go-assert-slices/actions/workflows/go.yml)

Handy and reusable functions that can I have been using when writing tests for Multi-Element Structures aka slices. All functions accept `*testing.T` and you can use them directly in your unit tests.

At the moment the following functions are available:
* `MatchExactly`: The slices have to match _exactly_, including size and order of elements,
* `MatchElements`: The slices have to match in size, but the order of the elements can be different.

:warning: If you're already using https://github.com/stretchr/testify, they already provide an `ElementsMatch` function that does something very similar

## Usage
Import the package alongside `testing`
```golang
import (
    "testing"
    slices "github.com/kampanosg/go-assert-slices"
)
```

You can then use it in your tests
```golang
func TestExample(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	slices.MatchExactly(t, slice1, slice2)
}
```

## Future work
* Similar functions for `map`s.

## Building and Testing
A `Makefile` is provided that allows you to build and test the module
```
make build
make test
```