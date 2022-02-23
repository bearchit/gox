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
	config
}

type config struct {
	activation   ActivationOption
	softDeletion SoftDeletionOption
	lifespan     LifespanOption
}

type Option func(*config)

var (
	defaultActivationOption = ActivationOption{
		enabled:   true,
		fieldName: "activation",
	}
	defaultLifespanOption = LifespanOption{
		enabled:          true,
		startAtFieldName: "lifespan_start_at",
		endAtFieldName:   "lifespan_end_at",
	}
	defaultSoftDeletionOption = SoftDeletionOption{
		enabled:   true,
		fieldName: "deleted_at",
	}
)

func NewMixin(options ...Option) Mixin {
	cfg := config{}
	for _, opt := range options {
		opt(&cfg)
	}
	return Mixin{config: cfg}
}

func NewDefaultMixin() Mixin {
	return Mixin{
		config: config{
			activation:   defaultActivationOption,
			lifespan:     defaultLifespanOption,
			softDeletion: defaultSoftDeletionOption,
		},
	}
}

type ActivationOption struct {
	enabled    bool
	fieldName  string
	storageKey string
}

func (option *ActivationOption) SetFieldName(name string) *ActivationOption {
	option.fieldName = name
	return option
}

func (option *ActivationOption) SetStorageKey(key string) *ActivationOption {
	option.storageKey = key
	return option
}

func WithActivation() Option {
	return func(cfg *config) {
		cfg.activation = defaultActivationOption
	}
}

func WithActivationOption(option func(*ActivationOption)) Option {
	return func(cfg *config) {
		cfg.activation = defaultActivationOption
		option(&cfg.activation)
	}
}

type LifespanOption struct {
	enabled           bool
	startAtFieldName  string
	endAtFieldName    string
	startAtStorageKey string
	endAtStorageKey   string
}

func (option *LifespanOption) SetFieldNames(startAt, endAt string) *LifespanOption {
	option.startAtFieldName, option.endAtFieldName = startAt, endAt
	return option
}

func (option *LifespanOption) SetStorageKeys(startAt, endAt string) *LifespanOption {
	option.startAtStorageKey, option.endAtStorageKey = startAt, endAt
	return option
}

func WithLifespan() Option {
	return func(cfg *config) {
		cfg.lifespan = defaultLifespanOption
	}
}

func WithLifespanOption(option func(*LifespanOption)) Option {
	return func(cfg *config) {
		cfg.lifespan = defaultLifespanOption
		option(&cfg.lifespan)
	}
}

type SoftDeletionOption struct {
	enabled    bool
	fieldName  string
	storageKey string
}

func (option *SoftDeletionOption) SetFieldName(name string) *SoftDeletionOption {
	option.fieldName = name
	return option
}

func (option *SoftDeletionOption) SetStorageKey(key string) *SoftDeletionOption {
	option.storageKey = key
	return option
}

func WithSoftDeletion() Option {
	return func(cfg *config) {
		cfg.softDeletion = defaultSoftDeletionOption
	}
}

func WithSoftDeletionOption(option func(*SoftDeletionOption)) Option {
	return func(cfg *config) {
		cfg.softDeletion = defaultSoftDeletionOption
		option(&cfg.softDeletion)
	}
}

func (m Mixin) Fields() []ent.Field {
	fields := make([]ent.Field, 0)

	if cfg := m.activation; cfg.enabled {
		f := field.Enum(cfg.fieldName).
			GoType(Activation("")).
			Default(Activated.String())
		if cfg.storageKey != "" {
			f.StorageKey(cfg.storageKey)
		}
		fields = append(fields, f)
	}

	if cfg := m.lifespan; cfg.enabled {
		fields = append(fields, []ent.Field{
			m.makeLifespanField(cfg.startAtFieldName, cfg.startAtStorageKey),
			m.makeLifespanField(cfg.endAtFieldName, cfg.endAtStorageKey),
		}...)
	}

	if cfg := m.softDeletion; cfg.enabled {
		f := field.Time(cfg.fieldName).
			Optional()
		if cfg.storageKey != "" {
			f.StorageKey(cfg.storageKey)
		}
		fields = append(fields, f)
	}

	return fields
}

func (m Mixin) Indexes() []ent.Index {
	indexes := make([]ent.Index, 0)
	if m.softDeletion.enabled {
		indexes = append(indexes, index.Fields("deleted_at"))
	}

	return indexes
}

func (m Mixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		Annotation{
			Activation:   m.activation.enabled,
			Lifespan:     m.lifespan.enabled,
			SoftDeletion: m.softDeletion.enabled,
		},
	}
}

func (m Mixin) makeLifespanField(fieldName, storageKey string) ent.Field {
	f := field.Time(fieldName).Optional()
	if storageKey != "" {
		f.StorageKey(storageKey)
	}
	return f
}
