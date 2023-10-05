package neo4j

import (
	"github.com/aliworkshop/configer"
	"github.com/aliworkshop/dbcore"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNeo_CommitTransaction(t *testing.T) {
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

	q := NewQuery()
	err = db.StartTransaction(q)
	assert.Nil(t, err)

	item, err := db.Get(q.WithModelFunc(func() dbcore.Modeler {
		return new(User)
	}).WithFilter("first_name", "John"))
	item.(*User).Age = 45

	err = db.Update(q.WithModelFunc(func() dbcore.Modeler {
		return new(User)
	}).WithFilter("first_name", "John").WithBody(item))
	assert.Nil(t, err)

	err = db.CommitTransaction(q)
	assert.Nil(t, err)
}
func TestNeo_RollbackTransaction(t *testing.T) {
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

	q := NewQuery()
	err = db.StartTransaction(q)
	assert.Nil(t, err)

	item, err := db.Get(q.WithModelFunc(func() dbcore.Modeler {
		return new(User)
	}).WithFilter("first_name", "John"))
	item.(*User).Age = 35

	err = db.Update(q.WithModelFunc(func() dbcore.Modeler {
		return new(User)
	}).WithFilter("first_name", "John").WithBody(item))
	assert.Nil(t, err)

	err = db.RollbackTransaction(q)
	assert.Nil(t, err)
}
