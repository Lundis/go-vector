package vector2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	a := NewFloat64(1, 2)
	b := NewFloat64(4, 3)
	res := Add(a, b)

	assert.Equal(t, res.X, 5.0)
	assert.Equal(t, res.Y, 5.0)
}

func TestSub(t *testing.T) {
	a := NewFloat64(1, 2)
	b := NewFloat64(4, 3)
	res := Sub(a, b)

	assert.Equal(t, res.X, -3.0)
	assert.Equal(t, res.Y, -1.0)
}

func TestMul(t *testing.T) {
	a := NewFloat64(1, 2)
	b := NewFloat64(4, 3)
	res := Mul(a, b)

	assert.Equal(t, res.X, 4.0)
	assert.Equal(t, res.Y, 6.0)
}

func TestDiv(t *testing.T) {
	a := NewFloat64(1, 2)
	b := NewFloat64(4, 3)
	res := Div(a, b)

	assert.Equal(t, res.X, 0.25)
	assert.Equal(t, res.Y, 0.6666666666666666)
}

func TestDistance(t *testing.T) {
	a := NewFloat64(1, 2)
	b := NewFloat64(4, 3)
	res := Distance(a, b)

	assert.Equal(t, res, 3.1622776601683795)
}

func TestDistanceSquared(t *testing.T) {
	a := NewFloat64(1, 2)
	b := NewFloat64(4, 3)
	res := DistanceSquared(a, b)

	assert.Equal(t, float64(3*3+1*1), res)
}

func TestDot(t *testing.T) {
	a := NewFloat64(1, 2)
	b := NewFloat64(4, 3)
	res := Dot(a, b)

	assert.Equal(t, res, 10.0)
}

func TestReflect(t *testing.T) {
	a := NewFloat64(2, 1)
	b := NewFloat64(6, 6)
	res := Reflect(a, b)

	assert.Equal(t, res.X, -66.0)
	assert.Equal(t, res.Y, -30.0)
}

func TestLerp(t *testing.T) {
	a := NewFloat64(0, 0)
	b := NewFloat64(10, 10)
	res := Lerp(a, b, 0.5)

	assert.Equal(t, res.X, 5.0)
	assert.Equal(t, res.Y, 5.0)
}

func TestMin(t *testing.T) {
	a := NewFloat64(-5, 10)
	b := NewFloat64(17, -13)
	c := Min(a, b)
	assert.Equal(t, -5.0, c.X)
	assert.Equal(t, -13.0, c.Y)
}

func TestMax(t *testing.T) {
	a := NewFloat64(-5, 10)
	b := NewFloat64(17, -13)
	c := Max(a, b)
	assert.Equal(t, 17.0, c.X)
	assert.Equal(t, 10.0, c.Y)
}

func TestIsBetweenInclusive(t *testing.T) {
	a := NewFloat32(1, 2)
	b := NewFloat64(3, 6)
	c := NewInt(5, 10)

	assert.Equal(t, false, IsBetweenInclusive(a, b, c))
	assert.Equal(t, true, IsBetweenInclusive(b, a, c))
	assert.Equal(t, false, IsBetweenInclusive(c, a, b))
}
