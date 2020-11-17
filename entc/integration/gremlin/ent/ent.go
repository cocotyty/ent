// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect/gremlin"
	"github.com/facebook/ent/dialect/gremlin/encoding/graphson"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/__"
)

// ent aliases to avoid import conflict in user's code.
type (
	Op         = ent.Op
	Hook       = ent.Hook
	Value      = ent.Value
	Query      = ent.Query
	Policy     = ent.Policy
	Mutator    = ent.Mutator
	Mutation   = ent.Mutation
	MutateFunc = ent.MutateFunc
)

// OrderFunc applies an ordering on the graph traversal.
type OrderFunc func(*dsl.Traversal)

// Asc applies the given fields in ASC order.
func Asc(fields ...string) OrderFunc {
	return func(tr *dsl.Traversal) {
		for _, f := range fields {
			tr.By(f, dsl.Incr)
		}
	}
}

// Desc applies the given fields in DESC order.
func Desc(fields ...string) OrderFunc {
	return func(tr *dsl.Traversal) {
		for _, f := range fields {
			tr.By(f, dsl.Decr)
		}
	}
}

// AggregateFunc applies an aggregation step on the group-by traversal/selector.
// It gets two labels as parameters. The first used in the `As` step for the predicate,
// and the second is an optional name for the next predicates (or for later usage).
type AggregateFunc func(string, string) (string, *dsl.Traversal)

// As is a pseudo aggregation function for renaming another other functions with custom names. For example:
//
//	GroupBy(field1, field2).
//	Aggregate(ent.As(ent.Sum(field1), "sum_field1"), (ent.As(ent.Sum(field2), "sum_field2")).
//	Scan(ctx, &v)
//
func As(fn AggregateFunc, end string) AggregateFunc {
	return func(start, _ string) (string, *dsl.Traversal) {
		return fn(start, end)
	}
}

// DefaultCountLabel is the default label name for the Count aggregation function.
// It should be used as the struct-tag for decoding, or a map key for interaction with the returned response.
// In order to "count" 2 or more fields and avoid conflicting, use the `ent.As(ent.Count(field), "custom_name")`
// function with custom name in order to override it.
const DefaultCountLabel = "count"

// Count applies the "count" aggregation function on each group.
func Count() AggregateFunc {
	return func(start, end string) (string, *dsl.Traversal) {
		if end == "" {
			end = DefaultCountLabel
		}
		return end, __.As(start).Count(dsl.Local).As(end)
	}
}

// DefaultMaxLabel is the default label name for the Max aggregation function.
// It should be used as the struct-tag for decoding, or a map key for interaction with the returned response.
// In order to "max" 2 or more fields and avoid conflicting, use the `ent.As(ent.Max(field), "custom_name")`
// function with custom name in order to override it.
const DefaultMaxLabel = "max"

// Max applies the "max" aggregation function on the given field of each group.
func Max(field string) AggregateFunc {
	return func(start, end string) (string, *dsl.Traversal) {
		if end == "" {
			end = DefaultMaxLabel
		}
		return end, __.As(start).Unfold().Values(field).Max().As(end)
	}
}

// DefaultMeanLabel is the default label name for the Mean aggregation function.
// It should be used as the struct-tag for decoding, or a map key for interaction with the returned response.
// In order to "mean" 2 or more fields and avoid conflicting, use the `ent.As(ent.Mean(field), "custom_name")`
// function with custom name in order to override it.
const DefaultMeanLabel = "mean"

// Mean applies the "mean" aggregation function on the given field of each group.
func Mean(field string) AggregateFunc {
	return func(start, end string) (string, *dsl.Traversal) {
		if end == "" {
			end = DefaultMeanLabel
		}
		return end, __.As(start).Unfold().Values(field).Mean().As(end)
	}
}

// DefaultMinLabel is the default label name for the Min aggregation function.
// It should be used as the struct-tag for decoding, or a map key for interaction with the returned response.
// In order to "min" 2 or more fields and avoid conflicting, use the `ent.As(ent.Min(field), "custom_name")`
// function with custom name in order to override it.
const DefaultMinLabel = "min"

// Min applies the "min" aggregation function on the given field of each group.
func Min(field string) AggregateFunc {
	return func(start, end string) (string, *dsl.Traversal) {
		if end == "" {
			end = DefaultMinLabel
		}
		return end, __.As(start).Unfold().Values(field).Min().As(end)
	}
}

// DefaultSumLabel is the default label name for the Sum aggregation function.
// It should be used as the struct-tag for decoding, or a map key for interaction with the returned response.
// In order to "sum" 2 or more fields and avoid conflicting, use the `ent.As(ent.Sum(field), "custom_name")`
// function with custom name in order to override it.
const DefaultSumLabel = "sum"

