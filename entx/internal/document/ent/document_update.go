// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bearchit/gox/entx/available"
	"github.com/bearchit/gox/entx/internal/document/ent/document"
	"github.com/bearchit/gox/entx/internal/document/ent/predicate"
)

// DocumentUpdate is the builder for updating Document entities.
type DocumentUpdate struct {
	config
	hooks    []Hook
	mutation *DocumentMutation
}

// Where appends a list predicates to the DocumentUpdate builder.
func (du *DocumentUpdate) Where(ps ...predicate.Document) *DocumentUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetActivation sets the "activation" field.
func (du *DocumentUpdate) SetActivation(a available.Activation) *DocumentUpdate {
	du.mutation.SetActivation(a)
	return du
}

// SetNillableActivation sets the "activation" field if the given value is not nil.
func (du *DocumentUpdate) SetNillableActivation(a *available.Activation) *DocumentUpdate {
	if a != nil {
		du.SetActivation(*a)
	}
	return du
}

// SetLifespanStartAt sets the "lifespan_start_at" field.
func (du *DocumentUpdate) SetLifespanStartAt(t time.Time) *DocumentUpdate {
	du.mutation.SetLifespanStartAt(t)
	return du
}

// SetNillableLifespanStartAt sets the "lifespan_start_at" field if the given value is not nil.
func (du *DocumentUpdate) SetNillableLifespanStartAt(t *time.Time) *DocumentUpdate {
	if t != nil {
		du.SetLifespanStartAt(*t)
	}
	return du
}

// ClearLifespanStartAt clears the value of the "lifespan_start_at" field.
func (du *DocumentUpdate) ClearLifespanStartAt() *DocumentUpdate {
	du.mutation.ClearLifespanStartAt()
	return du
}

// SetLifespanEndAt sets the "lifespan_end_at" field.
func (du *DocumentUpdate) SetLifespanEndAt(t time.Time) *DocumentUpdate {
	du.mutation.SetLifespanEndAt(t)
	return du
}

// SetNillableLifespanEndAt sets the "lifespan_end_at" field if the given value is not nil.
func (du *DocumentUpdate) SetNillableLifespanEndAt(t *time.Time) *DocumentUpdate {
	if t != nil {
		du.SetLifespanEndAt(*t)
	}
	return du
}

// ClearLifespanEndAt clears the value of the "lifespan_end_at" field.
func (du *DocumentUpdate) ClearLifespanEndAt() *DocumentUpdate {
	du.mutation.ClearLifespanEndAt()
	return du
}

// SetDeletedAt sets the "deleted_at" field.
func (du *DocumentUpdate) SetDeletedAt(t time.Time) *DocumentUpdate {
	du.mutation.SetDeletedAt(t)
	return du
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (du *DocumentUpdate) SetNillableDeletedAt(t *time.Time) *DocumentUpdate {
	if t != nil {
		du.SetDeletedAt(*t)
	}
	return du
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (du *DocumentUpdate) ClearDeletedAt() *DocumentUpdate {
	du.mutation.ClearDeletedAt()
	return du
}

// Mutation returns the DocumentMutation object of the builder.
func (du *DocumentUpdate) Mutation() *DocumentMutation {
	return du.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DocumentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		if err = du.check(); err != nil {
			return 0, err
		}
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DocumentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = du.check(); err != nil {
				return 0, err
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DocumentUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DocumentUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DocumentUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DocumentUpdate) check() error {
	if v, ok := du.mutation.Activation(); ok {
		if err := document.ActivationValidator(v); err != nil {
			return &ValidationError{Name: "activation", err: fmt.Errorf(`ent: validator failed for field "Document.activation": %w`, err)}
		}
	}
	return nil
}

func (du *DocumentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   document.Table,
			Columns: document.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: document.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Activation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: document.FieldActivation,
		})
	}
	if value, ok := du.mutation.LifespanStartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: document.FieldLifespanStartAt,
		})
	}
	if du.mutation.LifespanStartAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: document.FieldLifespanStartAt,
		})
	}
	if value, ok := du.mutation.LifespanEndAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: document.FieldLifespanEndAt,
		})
	}
	if du.mutation.LifespanEndAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: document.FieldLifespanEndAt,
		})
	}
	if value, ok := du.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: document.FieldDeletedAt,
		})
	}
	if du.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: document.FieldDeletedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{document.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DocumentUpdateOne is the builder for updating a single Document entity.
type DocumentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DocumentMutation
}

