# Go Neo4j Integration Package

This is a Go package for seamless integration with Neo4j, a popular graph database. With this package, you can easily connect to a Neo4j database, execute queries, transactions, and interact with the graph data in your Go applications.

## Features

- Connect to Neo4j databases using the official Neo4j Go driver.
- Execute Cypher queries and retrieve results.
- Manage transactions for data consistency.
- Simplify CRUD operations on nodes and relationships.
- Error handling and robust connection management.
## Installation

To use this package, you need to have Go installed. You can install it via:
```
go get -v github.com/aliworkshop/neo4j
```
## Usage

Here's a quick guide on how to use the package in your Go application
```
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aliworkshop/configer"
	"github.com/aliworkshop/dbcore"
	"github.com/aliworkshop/neo4j"
)

type User struct {
	Uuid      string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Mobile    string `json:"mobile"`
        Active    bool   `json:"active"`
}

func (User) TableName() string {
	return "users"
}

func main() {
        // Load Config
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

	// Initialize Neo4j connection
	db := neo4j.NewNeo4jRepository(registry)

	// Connect to the database
	err = db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Execute a query
	items, err := db.List(neo4j.NewQuery().WithModelFunc(func() dbcore.Modeler {
		return new(User)
	}).WithFilter("active", true))

	// Fetch the query result
	for _, item := range items.([]*User) {
		fmt.Println(item)
	}

	// Start a transaction
	q := neo4j.NewQuery()
	err = db.StartTransaction(q)
	if err != nil {
		log.Fatal(err)
	}

        // Fetch Data and change it
        item, err := db.Get(q.WithModelFunc(func() dbcore.Modeler {
		return new(User)
	}).WithFilter("first_name", "John"))
	item.(*User).Age = 45

	// Execute a transactional query
	err = db.Update(q.WithModelFunc(func() dbcore.Modeler {
		return new(User)
	}).WithFilter("first_name", "John").WithBody(item))
	if err != nil {
		log.Fatal(err)
	}

	// Commit the transaction
	err = db.CommitTransaction(q)
	if err != nil {
		log.Fatal(err)
	}
}
```

Please refer to the package documentation and examples for more detailed usage instructions.

## Contribution

Contributions are welcome! If you have suggestions, feature requests, or found a bug, please open an issue or submit a pull request.

## License

This package is licensed under the MIT License.

