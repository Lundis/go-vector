package vector2

import (
	"math"
)

// Self is useful if you want to build complex stuff with interfaces
func (v Vector2[T]) Self() Vector2[T] {
	return v
}

func (v Vector2[T]) Plus(other Vector2[T]) Vector2[T] {
	v.X += other.X
	v.Y += other.Y
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

func (v Vector2[T]) Minus(other Vector2[T]) Vector2[T] {
	v.X -= other.X
	v.Y -= other.Y
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

func (v Vector2[T]) Mul(other Vector2[T]) Vector2[T] {
	v.X *= other.X
	v.Y *= other.Y
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

func (v Vector2[T]) Div(other Vector2[T]) Vector2[T] {
	v.X /= other.X
	v.Y /= other.Y
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

func (v Vector2[T]) InvertY() Vector2[T] {
	v.Y = -v.Y
	return v
}

func (v Vector2[T]) Angle() float64 {
	return math.Atan2(float64(v.Y), float64(v.X))
}

func (v Vector2[T]) Abs() Vector2[T] {
	v.X = T(math.Abs(float64(v.X)))
	v.Y = T(math.Abs(float64(v.Y)))
	return v
}

func (v Vector2[T]) Clamp(low, high Vector2[T]) Vector2[T] {
	return Max(low, Min(v, high))
}

func (v Vector2[T]) Round() Vector2[T] {
	v.X = T(math.Round(float64(v.X)))
	v.Y = T(math.Round(float64(v.Y)))
	return v
}

func (v Vector2[T]) Floor() Vector2[T] {
	v.X = T(math.Floor(float64(v.X)))
	v.Y = T(math.Floor(float64(v.Y)))
	return v
}

func (v Vector2[T]) Ceil() Vector2[T] {
	v.X = T(math.Ceil(float64(v.X)))
	v.Y = T(math.Ceil(float64(v.Y)))
	return v
}

func (v Vector2[T]) Swap() Vector2[T] {
	v.X, v.Y = v.Y, v.X
	return v
}

func (v Vector2[T]) Perpendicular() Vector2[T] {
	v.X, v.Y = v.Y, -v.X
	return v
}

func (v Vector2[T]) WithX(value T) Vector2[T] {
	v.X = value
	return v
}

func (v Vector2[T]) WithY(value T) Vector2[T] {
	v.Y = value
	return v
}
