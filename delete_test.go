package neo4j

import (
	"github.com/aliworkshop/configer"
	"github.com/aliworkshop/dbcore"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNeo_Delete(t *testing.T) {
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
	defer db.Close()

	err = db.Delete(NewQuery().WithModelFunc(func() dbcore.Modeler {
		return new(User)
	}).WithFilter("FirstName", "John"))
	assert.Nil(t, err)
}
