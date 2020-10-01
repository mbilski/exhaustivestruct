package s

type Test struct {
	A string
	B int
	C float32
	D bool
}

func shouldPass() Test {
	return Test{
		A: "a",
		B: 1,
		C: 0.0,
		D: false,
	}
}

func shouldFailWithMissingFields() Test {
	return Test{ // want "missing fields"
		A: "a",
		B: 1,
	}
}
