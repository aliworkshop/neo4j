package neo4j

import (
	"github.com/aliworkshop/error"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func (n *neo) getTx(query Query) neo4j.Transaction {
	if query == nil {
		return nil
	}
	iTx := query.GetTransaction()
	if iTx != nil {
		tx := iTx.(neo4j.Transaction)
		return tx
	}
	return nil
}

func (n *neo) StartTransaction(query Query) (err error.ErrorModel) {
	session, e := n.db.Session(neo4j.AccessModeWrite)
	if e != nil {
		return error.Internal(e)
	}
	tx := n.getTx(query)
	if tx == nil {
		tx, e = session.BeginTransaction()
		if e != nil {
			return error.Internal(e)
		}
	}
	query.SetTransaction(tx)
	return nil
}

func (n *neo) CommitTransaction(query Query) (err error.ErrorModel) {
	tx := n.getTx(query)
	if tx == nil {
		return
	}
	defer tx.Close()
	return error.HandleError(tx.Commit())
}

func (n *neo) RollbackTransaction(query Query) (err error.ErrorModel) {
	tx := n.getTx(query)
	if tx == nil {
		return
	}
	defer tx.Close()
	return error.HandleError(tx.Rollback())
}
