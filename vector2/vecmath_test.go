package vector2

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

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

func TestDistanceToLine(t *testing.T) {
	/*
				 |           p2
		         |
				 b
		p3       |
				 a     p

	*/
	a := NewFloat64(0, 0)
	b := NewFloat64(0, 1)

	p := NewFloat64(1, 0)
	assert.Equal(t, 1.0, DistanceToLine(a, b, p))

	p2 := NewFloat64(2, 2)
	assert.Equal(t, 2.0, DistanceToLine(a, b, p2))

	p3 := NewFloat64(-0.5, 0.5)
	assert.Equal(t, 0.5, DistanceToLine(a, b, p3))
}

func TestSideOfLine(t *testing.T) {
	/*
				 |           p2
		         |
				 b
		p3       |
				 a     p
	*/
	a := NewFloat64(0, 0)
	b := NewFloat64(0, 1)

	p := NewFloat64(1, 0)
	assert.Less(t, SideOfPoint(a, b, p), 0.0)

	p2 := NewFloat64(2, 2)
	assert.Less(t, SideOfPoint(a, b, p2), 0.0)

	p3 := NewFloat64(-0.5, 0.5)
	assert.Greater(t, SideOfPoint(a, b, p3), 0.0)
}

func TestAngleBetween(t *testing.T) {
	up := NewFloat64(0, 1)
	upRight := NewFloat64(1, 1)
	right := NewFloat64(1, 0)
	down := NewFloat64(0, -1)
	downLeft := NewFloat64(-1, -1)

	assert.Equal(t, math.Pi/2, AngleBetween(right, up))
	assert.Equal(t, math.Pi/4, AngleBetween(right, upRight))
	assert.Equal(t, -math.Pi/2, AngleBetween(up, right))
	assert.Equal(t, math.Pi, AngleBetween(up, down))
	assert.Equal(t, math.Pi*3/4, AngleBetween(up, downLeft))
	assert.Equal(t, -math.Pi*3/4, AngleBetween(downLeft, up))

}
