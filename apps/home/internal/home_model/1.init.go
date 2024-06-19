package home_model

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

var db *mongo.Database
var logger *slog.Logger

// Init 初始化db
func Init(_db *mongo.Database, _logger *slog.Logger) {
	db = _db
	logger = _logger
}
