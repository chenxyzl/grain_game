package home_model

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

const maxBatchLoadMails = 100

// LoadMailsByUids 批量获取邮件
// 一次最多批量获取100封邮件
func LoadMailsByUids(uids []uint64) ([]*Mail, error) {
	if len(uids) == 0 {
		return nil, nil
	}
	if len(uids) > maxBatchLoadMails {
		logger.Warn("batch load mails by uids len too big", "max", maxBatchLoadMails, "now", len(uids))
		uids = uids[:maxBatchLoadMails]
	}
	filter := bson.M{"_id": bson.M{"$in": uids}}
	cursor, err := db.Collection("mail").Find(context.Background(), filter)
	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("batch load mails by id fail, uids:%v", uids))
	}
	var mails []*Mail
	// 遍历结果
	for cursor.Next(context.Background()) {
		var mail = &Mail{}
		if err := cursor.Decode(&mail); err != nil {
			return nil, fmt.Errorf("batch load mails, unmarshal mail from db err, uids:%v, err:%v", uids, err)
		}
		//
		mails = append(mails, mail)
	}
	if err = cursor.Err(); err != nil {
		return nil, fmt.Errorf("batch load mails, cursor err, uids:%v|err:%v", uids, err)
	}
	return mails, nil
}
