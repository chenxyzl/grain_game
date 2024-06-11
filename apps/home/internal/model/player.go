package model

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Game struct {
	bag *Bag `bson:"bag"`
}
type Player struct {
	base *Base
	game *Game
}

func (player *Player) Init(db *mongo.Database, uid uint64) error {
	filter := bson.D{{"_id", uid}}
	player.base = &Base{}
	err := db.Collection("player_base").FindOne(context.TODO(), filter).Decode(player.base)
	if err != nil {
		return errors.Join(err, fmt.Errorf("load player base data err, uid:%d", uid))
	}
	player.game = &Game{}
	err = db.Collection("player_game").FindOne(context.TODO(), filter).Decode(player.game)
	if err != nil {
		return errors.Join(err, fmt.Errorf("load player game data err, uid:%d", uid))
	}
	return nil
}
