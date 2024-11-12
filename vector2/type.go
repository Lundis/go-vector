package vector2

import (
	"fmt"
	"math"
	"math/rand/v2"
)

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

func NewFloat64Random[T Number](minValue, maxValue T) Vector2[float64] {
	spread := float64(maxValue - minValue)
	return NewFloat64(
		float64(minValue)+rand.Float64()*spread,
		float64(minValue)+rand.Float64()*spread,
	)
}

func (v Vector2[T]) AsFloat32() Vector2[float32] {
	return Vector2[float32]{X: float32(v.X), Y: float32(v.Y)}
}

func (v Vector2[T]) AsFloat64() Vector2[float64] {
	return Vector2[float64]{X: float64(v.X), Y: float64(v.Y)}
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
