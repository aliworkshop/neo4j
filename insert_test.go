package neo4j

import (
	"github.com/aliworkshop/configer"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type User struct {
	Id        uint64 `json:"id"`
	Uuid      string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Mobile    string `json:"mobile"`
}

func (User) TableName() string {
	return "users"
}

func TestNeo_Insert(t *testing.T) {
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

	uuid := uuid.New().String()
	u := &User{
		Id:        13,
		Uuid:      uuid,
		FirstName: "John",
		LastName:  "Cena",
		Age:       40,
		Mobile:    "09123456789",
	}

	err = db.Insert(NewQuery().WithBody(u))
	assert.Nil(t, err)
}
