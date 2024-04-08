package spectra

import "encoding/json"

const (
	Allow Effect = "allow"
	Deny  Effect = "deny"
)

type Effect string

func NewPolicy(expression Expression, effect Effect, permissions []string, description string) *Policy {
	return &Policy{
		expression:  expression,
		effect:      effect,
		permissions: permissions,
		description: description,
	}
}

type Policy struct {
	expression  Expression
	effect      Effect
	permissions []string
	description string
}

func (p *Policy) GetDescription() string {
	return p.description
}

func (p *Policy) GetFields() []FieldName {
	return p.GetFilter().GetFields()
}

func (p *Policy) GetPermissions() []string {
	return p.permissions
}

func (p *Policy) GetEffect() Effect {
	return p.effect
}

func (p *Policy) GetFilter() Expression {
	return p.expression
}

func (p *Policy) Apply(data Data) bool {
	return p.GetFilter().Evaluate(data)
}

func (p *Policy) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Description string     `json:"description"`
		Effect      Effect     `json:"effect"`
		Expression  Expression `json:"filter"`
		Permissions []string   `json:"permissions"`
	}{
		Description: p.description,
		Effect:      p.effect,
		Expression:  p.expression,
		Permissions: p.permissions,
	})
}
