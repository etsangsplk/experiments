package helper

func Exchange(a *int, b *int, c *int) {
	switch {
	case *a > *b:
		Swap(a, b)
	case *a > *c:
		Swap(a, c)

	}
}
