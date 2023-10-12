package neo4j

import (
	"fmt"
	"github.com/aliworkshop/configer"
	"github.com/aliworkshop/dbcore"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNeo_Get(t *testing.T) {
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
	err = db.Connect()
	assert.Nil(t, err)
	defer db.Close()

	item, err := db.Get(NewQuery().WithModelFunc(func() dbcore.Modeler {
		return new(User)
	}))
	assert.Nil(t, err)
	assert.NotNil(t, item)
}

func TestNeo_Count(t *testing.T) {
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
	err = db.Connect()
	assert.Nil(t, err)
	defer db.Close()

	count, err := db.Count(NewQuery().WithModelFunc(func() dbcore.Modeler {
		return new(User)
	}))
	assert.Nil(t, err)
	assert.NotNil(t, count)
}

func TestNeo_Followings(t *testing.T) {
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
	err = db.Connect()
	assert.Nil(t, err)
	defer db.Close()

	items, err := db.List(NewQuery().WithModelFunc(func() dbcore.Modeler {
		return new(User)
	}).WithQuery(`MATCH(u:User)-[:FOLLOWS]->(n:User) WHERE u.Name=$Name RETURN n`).
		WithCountQuery(`MATCH(u:User)-[:FOLLOWS]->(n:User) WHERE u.Name=$Name RETURN COUNT(n)`).
		WithParams("Name", "ali").
		WithSorts(dbcore.SortItem{Field: "id", Order: dbcore.DESC}).
		WithPage(1).WithPageSize(15))
	assert.Nil(t, err)
	for _, item := range items.([]*User) {
		fmt.Println(item)
	}
}

func TestNeo_Followers(t *testing.T) {
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
	err = db.Connect()
	assert.Nil(t, err)
	defer db.Close()

	items, err := db.List(NewQuery().WithModelFunc(func() dbcore.Modeler {
		return new(User)
	}).WithQuery(`
	MATCH(u:User)<-[:FOLLOWS]-(n:User) WHERE u.first_name=$Name RETURN n`).WithParams("Name", "John"))
	assert.Nil(t, err)
	for _, item := range items.([]*User) {
		fmt.Println(item)
	}
}
