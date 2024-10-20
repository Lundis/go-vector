# go-vector
A Vector library for 2D and 3D applications. Forked from [deeean/go-vector](https://github.com/deeean/go-vector), with the main change being that all vector parameters aren't pointers.

[![github-actions](https://github.com/Lundis/go-vector/actions/workflows/ci.yml/badge.svg)](https://github.com/Lundis/go-vector)
[![codecov](https://codecov.io/gh/Lundis/go-vector/graph/badge.svg?token=YAFDG1MSZD)](https://codecov.io/gh/Lundis/go-vector)

## Installation
```shell
go get -u github.com/Lundis/go-vector/[vector3|vector2]
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/Lundis/go-vector/vector3"
)

func main() {
	// basic usage
	a := vector3.New(1, 2, 3)
	b := vector3.New(1, 2, 3)
	fmt.Println(a.Add(b))

	// method chaining
	c := vector3.New(0, 3, 5).
		DivScalars(1, 2, 4).
		MulScalars(0, 5, 3).
		Magnitude()
	fmt.Println(c)
}
```