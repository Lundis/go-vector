# go-vector


## Installation
```shell
go get -u github.com/KimByungChan/go-vector
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/KimByungChan/go-vector/vector3"
)

func main() {
	a := vector3.New(1, 2, 3)
	b := vector3.New(1, 2, 3)
	fmt.Println(a.Add(b))
}
```

## Documentation
```go
func New(x float64, y float64, z float64) *Vector3

func (v *Vector3) Copy() *Vector3

func (v *Vector3) Set(x float64, y float64, z float64) *Vector3

func (v *Vector3) Add(other *Vector3) *Vector3

func (v *Vector3) Sub(other *Vector3) *Vector3

func (v *Vector3) Mul(other *Vector3) *Vector3

func (v *Vector3) MulByScalar(scalar float64) *Vector3

func (v *Vector3) Div(other *Vector3) *Vector3

func (v *Vector3) DivByScalar(scalar float64) *Vector3

func (v *Vector3) Distance(other *Vector3) float64

func (v *Vector3) Dot(other *Vector3) float64

func (v *Vector3) Cross(other *Vector3) *Vector3

func (v *Vector3) Lerp(other *Vector3, t float64) *Vector3

func (v *Vector3) Magnitude() float64

func (v *Vector3) Normalize() *Vector3

func (v *Vector3) Equals(other *Vector3) bool
```