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

// func uuidEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
// 	if !val.IsValid() || val.Type() != tUUID {
// 		return bsoncodec.ValueEncoderError{Name: "uuidEncodeValue", Types: []reflect.Type{tUUID}, Received: val}
// 	}
// 	b := val.Interface().(uuid.UUID)
// 	return vw.WriteBinaryWithSubtype(b[:], uuidSubtype)
// }

// func uuidDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
// 	if !val.CanSet() || val.Type() != tUUID {
// 		return bsoncodec.ValueDecoderError{Name: "uuidDecodeValue", Types: []reflect.Type{tUUID}, Received: val}
// 	}

// 	var data []byte
// 	var subtype byte
// 	var err error
// 	switch vrType := vr.Type(); vrType {
// 	case bsontype.Binary:
// 		data, subtype, err = vr.ReadBinary()
// 		if subtype != uuidSubtype {
// 			return fmt.Errorf("unsupported binary subtype %v for UUID", subtype)
// 		}
// 	case bsontype.Null:
// 		err = vr.ReadNull()
// 	case bsontype.Undefined:
// 		err = vr.ReadUndefined()
// 	default:
// 		return fmt.Errorf("cannot decode %v into a UUID", vrType)
// 	}

// 	if err != nil {
// 		return err
// 	}
// 	uuid2, err := uuid.FromBytes(data)
// 	if err != nil {
// 		return err
// 	}
// 	val.Set(reflect.ValueOf(uuid2))
// 	return nil
// }
