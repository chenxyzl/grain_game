// Code generated by https://github.com/chenxyzl/gsgen; DO NOT EDIT.
// gen_tools version: 1.1.1
// generate time: 2024-06-19 12:13:21
package model

import (
	"github.com/chenxyzl/gsgen/gsmodel"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Game) MarshalBSON() ([]byte, error) {
	var doc = bson.M{"bag": s.bag, "mails": s.mails}
	return bson.Marshal(doc)
}
func (s *Game) UnmarshalBSON(data []byte) error {
	doc := struct {
		Bag   *Bag   `bson:"bag"`
		Mails *Mails `bson:"mails"`
	}{}
	if err := bson.Unmarshal(data, &doc); err != nil {
		return err
	}
	s.SetBag(doc.Bag)
	s.SetMails(doc.Mails)
	return nil
}
func (s *Game) BuildBson(m bson.M, preKey string) {
	dirty := s.GetDirty()
	if dirty == 0 {
		return
	}
	if dirty&(1<<0) != 0 {
		if s.bag == nil {
			gsmodel.AddUnsetDirtyM(m, gsmodel.MakeBsonKey("bag", preKey))
		} else {
			s.bag.BuildBson(m, gsmodel.MakeBsonKey("bag", preKey))
		}
	}
	if dirty&(1<<1) != 0 {
		if s.mails == nil {
			gsmodel.AddUnsetDirtyM(m, gsmodel.MakeBsonKey("mails", preKey))
		} else {
			s.mails.BuildBson(m, gsmodel.MakeBsonKey("mails", preKey))
		}
	}
	return
}
func (s *Game) CleanDirty() {
	s.DirtyModel.CleanDirty()
	if s.bag != nil {
		s.bag.CleanDirty()
	}
	if s.mails != nil {
		s.mails.CleanDirty()
	}
}