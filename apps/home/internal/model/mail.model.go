package model

type Mail struct {
	uid      uint64  `bson:"uid"` //唯一id
	tid      uint64  `bson:"tid"` //模板id
	time     int64   `bson:"time"`
	fromUid  uint64  `bson:"from_uid"`
	fromName string  `bson:"from_name"`
	tile     string  `bson:"tile"`
	content  string  `bson:"content"`
	rewards  []TItem `bson:"rewards"`
	read     bool    `bson:"read"`
	received bool    `bson:"received"`
}
