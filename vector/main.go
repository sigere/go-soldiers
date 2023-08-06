package vector

type T [2]int

func (v T) Multiply(value int) T {
	return T{v[0] * value, v[1] * value}
}
