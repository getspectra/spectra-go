package spectra

import (
	"errors"
	"reflect"
)

type Operation string

const (
	eq    Operation = "="
	neq   Operation = "!="
	neq2  Operation = "<>"
	gt    Operation = ">"
	gte   Operation = ">="
	lt    Operation = "<"
	lte   Operation = "<="
	in    Operation = "in"
	nin   Operation = "nin"
	notIn Operation = "not_in"
)

var operations = map[string]Operation{
	"=":      eq,
	"!=":     neq,
	"<>":     neq2,
	">":      gt,
	">=":     gte,
	"<":      lt,
	"<=":     lte,
	"in":     in,
	"nin":    nin,
	"not_in": notIn,
}

func toOperation(key string) (Operation, error) {
	if operation, ok := operations[key]; ok {
		return operation, nil
	} else {
		return "", errors.New("operation not found")
	}
}

func isValidOperation(key string) bool {
	_, ok := operations[key]
	return ok
}

func (op Operation) operate(left, right any) bool {
	switch op {
	case eq:
		return left == right
	case neq, neq2:
		return left != right
	case gt, gte, lt, lte:
		return isSameType(left, right) && compare(left, right, op)
	case in:
		return isSameTypeAsSliceElement(left, right) && contains(left, right)
	case nin, notIn:
		return isSameTypeAsSliceElement(left, right) && !contains(left, right)
	}

	return false
}

func compare(left, right any, op Operation) bool {
	switch left.(type) {
	case int:
		return compareNumber(left.(int), right.(int), op)
	case int8:
		return compareNumber(left.(int8), right.(int8), op)
	case int16:
		return compareNumber(left.(int16), right.(int16), op)
	case int32:
		return compareNumber(left.(int32), right.(int32), op)
	case int64:
		return compareNumber(left.(int64), right.(int64), op)
	case uint:
		return compareNumber(left.(uint), right.(uint), op)
	case uint8:
		return compareNumber(left.(uint8), right.(uint8), op)
	case uint16:
		return compareNumber(left.(uint16), right.(uint16), op)
	case uint32:
		return compareNumber(left.(uint32), right.(uint32), op)
	case uint64:
		return compareNumber(left.(uint64), right.(uint64), op)
	case float32:
		return compareNumber(left.(float32), right.(float32), op)
	case float64:
		return compareNumber(left.(float64), right.(float64), op)
	default:
		return false
	}
}

func compareNumber[T Number](left, right T, op Operation) bool {
	switch op {
	case gt:
		return left > right
	case gte:
		return left >= right
	case lt:
		return left < right
	case lte:
		return left <= right
	}
	return false
}

func contains(left, right any) bool {
	rightValue := reflect.ValueOf(right)

	for i := 0; i < rightValue.Len(); i++ {
		if left == rightValue.Index(i).Interface() {
			return true
		}
	}

	return false
}

func isSameType(left, right any) bool {
	return reflect.TypeOf(left) == reflect.TypeOf(right)
}

func isSameTypeAsSliceElement(left any, right any) bool {
	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)
	return rightType.Kind() == reflect.Slice && leftType == rightType.Elem()
}
