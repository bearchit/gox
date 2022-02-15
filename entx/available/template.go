package available

import (
	"embed"
	"entgo.io/ent/entc/gen"
)

var (
	QueryAvailableTemplate = parseT("template/query_available.tmpl")

	AllTemplates = []*gen.Template{
		QueryAvailableTemplate,
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
