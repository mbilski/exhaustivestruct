package s

import "e"

type Embedded struct {
	E string
	F string
	g string
	H string
}

type Test struct {
	A string
	B int
	C float32
	D bool
}

type Test2 struct {
	Embedded
	External e.External
}

func shouldPass() Test {
	return Test{
		A: "a",
		B: 1,
		C: 0.0,
		D: false,
	}
}

func shouldPass2() Test2 {
	return Test2{
		External: e.External{
			A: "",
			B: "",
		},
		Embedded: Embedded{
			E: "",
			F: "",
			H: "",
		},
	}
}

func shouldFailWithMissingFields() Test {
	return Test{ // want "C is missing in Test"
		A: "a",
		B: 1,
		D: false,
	}
}

// Unchecked is a struct not listed in StructPatternList
type Unchecked struct {
	A int
	B int
}

func unchecked() {
	// This struct is not listed in StructPatternList so the linter won't complain that it's not filled out
	_ = Unchecked{
		A: 1,
	}
}

func shouldFailOnEmbedded() Test2 {
	return Test2{
		Embedded: Embedded{ // want "E, H are missing in Embedded"
			F: "",
		},
		External: e.External{
			A: "",
			B: "",
		},
	}
}

func shoildFailOnExternal() Test2 {
	return Test2{
		External: e.External{ // want "A is missing in External"
			B: "",
		},
		Embedded: Embedded{
			E: "",
			F: "",
			H: "",
		},
	}
}
