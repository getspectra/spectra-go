package spectra

func And(expressions []Expression) Expression {
	return &AndExpression{expressions: expressions}
}

func Or(expressions []Expression) Expression {
	return &OrExpression{expressions: expressions}
}

func Not(expression Expression) Expression {
	return &NotExpression{expression: expression}
}

func Eq(left FieldName, right FieldValue) Expression {
	return &BinaryExpression{
		left:      left,
		operation: eq,
		right:     right,
	}
}

func Neq(left FieldName, right FieldValue) Expression {
	return &BinaryExpression{
		left:      left,
		operation: neq,
		right:     right,
	}
}

func Neq2(left FieldName, right FieldValue) Expression {
	return &BinaryExpression{
		left:      left,
		operation: neq2,
		right:     right,
	}
}

func Gt(left FieldName, right FieldValue) Expression {
	return &BinaryExpression{
		left:      left,
		operation: gt,
		right:     right,
	}
}

func Gte(left FieldName, right FieldValue) Expression {
	return &BinaryExpression{
		left:      left,
		operation: gte,
		right:     right,
	}
}

func Lt(left FieldName, right FieldValue) Expression {
	return &BinaryExpression{
		left:      left,
		operation: lt,
		right:     right,
	}
}

func Lte(left FieldName, right FieldValue) Expression {
	return &BinaryExpression{
		left:      left,
		operation: lte,
		right:     right,
	}
}

func In(left FieldName, right FieldValue) Expression {
	return &BinaryExpression{
		left:      left,
		operation: in,
		right:     right,
	}
}

func Nin(left FieldName, right FieldValue) Expression {
	return &BinaryExpression{
		left:      left,
		operation: nin,
		right:     right,
	}
}

func NotIn(left FieldName, right FieldValue) Expression {
	return &BinaryExpression{
		left:      left,
		operation: notIn,
		right:     right,
	}
}

func NormalizeExpression(definition any) Expression {
	switch data := definition.(type) {
	case []any:
		operation, err := toOperation(data[1].(string))
		if err != nil {
			return nil
		}
		return &BinaryExpression{
			left:      FieldName(data[0].(string)),
			operation: operation,
			right:     data[2],
		}
	case map[string]any:
		for _, operator := range []string{"and", "or", "not"} {
			if expressions, ok := data[operator]; ok {
				switch operator {
				case "and", "or":
					normalizedExpressions := make([]Expression, len(expressions.([]any)))
					for i, expression := range expressions.([]any) {
						normalizedExpressions[i] = NormalizeExpression(expression)
					}
					if operator == "and" {
						return And(normalizedExpressions)
					}
					return Or(normalizedExpressions)
				case "not":
					return Not(NormalizeExpression(expressions))
				}
			}
		}
	}
	return nil
}
