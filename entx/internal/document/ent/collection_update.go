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
	"github.com/bearchit/gox/entx/internal/document/ent/collection"
	"github.com/bearchit/gox/entx/internal/document/ent/predicate"
)

// CollectionUpdate is the builder for updating Collection entities.
type CollectionUpdate struct {
	config
	hooks    []Hook
	mutation *CollectionMutation
}

// Where appends a list predicates to the CollectionUpdate builder.
func (cu *CollectionUpdate) Where(ps ...predicate.Collection) *CollectionUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetLifespanStartAt sets the "lifespan_start_at" field.
func (cu *CollectionUpdate) SetLifespanStartAt(t time.Time) *CollectionUpdate {
	cu.mutation.SetLifespanStartAt(t)
	return cu
}

// SetNillableLifespanStartAt sets the "lifespan_start_at" field if the given value is not nil.
func (cu *CollectionUpdate) SetNillableLifespanStartAt(t *time.Time) *CollectionUpdate {
	if t != nil {
		cu.SetLifespanStartAt(*t)
	}
	return cu
}

// ClearLifespanStartAt clears the value of the "lifespan_start_at" field.
func (cu *CollectionUpdate) ClearLifespanStartAt() *CollectionUpdate {
	cu.mutation.ClearLifespanStartAt()
	return cu
}

// SetLifespanEndAt sets the "lifespan_end_at" field.
func (cu *CollectionUpdate) SetLifespanEndAt(t time.Time) *CollectionUpdate {
	cu.mutation.SetLifespanEndAt(t)
	return cu
}

// SetNillableLifespanEndAt sets the "lifespan_end_at" field if the given value is not nil.
func (cu *CollectionUpdate) SetNillableLifespanEndAt(t *time.Time) *CollectionUpdate {
	if t != nil {
		cu.SetLifespanEndAt(*t)
	}
	return cu
}

// ClearLifespanEndAt clears the value of the "lifespan_end_at" field.
func (cu *CollectionUpdate) ClearLifespanEndAt() *CollectionUpdate {
	cu.mutation.ClearLifespanEndAt()
	return cu
}

// Mutation returns the CollectionMutation object of the builder.
func (cu *CollectionUpdate) Mutation() *CollectionMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CollectionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CollectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CollectionUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CollectionUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CollectionUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CollectionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   collection.Table,
			Columns: collection.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: collection.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.LifespanStartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: collection.FieldLifespanStartAt,
		})
	}
	if cu.mutation.LifespanStartAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: collection.FieldLifespanStartAt,
		})
	}
	if value, ok := cu.mutation.LifespanEndAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: collection.FieldLifespanEndAt,
		})
	}
	if cu.mutation.LifespanEndAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: collection.FieldLifespanEndAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{collection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CollectionUpdateOne is the builder for updating a single Collection entity.
type CollectionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CollectionMutation
}

// SetLifespanStartAt sets the "lifespan_start_at" field.
func (cuo *CollectionUpdateOne) SetLifespanStartAt(t time.Time) *CollectionUpdateOne {
	cuo.mutation.SetLifespanStartAt(t)
	return cuo
}

// SetNillableLifespanStartAt sets the "lifespan_start_at" field if the given value is not nil.
func (cuo *CollectionUpdateOne) SetNillableLifespanStartAt(t *time.Time) *CollectionUpdateOne {
	if t != nil {
		cuo.SetLifespanStartAt(*t)
	}
	return cuo
}

// ClearLifespanStartAt clears the value of the "lifespan_start_at" field.
func (cuo *CollectionUpdateOne) ClearLifespanStartAt() *CollectionUpdateOne {
	cuo.mutation.ClearLifespanStartAt()
	return cuo
}

// SetLifespanEndAt sets the "lifespan_end_at" field.
func (cuo *CollectionUpdateOne) SetLifespanEndAt(t time.Time) *CollectionUpdateOne {
	cuo.mutation.SetLifespanEndAt(t)
	return cuo
}

// SetNillableLifespanEndAt sets the "lifespan_end_at" field if the given value is not nil.
func (cuo *CollectionUpdateOne) SetNillableLifespanEndAt(t *time.Time) *CollectionUpdateOne {
	if t != nil {
		cuo.SetLifespanEndAt(*t)
	}
	return cuo
}

// ClearLifespanEndAt clears the value of the "lifespan_end_at" field.
func (cuo *CollectionUpdateOne) ClearLifespanEndAt() *CollectionUpdateOne {
	cuo.mutation.ClearLifespanEndAt()
	return cuo
}

// Mutation returns the CollectionMutation object of the builder.
func (cuo *CollectionUpdateOne) Mutation() *CollectionMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CollectionUpdateOne) Select(field string, fields ...string) *CollectionUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Collection entity.
func (cuo *CollectionUpdateOne) Save(ctx context.Context) (*Collection, error) {
	var (
		err  error
		node *Collection
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CollectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CollectionUpdateOne) SaveX(ctx context.Context) *Collection {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CollectionUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CollectionUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CollectionUpdateOne) sqlSave(ctx context.Context) (_node *Collection, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   collection.Table,
			Columns: collection.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: collection.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Collection.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, collection.FieldID)
		for _, f := range fields {
			if !collection.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != collection.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.LifespanStartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: collection.FieldLifespanStartAt,
		})
	}
	if cuo.mutation.LifespanStartAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: collection.FieldLifespanStartAt,
		})
	}
	if value, ok := cuo.mutation.LifespanEndAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: collection.FieldLifespanEndAt,
		})
	}
	if cuo.mutation.LifespanEndAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: collection.FieldLifespanEndAt,
		})
	}
	_node = &Collection{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{collection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
