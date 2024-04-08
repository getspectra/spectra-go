package spectra

import "encoding/json"

type AndExpression struct {
	expressions []Expression
}

func (a *AndExpression) GetFields() []FieldName {
	fields := make(map[FieldName]bool)
	for _, expression := range a.expressions {
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

func (a *AndExpression) Evaluate(data Data) bool {
	for _, expression := range a.expressions {
		if !expression.Evaluate(data) {
			return false
		}
	}

	return true
}

func (a *AndExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		And []Expression `json:"and"`
	}{
		And: a.expressions,
	})
}
