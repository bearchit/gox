// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/bearchit/gox/entx/internal/document/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	documentMixin := schema.Document{}.Mixin()
	documentMixinFields0 := documentMixin[0].Fields()
	_ = documentMixinFields0
	documentFields := schema.Document{}.Fields()
	_ = documentFields
	revisionMixin := schema.Revision{}.Mixin()
	revisionMixinFields0 := revisionMixin[0].Fields()
	_ = revisionMixinFields0
	revisionFields := schema.Revision{}.Fields()
	_ = revisionFields
}
