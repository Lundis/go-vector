package vector2

import "math"

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

func Cross[T1, T2 Number](v Vector2[T1], other Vector2[T2]) T1 {
	return v.X*T1(other.Y) - v.Y*T1(other.X)
}

func Dot[T1, T2 Number](v Vector2[T1], other Vector2[T2]) T1 {
	return v.X*T1(other.X) + v.Y*T1(other.Y)
}

func Reflect[T1, T2 Number](v Vector2[T1], other Vector2[T2]) Vector2[T1] {
	factor := T1(-2.0) * Dot(v, other)
	return Vector2[T1]{
		X: factor*v.X + T1(other.X),
		Y: factor*v.Y + T1(other.Y),
	}
}

// DistanceToLine calculates the distance from a point P to a line defined by two points A->B.
func DistanceToLine[T Number](a, b, p Vector2[T]) T {
	ab := b.Minus(a)
	ap := p.Minus(a)

	cross := Cross(ab, ap)

	if cross < 0 {
		cross = -cross
	}

	return cross / ab.Magnitude()
}

// SideOfPoint calculates which side of the line A->B the point P lies on. Check the sign of the response.
func SideOfPoint[T Number](a, b, p Vector2[T]) T {
	ab := b.Minus(a)
	ap := p.Minus(a)

	return Cross(ab, ap)
}

// AngleBetween calculates the angle between two vectors, returning a value in the range [-Pi, Pi].
func AngleBetween[T Number](v1, v2 Vector2[T]) float64 {
	angle := v2.Angle() - v1.Angle()
	if angle > math.Pi {
		angle -= 2 * math.Pi
	} else if angle <= -math.Pi {
		angle += 2 * math.Pi
	}
	return angle
}
