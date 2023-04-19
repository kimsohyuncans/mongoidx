//go:build integration

package models

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/mongodb/mongo-tools/mongorestore"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

var connection = Connection{
	ID:  "1",
	URI: "mongodb://root:example@localhost:27017",
}

func TestMain(m *testing.M) {
	// restoreTestData()
	os.Exit(m.Run())
}

func restoreTestData() {
	mongorestoreArgs := []string{"--nsFrom=datasets.*", "--nsTo=mflix.*", "--drop", "--gzip", " --quiet", "--archive=../../testdata/sample_dataset.gz", "--quiet", "--authenticationDatabase=admin", connection.URI}
	opts, err := mongorestore.ParseOptions(mongorestoreArgs, "built-without-version-string", "build-without-git-commit")
	if err != nil {
		log.Fatal("Error parsing mongorestore args: ", err)
	}

	executor, err := mongorestore.New(opts)
	if err != nil {
		log.Fatal("Error creating mongorestore executor: ", err)
	}

	result := executor.Restore()
	if result.Err != nil {
		log.Printf("Failed: %v", result.Err)
	}

	defer executor.Close()

	if executor.ToolOptions.WriteConcern.Acknowledged() {
		log.Printf("%v document(s) restored successfully. %v document(s) failed to restore.", result.Successes, result.Failures)
	} else {
		log.Printf("done")
	}

	if result.Err != nil {
		log.Fatalf("Error restoring data: %s", err)
	}
}

func simulateRealProductionQuery() {
	ctx := context.TODO()
	cursor, err := connection.driver.Database("mflix").Collection("movies").Find(ctx, bson.M{"title": "The Godfather"})
	if err != nil {
		log.Fatal(err)
	}

	var movies []map[string]interface{}
	if err = cursor.All(ctx, &movies); err != nil {
		log.Fatal(err)
	}

	for _, movie := range movies {
		if _, err := connection.driver.Database("mflix").Collection("comments").Find(ctx, bson.M{"movie_id": movie["_id"]}); err != nil {
			log.Fatal(err)
		}
	}

}

func TestEnableProfiling(t *testing.T) {
	ctx := context.Background()
	assert.NoError(t, connection.Connect(ctx))
	assert.NoError(t, connection.Ping(ctx))

	assert.NoError(t, connection.StartProfiling(ctx, "mongoidx"))

	profiling, err := connection.Profiling(ctx, "mongoidx")
	assert.Equal(t, true, profiling)
	assert.NoError(t, err)

	assert.NoError(t, connection.StopProfiling(ctx, "mongoidx"))

	profiling, err = connection.Profiling(ctx, "mongoidx")
	assert.Equal(t, false, profiling)
	assert.NoError(t, err)
	assert.NoError(t, connection.Disconnect(ctx))
}

func TestListDatabases(t *testing.T) {
	ctx := context.Background()
	assert.NoError(t, connection.Connect(ctx))
	assert.NoError(t, connection.Ping(ctx))

	_, err := connection.ListDatabases(ctx)
	assert.NoError(t, err)

	assert.NoError(t, connection.Disconnect(ctx))
}

func TestGetProfileData(t *testing.T) {
	ctx := context.Background()

	assert.NoError(t, connection.Connect(ctx))
	assert.NoError(t, connection.Ping(ctx))

	c, err := connection.GetProfileData(ctx, "mflix")
	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.NoError(t, connection.StopProfiling(ctx, "mflix"))
	assert.NoError(t, connection.Disconnect(ctx))
}
