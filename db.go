package neo4j

import (
	"github.com/aliworkshop/configer"
	"github.com/aliworkshop/error"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type neo struct {
	cfg config
	db  neo4j.Driver
}

func NewNeo4jRepository(configRegistry configer.Registry) Repository {
	n := new(neo)
	err := configRegistry.Unmarshal(&n.cfg)
	if err != nil {
		panic(err)
	}

	return n
}

func (n *neo) DB() any {
	return n.db
}

func (n *neo) Initialize() error.ErrorModel {
	driver, err := neo4j.NewDriver(n.cfg.Address, neo4j.BasicAuth(n.cfg.Username, n.cfg.Password, ""), func(config *neo4j.Config) {
		config.Encrypted = n.cfg.Tls
		config.MaxConnectionPoolSize = n.cfg.MaxConnectionPoolSize
	})
	if err != nil {
		return error.Internal(err)
	}

	n.db = driver
	return nil
}
