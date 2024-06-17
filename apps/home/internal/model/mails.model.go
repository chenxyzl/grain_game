package model

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

type Mail struct {
	uid         uint64            `bson:"_id"`    //唯一id
	tid         uint64            `bson:"tid"`    //模板编号，决定以下参数如何使用
	toUid       uint64            `bson:"to_uid"` //是否需要？
	fromUid     uint64            `bson:"from_uid"`
	fromName    string            `bson:"from_name"`
	title       string            `bson:"title"`
	content     string            `bson:"content"`
	rewards     []TItem           `bson:"rewards"`
	params      map[string]string `bson:"params"`
	createTime  int64             `bson:"create_time"`
	readTime    bool              `bson:"read_time"`    //存在则表示-已读
	getTime     bool              `bson:"get_time"`     //存在则表示-已领取
	deletedTime bool              `bson:"deleted_time"` //存在则表示-已删除
}

type Mails struct {
	mailIds            []uint64 `bson:"mail_ids"`              //右键id-客户端更具邮件id来再次请求邮件内容-避免登录时候db压力集中,分散登录返回包体的大小
	lastGlobalMailTime uint64   `bson:"last_global_mail_time"` //上一次的全局邮件自增id
}

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
