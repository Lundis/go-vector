package vector2

import (
	"fmt"
	"math"
)

type SignedInteger interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	SignedInteger | Float
}

type Vector2[T Number] struct {
	X, Y T
}

func NewInt[T Number](x, y T) Vector2[int] {
	return Vector2[int]{X: int(x), Y: int(y)}
}

func NewFloat32[T Number](x, y T) Vector2[float32] {
	return Vector2[float32]{X: float32(x), Y: float32(y)}
}

func NewFloat64[T Number](x, y T) Vector2[float64] {
	return Vector2[float64]{X: float64(x), Y: float64(y)}
}

func NewFloat64Polar[T1, T2 Number](angle T1, radius T2) Vector2[float64] {
	return Vector2[float64]{
		X: float64(radius) * math.Cos(float64(angle)),
		Y: float64(radius) * math.Sin(float64(angle)),
	}
}

func Add[T1, T2 Number](v Vector2[T1], other ...Vector2[T2]) Vector2[T1] {
	for _, o := range other {
		v.X += T1(o.X)
		v.Y += T1(o.Y)
	}
	return v
}

func (v Vector2[T]) AddScalar(scalar T) Vector2[T] {
	return v.AddScalars(scalar, scalar)
}

func (v Vector2[T]) AddScalars(x, y T) Vector2[T] {
	v.X += x
	v.Y += y
	return v
}

func Sub[T1, T2 Number](v Vector2[T1], other ...Vector2[T2]) Vector2[T1] {
	for _, o := range other {
		v.X -= T1(o.X)
		v.Y -= T1(o.Y)
	}
	return v
}

func (v Vector2[T]) SubScalar(scalar T) Vector2[T] {
	return v.SubScalars(scalar, scalar)
}

func (v Vector2[T]) SubScalars(x, y T) Vector2[T] {
	v.X -= x
	v.Y -= y
	return v
}

func Mul[T1, T2 Number](v Vector2[T1], other ...Vector2[T2]) Vector2[T1] {
	for _, o := range other {
		v.X *= T1(o.X)
		v.Y *= T1(o.Y)
	}
	return v
}

func (v Vector2[T]) MulScalar(scalar T) Vector2[T] {
	return v.MulScalars(scalar, scalar)
}

func (v Vector2[T]) MulScalars(x, y T) Vector2[T] {
	v.X *= x
	v.Y *= y
	return v
}

func Div[T1, T2 Number](v Vector2[T1], other ...Vector2[T2]) Vector2[T1] {
	for _, o := range other {
		v.X /= T1(o.X)
		v.Y /= T1(o.Y)
	}
	return v
}

func (v Vector2[T]) DivScalar(scalar T) Vector2[T] {
	return v.DivScalars(scalar, scalar)
}

func (v Vector2[T]) DivScalars(x, y T) Vector2[T] {
	v.X /= x
	v.Y /= y
	return v
}

func Distance[T1, T2 Number](v Vector2[T1], other Vector2[T2]) T1 {
	dx := float64(v.X) - float64(other.X)
	dy := float64(v.Y) - float64(other.Y)
	return T1(math.Sqrt(dx*dx + dy*dy))
}

func DistanceSquared[T1, T2 Number](v Vector2[T1], other Vector2[T2]) T1 {
	dx := float64(v.X) - float64(other.X)
	dy := float64(v.Y) - float64(other.Y)
	return T1(dx*dx + dy*dy)
}

func Dot[T1, T2 Number](v Vector2[T1], other Vector2[T2]) T1 {
	return v.X*T1(other.X) + v.Y*T1(other.Y)
}

func Lerp[T1, T2, T Number](v Vector2[T1], other Vector2[T2], t T) Vector2[T1] {
	return Vector2[T1]{
		X: v.X + (T1(other.X)-v.X)*T1(t),
		Y: v.Y + (T1(other.Y)-v.Y)*T1(t),
	}
}

func (v Vector2[T]) Magnitude() T {
	return T(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v Vector2[T]) Normalize() Vector2[T] {
	m := v.Magnitude()

	if m > 0.0 {
		return v.DivScalar(m)
	} else {
		return v
	}
}

func Reflect[T1, T2 Number](v Vector2[T1], other Vector2[T2]) Vector2[T1] {
	dummy := v.X
	dummy = T1(-2.0)
	factor := dummy * Dot(v, other)
	return Vector2[T1]{
		X: factor*v.X + T1(other.X),
		Y: factor*v.Y + T1(other.Y),
	}
}

func (v Vector2[T]) Equals(other Vector2[T]) bool {
	return v.X == other.X && v.Y == other.Y
}

func (v Vector2[T]) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v Vector2[T]) String() string {
	return fmt.Sprintf("Vector2(%v, %v)", v.X, v.Y)
}

func (v Vector2[T]) AsFloat32() Vector2[float32] {
	return Vector2[float32]{X: float32(v.X), Y: float32(v.Y)}
}

func (v Vector2[T]) AsFloat64() Vector2[float64] {
	return Vector2[float64]{X: float64(v.X), Y: float64(v.Y)}
}
