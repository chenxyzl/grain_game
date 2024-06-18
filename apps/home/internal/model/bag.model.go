package model

import "github.com/chenxyzl/gsgen/gsmodel"

type Bag struct {
	gsmodel.DirtyModel `bson:"-"`
	items              *gsmodel.DMap[uint64, *AItem] `bson:"items"`
}
