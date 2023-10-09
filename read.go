package neo4j

import (
	"encoding/json"
	"fmt"
	"github.com/aliworkshop/error"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"reflect"
	"strings"
)

func (n *neo) Count(query Query) (count uint64, err error.ErrorModel) {
	session, e := n.db.Session(neo4j.AccessModeRead)
	if e != nil {
		return 0, error.Internal(e)
	}
	defer session.Close()
	if len(query.GetQuery()) > 0 {
		result, ee := session.Run(query.GetQuery(), query.GetParams())
		if ee != nil {
			return 0, error.Internal(ee)
		}
		for result.Next() {
			record := result.Record()
			if c, ok := record.GetByIndex(0).(int64); ok {
				count = uint64(c)
			}
		}
		return count, nil
	}
	_, params := structToMapParam(query.GetFilter())
	typ := reflect.TypeOf(query.GetModel())
	nodeName := typ.Elem().Name()
	q := fmt.Sprintf("MATCH (n:%s)", nodeName)
	if len(query.GetFilter()) > 0 {
		q += " WHERE"
	}
	for k := range query.GetFilter() {
		q += fmt.Sprintf(" n.%s = $%s AND", k, k)
	}
	if len(query.GetFilter()) > 0 {
		q = q[:len(q)-4]
	}
	q += fmt.Sprint(" RETURN COUNT(n)")
	result, ee := session.Run(q, params)
	if ee != nil {
		return 0, error.Internal(ee)
	}
	if result.Err() != nil {
		return 0, error.Internal(result.Err())
	}
	for result.Next() {
		record := result.Record()
		if c, ok := record.GetByIndex(0).(int64); ok {
			count = uint64(c)
		}
	}
	return count, nil
}

func (n *neo) List(query Query) (items any, err error.ErrorModel) {
	session, e := n.db.Session(neo4j.AccessModeRead)
	if e != nil {
		return nil, error.Internal(e)
	}
	defer session.Close()
	offset := (query.GetPage() - 1) * query.GetPageSize()
	if len(query.GetQuery()) > 0 {
		q := strings.ReplaceAll(query.GetQuery(), "\n", " ")
		q = strings.ReplaceAll(q, "\t", "")
		sort, err := n.sort(query, query.GetPrefix())
		if err != nil {
			return nil, err
		}
		if len(sort) > 0 {
			q += fmt.Sprintf(" ORDER BY %s", sort)
		}
		q += fmt.Sprintf(" SKIP %d LIMIT %d", offset, query.GetPageSize())
		result, ee := session.Run(query.GetQuery(), query.GetParams())
		if ee != nil {
			return nil, error.Internal(ee)
		}
		typ := reflect.TypeOf(query.GetModel())
		slice := reflect.New(reflect.SliceOf(typ)).Elem()
		for result.Next() {
			record := result.Record()
			node := record.GetByIndex(0)
			if v, ok := node.(neo4j.Node); ok {
				b, e := json.Marshal(v.Props())
				if e != nil {
					return nil, error.Internal(e)
				}
				elm := reflect.New(typ.Elem())
				e = json.Unmarshal(b, elm.Interface())
				if e != nil {
					return nil, error.HandleError(e)
				}
				slice = reflect.Append(slice, elm)
			}
		}
		return slice.Interface(), nil
	}
	_, params := structToMapParam(query.GetFilter())
	typ := reflect.TypeOf(query.GetModel())
	nodeName := typ.Elem().Name()
	q := fmt.Sprintf("MATCH (n:%s)", nodeName)
	if len(query.GetFilter()) > 0 {
		q += " WHERE"
	}
	for k := range query.GetFilter() {
		q += fmt.Sprintf(" n.%s = $%s AND", k, k)
	}
	if len(query.GetFilter()) > 0 {
		q = q[:len(q)-4]
	}
	q += fmt.Sprintf(" RETURN n")
	sort, err := n.sort(query, "n")
	if err != nil {
		return nil, err
	}
	if len(sort) > 0 {
		q += fmt.Sprintf(" ORDER BY %s", sort)
	}
	q += fmt.Sprintf(" SKIP %d LIMIT %d", offset, query.GetPageSize())
	result, ee := session.Run(q, params)
	if ee != nil {
		return nil, error.Internal(ee)
	}
	if result.Err() != nil {
		return nil, error.Internal(result.Err())
	}
	slice := reflect.New(reflect.SliceOf(typ)).Elem()
	for result.Next() {
		record := result.Record()
		node := record.GetByIndex(0)
		if v, ok := node.(neo4j.Node); ok {
			b, e := json.Marshal(v.Props())
			if e != nil {
				return nil, error.Internal(e)
			}
			elm := reflect.New(typ.Elem())
			e = json.Unmarshal(b, elm.Interface())
			if e != nil {
				return nil, error.HandleError(e)
			}
			slice = reflect.Append(slice, elm)
		}
	}
	return slice.Interface(), nil
}

func (n *neo) Get(query Query) (item any, err error.ErrorModel) {
	session, e := n.db.Session(neo4j.AccessModeRead)
	if e != nil {
		return nil, error.Internal(e)
	}
	defer session.Close()
	if len(query.GetQuery()) > 0 {
		result, ee := session.Run(query.GetQuery(), query.GetParams())
		if ee != nil {
			return nil, error.Internal(ee)
		}
		typ := reflect.TypeOf(query.GetModel())
		elm := reflect.ValueOf(query.GetModel())
		if elm.Kind() != reflect.Ptr {
			elm = reflect.New(typ.Elem())
		}
		for result.Next() {
			record := result.Record()
			node := record.GetByIndex(0)
			if v, ok := node.(neo4j.Node); ok {
				b, e := json.Marshal(v.Props())
				if e != nil {
					return nil, error.Internal(e)
				}
				e = json.Unmarshal(b, elm.Interface())
				if e != nil {
					return nil, error.HandleError(e)
				}
			}
		}
		return elm.Interface(), nil
	}
	_, params := structToMapParam(query.GetFilter())
	typ := reflect.TypeOf(query.GetModel())
	nodeName := typ.Elem().Name()
	q := fmt.Sprintf("MATCH (n:%s)", nodeName)
	if len(query.GetFilter()) > 0 {
		q += " WHERE"
	}
	for k := range query.GetFilter() {
		q += fmt.Sprintf(" n.%s = $%s AND", k, k)
	}
	if len(query.GetFilter()) > 0 {
		q = q[:len(q)-4]
	}
	q += fmt.Sprintf(" RETURN n LIMIT 1")
	result, ee := session.Run(q, params)
	if ee != nil {
		return nil, error.Internal(ee)
	}
	if result.Err() != nil {
		return nil, error.Internal(result.Err())
	}
	elm := reflect.ValueOf(query.GetModel())
	if elm.Kind() != reflect.Ptr {
		elm = reflect.New(typ.Elem())
	}
	for result.Next() {
		record := result.Record()
		node := record.GetByIndex(0)
		if v, ok := node.(neo4j.Node); ok {
			b, e := json.Marshal(v.Props())
			if e != nil {
				return nil, error.Internal(e)
			}
			e = json.Unmarshal(b, elm.Interface())
			if e != nil {
				return nil, error.HandleError(e)
			}
		}
	}
	return elm.Interface(), nil
}

func (n *neo) Exist(query Query) (bool, error.ErrorModel) {
	count, err := n.Count(query)
	return count > 0, err
}
