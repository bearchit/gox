package available_test

import (
	"testing"

	"entgo.io/ent"
	"github.com/bearchit/gox/entx/available"
	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	t.Parallel()

	type wants struct {
		fieldName  string
		storageKey string
	}

	type testCase struct {
		given available.Mixin
		wants wants
	}

	tests := []struct {
		name  string
		cases []testCase
	}{
		{
			"Available",
			[]testCase{
				{available.NewDefaultMixin(), wants{"activation", ""}},
				{available.NewMixin(available.WithActivation()), wants{"activation", ""}},
				{available.NewMixin(available.WithActivationOption(func(ao *available.ActivationOption) {
					ao.SetFieldName("state")
				})), wants{"state", ""}},
				{available.NewMixin(available.WithActivationOption(func(ao *available.ActivationOption) {
					ao.SetStorageKey("state")
				})), wants{"activation", "state"}},
				{available.NewMixin(available.WithActivationOption(func(ao *available.ActivationOption) {
					ao.SetFieldName("activation").
						SetStorageKey("state")
				})), wants{"activation", "state"}},
			},
		},
		{
			"Lifespan",
			[]testCase{
				{available.NewDefaultMixin(), wants{"lifespan_start_at", ""}},
				{available.NewDefaultMixin(), wants{"lifespan_end_at", ""}},
				{available.NewMixin(available.WithLifespan()), wants{"lifespan_start_at", ""}},
				{available.NewMixin(available.WithLifespan()), wants{"lifespan_end_at", ""}},
				{available.NewMixin(available.WithLifespanOption(func(lo *available.LifespanOption) {
					lo.SetFieldNames("start_at", "end_at")
				})), wants{"start_at", ""}},
				{available.NewMixin(available.WithLifespanOption(func(lo *available.LifespanOption) {
					lo.SetFieldNames("start_at", "end_at")
				})), wants{"end_at", ""}},
				{available.NewMixin(available.WithLifespanOption(func(lo *available.LifespanOption) {
					lo.SetStorageKeys("start_at", "end_at")
				})), wants{"lifespan_start_at", "start_at"}},
				{available.NewMixin(available.WithLifespanOption(func(lo *available.LifespanOption) {
					lo.SetStorageKeys("start_at", "end_at")
				})), wants{"lifespan_end_at", "end_at"}},
				{available.NewMixin(available.WithLifespanOption(func(lo *available.LifespanOption) {
					lo.SetFieldNames("start_at", "end_at").
						SetStorageKeys("start_at", "end_at")
				})), wants{"start_at", "start_at"}},
				{available.NewMixin(available.WithLifespanOption(func(lo *available.LifespanOption) {
					lo.SetFieldNames("start_at", "end_at").
						SetStorageKeys("start_at", "end_at")
				})), wants{"end_at", "end_at"}},
			},
		},
		{
			"SoftDeletion",
			[]testCase{
				{available.NewDefaultMixin(), wants{"deleted_at", ""}},
				{available.NewMixin(available.WithSoftDeletion()), wants{"deleted_at", ""}},
				{available.NewMixin(available.WithSoftDeletionOption(func(sdo *available.SoftDeletionOption) {
					sdo.SetFieldName("deleted")
				})), wants{"deleted", ""}},
				{available.NewMixin(available.WithSoftDeletionOption(func(sdo *available.SoftDeletionOption) {
					sdo.SetStorageKey("deleted")
				})), wants{"deleted_at", "deleted"}},
				{available.NewMixin(available.WithSoftDeletionOption(func(sdo *available.SoftDeletionOption) {
					sdo.SetFieldName("deleted_at").
						SetStorageKey("deleted")
				})), wants{"deleted_at", "deleted"}},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			for _, tc := range tt.cases {
				tc := tc
				t.Run("", func(t *testing.T) {
					t.Parallel()
					assert.True(t, hasField(tc.given.Fields(), tc.wants.fieldName, tc.wants.storageKey))
				})
			}
		})
	}
}

func hasField(fields []ent.Field, name, storageKey string) bool {
	for _, f := range fields {
		desc := f.Descriptor()
		if desc.Name == name && desc.StorageKey == storageKey {
			return true
		}
	}
	return false
}
