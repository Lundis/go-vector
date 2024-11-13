package vector2

import (
	"math"
)

func Add[T1, T2 Number](v Vector2[T1], other ...Vector2[T2]) Vector2[T1] {
	for _, o := range other {
		v.X += T1(o.X)
		v.Y += T1(o.Y)
	}
	return v
}

func Sub[T1, T2 Number](v Vector2[T1], other ...Vector2[T2]) Vector2[T1] {
	for _, o := range other {
		v.X -= T1(o.X)
		v.Y -= T1(o.Y)
	}
	return v
}

func Mul[T1, T2 Number](v Vector2[T1], other ...Vector2[T2]) Vector2[T1] {
	for _, o := range other {
		v.X *= T1(o.X)
		v.Y *= T1(o.Y)
	}
	return v
}

func Div[T1, T2 Number](v Vector2[T1], other ...Vector2[T2]) Vector2[T1] {
	for _, o := range other {
		v.X /= T1(o.X)
		v.Y /= T1(o.Y)
	}
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

func Reflect[T1, T2 Number](v Vector2[T1], other Vector2[T2]) Vector2[T1] {
	dummy := v.X
	dummy = T1(-2.0)
	factor := dummy * Dot(v, other)
	return Vector2[T1]{
		X: factor*v.X + T1(other.X),
		Y: factor*v.Y + T1(other.Y),
	}
}

func IsBetweenInclusive[T1, T2, T3 Number](p Vector2[T1], left Vector2[T2], right Vector2[T3]) bool {
	return float64(left.X) <= float64(p.X) && float64(p.X) <= float64(right.X) &&
		float64(left.Y) <= float64(p.Y) && float64(p.Y) <= float64(right.Y)
}

func Max[T1, T2 Number](v Vector2[T1], other ...Vector2[T2]) Vector2[T1] {
	for _, o := range other {
		v.X = max(v.X, T1(o.X))
		v.Y = max(v.Y, T1(o.Y))
	}
	return v
}

func Min[T1, T2 Number](v Vector2[T1], other ...Vector2[T2]) Vector2[T1] {
	for _, o := range other {
		v.X = min(v.X, T1(o.X))
		v.Y = min(v.Y, T1(o.Y))
	}
	return v
}
