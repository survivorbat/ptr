package ptr

// Ptr returns a pointer for the given object
func Ptr[T any](in T) *T {
	return &in
}
