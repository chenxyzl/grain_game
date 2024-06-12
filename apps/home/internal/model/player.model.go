package model

type Game struct {
	bag *Bag `bson:"bag"`
}

type Player struct {
	uid  uint64 `bson:"_id"`
	base *Base  `bson:"base"`
	game *Game  `bson:"game"`
}
