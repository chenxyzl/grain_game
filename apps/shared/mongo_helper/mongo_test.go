package mongo_helper

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
	"testing"
)

func TestMongo(t *testing.T) {
	// Rest of the code will go here
	// Set _client options 设置连接参数
	//clientOptions := options.Client().ApplyURI("mongodb://root:xxxxxxxxxxxxxxxxxxxxxxx@172.20.52.158:41134/?connect=direct;authSource=admin")
	url := "mongodb+srv://ichenzhl:<password>@cluster0.feqwf3z.mongodb.net/?retryWrites=true&w=majority"

	err := Init(url, "test", 10, slog.Default())
	if err != nil {
		t.Error(err)
	}
	defer Close()

	filter := bson.D{{"_id", 1}}
	type TestStruct struct {
		Id   uint64 `bson:"_id,omitempty"`
		Name string `bson:"name"`
	}
	var res = TestStruct{}
	err = GetColByName("test").FindOne(context.Background(), filter).Decode(&res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			res.Id = 1
			res.Name = "aaa"
			_, err = GetColByName("test").InsertOne(context.TODO(), res)
			if err != nil {
				t.Error(err)
			}
			err = GetColByName("test").FindOne(context.TODO(), filter).Decode(&res)
			if err != nil {
				t.Error(err)
			}
			if res.Name != "aaa" {
				t.Error()
			}
		} else {
			t.Error(err)
		}
	}
	fmt.Println(res)

	data, err := Transaction(func() (interface{}, error) {
		res := &TestStruct{}
		err = GetColByName("test").FindOne(context.Background(), filter).Decode(&res)
		if err != nil {
			t.Error(err)
		}
		res.Name = "bbb"
		_, err = GetColByName("test").ReplaceOne(context.TODO(), filter, res)
		if err != nil {
			t.Error(err)
		}
		err = GetColByName("test").FindOne(context.TODO(), filter).Decode(&res)
		if err != nil {
			t.Error(err)
		}
		if res.Name != "bbb" {
			t.Error()
		}
		return res, nil
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}
