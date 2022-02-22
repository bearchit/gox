package schema

import (
	"entgo.io/ent"
	"github.com/bearchit/gox/entx/available"
)

// Revision holds the schema definition for the Revision entity.
type Revision struct {
	ent.Schema
}

func (Revision) Mixin() []ent.Mixin {
	return []ent.Mixin{
		available.NewMixin(available.WithLifespan(false)),
	}
}

// Fields of the Revision.
func (Revision) Fields() []ent.Field {
	return nil
}

// Edges of the Revision.
func (Revision) Edges() []ent.Edge {
	return nil
}
