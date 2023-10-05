package neo4j

import (
	"github.com/aliworkshop/error"
)

type Repository interface {
	Initialize() error.ErrorModel
	DB() any
	Close() error.ErrorModel
	Ping() error.ErrorModel

	Count(query Query) (count uint64, err error.ErrorModel)
	List(query Query) (items any, err error.ErrorModel)
	Get(query Query) (item any, err error.ErrorModel)
	Exist(query Query) (exists bool, err error.ErrorModel)

	Insert(query Query) error.ErrorModel
	Update(query Query) error.ErrorModel
	Delete(query Query) error.ErrorModel
	Exec(query Query) error.ErrorModel

	StartTransaction(query Query) (err error.ErrorModel)
	CommitTransaction(query Query) (err error.ErrorModel)
	RollbackTransaction(query Query) (err error.ErrorModel)
}
