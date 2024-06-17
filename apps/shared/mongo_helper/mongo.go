package mongo_helper

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"log/slog"
	"sync/atomic"
)

var _init int32 = 0
var _url, dbName string
var _client *mongo.Client
var _db *mongo.Database
var _logger *slog.Logger

func GetColByName(colName string) *mongo.Collection {
	if _db == nil {
		panic(fmt.Sprintf("mongo dn is nil, colName:%v", colName))
	}
	return _db.Collection(colName)
}

func Init(url string, dbName string, poolSize uint64, logger *slog.Logger) error {
	if atomic.CompareAndSwapInt32(&_init, 0, 1) {
		return fmt.Errorf("repeated init mongo err, url:%v", _url)
	}
	// Rest of the code will go here
	// Set _client options 设置连接参数
	//clientOptions := options.Client().ApplyURI("mongodb://root:xxxxxxxxxxxxxxxxxxxxxxx@172.20.52.158:41134/?connect=direct;authSource=admin")
	clientOptions := options.Client().ApplyURI(url).SetMaxPoolSize(poolSize)

	// Connect to MongoDB 连接数据库
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return errors.Join(err, fmt.Errorf("connect to mongo err, url:%v", url))
	}
	// Check the connection 测试连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return errors.Join(err, fmt.Errorf("ping to mongo connect err, url:%v", url))
	}
	_client = client
	_db = client.Database(dbName)
	_url = url
	dbName = dbName
	_logger = logger

	_logger.Warn("connect to mongo success", "url", url, "dbName", dbName)
	return nil
}

func Close() {
	if !atomic.CompareAndSwapInt32(&_init, 1, 0) {
		return
	}
	if _client != nil {
		return
	}
	_logger.Warn("mongo stop...")
	err := _client.Disconnect(context.TODO())
	if err != nil {
		_logger.Warn("mongo stop success", "url", _url, "dbName", dbName)
	} else {
		_logger.Error("mongo stop err", "url", _url, "dbName", dbName, "err", err)
	}
	_db = nil
	_client = nil
}

func Transaction(f func() (interface{}, error)) (interface{}, error) {
	if _client == nil {
		return nil, fmt.Errorf("mongo client not connect")
	}
	//
	opts := options.Transaction().SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	// Start a transaction
	session, err := _client.StartSession()
	if err != nil {
		_logger.Error("Error starting a session:", err)
		return nil, err
	}
	defer session.EndSession(context.Background())
	// Define the transaction callback function
	callback := func(sessionContext mongo.SessionContext) (interface{}, error) {
		return f()
	}
	// Execute the transaction
	return session.WithTransaction(context.Background(), callback, opts)
}
