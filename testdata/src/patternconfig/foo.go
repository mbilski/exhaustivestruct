package patternconfig

// Unchecked is a struct not listed in StructPatternList
type Unchecked struct {
	A string
}

type Checked struct {
	A string
}

type AnotherChecked struct {
	A string
}

func foo() {
	_ = Unchecked{}
	_ = Checked{}        // want "A is missing in Checked"
	_ = AnotherChecked{} // want "A is missing in AnotherChecked"
}
