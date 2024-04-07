package spectra

type NotExpression struct {
	expression Expression
}

func (n *NotExpression) GetFields() []FieldName {
	return n.expression.GetFields()
}

func (n *NotExpression) Evaluate(data Data) bool {
	return !n.expression.Evaluate(data)
}

func (n *NotExpression) JsonSerialize() string {
	//TODO implement me
	panic("implement me")
}
