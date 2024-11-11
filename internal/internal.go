package internal

import (
	"encoding/json"
)

type Builder struct {
	Context    []string            `json:"context"`
	Variables  map[string]string   `json:"variables"`
	Templates  map[string]Template `json:"templates"`
	Conditions []Condition         `json:"conditions"`
}

func NewBuilder() *Builder {
	return &Builder{
		Variables: make(map[string]string),
		Templates: make(map[string]Template),
	}
}

func (b *Builder) AddContext(ctx string) *Builder {
	b.Context = append(b.Context, ctx)
	return b
}

func (b *Builder) AddVariable(key, value string) *Builder {
	b.Variables[key] = value
	return b
}

func (b *Builder) AddTemplate(name string, template Template) *Template {
	b.Templates[name] = template
	return &template
}

func (b *Builder) AddCondition(If, Do string) *Builder {
	b.Conditions = append(b.Conditions, NewCondition(If, Do))
	return b
}

func (b *Builder) Build() []byte {
	resp, _ := json.MarshalIndent(b, "", "  ")
	return resp
}

type Template []string

func NewTemplate() Template {
	return make(Template, 0)
}

func (t Template) AddLine(line string) Template {
	t = append(t, line)
	return t
}

type Condition struct {
	If string `json:"if"`
	Do string `json:"do"`
}

func NewCondition(If, Do string) Condition {
	return Condition{
		If: If,
		Do: Do,
	}
}
