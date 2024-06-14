package model

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (player *Player) Init(db *mongo.Database) error {
	filter := bson.D{{"_id", player.uid}}
	err := db.Collection("player").FindOne(context.TODO(), filter).Decode(player)
	if err != nil {
		return errors.Join(err, fmt.Errorf("load player base data err, uid:%d", player.uid))
	}
	return nil
}