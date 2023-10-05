package neo4j

import (
	"github.com/aliworkshop/configer"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNeo_Ping(t *testing.T) {
	registry := configer.New()
	registry.SetConfigType("yaml")
	f, err := os.Open("./config.sample.yaml")
	if err != nil {
		panic("cannot read config: " + err.Error())
	}
	err = registry.ReadConfig(f)
	if err != nil {
		panic("cannot read config" + err.Error())
	}

	db := NewNeo4jRepository(registry)
	err = db.Initialize()
	assert.Nil(t, err)

	err = db.Ping()
	assert.Nil(t, err)
}