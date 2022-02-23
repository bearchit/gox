package schema

import (
	"entgo.io/ent"
	"github.com/bearchit/gox/entx/available"
)

// Document holds the schema definition for the Document entity.
type Document struct {
	ent.Schema
}

func (Document) Mixin() []ent.Mixin {
	return []ent.Mixin{
		available.NewMixin(available.WithAll()),
	}
}

// Fields of the Document.
func (Document) Fields() []ent.Field {
	return nil
}

// Edges of the Document.
func (Document) Edges() []ent.Edge {
	return nil
}
