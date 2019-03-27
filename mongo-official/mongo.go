package mongo-official

import (
	"github.com/bolket/mongostore"
	gsessions "github.com/gorilla/sessions"
	"github.com/bolket/sessions"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	sessions.Store
}

func NewStore(c *mongo.Collection, maxAge int32, ensureTTL bool, keyPairs ...[]byte) Store {
	return &store{mongostore.NewMongoStore(c, maxAge, ensureTTL, keyPairs...)}
}

type store struct {
	*mongostore.MongoStore
}

func (c *store) Options(options sessions.Options) {
	c.MongoStore.Options = &gsessions.Options{
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}
}
