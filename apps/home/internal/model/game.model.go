package model

import "github.com/chenxyzl/gsgen/gsmodel"

type Game struct {
	gsmodel.DirtyModel `bson:"-"`
	bag                *Bag   `bson:"bag"`
	mails              *Mails `bson:"mails"`
}
