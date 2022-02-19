package available

import (
	"embed"
	"entgo.io/ent/entc/gen"
)

var (
	QueryAvailableTemplate = parseT("template/query_available.tmpl")
	LifespanTemplate       = parseT("template/lifespan.tmpl")
	AvailabilityTemplate   = parseT("template/availability.tmpl")

	AllTemplates = []*gen.Template{
		QueryAvailableTemplate,
		LifespanTemplate,
		AvailabilityTemplate,
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
