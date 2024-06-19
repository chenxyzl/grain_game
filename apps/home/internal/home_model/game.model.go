package home_model

import (
	"github.com/chenxyzl/gsgen/gsmodel"
	"grain_game/apps/common_model"
)

type Game struct {
	gsmodel.DirtyModel `bson:"-"`
	bag                *Bag     `bson:"bag"`
	mailbox            *Mailbox `bson:"mailbox"`
}

// Bag 背包
type Bag struct {
	gsmodel.DirtyModel `bson:"-"`
	items              *gsmodel.DMap[uint64, *common_model.AItem] `bson:"items"`
}

// Mail 邮件内容
type Mail struct {
	gsmodel.DirtyModel `bson:"-"`
	uid                uint64                              `bson:"_id"`    //唯一id
	tid                uint64                              `bson:"tid"`    //模板编号，决定以下参数如何使用
	toUid              uint64                              `bson:"to_uid"` //是否需要？
	fromUid            uint64                              `bson:"from_uid"`
	fromName           string                              `bson:"from_name"`
	title              string                              `bson:"title"`
	content            string                              `bson:"content"`
	rewards            *gsmodel.DList[*common_model.AItem] `bson:"rewards"`
	params             *gsmodel.DMap[string, string]       `bson:"params"`
	createTime         int64                               `bson:"create_time"`
	readTime           bool                                `bson:"read_time"`    //存在则表示-已读
	getTime            bool                                `bson:"get_time"`     //存在则表示-已领取
	deletedTime        bool                                `bson:"deleted_time"` //存在则表示-已删除
}

// Mailbox 邮箱
type Mailbox struct {
	gsmodel.DirtyModel `bson:"-"`
	mailIds            *gsmodel.DList[uint64] `bson:"mail_ids"`              //右键id-客户端更具邮件id来再次请求邮件内容-避免登录时候db压力集中,分散登录返回包体的大小
	lastGlobalMailTime uint64                 `bson:"last_global_mail_time"` //上一次的全局邮件自增id
}
