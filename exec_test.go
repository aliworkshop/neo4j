package neo4j

import (
	"github.com/aliworkshop/configer"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNeo_Follow(t *testing.T) {
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

	err = db.Exec(NewQuery().WithQuery(`
	MATCH (u:User), (p:User) WHERE u.Name=$Name1 and p.first_name=$Name2
	MERGE (u)-[:FOLLOWS]->(p)`).
		WithParams("Name1", "ali").
		WithParams("Name2", "John"))
	assert.Nil(t, err)
}
