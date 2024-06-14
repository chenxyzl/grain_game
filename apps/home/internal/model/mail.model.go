package model

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
	mailIds []uint64 `bson:"mail_ids"` //右键id-客户端更具邮件id来再次请求邮件内容-避免登录时候db压力集中,分散登录返回包体的大小
}