package spectra

type DataLoader interface {
	Load(fields []FieldName) (Data, error)
}

type DataLoaderFunc func(fields []FieldName) (Data, error)

func (f DataLoaderFunc) Load(fields []FieldName) (Data, error) {
	return f(fields)
}
