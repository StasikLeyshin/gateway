package mongo

import (
	"context"
	"errors"
	"gateway/internal/repository/database/mongo/errmongo"
	"gateway/pkg/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

type (
	Config struct {
		ConnString   string
		DatabaseName string
	}

	Client struct {
		log    log.Logger
		client *mongo.Client

		Database *database
	}

	Collection struct {
		collection *mongo.Collection
	}

	database struct {
		db             *mongo.Database
		collectionsMux sync.RWMutex
		collections    map[string]*Collection
	}
)

func NewClientMongo() *Client {
	return &Client{}
}

func (c *Client) Configure(ctx context.Context, config *Config, init bool) error {
	// Set client options
	clientOptions := options.Client().ApplyURI(config.ConnString)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		c.log.WithError(err).Error("failed to connect to mongodb")
		return err
	}

	// Check the connection
	err = CheckConnection(client)
	if err != nil {
		c.log.WithError(err).Error("failed to connect to mongodb")
		return err
	}

	c.client = client

	c.Database = newDatabase(c.client.Database(config.DatabaseName))

	return nil
}

func (c *Client) Start(ctx context.Context) error {

	return nil
}

func (c *Client) Stop(ctx context.Context) error {
	if err := c.client.Disconnect(ctx); err != nil {
		c.log.WithError(err).Error("failed to close connection to mongodb")
		return err
	}

	return nil
}

func CheckConnection(client *mongo.Client) error {
	err := client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	return nil
}

func newDatabase(db *mongo.Database) *database {
	return &database{
		db:          db,
		collections: make(map[string]*Collection),
	}
}

func (d *database) getCollection(name string) *Collection {
	d.collectionsMux.RLock()
	defer d.collectionsMux.RUnlock()

	return d.collections[name]
}

func (d *database) addCollection(name string) *Collection {
	d.collectionsMux.Lock()
	defer d.collectionsMux.Unlock()

	d.collections[name] = &Collection{
		d.db.Collection(name),
	}

	return d.collections[name]
}

func (c *Collection) createIndices(ctx context.Context) error {
	_, err := c.collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{"expireAt", 1}},
		Options: options.Index().SetExpireAfterSeconds(0),
	})
	return err
}

func (c *Collection) Insert(ctx context.Context, object any) error {
	_, err := c.collection.InsertOne(ctx, object)
	return err
}

func (c *Collection) FindByID(ctx context.Context, id string) (map[string]any, error) {
	filter := bson.D{{"_id", id}}

	var result map[string]any

	err := c.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return nil, errmongo.ErrNotFound
		default:
			return nil, err
		}
	}

	return result, nil
}

func (c *Collection) Update(ctx context.Context, id string, object *map[string]any) error {
	filter := bson.D{{"_id", id}}
	result, err := c.collection.ReplaceOne(ctx, filter, object)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errmongo.ErrNotFound
	}
	return nil
}

func (c *Collection) DeleteByID(ctx context.Context, id string) error {
	filter := bson.D{{"_id", id}}
	result, err := c.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errmongo.ErrNotFound
	}
	return nil
}
