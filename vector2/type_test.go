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

func TestNewFloat64Random(t *testing.T) {
	for i := 0.0; i < 100; i++ {
		a := NewFloat64Random(-1, i)
		assert.True(t, -1 <= a.X && a.X <= i)
		assert.True(t, -1 <= a.Y && a.Y <= i)
	}
}

func TestSelf(t *testing.T) {
	a := NewFloat64(-1, -2)
	res := a.Self()
	assert.Equal(t, a, res)
}

func TestEquals(t *testing.T) {
	a := NewFloat64(5, 3)
	b := NewFloat64(4, 3)

	assert.False(t, a.Equals(b))
}

func TestString(t *testing.T) {
	a := NewFloat64(3, 6)
	assert.Equal(t, "Vector2(3, 6)", a.String())
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

type MyAlias = Vector2[float64]

func TestAlias(t *testing.T) {
	a := MyAlias{
		X: 1,
		Y: 2,
	}
	b := NewFloat64(10, 20)
	c := Add(a, b)
	assert.Equal(t, 11.0, c.X)
	assert.Equal(t, 22.0, c.Y)
}
