// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/MONAKA0721/hokkori/ent/category"
	"github.com/MONAKA0721/hokkori/ent/draft"
	"github.com/MONAKA0721/hokkori/ent/hashtag"
	"github.com/MONAKA0721/hokkori/ent/predicate"
	"github.com/MONAKA0721/hokkori/ent/user"
	"github.com/MONAKA0721/hokkori/ent/work"
)

// DraftQuery is the builder for querying Draft entities.
type DraftQuery struct {
	config
	limit        *int
	offset       *int
	unique       *bool
	order        []OrderFunc
	fields       []string
	predicates   []predicate.Draft
	withOwner    *UserQuery
	withHashtags *HashtagQuery
	withWork     *WorkQuery
	withCategory *CategoryQuery
	withFKs      bool
	modifiers    []func(*sql.Selector)
	loadTotal    []func(context.Context, []*Draft) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DraftQuery builder.
func (dq *DraftQuery) Where(ps ...predicate.Draft) *DraftQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit adds a limit step to the query.
func (dq *DraftQuery) Limit(limit int) *DraftQuery {
	dq.limit = &limit
	return dq
}

// Offset adds an offset step to the query.
func (dq *DraftQuery) Offset(offset int) *DraftQuery {
	dq.offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DraftQuery) Unique(unique bool) *DraftQuery {
	dq.unique = &unique
	return dq
}

// Order adds an order step to the query.
func (dq *DraftQuery) Order(o ...OrderFunc) *DraftQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryOwner chains the current query on the "owner" edge.
func (dq *DraftQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(draft.Table, draft.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, draft.OwnerTable, draft.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHashtags chains the current query on the "hashtags" edge.
func (dq *DraftQuery) QueryHashtags() *HashtagQuery {
	query := &HashtagQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(draft.Table, draft.FieldID, selector),
			sqlgraph.To(hashtag.Table, hashtag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, draft.HashtagsTable, draft.HashtagsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWork chains the current query on the "work" edge.
func (dq *DraftQuery) QueryWork() *WorkQuery {
	query := &WorkQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(draft.Table, draft.FieldID, selector),
			sqlgraph.To(work.Table, work.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, draft.WorkTable, draft.WorkColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCategory chains the current query on the "category" edge.
func (dq *DraftQuery) QueryCategory() *CategoryQuery {
	query := &CategoryQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(draft.Table, draft.FieldID, selector),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, draft.CategoryTable, draft.CategoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Draft entity from the query.
// Returns a *NotFoundError when no Draft was found.
func (dq *DraftQuery) First(ctx context.Context) (*Draft, error) {
	nodes, err := dq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{draft.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DraftQuery) FirstX(ctx context.Context) *Draft {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Draft ID from the query.
// Returns a *NotFoundError when no Draft ID was found.
func (dq *DraftQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{draft.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DraftQuery) FirstIDX(ctx context.Context) int {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Draft entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Draft entity is found.
// Returns a *NotFoundError when no Draft entities are found.
func (dq *DraftQuery) Only(ctx context.Context) (*Draft, error) {
	nodes, err := dq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{draft.Label}
	default:
		return nil, &NotSingularError{draft.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DraftQuery) OnlyX(ctx context.Context) *Draft {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Draft ID in the query.
// Returns a *NotSingularError when more than one Draft ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DraftQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{draft.Label}
	default:
		err = &NotSingularError{draft.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DraftQuery) OnlyIDX(ctx context.Context) int {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Drafts.
func (dq *DraftQuery) All(ctx context.Context) ([]*Draft, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dq *DraftQuery) AllX(ctx context.Context) []*Draft {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Draft IDs.
func (dq *DraftQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dq.Select(draft.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DraftQuery) IDsX(ctx context.Context) []int {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DraftQuery) Count(ctx context.Context) (int, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DraftQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DraftQuery) Exist(ctx context.Context) (bool, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DraftQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DraftQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DraftQuery) Clone() *DraftQuery {
	if dq == nil {
		return nil
	}
	return &DraftQuery{
		config:       dq.config,
		limit:        dq.limit,
		offset:       dq.offset,
		order:        append([]OrderFunc{}, dq.order...),
		predicates:   append([]predicate.Draft{}, dq.predicates...),
		withOwner:    dq.withOwner.Clone(),
		withHashtags: dq.withHashtags.Clone(),
		withWork:     dq.withWork.Clone(),
		withCategory: dq.withCategory.Clone(),
		// clone intermediate query.
		sql:    dq.sql.Clone(),
		path:   dq.path,
		unique: dq.unique,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DraftQuery) WithOwner(opts ...func(*UserQuery)) *DraftQuery {
	query := &UserQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withOwner = query
	return dq
}

// WithHashtags tells the query-builder to eager-load the nodes that are connected to
// the "hashtags" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DraftQuery) WithHashtags(opts ...func(*HashtagQuery)) *DraftQuery {
	query := &HashtagQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withHashtags = query
	return dq
}

// WithWork tells the query-builder to eager-load the nodes that are connected to
// the "work" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DraftQuery) WithWork(opts ...func(*WorkQuery)) *DraftQuery {
	query := &WorkQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withWork = query
	return dq
}

// WithCategory tells the query-builder to eager-load the nodes that are connected to
// the "category" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DraftQuery) WithCategory(opts ...func(*CategoryQuery)) *DraftQuery {
	query := &CategoryQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withCategory = query
	return dq
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
//	client.Draft.Query().
//		GroupBy(draft.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DraftQuery) GroupBy(field string, fields ...string) *DraftGroupBy {
	grbuild := &DraftGroupBy{config: dq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dq.sqlQuery(ctx), nil
	}
	grbuild.label = draft.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
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
//	client.Draft.Query().
//		Select(draft.FieldCreateTime).
//		Scan(ctx, &v)
func (dq *DraftQuery) Select(fields ...string) *DraftSelect {
	dq.fields = append(dq.fields, fields...)
	selbuild := &DraftSelect{DraftQuery: dq}
	selbuild.label = draft.Label
	selbuild.flds, selbuild.scan = &dq.fields, selbuild.Scan
	return selbuild
}

func (dq *DraftQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dq.fields {
		if !draft.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DraftQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Draft, error) {
	var (
		nodes       = []*Draft{}
		withFKs     = dq.withFKs
		_spec       = dq.querySpec()
		loadedTypes = [4]bool{
			dq.withOwner != nil,
			dq.withHashtags != nil,
			dq.withWork != nil,
			dq.withCategory != nil,
		}
	)
	if dq.withOwner != nil || dq.withWork != nil || dq.withCategory != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, draft.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Draft).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Draft{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withOwner; query != nil {
		if err := dq.loadOwner(ctx, query, nodes, nil,
			func(n *Draft, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	if query := dq.withHashtags; query != nil {
		if err := dq.loadHashtags(ctx, query, nodes,
			func(n *Draft) { n.Edges.Hashtags = []*Hashtag{} },
			func(n *Draft, e *Hashtag) { n.Edges.Hashtags = append(n.Edges.Hashtags, e) }); err != nil {
			return nil, err
		}
	}
	if query := dq.withWork; query != nil {
		if err := dq.loadWork(ctx, query, nodes, nil,
			func(n *Draft, e *Work) { n.Edges.Work = e }); err != nil {
			return nil, err
		}
	}
	if query := dq.withCategory; query != nil {
		if err := dq.loadCategory(ctx, query, nodes, nil,
			func(n *Draft, e *Category) { n.Edges.Category = e }); err != nil {
			return nil, err
		}
	}
	for i := range dq.loadTotal {
		if err := dq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DraftQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*Draft, init func(*Draft), assign func(*Draft, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Draft)
	for i := range nodes {
		if nodes[i].user_drafts == nil {
			continue
		}
		fk := *nodes[i].user_drafts
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_drafts" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (dq *DraftQuery) loadHashtags(ctx context.Context, query *HashtagQuery, nodes []*Draft, init func(*Draft), assign func(*Draft, *Hashtag)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Draft)
	nids := make(map[int]map[*Draft]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(draft.HashtagsTable)
		s.Join(joinT).On(s.C(hashtag.FieldID), joinT.C(draft.HashtagsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(draft.HashtagsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(draft.HashtagsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]interface{}, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]interface{}{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []interface{}) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Draft]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "hashtags" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (dq *DraftQuery) loadWork(ctx context.Context, query *WorkQuery, nodes []*Draft, init func(*Draft), assign func(*Draft, *Work)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Draft)
	for i := range nodes {
		if nodes[i].work_drafts == nil {
			continue
		}
		fk := *nodes[i].work_drafts
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(work.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "work_drafts" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (dq *DraftQuery) loadCategory(ctx context.Context, query *CategoryQuery, nodes []*Draft, init func(*Draft), assign func(*Draft, *Category)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Draft)
	for i := range nodes {
		if nodes[i].draft_category == nil {
			continue
		}
		fk := *nodes[i].draft_category
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(category.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "draft_category" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dq *DraftQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	_spec.Node.Columns = dq.fields
	if len(dq.fields) > 0 {
		_spec.Unique = dq.unique != nil && *dq.unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DraftQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dq *DraftQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   draft.Table,
			Columns: draft.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: draft.FieldID,
			},
		},
		From:   dq.sql,
		Unique: true,
	}
	if unique := dq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, draft.FieldID)
		for i := range fields {
			if fields[i] != draft.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DraftQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(draft.Table)
	columns := dq.fields
	if len(columns) == 0 {
		columns = draft.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.unique != nil && *dq.unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DraftGroupBy is the group-by builder for Draft entities.
type DraftGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DraftGroupBy) Aggregate(fns ...AggregateFunc) *DraftGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dgb *DraftGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dgb.path(ctx)
	if err != nil {
		return err
	}
	dgb.sql = query
	return dgb.sqlScan(ctx, v)
}

func (dgb *DraftGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dgb.fields {
		if !draft.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dgb *DraftGroupBy) sqlQuery() *sql.Selector {
	selector := dgb.sql.Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dgb.fields)+len(dgb.fns))
		for _, f := range dgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dgb.fields...)...)
}

// DraftSelect is the builder for selecting fields of Draft entities.
type DraftSelect struct {
	*DraftQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DraftSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	ds.sql = ds.DraftQuery.sqlQuery(ctx)
	return ds.sqlScan(ctx, v)
}

func (ds *DraftSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ds.sql.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
