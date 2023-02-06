// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/card"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/entc/integration/ent/spec"
	"entgo.io/ent/schema/field"
)

// SpecUpdate is the builder for updating Spec entities.
type SpecUpdate struct {
	config
	hooks     []Hook
	mutation  *SpecMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SpecUpdate builder.
func (su *SpecUpdate) Where(ps ...predicate.Spec) *SpecUpdate {
	su.mutation.Where(ps...)
	return su
}

// AddCardIDs adds the "card" edge to the Card entity by IDs.
func (su *SpecUpdate) AddCardIDs(ids ...int) *SpecUpdate {
	su.mutation.AddCardIDs(ids...)
	return su
}

// AddCard adds the "card" edges to the Card entity.
func (su *SpecUpdate) AddCard(c ...*Card) *SpecUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.AddCardIDs(ids...)
}

// Mutation returns the SpecMutation object of the builder.
func (su *SpecUpdate) Mutation() *SpecMutation {
	return su.mutation
}

// ClearCard clears all "card" edges to the Card entity.
func (su *SpecUpdate) ClearCard() *SpecUpdate {
	su.mutation.ClearCard()
	return su
}

// RemoveCardIDs removes the "card" edge to Card entities by IDs.
func (su *SpecUpdate) RemoveCardIDs(ids ...int) *SpecUpdate {
	su.mutation.RemoveCardIDs(ids...)
	return su
}

// RemoveCard removes "card" edges to Card entities.
func (su *SpecUpdate) RemoveCard(c ...*Card) *SpecUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.RemoveCardIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SpecUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, SpecMutation](ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SpecUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SpecUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SpecUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (su *SpecUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SpecUpdate {
	su.modifiers = append(su.modifiers, modifiers...)
	return su
}

func (su *SpecUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(spec.Table, spec.Columns, sqlgraph.NewFieldSpec(spec.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if su.mutation.CardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   spec.CardTable,
			Columns: spec.CardPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedCardIDs(); len(nodes) > 0 && !su.mutation.CardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   spec.CardTable,
			Columns: spec.CardPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.CardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   spec.CardTable,
			Columns: spec.CardPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(su.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{spec.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SpecUpdateOne is the builder for updating a single Spec entity.
type SpecUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SpecMutation
	modifiers []func(*sql.UpdateBuilder)
}

// AddCardIDs adds the "card" edge to the Card entity by IDs.
func (suo *SpecUpdateOne) AddCardIDs(ids ...int) *SpecUpdateOne {
	suo.mutation.AddCardIDs(ids...)
	return suo
}

// AddCard adds the "card" edges to the Card entity.
func (suo *SpecUpdateOne) AddCard(c ...*Card) *SpecUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.AddCardIDs(ids...)
}

// Mutation returns the SpecMutation object of the builder.
func (suo *SpecUpdateOne) Mutation() *SpecMutation {
	return suo.mutation
}

// ClearCard clears all "card" edges to the Card entity.
func (suo *SpecUpdateOne) ClearCard() *SpecUpdateOne {
	suo.mutation.ClearCard()
	return suo
}

// RemoveCardIDs removes the "card" edge to Card entities by IDs.
func (suo *SpecUpdateOne) RemoveCardIDs(ids ...int) *SpecUpdateOne {
	suo.mutation.RemoveCardIDs(ids...)
	return suo
}

// RemoveCard removes "card" edges to Card entities.
func (suo *SpecUpdateOne) RemoveCard(c ...*Card) *SpecUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.RemoveCardIDs(ids...)
}

// Where appends a list predicates to the SpecUpdate builder.
func (suo *SpecUpdateOne) Where(ps ...predicate.Spec) *SpecUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SpecUpdateOne) Select(field string, fields ...string) *SpecUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Spec entity.
func (suo *SpecUpdateOne) Save(ctx context.Context) (*Spec, error) {
	return withHooks[*Spec, SpecMutation](ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SpecUpdateOne) SaveX(ctx context.Context) *Spec {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SpecUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SpecUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suo *SpecUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SpecUpdateOne {
	suo.modifiers = append(suo.modifiers, modifiers...)
	return suo
}

func (suo *SpecUpdateOne) sqlSave(ctx context.Context) (_node *Spec, err error) {
	_spec := sqlgraph.NewUpdateSpec(spec.Table, spec.Columns, sqlgraph.NewFieldSpec(spec.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Spec.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, spec.FieldID)
		for _, f := range fields {
			if !spec.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != spec.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if suo.mutation.CardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   spec.CardTable,
			Columns: spec.CardPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedCardIDs(); len(nodes) > 0 && !suo.mutation.CardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   spec.CardTable,
			Columns: spec.CardPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.CardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   spec.CardTable,
			Columns: spec.CardPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(suo.modifiers...)
	_node = &Spec{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{spec.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
