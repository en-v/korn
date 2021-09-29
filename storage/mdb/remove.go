package mdb

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (self *MDB) Remove(collectionName string, id string) error {
	col := self.cols[collectionName]
	if col == nil {
		return errors.New("Mongo collection not found, name = " + collectionName)
	}
	
	res, err := col.DeleteOne(context.TODO(), bson.M{"_id": id}, nil)
	if err != nil {
		return errors.Wrap(err, "Delete One From Collection")
	}

	if res.DeletedCount != 1 {
		return errors.New("Object is not really deleted")
	}

	return nil
}
