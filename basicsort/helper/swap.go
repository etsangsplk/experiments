package helper

// helper swap function.
func Swap(a *int, b *int) {
	*a, *b = *b, *a
}
