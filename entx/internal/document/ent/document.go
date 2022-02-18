// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/bearchit/gox/entx/available/activation"
	"github.com/bearchit/gox/entx/internal/document/ent/document"
)

// Document is the model entity for the Document schema.
type Document struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Activation holds the value of the "activation" field.
	Activation activation.Activation `json:"activation,omitempty"`
	// LifespanStartAt holds the value of the "lifespan_start_at" field.
	LifespanStartAt time.Time `json:"lifespan_start_at,omitempty"`
	// LifespanEndAt holds the value of the "lifespan_end_at" field.
	LifespanEndAt time.Time `json:"lifespan_end_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Document) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case document.FieldID:
			values[i] = new(sql.NullInt64)
		case document.FieldActivation:
			values[i] = new(sql.NullString)
		case document.FieldLifespanStartAt, document.FieldLifespanEndAt, document.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Document", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Document fields.
func (d *Document) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case document.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case document.FieldActivation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field activation", values[i])
			} else if value.Valid {
				d.Activation = activation.Activation(value.String)
			}
		case document.FieldLifespanStartAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field lifespan_start_at", values[i])
			} else if value.Valid {
				d.LifespanStartAt = value.Time
			}
		case document.FieldLifespanEndAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field lifespan_end_at", values[i])
			} else if value.Valid {
				d.LifespanEndAt = value.Time
			}
		case document.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				d.DeletedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Document.
// Note that you need to call Document.Unwrap() before calling this method if this Document
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Document) Update() *DocumentUpdateOne {
	return (&DocumentClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Document entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Document) Unwrap() *Document {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Document is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Document) String() string {
	var builder strings.Builder
	builder.WriteString("Document(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", activation=")
	builder.WriteString(fmt.Sprintf("%v", d.Activation))
	builder.WriteString(", lifespan_start_at=")
	builder.WriteString(d.LifespanStartAt.Format(time.ANSIC))
	builder.WriteString(", lifespan_end_at=")
	builder.WriteString(d.LifespanEndAt.Format(time.ANSIC))
	builder.WriteString(", deleted_at=")
	builder.WriteString(d.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Documents is a parsable slice of Document.
type Documents []*Document

func (d Documents) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}