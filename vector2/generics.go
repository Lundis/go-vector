package vector2

type SignedInteger interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	SignedInteger | Float
}
