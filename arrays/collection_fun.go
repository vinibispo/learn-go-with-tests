package main

func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
	var result = initialValue

	for _, x := range collection {
		result = f(result, x)
	}
	return result
}

func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
	for _, v := range items {
		if predicate(v) {
			return v, true
		}
	}
	return
}
