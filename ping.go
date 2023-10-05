package neo4j

import (
	"github.com/aliworkshop/error"
)

func (n *neo) Ping() error.ErrorModel {
	return error.HandleError(n.db.VerifyConnectivity())
}
