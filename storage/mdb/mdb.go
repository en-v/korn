package mdb

import (
	"context"
	"time"

	"github.com/en-v/log"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MDB struct {
	ready  bool
	client *mongo.Client
	db     *mongo.Database
	cols   map[string]*mongo.Collection
}

func Make(dbname string, connectionStr string) (*MDB, error) {
	uri := connectionStr + "/" + dbname
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.Wrap(err, "Make MDB")
	}

	db := client.Database(dbname)

	pingctx, pingcf := context.WithTimeout(context.TODO(), time.Second*5)
	defer pingcf()
	err = client.Ping(pingctx, readpref.Primary())
	if err != nil {
		log.Error(err)
		return nil, errors.Wrap(err, "Ping MongoDb")
	}

	return &MDB{
		db:     db,
		client: client,
		cols:   map[string]*mongo.Collection{},
		ready:  true,
	}, nil
}

func (self *MDB) Shutdown() {
	self.cols = map[string]*mongo.Collection{}
	err := self.client.Disconnect(context.TODO())
	if err != nil {
		log.Error(err)
	}
}

func (self *MDB) IsReady() bool {
	return self.ready
}

func (self *MDB) Prepare(colName string) error {
	self.cols[colName] = self.db.Collection(colName, nil)
	return nil
}
