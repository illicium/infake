package infake

import (
	"bytes"
	"text/template"
)

type StringTemplate struct {
	*template.Template
}

func (t StringTemplate) Execute(data interface{}) (string, error) {
	var b bytes.Buffer

	err := t.Template.Execute(&b, data)

	if err != nil {
		return "", err
	}

	return b.String(), nil
}
