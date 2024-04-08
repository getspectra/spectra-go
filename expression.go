package spectra

type Data map[FieldName]any

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// Definition is a type in map[string]any or []any
type Definition any

type Expression interface {
	Evaluate(data Data) bool
	GetFields() []FieldName
	MarshalJSON() ([]byte, error)
}