// SetActivation sets the "activation" field.
func (duo *DocumentUpdateOne) SetActivation(a available.Activation) *DocumentUpdateOne {
	duo.mutation.SetActivation(a)
	return duo
}

// SetNillableActivation sets the "activation" field if the given value is not nil.
func (duo *DocumentUpdateOne) SetNillableActivation(a *available.Activation) *DocumentUpdateOne {
	if a != nil {
		duo.SetActivation(*a)
	}
	return duo
}

// SetLifespanStartAt sets the "lifespan_start_at" field.
func (duo *DocumentUpdateOne) SetLifespanStartAt(t time.Time) *DocumentUpdateOne {
	duo.mutation.SetLifespanStartAt(t)
	return duo
}

// SetNillableLifespanStartAt sets the "lifespan_start_at" field if the given value is not nil.
func (duo *DocumentUpdateOne) SetNillableLifespanStartAt(t *time.Time) *DocumentUpdateOne {
	if t != nil {
		duo.SetLifespanStartAt(*t)
	}
	return duo
}

// ClearLifespanStartAt clears the value of the "lifespan_start_at" field.
func (duo *DocumentUpdateOne) ClearLifespanStartAt() *DocumentUpdateOne {
	duo.mutation.ClearLifespanStartAt()
	return duo
}

// SetLifespanEndAt sets the "lifespan_end_at" field.
func (duo *DocumentUpdateOne) SetLifespanEndAt(t time.Time) *DocumentUpdateOne {
	duo.mutation.SetLifespanEndAt(t)
	return duo
}

// SetNillableLifespanEndAt sets the "lifespan_end_at" field if the given value is not nil.
func (duo *DocumentUpdateOne) SetNillableLifespanEndAt(t *time.Time) *DocumentUpdateOne {
	if t != nil {
		duo.SetLifespanEndAt(*t)
	}
	return duo
}

// ClearLifespanEndAt clears the value of the "lifespan_end_at" field.
func (duo *DocumentUpdateOne) ClearLifespanEndAt() *DocumentUpdateOne {
	duo.mutation.ClearLifespanEndAt()
	return duo
}

// SetDeletedAt sets the "deleted_at" field.
func (duo *DocumentUpdateOne) SetDeletedAt(t time.Time) *DocumentUpdateOne {
	duo.mutation.SetDeletedAt(t)
	return duo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (duo *DocumentUpdateOne) SetNillableDeletedAt(t *time.Time) *DocumentUpdateOne {
	if t != nil {
		duo.SetDeletedAt(*t)
	}
	return duo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (duo *DocumentUpdateOne) ClearDeletedAt() *DocumentUpdateOne {
	duo.mutation.ClearDeletedAt()
	return duo
}

// Mutation returns the DocumentMutation object of the builder.
func (duo *DocumentUpdateOne) Mutation() *DocumentMutation {
	return duo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DocumentUpdateOne) Select(field string, fields ...string) *DocumentUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Document entity.
func (duo *DocumentUpdateOne) Save(ctx context.Context) (*Document, error) {
	var (
		err  error
		node *Document
	)
	if len(duo.hooks) == 0 {
		if err = duo.check(); err != nil {
			return nil, err
		}
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DocumentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = duo.check(); err != nil {
				return nil, err
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DocumentUpdateOne) SaveX(ctx context.Context) *Document {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DocumentUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DocumentUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DocumentUpdateOne) check() error {
	if v, ok := duo.mutation.Activation(); ok {
		if err := document.ActivationValidator(v); err != nil {
			return &ValidationError{Name: "activation", err: fmt.Errorf(`ent: validator failed for field "Document.activation": %w`, err)}
		}
	}
	return nil
}

func (duo *DocumentUpdateOne) sqlSave(ctx context.Context) (_node *Document, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   document.Table,
			Columns: document.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: document.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Document.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, document.FieldID)
		for _, f := range fields {
			if !document.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != document.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Activation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: document.FieldActivation,
		})
	}
	if value, ok := duo.mutation.LifespanStartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: document.FieldLifespanStartAt,
		})
	}
	if duo.mutation.LifespanStartAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: document.FieldLifespanStartAt,
		})
	}
	if value, ok := duo.mutation.LifespanEndAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: document.FieldLifespanEndAt,
		})
	}
	if duo.mutation.LifespanEndAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: document.FieldLifespanEndAt,
		})
	}
	if value, ok := duo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: document.FieldDeletedAt,
		})
	}
	if duo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: document.FieldDeletedAt,
		})
	}
	_node = &Document{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{document.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
