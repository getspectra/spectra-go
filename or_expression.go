package spectra

import "encoding/json"

type OrExpression struct {
	expressions []Expression
}

func (o *OrExpression) GetFields() []FieldName {
	fields := make(map[FieldName]bool)
	for _, expression := range o.expressions {
		for _, field := range expression.GetFields() {
			fields[field] = true
		}
	}
	output := make([]FieldName, 0, len(fields))
	for field := range fields {
		output = append(output, field)
	}
	return output
}

func (o *OrExpression) Evaluate(data Data) bool {
	for _, expression := range o.expressions {
		if expression.Evaluate(data) {
			return true
		}
	}

	return false
}

func (o *OrExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Or []Expression `json:"or"`
	}{
		Or: o.expressions,
	})
}
