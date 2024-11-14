package mongo

import (
	"context"
	"gateway/pkg/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	database struct {
		db          *mongo.Database
		collections map[string]*mongo.Collection
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
		collections: make(map[string]*mongo.Collection),
	}
}
