package spectra

type OrExpression struct {
	expressions []Expression
}

func (o *OrExpression) GetFields() []FieldName {
	fields := make([]FieldName, 0)
	for _, expression := range o.expressions {
		fields = append(fields, expression.GetFields()...)
	}
	return fields
}

func (o *OrExpression) Evaluate(data Data) bool {
	for _, expression := range o.expressions {
		if expression.Evaluate(data) {
			return true
		}
	}

	return false
}

func (o *OrExpression) JsonSerialize() string {
	//TODO implement me
	panic("implement me")
}
