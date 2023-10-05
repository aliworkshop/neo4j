package neo4j

import (
	"fmt"
	"github.com/aliworkshop/error"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"reflect"
)

func (n *neo) Update(query Query) error.ErrorModel {
	session, e := n.db.Session(neo4j.AccessModeWrite)
	if e != nil {
		return error.Internal(e)
	}
	defer session.Close()
	bodyMap, params := structToMapParam(query.GetBody())
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
	q += " SET "
	for k, v := range bodyMap {
		q += fmt.Sprintf("n.%s = %s, ", k, v)
	}
	q = q[:len(q)-2]
	q += fmt.Sprint(" RETURN n")
	var result neo4j.Result
	if tx := n.getTx(query); tx != nil {
		result, e = tx.Run(q, params)
	} else {
		result, e = session.Run(q, params)
	}
	if e != nil {
		return error.Internal(e)
	}
	if result.Err() != nil {
		return error.Internal(result.Err())
	}
	return nil
}
