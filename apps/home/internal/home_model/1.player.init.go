package home_model

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func (player *Player) Init() error {
	filter := bson.D{{"_id", player.uid}}
	err := db.Collection("player").FindOne(context.TODO(), filter).Decode(player)
	if err != nil {
		return errors.Join(err, fmt.Errorf("load player base data err, uid:%d", player.uid))
	}
	//如果出现大key,考虑吧数据拆分后组合到model中
	return nil
}
