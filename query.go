package neo4j

import (
	"github.com/aliworkshop/dbcore"
)

type Query interface {
	GetSort() (sort []dbcore.SortItem)
	SetBody(body any)
	GetBody() (body any)
	GetModel() (instance dbcore.Modeler)
	SetModelFunc(func() dbcore.Modeler)
	SetTransaction(transaction any)
	GetTransaction() (transaction any)
	SetPageSize(pageSize int)
	GetPageSize() int
	SetPage(page int)
	GetPage() (page int)
	GetFilter() map[string]any
	SetQuery(query string)
	GetQuery() string
	GetCountQuery() string
	GetParams() map[string]any
	GetPrefix() string

	WithModelFunc(func() dbcore.Modeler) Query
	WithBody(body any) Query
	WithPage(page int) Query
	WithPageSize(pageSize int) Query
	WithSorts(sort ...dbcore.SortItem) Query
	WithTransaction(transaction any) Query
	WithQuery(query string) Query
	WithCountQuery(query string) Query
	WithFilter(key string, value any) Query
	WithParams(key string, value any) Query
	WithPrefix(prefix string) Query
}

type ModelFunc func() dbcore.Modeler

type query struct {
	query       string
	countQuery  string
	params      map[string]any
	modelFunc   ModelFunc
	transaction any
	pageSize    int
	page        int
	sortItem    []dbcore.SortItem
	body        any
	filters     map[string]any
	prefix      string
}

func NewQuery() Query {
	return &query{
		params:  make(map[string]any),
		filters: make(map[string]any),
	}
}

func (q *query) GetSort() (sort []dbcore.SortItem) {
	return q.sortItem
}

func (q *query) SetBody(body any) {
	q.body = body
}

func (q *query) GetBody() (body any) {
	return q.body
}

func (q *query) GetModel() (instance dbcore.Modeler) {
	return q.modelFunc()
}

func (q *query) SetModelFunc(f func() dbcore.Modeler) {
	q.modelFunc = f
}

func (q *query) SetTransaction(transaction any) {
	q.transaction = transaction
}

func (q *query) GetTransaction() (transaction any) {
	return q.transaction
}

func (q *query) SetPageSize(pageSize int) {
	q.pageSize = pageSize
}

func (q *query) GetPageSize() int {
	if q.pageSize == 0 {
		q.pageSize = 10
	}
	return q.pageSize
}

func (q *query) SetPage(page int) {
	q.page = page
}

func (q *query) GetPage() (page int) {
	if q.page == 0 {
		q.page = 1
	}
	return q.page
}

func (q *query) GetFilter() map[string]any {
	return q.filters
}

func (q *query) SetQuery(query string) {
	q.query = query
}

func (q *query) GetQuery() string {
	return q.query
}

func (q *query) GetCountQuery() string {
	return q.countQuery
}

func (q *query) GetParams() map[string]any {
	return q.params
}

func (q *query) GetPrefix() string {
	return q.prefix
}

func (q *query) WithModelFunc(f func() dbcore.Modeler) Query {
	q.modelFunc = f
	return q
}

func (q *query) WithBody(body any) Query {
	q.body = body
	return q
}

func (q *query) WithPage(page int) Query {
	q.page = page
	return q
}

func (q *query) WithPageSize(pageSize int) Query {
	q.pageSize = pageSize
	return q
}

func (q *query) WithSorts(sort ...dbcore.SortItem) Query {
	for _, s := range sort {
		q.sortItem = append(q.sortItem, dbcore.SortItem{Field: s.Field, Order: s.Order})
	}
	return q
}

func (q *query) WithTransaction(transaction any) Query {
	q.transaction = transaction
	return q
}

func (q *query) WithQuery(query string) Query {
	q.query = query
	return q
}

func (q *query) WithCountQuery(query string) Query {
	q.countQuery = query
	return q
}

func (q *query) WithFilter(key string, value any) Query {
	q.filters[key] = value
	return q
}

func (q *query) WithParams(key string, value any) Query {
	q.params[key] = value
	return q
}

func (q *query) WithPrefix(prefix string) Query {
	q.prefix = prefix
	return q
}
