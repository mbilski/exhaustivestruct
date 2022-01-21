package basic

import (
	"fmt"

	"e"
)

type EmbeddedStruct struct {
	E string
	F string
	g string
	H string
}

type FlatStruct struct {
	A string
	B int
	C float32
	D bool
}

type NestStruct struct {
	EmbeddedStruct
	External e.External
}

type DefStruct FlatStruct

type DeepDefStruct DefStruct

type AliasStruct = FlatStruct

func returnFlatNoMissingFields() FlatStruct {
	return FlatStruct{
		A: "a",
		B: 1,
		C: 0.0,
		D: false,
	}
}

func returnDefNoMissingFields() DefStruct {
	return DefStruct{
		A: "a",
		B: 1,
		C: 0.0,
		D: false,
	}
}

func returnDeepDefNoMissingFields() DeepDefStruct {
	return DeepDefStruct{
		A: "a",
		B: 1,
		C: 0.0,
		D: false,
	}
}

func returnAliasNoMissingFields() AliasStruct {
	return AliasStruct{
		A: "a",
		B: 1,
		C: 0.0,
		D: false,
	}
}

func returnNestNoMissingFields() NestStruct {
	return NestStruct{
		External: e.External{
			A: "",
			B: "",
		},
		EmbeddedStruct: EmbeddedStruct{
			E: "",
			F: "",
			H: "",
			g: "",
		},
	}
}

// Empty structs in return statements are ignored if also returning an error
func returnEmptyFlatWithError() (FlatStruct, error) {
	return FlatStruct{}, fmt.Errorf("error")
}

func returnEmptyNestWithError2() (NestStruct, error) {
	return NestStruct{}, fmt.Errorf("error")
}

func assignEmptyStruct() {
	_ = FlatStruct{} // want "A, B, C, D are missing in FlatStruct"
}

func returnEmptyStruct() FlatStruct {
	return FlatStruct{} // want "A, B, C, D are missing in FlatStruct"
}

func returnEmptyDefStruct() DefStruct {
	return DefStruct{} // want "A, B, C, D are missing in DefStruct"
}

func returnEmptyDeepDefStruct() DeepDefStruct {
	return DeepDefStruct{} // want "A, B, C, D are missing in DeepDefStruct"
}

func returnEmptyAliasStruct() AliasStruct {
	return AliasStruct{} // want "A, B, C, D are missing in AliasStruct"
}

func returnNoNames() FlatStruct {
	return FlatStruct{"", 0, 0, false}
}

// Empty structs in return statements are not ignored if returning nil error
func returnEmptyFlatWithNilError() (FlatStruct, error) {
	return FlatStruct{}, nil // want "A, B, C, D are missing in FlatStruct"
}

func returnEmptyNestWithNilError() (NestStruct, error) {
	return NestStruct{}, nil // want "EmbeddedStruct, External are missing in NestStruct"
}

// Empty structs as an inner field in the nest struct are not ignored even if there is a return error statement.
func returnEmptyInnerWithError() (NestStruct, error) {
	return NestStruct{
		EmbeddedStruct: EmbeddedStruct{}, // want "E, F, g, H are missing in EmbeddedStruct"
		External: e.External{
			A: "",
			B: "",
		},
	}, fmt.Errorf("error")
}

func returnFlatMissingFields() FlatStruct {
	return FlatStruct{ // want "C is missing in FlatStruct"
		A: "a",
		B: 1,
		D: false,
	}
}

func returnNestMissingFields() NestStruct {
	return NestStruct{ // want "External is missing in NestStruct"
		EmbeddedStruct: EmbeddedStruct{
			E: "",
			F: "",
			H: "",
			g: "",
		},
	}
}

func returnMissingFieldsInEmbedded() NestStruct {
	return NestStruct{
		EmbeddedStruct: EmbeddedStruct{ // want "E, g, H are missing in EmbeddedStruct"
			F: "",
		},
		External: e.External{
			A: "",
			B: "",
		},
	}
}

func returnMissingFieldsInExternal() NestStruct {
	return NestStruct{
		External: e.External{ // want "A is missing in External"
			B: "",
		},
		EmbeddedStruct: EmbeddedStruct{
			E: "",
			F: "",
			H: "",
			g: "",
		},
	}
}

func returnMissingFieldsInClosure() (func() FlatStruct, error) {
	return func() FlatStruct {
		return FlatStruct{} // want "A, B, C, D are missing in FlatStruct"
	}, fmt.Errorf("x")
}

func returnMissingFieldsInClosureWithError() (func() (FlatStruct, error), error) {
	return func() (FlatStruct, error) {
		return FlatStruct{}, fmt.Errorf("y")
	}, fmt.Errorf("x")
}

func returnMissingFieldsInClosureWithNilError() (func() (FlatStruct, error), error) {
	return func() (FlatStruct, error) {
		return FlatStruct{}, nil // want "A, B, C, D are missing in FlatStruct"
	}, fmt.Errorf("x")
}
