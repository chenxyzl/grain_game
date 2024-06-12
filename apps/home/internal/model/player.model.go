package model

type Player struct {
	uid  uint64 `bson:"_id"`
	base *Base  `bson:"base"`
	game *Game  `bson:"game"`
}
