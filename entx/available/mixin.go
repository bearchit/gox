package available

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/bearchit/gox/entx/available/activation"
)

type Mixin struct {
	mixin.Schema
	option
}

type option struct {
	activation   bool
	lifespan     bool
	softDeletion bool
}

type OptionFunc func(*option)

func NewMixin(opts ...OptionFunc) Mixin {
	result := Mixin{
		option: option{
			activation:   true,
			lifespan:     true,
			softDeletion: true,
		},
	}
	for _, opt := range opts {
		opt(&result.option)
	}
	return result
}

func WithActivation(v bool) OptionFunc {
	return func(o *option) {
		o.activation = v
	}
}

func WithLifespan(v bool) OptionFunc {
	return func(o *option) {
		o.lifespan = v
	}
}

func WithSoftDeletion(v bool) OptionFunc {
	return func(o *option) {
		o.softDeletion = v
	}
}

func (m Mixin) Fields() []ent.Field {
	fields := make([]ent.Field, 0)
	if m.activation {
		fields = append(fields, []ent.Field{
			field.Enum("activation").
				GoType(activation.Activation("")).
				Default(activation.Activated.String()),
		}...)
	}
	if m.lifespan {
		fields = append(fields, []ent.Field{
			field.Time("lifespan_start_at").
				StorageKey("lifespan_start_at").
				Optional(),
			field.Time("lifespan_end_at").
				StorageKey("lifespan_end_at").
				Optional(),
		}...)
	}
	if m.softDeletion {
		fields = append(fields, []ent.Field{
			field.Time("deleted_at").
				Optional(),
		}...)
	}

	return fields
}

func (m Mixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		Annotation{
			Activation:   m.activation,
			Lifespan:     m.lifespan,
			SoftDeletion: m.softDeletion,
		},
	}
}
