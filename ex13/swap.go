package main

func Swap(a *int, b *int) {
	swap := *a
	*a = *b
	*b = swap
}
