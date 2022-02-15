package schema

import (
	"entgo.io/ent"
	"github.com/bearchit/gox/entx/available"
)

// Collection holds the schema definition for the Collection entity.
type Collection struct {
	ent.Schema
}

func (Collection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		available.NewMixin(
			available.WithActivation(false),
			available.WithSoftDeletion(false),
		),
	}
}

// Fields of the Collection.
func (Collection) Fields() []ent.Field {
	return nil
}

// Edges of the Collection.
func (Collection) Edges() []ent.Edge {
	return nil
}
