package neo4j

import (
	"github.com/aliworkshop/error"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"strings"
)

func (n *neo) Exec(query Query) error.ErrorModel {
	session, err := n.db.Session(neo4j.AccessModeWrite)
	if err != nil {
		return error.Internal(err)
	}
	defer session.Close()
	q := strings.ReplaceAll(query.GetQuery(), "\n", " ")
	q = strings.ReplaceAll(q, "\t", "")

	var result neo4j.Result
	if tx := n.getTx(query); tx != nil {
		result, err = tx.Run(q, query.GetParams())
	} else {
		result, err = session.Run(q, query.GetParams())
	}
	if err != nil {
		return error.Internal(err)
	}
	if result.Err() != nil {
		return error.Internal(result.Err())
	}
	return nil
}
