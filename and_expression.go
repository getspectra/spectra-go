package spectra

type AndExpression struct {
	expressions []Expression
}

func (a AndExpression) GetFields() []FieldName {
	fields := make([]FieldName, 0)
	for _, expression := range a.expressions {
		fields = append(fields, expression.GetFields()...)
	}
	return fields
}

func (a AndExpression) Evaluate(data Data) bool {
	for _, expression := range a.expressions {
		if !expression.Evaluate(data) {
			return false
		}
	}

	return true
}

func (a AndExpression) JsonSerialize() string {
	//TODO implement me
	panic("implement me")
}
