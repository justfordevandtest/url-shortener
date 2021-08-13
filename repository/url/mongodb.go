package url

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"shorturl/entity"
	"strings"
)

type MongoDBRepo struct {
	Ctx    context.Context
	Client *mongo.Client
	DB     *mongo.Database
	Coll   *mongo.Collection
	URI    string
	DBName string
}

func New(ctx context.Context, uri string, dbName string, collName string) (repo *MongoDBRepo, err error) {
	fullURI := fmt.Sprintf("%s/%s?authSource=admin", uri, dbName)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fullURI))
	if err != nil {
		return nil, err
	}

	repo = &MongoDBRepo{}
	repo.Ctx = ctx
	repo.URI = uri
	repo.DBName = dbName
	repo.Client = client
	repo.DB = client.Database(dbName)
	repo.Coll = repo.DB.Collection(collName)

	return repo, nil
}

func (repo *MongoDBRepo) List(page int, perPage int, filters map[string]interface{}) (total int, items []entity.URL, err error) {
	opts := repo.buildFindOpts(page, perPage)
	query := repo.buildFindQuery(filters)

	total, err = repo.Count(query)
	if err != nil {
		return 0, nil, err
	}

	cursor, err := repo.Coll.Find(repo.Ctx, query, opts)
	if err != nil {
		return 0, nil, err
	}
	defer func() { _ = cursor.Close(repo.Ctx) }()

	for cursor.Next(repo.Ctx) {
		var item entity.URL
		err = cursor.Decode(&item)
		if err != nil {
			return 0, nil, err
		}
		items = append(items, item)
	}

	return total, items, nil
}

func (repo *MongoDBRepo) Create(ent *entity.URL) (err error) {
	_, err = repo.Coll.InsertOne(repo.Ctx, ent)
	return err
}

func (repo *MongoDBRepo) Read(ID string) (ent *entity.URL, err error) {
	filters := repo.buildFindQuery(map[string]interface{}{"id": ID})
	ent = &entity.URL{}

	err = repo.Coll.FindOne(repo.Ctx, filters).Decode(ent)
	if err != nil {
		return nil, err
	}

	return ent, nil
}

func (repo *MongoDBRepo) IncrHit(ID string) (err error) {
	query := repo.buildFindQuery(map[string]interface{}{"id": ID})
	res, err := repo.Coll.UpdateOne(repo.Ctx, query, bson.M{"$inc": bson.M{"hitcount": 1}})
	if err != nil {
		return err
	}

	if res.MatchedCount < 1 {
		return errors.New("not found")
	}

	return nil
}

func (repo *MongoDBRepo) Delete(ID string) (err error) {
	query := repo.buildFindQuery(map[string]interface{}{"id": ID})
	res, err := repo.Coll.UpdateOne(repo.Ctx, query, bson.M{"$set": bson.M{"expired": 0}})
	if err != nil {
		return err
	}

	if res.MatchedCount < 1 {
		return errors.New("not found")
	}

	return nil
}

func (repo *MongoDBRepo) Count(filters bson.M) (total int, err error) {
	cnt, err := repo.Coll.CountDocuments(repo.Ctx, filters)
	if err != nil {
		return 0, err
	}
	return int(cnt), nil
}

func (repo *MongoDBRepo) CountID(id string) (total int, err error) {
	query := repo.buildFindQuery(map[string]interface{}{"id": id})
	cnt, err := repo.Coll.CountDocuments(repo.Ctx, query)
	if err != nil {
		return 0, err
	}
	return int(cnt), nil
}

func (repo *MongoDBRepo) buildFindQuery(filters map[string]interface{}) (query bson.M) {
	if filters == nil {
		return nil
	}
	if len(filters) < 1 {
		return nil
	}

	query = bson.M{}
	for field, val := range filters {
		field = strings.ToLower(field)
		switch field {
		case "url":
			query[field] = bson.M{
				"$regex":   val,
				"$options": "i",
			}
			break
		default:
			query[field] = val
			break
		}
	}

	return query
}

func (repo *MongoDBRepo) buildFindOpts(page int, perPage int) (opts *options.FindOptions) {
	skip := (page - 1) * perPage
	opts = options.Find()
	opts.SetSkip(int64(skip))

	if perPage > 0 {
		opts.SetLimit(int64(perPage))
	}

	return opts
}
