package helper

import (
	"context"
	"fmt"
	"time"

	"github.com/ariefsn/ngobrol/entities"
	"github.com/ariefsn/ngobrol/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoClient(address string) (client *mongo.Client, cancel context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))
	if err != nil {
		logger.Fatal(err, entities.M{
			"func": "helper.MongoClient",
		})
	}

	return
}

func RedisClient(env envDb) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", env.Host, env.Port),
		Password: env.Password,
		DB:       env.DbIndex,
	})

	ping := rdb.Ping(context.Background())

	if ping.Err() != nil {
		logger.Fatal(ping.Err(), entities.M{
			"func": "helper.RedisClient",
		})
	}

	return rdb
}
