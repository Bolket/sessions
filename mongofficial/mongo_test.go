package mongofficial

import (
	"context"
	"testing"

	"github.com/bolket/sessions"
	"github.com/bolket/sessions/tester"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoTestServer = "localhost:27017"

var newStore = func(_ *testing.T) sessions.Store {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoTestServer))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	c := client.Database("test").Collection("sessions")
	return NewStore(c, 3600, true, []byte("secret"))
}

func TestMongo_SessionGetSet(t *testing.T) {
	tester.GetSet(t, newStore)
}

func TestMongo_SessionDeleteKey(t *testing.T) {
	tester.DeleteKey(t, newStore)
}

func TestMongo_SessionFlashes(t *testing.T) {
	tester.Flashes(t, newStore)
}

func TestMongo_SessionClear(t *testing.T) {
	tester.Clear(t, newStore)
}

func TestMongo_SessionOptions(t *testing.T) {
	tester.Options(t, newStore)
}
