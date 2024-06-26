package mongo

import (
	"context"

	"github.com/ariefsn/ngobrol/entities"
	"github.com/ariefsn/ngobrol/helper"
	"github.com/ariefsn/ngobrol/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoMessageRepository struct {
	Db       *mongo.Database
	RoomRepo entities.RoomRepository
}

// Create implements entities.MessageRepository.
func (r *mongoMessageRepository) Create(ctx context.Context, payload *entities.MessageData) (*entities.MessageData, error) {
	_, err := r.Db.Collection(payload.TableName()).InsertOne(ctx, payload)

	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	return payload, nil
}

// Delete implements entities.MessageRepository.
func (r *mongoMessageRepository) Delete(ctx context.Context, id string) error {
	m := entities.MessageData{}
	res := r.Db.Collection(m.TableName()).FindOneAndDelete(ctx, bson.M{"_id": id})

	if res.Err() != nil {
		err := helper.ParseMongoError(res.Err())
		logger.Error(err)
	}

	return helper.ParseMongoError(res.Err())
}

// Gets implements entities.MessageRepository.
func (r *mongoMessageRepository) Gets(ctx context.Context, filter interface{}, skip int64, limit int64) ([]*entities.MessageData, int64, error) {
	m := entities.MessageData{}
	result := []*entities.MessageData{}
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
		Sort:  []helper.MongoSort{{SortField: "createdAt", SortBy: helper.SortByDesc}},
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
		var row entities.MessageData

		err = cur.Decode(&row)
		if err != nil {
			break
		}

		result = append(result, &row)
	}

	return result, count, nil
}

// GetByID implements entities.MessageRepository.
func (r *mongoMessageRepository) GetByID(ctx context.Context, id string) (*entities.MessageData, error) {
	result := entities.MessageData{}
	err := r.Db.Collection(result.TableName()).FindOne(ctx, bson.M{"_id": id}).Decode(&result)

	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	return &result, nil
}

// GetByCode implements entities.MessageRepository.
func (r *mongoMessageRepository) GetByCode(ctx context.Context, code string) (*entities.MessageData, error) {
	result := entities.MessageData{}
	err := r.Db.Collection(result.TableName()).FindOne(ctx, bson.M{"code": code}).Decode(&result)

	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	return &result, nil
}

func (r *mongoMessageRepository) update(ctx context.Context, filter bson.M, payload bson.M) *mongo.SingleResult {
	m := entities.MessageData{}
	upsert := true
	returnDoc := options.After

	return r.Db.Collection(m.TableName()).FindOneAndUpdate(ctx, filter, bson.M{
		"$set": payload,
	}, &options.FindOneAndUpdateOptions{
		ReturnDocument: &returnDoc,
		Upsert:         &upsert,
	})
}

// Update implements entities.MessageRepository.
// func (r *mongoMessageRepository) Update(ctx context.Context, id string, payload *entities.MessageData) (*entities.MessageData, error) {
// 	var data entities.MessageData

// 	res := r.update(ctx, bson.M{"_id": id}, bson.M{
// 		"firstName": payload.FirstName,
// 		"lastName":  payload.LastName,
// 		"email":     payload.Email,
// 		"image":     payload.Image,
// 		"updatedAt": time.Now(),
// 	})

// 	if res.Err() != nil {
// 		err := helper.ParseMongoError(res.Err())
// 		logger.Error(err)
// 		return nil, err
// 	}

// 	res.Decode(&data)

// 	return &data, nil
// }

// NewMongoMessageRepository will create an object that represent the Message.Repository interface
func NewMongoMessageRepository(database *mongo.Database, roomRepo entities.RoomRepository) entities.MessageRepository {
	return &mongoMessageRepository{
		Db:       database,
		RoomRepo: roomRepo,
	}
}
