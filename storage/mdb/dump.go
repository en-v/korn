package mdb

import (
	"context"

	"github.com/en-v/korn/inset"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (self *MDB) Dump(collectionName string, obj inset.InsetInterface) error {
	col := self.cols[collectionName]
	if col == nil {
		return errors.New("Mongo collection not found, name = " + collectionName)
	}

	res, err := col.UpdateByID(context.TODO(), obj.GetId(), bson.D{{"$set", obj}})
	if err != nil {
		return errors.Wrap(err, "Update By ID")
	}

	if res.MatchedCount == 0 {
		_, err := col.InsertOne(context.TODO(), obj)
		if err != nil {
			return errors.Wrap(err, "Insert One")
		}
		return nil
	}

	if res.MatchedCount == 1 {
		return nil
	}

	return errors.New("Something went wrong")
}
