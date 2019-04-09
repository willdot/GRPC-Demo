package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gocql/gocql"
)

// Session is a Cassandra session
var Session *gocql.Session

func init() {

	host := os.Getenv("DB_HOST")

	if host == "" {
		host = "127.0.0.1"
	}

	portString := os.Getenv("DB_PORT")

	var port int
	if portString != "" {
		port, _ = strconv.Atoi(portString)
	}
	var err error
	cluster := gocql.NewCluster(host)
	//cluster.ProtoVersion = 4
	cluster.Keyspace = "shippy"
	cluster.Port = port

	cluster.ConnectTimeout = time.Second * 10
	cluster.DisableInitialHostLookup = true

	fmt.Println("Connecting now")
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	fmt.Println(cluster.Port)
	fmt.Println("cassandra init done")

	keySpaceMeta, _ := Session.KeyspaceMetadata("shippy")

	if _, exists := keySpaceMeta.Tables["user"]; exists != true {
		Session.Query("CREATE TABLE user (id UUID, name text, email text, password text, company text, PRIMARY KEY(id))").Exec()
		Session.Query("create index UserEmailIndex on user(email)").Exec()
	}
}
