package pointer

func Of[Value any](v Value) *Value {
	return &v
}
