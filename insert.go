package neo4j

import (
	"encoding/json"
	"fmt"
	"github.com/aliworkshop/error"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"reflect"
	"strings"
)

func (n *neo) Insert(query Query) error.ErrorModel {
	session, err := n.db.Session(neo4j.AccessModeWrite)
	if err != nil {
		return error.Internal(err)
	}
	defer session.Close()
	bodyMap, params := structToMapParam(query.GetBody())
	marshalledBody, err := json.Marshal(bodyMap)
	if err != nil {
		return error.Internal(err)
	}

	body := string(marshalledBody)
	body = strings.Replace(body, "\"", "", -1)
	val := reflect.TypeOf(query.GetBody())
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	nodeName := val.Name()
	q := fmt.Sprintf("MERGE (n:%s %s) RETURN n", nodeName, body)
	var result neo4j.Result
	if tx := n.getTx(query); tx != nil {
		result, err = tx.Run(q, params)
	} else {
		result, err = session.Run(q, params)
	}
	if err != nil {
		return error.Internal(err)
	}
	if result.Err() != nil {
		return error.Internal(result.Err())
	}
	return nil
}
