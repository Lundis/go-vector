package vector2

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestPlus(t *testing.T) {
	a := NewFloat64(1, 2)
	b := NewFloat64(4, 3)
	res := a.Plus(b)

	assert.Equal(t, res.X, 5.0)
	assert.Equal(t, res.Y, 5.0)
}

func TestAddScalar(t *testing.T) {
	a := NewFloat64(1, 2)
	res := a.AddScalar(1.0)

	assert.Equal(t, res.X, 2.0)
	assert.Equal(t, res.Y, 3.0)
}

func TestAddScalars(t *testing.T) {
	a := NewFloat64(1, 2)
	res := a.AddScalars(1, 2)

	assert.Equal(t, res.X, 2.0)
	assert.Equal(t, res.Y, 4.0)
}

func TestMinus(t *testing.T) {
	a := NewFloat64(1, 2)
	b := NewFloat64(4, 3)
	res := a.Minus(b)

	assert.Equal(t, res.X, -3.0)
	assert.Equal(t, res.Y, -1.0)
}

func TestSubScalar(t *testing.T) {
	a := NewFloat64(1, 2)
	res := a.SubScalar(1.0)

	assert.Equal(t, res.X, 0.0)
	assert.Equal(t, res.Y, 1.0)
}

func TestSubScalars(t *testing.T) {
	a := NewFloat64(1, 2)
	res := a.SubScalars(1, 2)

	assert.Equal(t, res.X, 0.0)
	assert.Equal(t, res.Y, 0.0)
}

func TestMulMethod(t *testing.T) {
	a := NewFloat64(1, 2)
	b := NewFloat64(4, 3)
	res := a.Mul(b)

	assert.Equal(t, res.X, 4.0)
	assert.Equal(t, res.Y, 6.0)
}

func TestMulScalar(t *testing.T) {
	a := NewFloat64(1, 2)
	b := 4.0
	res := a.MulScalar(b)

	assert.Equal(t, res.X, 4.0)
	assert.Equal(t, res.Y, 8.0)
}

func TestMulScalars(t *testing.T) {
	a := NewFloat64(1, 2)
	res := a.MulScalars(1.0, 2.0)

	assert.Equal(t, res.X, 1.0)
	assert.Equal(t, res.Y, 4.0)
}

func TestDivMethod(t *testing.T) {
	a := NewFloat64(1, 2)
	b := NewFloat64(4, 3)
	res := a.Div(b)

	assert.Equal(t, res.X, 0.25)
	assert.Equal(t, res.Y, 0.6666666666666666)
}

func TestDivScalar(t *testing.T) {
	a := NewFloat64(1, 2)
	b := 4.0
	res := a.DivScalar(b)

	assert.Equal(t, res.X, 0.25)
	assert.Equal(t, res.Y, 0.5)
}

func TestDivScalars(t *testing.T) {
	a := NewFloat64(1, 2)
	res := a.DivScalars(1.0, 2.0)

	assert.Equal(t, res.X, 1.0)
	assert.Equal(t, res.Y, 1.0)
}

func TestMagnitude(t *testing.T) {
	a := NewFloat64(3, 2)
	res := a.Magnitude()

	assert.Equal(t, res, 3.605551275463989)
}

func TestNormalize(t *testing.T) {
	a := NewFloat64(3, 2)
	b := NewFloat64(0.000003, 0.000002)
	aa := a.Normalize()
	bb := b.Normalize()

	assert.True(t, aa.Equals(bb))
}

func TestNormalizeZero(t *testing.T) {
	a := NewFloat64(0, 0)
	aa := a.Normalize()

	assert.True(t, aa.Equals(a))
}

func TestInvertY(t *testing.T) {
	a := NewFloat64(1, 2)
	res := a.InvertY()
	assert.Equal(t, -2.0, res.Y)
}

func TestAngle(t *testing.T) {
	a := NewFloat64Polar(1, 17)
	res := a.Angle()
	assert.Equal(t, 1.0, res)

	a = NewFloat64Polar(1.84, 17)
	res = a.Angle()
	assert.Equal(t, 1.84, res)
}

func TestAbs(t *testing.T) {
	a := NewFloat64(-1, -2)
	res := a.Abs()
	assert.Equal(t, 1.0, res.X)
	assert.Equal(t, 2.0, res.Y)
}

func TestClamp(t *testing.T) {
	low := NewFloat64(-5, -13)
	high := NewFloat64(17, 10)
	a := NewFloat64(-100, -100).Clamp(low, high)
	b := NewFloat64(100, 100).Clamp(low, high)
	assert.Equal(t, low.X, a.X)
	assert.Equal(t, low.Y, a.Y)
	assert.Equal(t, high.X, b.X)
	assert.Equal(t, high.Y, b.Y)
}

func TestRound(t *testing.T) {
	a := NewFloat64(-3.4, 5.5).Round()
	assert.Equal(t, math.Round(-3.4), a.X)
	assert.Equal(t, math.Round(5.5), a.Y)
}

func TestCeil(t *testing.T) {
	a := NewFloat64(-3.4, 5.5).Ceil()
	assert.Equal(t, math.Ceil(-3.4), a.X)
	assert.Equal(t, math.Ceil(5.5), a.Y)
}

func TestFloor(t *testing.T) {
	a := NewFloat64(-3.4, 5.5).Floor()
	assert.Equal(t, math.Floor(-3.4), a.X)
	assert.Equal(t, math.Floor(5.5), a.Y)
}

func TestSwap(t *testing.T) {
	a := NewFloat64(-3.4, 5.5).Swap()
	assert.Equal(t, -3.4, a.Y)
	assert.Equal(t, 5.5, a.X)
}

func TestPerpendicular(t *testing.T) {
	a := NewFloat64(-3.4, 5.5).Perpendicular()
	assert.Equal(t, 3.4, a.Y)
	assert.Equal(t, 5.5, a.X)
}

func TestWithX(t *testing.T) {
	a := NewFloat64(-3.4, 5.5).WithX(478.73)
	assert.Equal(t, 478.73, a.X)
	assert.Equal(t, 5.5, a.Y)
}

func TestWithY(t *testing.T) {
	a := NewFloat64(-3.4, 5.5).WithY(478.73)
	assert.Equal(t, -3.4, a.X)
	assert.Equal(t, 478.73, a.Y)
}

func TestComponents(t *testing.T) {
	x, y := NewFloat64(-3.4, 5.5).Components()
	assert.Equal(t, -3.4, x)
	assert.Equal(t, 5.5, y)
}
