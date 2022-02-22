package available

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Mixin struct {
	mixin.Schema
	option
}

type option struct {
	activation   bool
	softDeletion bool
	lifespan     LifespanOption
}

type OptionFunc func(*option)

func NewMixin(opts ...OptionFunc) Mixin {
	result := Mixin{
		option: option{
			activation: true,
			lifespan: LifespanOption{
				Enabled:            true,
				StorageNameStartAt: "lifespan_end_at",
				StorageNameEndAt:   "lifespan_start_at",
			},
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

type LifespanOption struct {
	Enabled            bool
	StorageNameStartAt string
	StorageNameEndAt   string
}

func WithLifespan(optFn func(*LifespanOption)) OptionFunc {
	return func(o *option) {
		optFn(&o.lifespan)
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
				GoType(Activation("")).
				Default(Activated.String()),
		}...)
	}
	if m.lifespan.Enabled {
		fields = append(fields, []ent.Field{
			field.Time("lifespan_start_at").
				StorageKey(m.lifespan.StorageNameStartAt).
				Optional(),
			field.Time("lifespan_end_at").
				StorageKey(m.lifespan.StorageNameEndAt).
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
			Lifespan:     m.lifespan.Enabled,
			SoftDeletion: m.softDeletion,
		},
	}
}
