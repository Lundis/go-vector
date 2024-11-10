package vector2

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestNewInt(t *testing.T) {
	a := NewInt(2, 4)

	assert.Equal(t, a.X, 2)
	assert.Equal(t, a.Y, 4)
}

func TestNewFloat32(t *testing.T) {
	a := NewFloat32(2, 4)

	assert.Equal(t, a.X, float32(2.0))
	assert.Equal(t, a.Y, float32(4.0))
}

func TestNewFloat64(t *testing.T) {
	a := NewFloat64(2, 4)

	assert.Equal(t, a.X, 2.0)
	assert.Equal(t, a.Y, 4.0)
}
func TestNewFloat64Polar(t *testing.T) {
	a := NewFloat64Polar(0, 1000) // right

	assert.Equal(t, 1000.0, math.Round(a.X))
	assert.Equal(t, 0.0, math.Round(a.Y))

	b := NewFloat64Polar(0.25*2.0*math.Pi, 1000) // up
	assert.Equal(t, 0.0, math.Round(b.X))
	assert.Equal(t, 1000.0, math.Round(b.Y))
}

func TestAdd(t *testing.T) {
	a := NewFloat64(1, 2)
	b := NewFloat64(4, 3)
	res := Add(a, b)

	assert.Equal(t, res.X, 5.0)
	assert.Equal(t, res.Y, 5.0)
}

func TestVector2_AddScalar(t *testing.T) {
	a := NewFloat64(1, 2)
	res := a.AddScalar(1.0)

	assert.Equal(t, res.X, 2.0)
	assert.Equal(t, res.Y, 3.0)
}

func TestVector2_AddScalars(t *testing.T) {
	a := NewFloat64(1, 2)
	res := a.AddScalars(1, 2)

	assert.Equal(t, res.X, 2.0)
	assert.Equal(t, res.Y, 4.0)
}

func TestVector2_Sub(t *testing.T) {
	a := NewFloat64(1, 2)
	b := NewFloat64(4, 3)
	res := Sub(a, b)

	assert.Equal(t, res.X, -3.0)
	assert.Equal(t, res.Y, -1.0)
}

func TestVector2_SubScalar(t *testing.T) {
	a := NewFloat64(1, 2)
	res := a.SubScalar(1.0)

	assert.Equal(t, res.X, 0.0)
	assert.Equal(t, res.Y, 1.0)
}

func TestVector2_SubScalars(t *testing.T) {
	a := NewFloat64(1, 2)
	res := a.SubScalars(1, 2)

	assert.Equal(t, res.X, 0.0)
	assert.Equal(t, res.Y, 0.0)
}

func TestMul(t *testing.T) {
	a := NewFloat64(1, 2)
	b := NewFloat64(4, 3)
	res := Mul(a, b)

	assert.Equal(t, res.X, 4.0)
	assert.Equal(t, res.Y, 6.0)
}

func TestVector2_MulScalar(t *testing.T) {
	a := NewFloat64(1, 2)
	b := 4.0
	res := a.MulScalar(b)

	assert.Equal(t, res.X, 4.0)
	assert.Equal(t, res.Y, 8.0)
}

func TestVector2_MulScalars(t *testing.T) {
	a := NewFloat64(1, 2)
	res := a.MulScalars(1.0, 2.0)

	assert.Equal(t, res.X, 1.0)
	assert.Equal(t, res.Y, 4.0)
}

func TestVector2_Div(t *testing.T) {
	a := NewFloat64(1, 2)
	b := NewFloat64(4, 3)
	res := Div(a, b)

	assert.Equal(t, res.X, 0.25)
	assert.Equal(t, res.Y, 0.6666666666666666)
}

func TestVector2_DivScalar(t *testing.T) {
	a := NewFloat64(1, 2)
	b := 4.0
	res := a.DivScalar(b)

	assert.Equal(t, res.X, 0.25)
	assert.Equal(t, res.Y, 0.5)
}

func TestVector2_DivScalars(t *testing.T) {
	a := NewFloat64(1, 2)
	res := a.DivScalars(1.0, 2.0)

	assert.Equal(t, res.X, 1.0)
	assert.Equal(t, res.Y, 1.0)
}

func TestVector2_Distance(t *testing.T) {
	a := NewFloat64(1, 2)
	b := NewFloat64(4, 3)
	res := Distance(a, b)

	assert.Equal(t, res, 3.1622776601683795)
}

func TestVector2_Magnitude(t *testing.T) {
	a := NewFloat64(3, 2)
	res := a.Magnitude()

	assert.Equal(t, res, 3.605551275463989)
}

func TestVector2_Normalize(t *testing.T) {
	a := NewFloat64(3, 2)
	b := NewFloat64(0.000003, 0.000002)
	aa := a.Normalize()
	bb := b.Normalize()

	assert.True(t, aa.Equals(bb))
}

func TestVector2_NormalizeZero(t *testing.T) {
	a := NewFloat64(0, 0)
	aa := a.Normalize()

	assert.True(t, aa.Equals(a))
}

func TestVector2_Equals(t *testing.T) {
	a := NewFloat64(5, 3)
	b := NewFloat64(4, 3)

	assert.False(t, a.Equals(b))
}

func TestVector2_String(t *testing.T) {
	a := NewFloat64(3, 6)
	assert.Equal(t, "Vector2(3, 6)", a.String())
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

	assert.Equal(t, res, 3*3+1*1)
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

func TestIsZero(t *testing.T) {
	assert.Equal(t, NewFloat64(0, 0).IsZero(), true)
	assert.Equal(t, NewFloat64(0.000001, 0.000001).IsZero(), false)
	assert.Equal(t, NewFloat64(1, 0).IsZero(), false)
	assert.Equal(t, NewFloat64(0, 1).IsZero(), false)
}

func TestAsFloat32(t *testing.T) {
	a := NewFloat64(1, 2)
	res := a.AsFloat32()

	assert.Equal(t, res.X, float32(1.0))
	assert.Equal(t, res.Y, float32(2.0))
}

func TestAsFloat64(t *testing.T) {
	a := NewFloat32(1, 2)
	res := a.AsFloat64()

	assert.Equal(t, res.X, 1.0)
	assert.Equal(t, res.Y, 2.0)
}
