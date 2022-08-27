package font

//go:generate enumer -type=Type -output=type.gen.go

type Type int

const (
	MonoSpaced Type = iota + 1
	Propotional
)
