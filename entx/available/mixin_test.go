package available_test

import (
	"entgo.io/ent"
)

func hasFields(fields []ent.Field, fieldNames ...string) bool {
	matches := make([]bool, 0)
	for _, f := range fields {
		for _, fn := range fieldNames {
			if f.Descriptor().Name == fn {
				matches = append(matches, true)
			}
		}
	}
	return len(matches) == len(fieldNames)
}

//func TestLifespanOption(t *testing.T) {
//	tests := []struct {
//		given available.Mixin
//	}{}
//
//}
