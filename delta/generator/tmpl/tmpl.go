package tmpl

import (
	"io"
	"text/template"
)

// Template ...
type Template struct {
	Destination string
	Source      string
	Data        interface{}
	IsOverWrite bool
}

// Templates ...
type Templates []Template

// ExecuteTemplate ...
func (t *Template) ExecuteTemplate(w io.Writer, content []byte, funcMap template.FuncMap, data interface{}) error {
	tmpl, err := template.New(t.Destination).Funcs(funcMap).Parse(string(content))
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}
