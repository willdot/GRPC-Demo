package main

import (
	"fmt"

	"github.com/gocql/gocql"
)

// Session is a Cassandra session
var Session *gocql.Session

func init() {
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Port = 9042
	cluster.Keyspace = "shippy"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	fmt.Println("cassandra init done")
}
