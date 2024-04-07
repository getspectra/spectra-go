package spectra

import (
	"errors"
	"reflect"
)

func Validate(policies []*Policy, dataLoader any, permission string) bool {
	relevantPolicies := getRelevantPolicies(policies, permission)
	fields := getRequiredFieldsFromPolicies(relevantPolicies)
	data, err := loadAllNecessaryData(dataLoader, fields)
	if err != nil {
		return false
	}
	denyPolicies, allowPolicies := bisectPoliciesIntoDenyAndAllowPolicies(relevantPolicies)
	for _, policy := range denyPolicies {
		if policy.Apply(data) {
			return false
		}
	}
	for _, policy := range allowPolicies {
		if policy.Apply(data) {
			return true
		}
	}
	return false
}

func getRelevantPolicies(policies []*Policy, permission string) []*Policy {
	var relevantPolicies []*Policy
	for _, policy := range policies {
		for _, policyPermission := range policy.permissions {
			if policyPermission == permission {
				relevantPolicies = append(relevantPolicies, policy)
				break
			}
		}
	}
	return relevantPolicies
}

func getRequiredFieldsFromPolicies(policies []*Policy) []FieldName {
	var fields []FieldName
	for _, policy := range policies {
		for _, field := range policy.GetFields() {
			if !contains(fields, field) {
				fields = append(fields, field)
			}
		}
	}
	return fields
}

func loadAllNecessaryData(dataLoader any, fieldsToLoad []FieldName) (Data, error) {
	var data Data
	var err error

	switch loader := dataLoader.(type) {
	case DataLoader:
		data, err = loader.Load(fieldsToLoad)
	case DataLoaderFunc:
		data, err = loader(fieldsToLoad)
	default:
		return nil, errors.New("unsupported dataLoader type")
	}

	if err != nil {
		return nil, err
	}

	if data == nil || reflect.ValueOf(data).Kind() != reflect.Map {
		return nil, errors.New("expected data to be a map")
	}

	return data, nil
}

func bisectPoliciesIntoDenyAndAllowPolicies(policies []*Policy) ([]*Policy, []*Policy) {
	var denyPolicies []*Policy
	var allowPolicies []*Policy

	for _, policy := range policies {
		if policy.effect == Deny {
			denyPolicies = append(denyPolicies, policy)
		} else {
			allowPolicies = append(allowPolicies, policy)
		}
	}

	return denyPolicies, allowPolicies
}