// Sum applies the "sum" aggregation function on the given field of each group.
func Sum(field string) AggregateFunc {
	return func(start, end string) (string, *dsl.Traversal) {
		if end == "" {
			end = DefaultSumLabel
		}
		return end, __.As(start).Unfold().Values(field).Sum().As(end)
	}
}

// ValidationError returns when validating a field fails.
type ValidationError struct {
	Name string // Field or edge name.
	err  error
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	return e.err.Error()
}

// Unwrap implements the errors.Wrapper interface.
func (e *ValidationError) Unwrap() error {
	return e.err
}

// IsValidationError returns a boolean indicating whether the error is a validaton error.
func IsValidationError(err error) bool {
	if err == nil {
		return false
	}
	var e *ValidationError
	return errors.As(err, &e)
}

// NotFoundError returns when trying to fetch a specific entity and it was not found in the database.
type NotFoundError struct {
	label string
}

// Error implements the error interface.
func (e *NotFoundError) Error() string {
	return "ent: " + e.label + " not found"
}

// IsNotFound returns a boolean indicating whether the error is a not found error.
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	var e *NotFoundError
	return errors.As(err, &e)
}

// MaskNotFound masks not found error.
func MaskNotFound(err error) error {
	if IsNotFound(err) {
		return nil
	}
	return err
}

// NotSingularError returns when trying to fetch a singular entity and more then one was found in the database.
type NotSingularError struct {
	label string
}

// Error implements the error interface.
func (e *NotSingularError) Error() string {
	return "ent: " + e.label + " not singular"
}

// IsNotSingular returns a boolean indicating whether the error is a not singular error.
func IsNotSingular(err error) bool {
	if err == nil {
		return false
	}
	var e *NotSingularError
	return errors.As(err, &e)
}

// NotLoadedError returns when trying to get a node that was not loaded by the query.
type NotLoadedError struct {
	edge string
}

// Error implements the error interface.
func (e *NotLoadedError) Error() string {
	return "ent: " + e.edge + " edge was not loaded"
}

// IsNotLoaded returns a boolean indicating whether the error is a not loaded error.
func IsNotLoaded(err error) bool {
	if err == nil {
		return false
	}
	var e *NotLoadedError
	return errors.As(err, &e)
}

// ConstraintError returns when trying to create/update one or more entities and
// one or more of their constraints failed. For example, violation of edge or
// field uniqueness.
type ConstraintError struct {
	msg  string
	wrap error
}

// Error implements the error interface.
func (e ConstraintError) Error() string {
	return "ent: constraint failed: " + e.msg
}

// Unwrap implements the errors.Wrapper interface.
func (e *ConstraintError) Unwrap() error {
	return e.wrap
}

// IsConstraintError returns a boolean indicating whether the error is a constraint failure.
func IsConstraintError(err error) bool {
	if err == nil {
		return false
	}
	var e *ConstraintError
	return errors.As(err, &e)
}

// Code implements the dsl.Node interface.
func (e ConstraintError) Code() (string, []interface{}) {
	return strconv.Quote(e.prefix() + e.msg), nil
}

func (e *ConstraintError) UnmarshalGraphson(b []byte) error {
	var v [1]*string
	if err := graphson.Unmarshal(b, &v); err != nil {
		return err
	}
	if v[0] == nil {
		return fmt.Errorf("ent: missing string value")
	}
	if !strings.HasPrefix(*v[0], e.prefix()) {
		return fmt.Errorf("ent: invalid string for error: %s", *v[0])
	}
	e.msg = strings.TrimPrefix(*v[0], e.prefix())
	return nil
}

// prefix returns the prefix used for gremlin constants.
func (ConstraintError) prefix() string { return "Error: " }

// NewErrUniqueField creates a constraint error for unique fields.
func NewErrUniqueField(label, field string, v interface{}) *ConstraintError {
	return &ConstraintError{msg: fmt.Sprintf("field %s.%s with value: %#v", label, field, v)}
}

// NewErrUniqueEdge creates a constraint error for unique edges.
func NewErrUniqueEdge(label, edge, id string) *ConstraintError {
	return &ConstraintError{msg: fmt.Sprintf("edge %s.%s with id: %#v", label, edge, id)}
}

// isConstantError indicates if the given response holds a gremlin constant containing an error.
func isConstantError(r *gremlin.Response) (*ConstraintError, bool) {
	e := &ConstraintError{}
	if err := graphson.Unmarshal(r.Result.Data, e); err != nil {
		return nil, false
	}
	return e, true
}
