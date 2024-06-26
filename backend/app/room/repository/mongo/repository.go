package mongo

import (
	"context"
	"time"

	"github.com/ariefsn/ngobrol/entities"
	"github.com/ariefsn/ngobrol/helper"
	"github.com/ariefsn/ngobrol/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRoomRepository struct {
	Db       *mongo.Database
	UserRepo entities.UserRepository
}

// Create implements entities.RoomRepository.
func (r *mongoRoomRepository) Create(ctx context.Context, payload *entities.RoomData) (*entities.RoomDataDetails, error) {
	_, err := r.Db.Collection(payload.TableName()).InsertOne(ctx, payload)

	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	return &entities.RoomDataDetails{}, nil
}

// Delete implements entities.RoomRepository.
func (r *mongoRoomRepository) Delete(ctx context.Context, id string) error {
	m := entities.RoomData{}
	res := r.Db.Collection(m.TableName()).FindOneAndDelete(ctx, bson.M{"_id": id})

	if res.Err() != nil {
		err := helper.ParseMongoError(res.Err())
		logger.Error(err)
	}

	return helper.ParseMongoError(res.Err())
}

// Gets implements entities.RoomRepository.
func (r *mongoRoomRepository) Gets(ctx context.Context, filter interface{}, skip int64, limit int64) ([]*entities.RoomData, int64, error) {
	m := entities.RoomData{}
	result := []*entities.RoomData{}
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
		var row entities.RoomData

		err = cur.Decode(&row)
		if err != nil {
			break
		}

		result = append(result, &row)
	}

	return result, count, nil
}

// GetByID implements entities.RoomRepository.
func (r *mongoRoomRepository) GetByID(ctx context.Context, id string) (*entities.RoomDataDetails, error) {
	result := entities.RoomDataDetails{}
	err := r.Db.Collection(result.TableName()).FindOne(ctx, bson.M{"_id": id}).Decode(&result)

	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	return &result, nil
}

// GetByCode implements entities.RoomRepository.
func (r *mongoRoomRepository) GetByCode(ctx context.Context, code string) (*entities.RoomDataDetails, error) {
	result := entities.RoomData{}
	err := r.Db.Collection(result.TableName()).FindOne(ctx, bson.M{"code": code}).Decode(&result)

	if err != nil {
		e := helper.ParseMongoError(err)
		logger.Error(e)
		return nil, e
	}

	resultDetails := entities.RoomDataDetails{
		Id:    result.Id,
		Code:  result.Code,
		Image: result.Image,
		Audit: result.Audit,
	}

	users, _ := r.UserRepo.GetByIDs(ctx, result.Users)
	resultDetails.Users = users

	return &resultDetails, nil
}

func (r *mongoRoomRepository) update(ctx context.Context, filter bson.M, payload bson.M) *mongo.SingleResult {
	m := entities.RoomData{}
	upsert := true
	returnDoc := options.After

	return r.Db.Collection(m.TableName()).FindOneAndUpdate(ctx, filter, bson.M{
		"$set": payload,
	}, &options.FindOneAndUpdateOptions{
		ReturnDocument: &returnDoc,
		Upsert:         &upsert,
	})
}

// Update implements entities.RoomRepository.
func (r *mongoRoomRepository) Update(ctx context.Context, id string, payload *entities.RoomData) (*entities.RoomData, error) {
	var data entities.RoomData

	res := r.update(ctx, bson.M{"_id": id}, bson.M{
		"code":      payload.Code,
		"image":     payload.Image,
		"users":     payload.Users,
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

// NewMongoRoomRepository will create an object that represent the Room.Repository interface
func NewMongoRoomRepository(database *mongo.Database, userRepo entities.UserRepository) entities.RoomRepository {
	return &mongoRoomRepository{
		Db:       database,
		UserRepo: userRepo,
	}
}
