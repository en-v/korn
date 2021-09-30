package korn

import (
	"errors"
	"time"

	"github.com/en-v/korn/holder"
	"github.com/en-v/log"
)

type Inset struct {
	holder  holder.IHolder
	Id      string    `bson:"_id" json:"id" required:"true"`
	Updated time.Time `bson:"_updated" json:"_updated" required:"true"`
}

func (self *Inset) Commit() error {

	self.Updated = time.Now()

	if self.holder == nil {
		return errors.New("Holder Is Null.")
	}

	err := self.holder.Work(self.Id)
	if err != nil {
		return err
	}

	return nil
}

func (self *Inset) Link(hldr interface{}) {
	self.holder = hldr.(holder.IHolder)
}

func (self *Inset) SetId(id string) {
	self.Id = id
}

func (self *Inset) GetId() string {
	return self.Id
}

func (self *Inset) Clone() interface{} {
	clone := *self
	log.Debug(&self)
	return &clone
}

/*
// Custom JSON marshallind and unmarshalling

func (self *Inset) MarshalJSON() ([]byte, error) {
	return []byte("\"" + self.GetId() + "\""), nil
}

func (self *Inset) UnmarshalJSON(data []byte) error {

	d, err := hex.DecodeString(string(data[1 : len(data)-1]))
	if err != nil {
		log.Error(err)
		return err
	}
	self.Id = string(d)
	return nil
}

// Custom BSON marshallind and unmarshalling

func (self *Inset) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(self.GetId())
}

func (self *Inset) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	d, err := hex.DecodeString(string(data[4 : len(data)-1]))
	if err != nil {
		log.Error(err)
		return err
	}
	self.Id = string(d)
	return nil
}
*/
