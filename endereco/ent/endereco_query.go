// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"endereco/ent/cep"
	"endereco/ent/endereco"
	"endereco/ent/predicate"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EnderecoQuery is the builder for querying Endereco entities.
type EnderecoQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Endereco
	// eager-loading edges.
	withCep *CepQuery
	withFKs bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EnderecoQuery builder.
func (eq *EnderecoQuery) Where(ps ...predicate.Endereco) *EnderecoQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit adds a limit step to the query.
func (eq *EnderecoQuery) Limit(limit int) *EnderecoQuery {
	eq.limit = &limit
	return eq
}

// Offset adds an offset step to the query.
func (eq *EnderecoQuery) Offset(offset int) *EnderecoQuery {
	eq.offset = &offset
	return eq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eq *EnderecoQuery) Unique(unique bool) *EnderecoQuery {
	eq.unique = &unique
	return eq
}

// Order adds an order step to the query.
func (eq *EnderecoQuery) Order(o ...OrderFunc) *EnderecoQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// QueryCep chains the current query on the "cep" edge.
func (eq *EnderecoQuery) QueryCep() *CepQuery {
	query := &CepQuery{config: eq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(endereco.Table, endereco.FieldID, selector),
			sqlgraph.To(cep.Table, cep.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, endereco.CepTable, endereco.CepColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Endereco entity from the query.
// Returns a *NotFoundError when no Endereco was found.
func (eq *EnderecoQuery) First(ctx context.Context) (*Endereco, error) {
	nodes, err := eq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{endereco.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *EnderecoQuery) FirstX(ctx context.Context) *Endereco {
	node, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Endereco ID from the query.
// Returns a *NotFoundError when no Endereco ID was found.
func (eq *EnderecoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{endereco.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eq *EnderecoQuery) FirstIDX(ctx context.Context) int {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Endereco entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Endereco entity is not found.
// Returns a *NotFoundError when no Endereco entities are found.
func (eq *EnderecoQuery) Only(ctx context.Context) (*Endereco, error) {
	nodes, err := eq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{endereco.Label}
	default:
		return nil, &NotSingularError{endereco.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *EnderecoQuery) OnlyX(ctx context.Context) *Endereco {
	node, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Endereco ID in the query.
// Returns a *NotSingularError when exactly one Endereco ID is not found.
// Returns a *NotFoundError when no entities are found.
func (eq *EnderecoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{endereco.Label}
	default:
		err = &NotSingularError{endereco.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *EnderecoQuery) OnlyIDX(ctx context.Context) int {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Enderecos.
func (eq *EnderecoQuery) All(ctx context.Context) ([]*Endereco, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return eq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (eq *EnderecoQuery) AllX(ctx context.Context) []*Endereco {
	nodes, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Endereco IDs.
func (eq *EnderecoQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := eq.Select(endereco.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *EnderecoQuery) IDsX(ctx context.Context) []int {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *EnderecoQuery) Count(ctx context.Context) (int, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return eq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (eq *EnderecoQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *EnderecoQuery) Exist(ctx context.Context) (bool, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return eq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *EnderecoQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EnderecoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *EnderecoQuery) Clone() *EnderecoQuery {
	if eq == nil {
		return nil
	}
	return &EnderecoQuery{
		config:     eq.config,
		limit:      eq.limit,
		offset:     eq.offset,
		order:      append([]OrderFunc{}, eq.order...),
		predicates: append([]predicate.Endereco{}, eq.predicates...),
		withCep:    eq.withCep.Clone(),
		// clone intermediate query.
		sql:  eq.sql.Clone(),
		path: eq.path,
	}
}

// WithCep tells the query-builder to eager-load the nodes that are connected to
// the "cep" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EnderecoQuery) WithCep(opts ...func(*CepQuery)) *EnderecoQuery {
	query := &CepQuery{config: eq.config}
	for _, opt := range opts {
		opt(query)
	}
	eq.withCep = query
	return eq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Endereco.Query().
//		GroupBy(endereco.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (eq *EnderecoQuery) GroupBy(field string, fields ...string) *EnderecoGroupBy {
	group := &EnderecoGroupBy{config: eq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return eq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.Endereco.Query().
//		Select(endereco.FieldCreateTime).
//		Scan(ctx, &v)
//
func (eq *EnderecoQuery) Select(fields ...string) *EnderecoSelect {
	eq.fields = append(eq.fields, fields...)
	return &EnderecoSelect{EnderecoQuery: eq}
}

func (eq *EnderecoQuery) prepareQuery(ctx context.Context) error {
	for _, f := range eq.fields {
		if !endereco.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *EnderecoQuery) sqlAll(ctx context.Context) ([]*Endereco, error) {
	var (
		nodes       = []*Endereco{}
		withFKs     = eq.withFKs
		_spec       = eq.querySpec()
		loadedTypes = [1]bool{
			eq.withCep != nil,
		}
	)
	if eq.withCep != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, endereco.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Endereco{config: eq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := eq.withCep; query != nil {
		ids := make([]int32, 0, len(nodes))
		nodeids := make(map[int32][]*Endereco)
		for i := range nodes {
			if nodes[i].cep_enderecos == nil {
				continue
			}
			fk := *nodes[i].cep_enderecos
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(cep.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "cep_enderecos" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Cep = n
			}
		}
	}

	return nodes, nil
}

func (eq *EnderecoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	_spec.Node.Columns = eq.fields
	if len(eq.fields) > 0 {
		_spec.Unique = eq.unique != nil && *eq.unique
	}
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *EnderecoQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := eq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (eq *EnderecoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   endereco.Table,
			Columns: endereco.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: endereco.FieldID,
			},
		},
		From:   eq.sql,
		Unique: true,
	}
	if unique := eq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := eq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, endereco.FieldID)
		for i := range fields {
			if fields[i] != endereco.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eq *EnderecoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(endereco.Table)
	columns := eq.fields
	if len(columns) == 0 {
		columns = endereco.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eq.unique != nil && *eq.unique {
		selector.Distinct()
	}
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector)
	}
	if offset := eq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EnderecoGroupBy is the group-by builder for Endereco entities.
type EnderecoGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *EnderecoGroupBy) Aggregate(fns ...AggregateFunc) *EnderecoGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the group-by query and scans the result into the given value.
func (egb *EnderecoGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := egb.path(ctx)
	if err != nil {
		return err
	}
	egb.sql = query
	return egb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (egb *EnderecoGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := egb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (egb *EnderecoGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(egb.fields) > 1 {
		return nil, errors.New("ent: EnderecoGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := egb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (egb *EnderecoGroupBy) StringsX(ctx context.Context) []string {
	v, err := egb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (egb *EnderecoGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = egb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{endereco.Label}
	default:
		err = fmt.Errorf("ent: EnderecoGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (egb *EnderecoGroupBy) StringX(ctx context.Context) string {
	v, err := egb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (egb *EnderecoGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(egb.fields) > 1 {
		return nil, errors.New("ent: EnderecoGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := egb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (egb *EnderecoGroupBy) IntsX(ctx context.Context) []int {
	v, err := egb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (egb *EnderecoGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = egb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{endereco.Label}
	default:
		err = fmt.Errorf("ent: EnderecoGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (egb *EnderecoGroupBy) IntX(ctx context.Context) int {
	v, err := egb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (egb *EnderecoGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(egb.fields) > 1 {
		return nil, errors.New("ent: EnderecoGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := egb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (egb *EnderecoGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := egb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (egb *EnderecoGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = egb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{endereco.Label}
	default:
		err = fmt.Errorf("ent: EnderecoGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (egb *EnderecoGroupBy) Float64X(ctx context.Context) float64 {
	v, err := egb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (egb *EnderecoGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(egb.fields) > 1 {
		return nil, errors.New("ent: EnderecoGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := egb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (egb *EnderecoGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := egb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (egb *EnderecoGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = egb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{endereco.Label}
	default:
		err = fmt.Errorf("ent: EnderecoGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (egb *EnderecoGroupBy) BoolX(ctx context.Context) bool {
	v, err := egb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (egb *EnderecoGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range egb.fields {
		if !endereco.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := egb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := egb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (egb *EnderecoGroupBy) sqlQuery() *sql.Selector {
	selector := egb.sql.Select()
	aggregation := make([]string, 0, len(egb.fns))
	for _, fn := range egb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(egb.fields)+len(egb.fns))
		for _, f := range egb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(egb.fields...)...)
}

// EnderecoSelect is the builder for selecting fields of Endereco entities.
type EnderecoSelect struct {
	*EnderecoQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (es *EnderecoSelect) Scan(ctx context.Context, v interface{}) error {
	if err := es.prepareQuery(ctx); err != nil {
		return err
	}
	es.sql = es.EnderecoQuery.sqlQuery(ctx)
	return es.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (es *EnderecoSelect) ScanX(ctx context.Context, v interface{}) {
	if err := es.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (es *EnderecoSelect) Strings(ctx context.Context) ([]string, error) {
	if len(es.fields) > 1 {
		return nil, errors.New("ent: EnderecoSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := es.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (es *EnderecoSelect) StringsX(ctx context.Context) []string {
	v, err := es.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (es *EnderecoSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = es.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{endereco.Label}
	default:
		err = fmt.Errorf("ent: EnderecoSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (es *EnderecoSelect) StringX(ctx context.Context) string {
	v, err := es.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (es *EnderecoSelect) Ints(ctx context.Context) ([]int, error) {
	if len(es.fields) > 1 {
		return nil, errors.New("ent: EnderecoSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := es.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (es *EnderecoSelect) IntsX(ctx context.Context) []int {
	v, err := es.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (es *EnderecoSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = es.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{endereco.Label}
	default:
		err = fmt.Errorf("ent: EnderecoSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (es *EnderecoSelect) IntX(ctx context.Context) int {
	v, err := es.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (es *EnderecoSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(es.fields) > 1 {
		return nil, errors.New("ent: EnderecoSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := es.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (es *EnderecoSelect) Float64sX(ctx context.Context) []float64 {
	v, err := es.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (es *EnderecoSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = es.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{endereco.Label}
	default:
		err = fmt.Errorf("ent: EnderecoSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (es *EnderecoSelect) Float64X(ctx context.Context) float64 {
	v, err := es.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (es *EnderecoSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(es.fields) > 1 {
		return nil, errors.New("ent: EnderecoSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := es.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (es *EnderecoSelect) BoolsX(ctx context.Context) []bool {
	v, err := es.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (es *EnderecoSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = es.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{endereco.Label}
	default:
		err = fmt.Errorf("ent: EnderecoSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (es *EnderecoSelect) BoolX(ctx context.Context) bool {
	v, err := es.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (es *EnderecoSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := es.sql.Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
