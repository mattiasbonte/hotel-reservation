package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/fulltimegodev/hotel-reservation/db"
	"github.com/fulltimegodev/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testdburi = "mongodb://localhost:27018"
	dbname    = "hotel-reservation-test"
)

type testdb struct {
	db.UserStore
}

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdburi))
	if err != nil {
		log.Fatal(err)
	}
	return &testdb{
		UserStore: db.NewMongoUserStore(client, dbname),
	}
}

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	route := "/users"
	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	app.Post(route, userHandler.HandlePostUser)

	params := types.CreateUserParams{
		Email:     "some@foo.com",
		FirstName: "James",
		LastName:  "Foo",
		Password:  "loliwolliedrollie",
	}
	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", route, bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	var user types.User
	json.NewDecoder(res.Body).Decode(&user)

	if len(user.ID) == 0 {
		t.Errorf("expecting a user id to be set")
	} else {
		t.Logf("user id: %s", user.ID)
	}
	if len(user.EncryptedPassword) > 0 {
		t.Errorf("expected the EncryptedPassword not to be included in the json response")
	} else {
		t.Logf("user EncryptedPassword: %s", user.EncryptedPassword)
	}

	if user.FirstName != params.FirstName {
		t.Errorf("expected username %s but got %s", params.FirstName, user.FirstName)
	} else {
		t.Logf("user FirstName: %s", user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("expected username %s but got %s", params.LastName, user.LastName)
	} else {
		t.Logf("user LastName: %s", user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("expected username %s but got %s", params.Email, user.Email)
	} else {
		t.Logf("user Email: %s", user.Email)
	}
}
