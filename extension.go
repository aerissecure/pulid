package pulid

import (
	_ "embed"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

// Extension extends ent with tracking history capabilities.
type Extension struct {
	entc.DefaultExtension
}

// NewExtension creates a new Extension.
func NewExtension() *Extension {
	return &Extension{}
}

//go:embed templates/pulid.tmpl
var tmpl string

// Templates of the extension.
func (*Extension) Templates() []*gen.Template {
	t := gen.NewTemplate("pulid")
	return []*gen.Template{
		gen.MustParse(t.Parse(tmpl)),
	}
}
