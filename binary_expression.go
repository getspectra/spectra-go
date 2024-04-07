package spectra

import (
	"time"
)

type FieldName string

// FieldValue is a type in Value or RefValue
type FieldValue any
type Value interface {
	string | bool | Number | time.Time | []any
}
type RefValue struct {
	ref FieldName
}

type BinaryExpression struct {
	left      FieldName
	operation Operation
	right     FieldValue
}

func (b *BinaryExpression) GetFields() []FieldName {
	fields := make([]FieldName, 1)
	fields[0] = b.left
	if rightRef := b.resolveRightRef(); rightRef != "" {
		fields = append(fields, rightRef)
	}
	return fields
}

func (b *BinaryExpression) resolveRightRef() FieldName {
	switch v := b.right.(type) {
	case *RefValue:
		return v.ref
	case RefValue:
		return v.ref
	default:
		return ""
	}
}

func (b *BinaryExpression) Evaluate(data Data) bool {
	leftValue := data[b.left]
	rightValue := b.resolveRightValue(data)
	return b.operation.operate(leftValue, rightValue)
}

func (b *BinaryExpression) resolveRightValue(data Data) any {
	if rightRef := b.resolveRightRef(); rightRef != "" {
		return data[rightRef]
	}
	return b.right
}

func (b *BinaryExpression) JsonSerialize() string {
	//TODO implement me
	panic("implement me")
}
