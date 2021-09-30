package mdb

import (
	"context"
	"reflect"

	"github.com/en-v/korn/inset"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (self *MDB) Restore(collectionName string, reft reflect.Type) (map[string]interface{}, error) {
	col := self.cols[collectionName]
	if col == nil {
		return nil, errors.New("Mongo collection not found, name = " + collectionName)
	}

	cnt, err := col.CountDocuments(context.TODO(), bson.M{}, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Count Request")
	}

	cursor, err := col.Find(context.TODO(), bson.M{}, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Fing Request")
	}

	res := make(map[string]interface{}, int(cnt))
	ctx := context.TODO()

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		item := reflect.New(reft).Interface()
		
		err = cursor.Decode(item)
		if err != nil {
			return nil, errors.Wrap(err, "Decode")
		}

		iset, cast := item.(inset.InsetInterface)
		if !cast {
			return nil, errors.New("Restored object cant to be casted to InsetInterface")
		}
		res[iset.GetId()] = item

	}
	return res, nil
}
