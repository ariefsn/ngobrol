package mongo

import (
	"context"
	"time"

	"github.com/ariefsn/ngobrol/entities"
	"github.com/ariefsn/ngobrol/helper"
	"github.com/ariefsn/ngobrol/logger"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoUserRepository struct {
	Db  *mongo.Database
	Rdb *redis.Client
}

// GetByIDs implements entities.UserRepository.
func (r *mongoUserRepository) GetByIDs(ctx context.Context, ids []string) ([]*entities.UserData, error) {
	single := entities.UserData{}
	result := []*entities.UserData{}
	cur, err := r.Db.Collection(single.TableName()).Find(ctx, bson.M{"_id": bson.M{"$in": ids}})

	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	for cur.Next(ctx) {
		var row entities.UserData

		err = cur.Decode(&row)
		if err != nil {
			break
		}

		result = append(result, &row)
	}

	return result, nil
}

// Create implements entities.UserRepository.
func (r *mongoUserRepository) Create(ctx context.Context, payload *entities.UserData) (*entities.UserData, error) {
	_, err := r.Db.Collection(payload.TableName()).InsertOne(ctx, payload)

	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	return payload, nil
}

// Delete implements entities.UserRepository.
func (r *mongoUserRepository) Delete(ctx context.Context, id string) error {
	m := entities.UserData{}
	res := r.Db.Collection(m.TableName()).FindOneAndDelete(ctx, bson.M{"_id": id})

	if res.Err() != nil {
		err := helper.ParseMongoError(res.Err())
		logger.Error(err)
	}

	return helper.ParseMongoError(res.Err())
}

// Gets implements entities.UserRepository.
func (r *mongoUserRepository) Gets(ctx context.Context, filter interface{}, skip int64, limit int64) ([]*entities.UserData, int64, error) {
	m := entities.UserData{}
	result := []*entities.UserData{}
	coll := r.Db.Collection(m.TableName())

	if filter == nil {
		filter = entities.M{}
	}

	_bson, _ := helper.ToBsonM(filter)

	filterBson := bson.M{}

	for k, v := range _bson {
	SWITCH:
		switch t := v.(type) {
		case string:
			_f := helper.MongoFilter(helper.FoContains, k, t)
			filterBson[k] = _f[k]
			break SWITCH
		default:
			filterBson[k] = v
		}
	}

	count, err := coll.CountDocuments(ctx, filterBson)

	if err != nil && err != mongo.ErrNilDocument {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return result, 0, e
	}

	pipe := helper.MongoPipe(helper.MongoAggregate{
		Match: filterBson,
		Sort:  nil,
		Skip:  &skip,
		Limit: &limit,
	})

	cur, err := coll.Aggregate(ctx, pipe)

	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return result, 0, e
	}

	for cur.Next(ctx) {
		var row entities.UserData

		err = cur.Decode(&row)
		if err != nil {
			break
		}

		result = append(result, &row)
	}

	return result, count, nil
}

// GetByID implements entities.UserRepository.
func (r *mongoUserRepository) GetByID(ctx context.Context, id string) (*entities.UserData, error) {
	result := entities.UserData{}
	err := r.Db.Collection(result.TableName()).FindOne(ctx, bson.M{"_id": id}).Decode(&result)

	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	return &result, nil
}

// GetByID implements entities.UserRepository.
func (r *mongoUserRepository) GetByEmail(ctx context.Context, email string) (*entities.UserData, error) {
	result := entities.UserData{}
	err := r.Db.Collection(result.TableName()).FindOne(ctx, bson.M{"email": email}).Decode(&result)

	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	return &result, nil
}

func (r *mongoUserRepository) update(ctx context.Context, filter bson.M, payload bson.M) *mongo.SingleResult {
	m := entities.UserData{}
	upsert := true
	returnDoc := options.After

	return r.Db.Collection(m.TableName()).FindOneAndUpdate(ctx, filter, bson.M{
		"$set": payload,
	}, &options.FindOneAndUpdateOptions{
		ReturnDocument: &returnDoc,
		Upsert:         &upsert,
	})
}

// Update implements entities.UserRepository.
func (r *mongoUserRepository) Update(ctx context.Context, id string, payload *entities.UserData) (*entities.UserData, error) {
	var data entities.UserData

	res := r.update(ctx, bson.M{"_id": id}, bson.M{
		"firstName": payload.FirstName,
		"lastName":  payload.LastName,
		"email":     payload.Email,
		"image":     payload.Image,
		"updatedAt": time.Now(),
	})

	if res.Err() != nil {
		err := helper.ParseMongoError(res.Err())
		logger.Error(err)
		return nil, err
	}

	res.Decode(&data)

	return &data, nil
}

// NewMongoUserRepository will create an object that represent the User.Repository interface
func NewMongoUserRepository(database *mongo.Database) entities.UserRepository {
	return &mongoUserRepository{
		Db: database,
	}
}
