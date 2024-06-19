package home_model

import "github.com/chenxyzl/gsgen/gsmodel"

type Player struct {
	gsmodel.DirtyModel `bson:"-"`
	uid                uint64 `bson:"_id"`
	base               *Base  `bson:"base"`
	game               *Game  `bson:"game"`
}
