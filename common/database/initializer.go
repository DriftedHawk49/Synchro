package database

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

/*
 * This package initilizes a database connection and provides a client for
 * database operations.
 */

func New(connUri string) (*mongo.Client, error) {
	return mongo.Connect(options.Client().ApplyURI(connUri))
}
