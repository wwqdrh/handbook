package mongodb

import (
	"context"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Item represents an item at
// a shop
type Item struct {
	Name  string
	Price int64
}

// Storage is our storage interface
// We'll implement it with Mongo
// storage
type Storage interface {
	GetByName(context.Context, string) (*Item, error)
	Put(context.Context, *Item) error
}

// MongoStorage implements our storage interface
type MongoStorage struct {
	*mgo.Session
	DB         string
	Collection string
}

// NewMongoStorage initializes a MongoStorage
func NewMongoStorage(connection, db, collection string) (*MongoStorage, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	ms := MongoStorage{
		Session:    session,
		DB:         db,
		Collection: collection,
	}
	return &ms, nil
}

// GetByName queries mongodb for an item with
// the correct name
func (m *MongoStorage) GetByName(ctx context.Context, name string) (*Item, error) {
	c := m.Session.DB(m.DB).C(m.Collection)
	var i Item
	if err := c.Find(bson.M{"name": name}).One(&i); err != nil {
		return nil, err
	}

	return &i, nil
}

// Put adds an item to our mongo instance
func (m *MongoStorage) Put(ctx context.Context, i *Item) error {
	c := m.Session.DB(m.DB).C(m.Collection)
	return c.Insert(i)
}
