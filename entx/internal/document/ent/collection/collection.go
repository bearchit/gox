// Code generated by entc, DO NOT EDIT.

package collection

const (
	// Label holds the string label denoting the collection type in the database.
	Label = "collection"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLifespanStartAt holds the string denoting the lifespan_start_at field in the database.
	FieldLifespanStartAt = "lifespan_start_at"
	// FieldLifespanEndAt holds the string denoting the lifespan_end_at field in the database.
	FieldLifespanEndAt = "lifespan_end_at"
	// Table holds the table name of the collection in the database.
	Table = "collections"
)

// Columns holds all SQL columns for collection fields.
var Columns = []string{
	FieldID,
	FieldLifespanStartAt,
	FieldLifespanEndAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
