package main

func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
	var result = initialValue

	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}

func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
	for _, item := range items {
		if predicate(item) {
			return item, true
		}
	}
	return
}
