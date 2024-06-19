package home_model

import "github.com/chenxyzl/gsgen/gsmodel"

type Base struct {
	gsmodel.DirtyModel `bson:"-"`
	name               string `bson:"name"`
}
