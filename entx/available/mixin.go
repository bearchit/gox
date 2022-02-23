package available

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
	option := option{}
	for _, opt := range opts {
		opt(&option)
	}
	return Mixin{option: option}
}

func WithAll() OptionFunc {
	return func(o *option) {
		o.activation = true
		o.lifespan.enabled = true
		o.softDeletion = true
	}
}

func WithActivation() OptionFunc {
	return func(o *option) {
		o.activation = true
	}
}

type LifespanOption struct {
	enabled            bool
	startAtFieldName   string
	endAtFieldName     string
	startAtStorageName string
	endAtStorageName   string
}

var defaultLifespanOption = LifespanOption{
	enabled:            true,
	startAtStorageName: "lifespan_start_at",
}

func WithLifespan() OptionFunc {
	return func(o *option) {
		o.lifespan.enabled = true
	}
}

func WithLifespanOption(optFn func(*LifespanOption)) OptionFunc {
	return func(o *option) {
		optFn(&o.lifespan)
	}
}

func WithSoftDeletion() OptionFunc {
	return func(o *option) {
		o.softDeletion = true
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
	if m.lifespan.enabled {
		fields = append(fields, []ent.Field{
			field.Time("lifespan_start_at").
				StorageKey(m.lifespan.startAtStorageName).
				Optional(),
			field.Time("lifespan_end_at").
				StorageKey(m.lifespan.endAtStorageName).
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

func (m Mixin) Indexes() []ent.Index {
	indexes := make([]ent.Index, 0)
	if m.softDeletion {
		indexes = append(indexes, index.Fields("deleted_at"))
	}

	return indexes
}

func (m Mixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		Annotation{
			Activation:   m.activation,
			Lifespan:     m.lifespan.enabled,
			SoftDeletion: m.softDeletion,
		},
	}
}
