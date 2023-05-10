package main

import "fmt"

type IDBConnection interface {
	Connect()
}

type DBConnection struct {
	Db IDBConnection
}

func (conn *DBConnection) Connect() {
	conn.Db.Connect()
}

type MySqlConnection struct {
	ConnectionString string
}

func (conn *MySqlConnection) Connect() {
	fmt.Println("MySql " + conn.ConnectionString)
}

func NewMySqlConnection(connectionString string) *MySqlConnection {
	return &MySqlConnection{ConnectionString: connectionString}
}

type PostgresConnection struct {
	ConnectionString string
}

func (conn *PostgresConnection) Connect() {
	fmt.Println("Postgres " + conn.ConnectionString)
}

type MongoDBConnection struct {
	ConnectionString string
}

func (conn *MongoDBConnection) Connect() {
	fmt.Println("MongoDB " + conn.ConnectionString)
}

func main() {
	mysql := NewMySqlConnection("MySql connection string")
	connMySql := DBConnection{Db: mysql}
	connMySql.Connect()

	pg := &PostgresConnection{ConnectionString: "Postgres connection string"}
	connPostgres := DBConnection{Db: pg}
	connPostgres.Connect()

	mongo := &MongoDBConnection{ConnectionString: "MongoDB connection string"}
	connMongoDb := DBConnection{Db: mongo}
	connMongoDb.Connect()
}
