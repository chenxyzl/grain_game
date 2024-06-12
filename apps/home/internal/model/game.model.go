package model

type Game struct {
	bag   *Bag   `bson:"bag"`
	mails *Mails `bson:"mails"`
}
