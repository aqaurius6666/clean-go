package odm

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/aqaurius6666/clean-go/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// var (
// 	tUUID       = reflect.TypeOf(uuid.UUID{})
// 	uuidSubtype = byte(0x04)
// )

type ODMRepository struct {
	DB *mongo.Database
}

func ConnectMongoDB(cfg config.DBConfig) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if cfg.DSN == "" {
		cfg.DSN = (&url.URL{
			Scheme:   cfg.Scheme,
			User:     url.UserPassword(cfg.User, cfg.Pass),
			Host:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			Path:     cfg.Name,
			RawQuery: cfg.Query,
		}).String()
	}
	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			log.Print(evt.Command)
		},
	}
	registry := bson.NewRegistryBuilder().
		// RegisterTypeEncoder(tUUID, bsoncodec.ValueEncoderFunc(uuidEncodeValue)).
		// RegisterTypeDecoder(tUUID, bsoncodec.ValueDecoderFunc(uuidDecodeValue)).
		Build()
	opts := options.Client().ApplyURI(cfg.DSN)
	opts = opts.SetRegistry(registry)
	opts = opts.SetMonitor(cmdMonitor)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return client.Database(cfg.Name), nil
}

func (s *ODMRepository) Migrate(ctx context.Context) error {
	return nil
}
