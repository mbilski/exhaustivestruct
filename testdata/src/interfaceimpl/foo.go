package interfaceimpl

type MyWriter interface {
	Write()
}

type SuperWriter struct {
	A string
	B string
}

func (w SuperWriter) Write() {
	// nop
}

var _ MyWriter = SuperWriter{}
var _ MyWriter = &SuperWriter{}
var _ MyWriter = (*SuperWriter)(nil)

var _ MyWriter = SuperWriter{ // want "B is missing in SuperWriter"
	A: "x",
}

var global1 MyWriter = SuperWriter{}  // want "A, B are missing in SuperWriter"
var global2 MyWriter = &SuperWriter{} // want "A, B are missing in SuperWriter"
