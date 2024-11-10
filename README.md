# go-vector
A Vector library for 2D applications. Forked from [deeean/go-vector](https://github.com/deeean/go-vector), with the main change being that the vector parameters aren't pointers, and the struct is generically typed.

[![github-actions](https://github.com/Lundis/go-vector/actions/workflows/ci.yml/badge.svg)](https://github.com/Lundis/go-vector)
[![codecov](https://codecov.io/gh/Lundis/go-vector/graph/badge.svg?token=YAFDG1MSZD)](https://codecov.io/gh/Lundis/go-vector)

## Installation
```shell
go get -u github.com/Lundis/go-vector/vector2
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/Lundis/go-vector/vector2"
)

func main() {
	// basic usage
	a := vector2.NewFloat64(1, 2)
	b := vector2.NewFloat32(3, 4)
	fmt.Println(vector2.Add(a, b))
}
```