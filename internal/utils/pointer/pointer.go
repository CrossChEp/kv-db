package pointer

func Deref[T any](in *T) T {
	if in == nil {
		var out T

		return out
	}

	return *in
}
