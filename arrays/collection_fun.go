package main

func Reduce[A any](collection []A, f func(A, A) A, initialValue A) A {
	var result = initialValue

	for _, x := range collection {
		result = f(result, x)
	}
	return result
}
