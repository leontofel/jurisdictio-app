package model

type Entity interface {
	// Define methods that are common for all entities
	TableName() string
}
