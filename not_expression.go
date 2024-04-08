package spectra

import "encoding/json"

type NotExpression struct {
	expression Expression
}

func (n *NotExpression) GetFields() []FieldName {
	return n.expression.GetFields()
}

func (n *NotExpression) Evaluate(data Data) bool {
	return !n.expression.Evaluate(data)
}

func (n *NotExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Not Expression `json:"not"`
	}{
		Not: n.expression,
	})
}
