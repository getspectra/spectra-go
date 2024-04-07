package spectra

func isValidExpressionTuple(definition []any) bool {
	return len(definition) == 3 && isString(definition[0]) && isString(definition[1]) && isValidOperation(definition[1].(string))
}

func isString(value any) bool {
	_, ok := value.(string)
	return ok
}
