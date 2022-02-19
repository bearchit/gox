package available

import (
	"embed"
	"entgo.io/ent/entc/gen"
)

var (
	QueryAvailableTemplate = parseT("template/query_available.tmpl")
	LifespanTemplate       = parseT("template/lifespan.tmpl")

	AllTemplates = []*gen.Template{
		QueryAvailableTemplate,
		LifespanTemplate,
	}

	//go:embed template/*
	templates embed.FS
)

func parseT(path string) *gen.Template {
	return gen.MustParse(
		gen.NewTemplate(path).
			ParseFS(templates, path),
	)
}
