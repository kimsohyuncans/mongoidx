//go:build integration

package models

import (
	"context"
	"log"
	"math/rand"
	"os"
	"sync"
	"testing"

	"github.com/mongodb/mongo-tools/mongorestore"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

var connection = Connection{
	ID:  "1",
	URI: "mongodb://root:example@localhost:27017/mflix?authSource=admin",
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

func simulateRealProductionQuery() error {
	ctx := context.Background()

	titles, err := connection.driver.Database("mflix").Collection("movies").Distinct(ctx, "title", bson.M{})
	if err != nil {
		return err
	}

	genres, err := connection.driver.Database("mflix").Collection("movies").Distinct(ctx, "title", bson.M{})
	if err != nil {
		return err
	}

	numRequests := 1000
	var wg sync.WaitGroup

	// simulate thousand users search a movies with random title and genres
	for i := 0; i <= numRequests; i++ {
		title := titles[rand.Intn(len(titles))]
		genre := genres[rand.Intn(len(genres))]
		wg.Add(1)
		go func(searchTitle, searchGenre string) {
			_, err := connection.driver.Database("mflix").Collection("movies").Find(ctx, bson.M{"title": "The Godfather"})
			if err != nil {
				log.Fatal(err)
			}

		}(title.(string), genre.(string))
	}

	return nil
}

func TestEnableProfiling(t *testing.T) {
	ctx := context.Background()
	assert.NoError(t, connection.Connect(ctx))
	assert.NoError(t, connection.Ping(ctx))

	assert.NoError(t, connection.StartProfiling(ctx, "mflix"))
	assert.NoError(t, simulateRealProductionQuery())

	profiling, err := connection.Profiling(ctx, "mflix")
	assert.Equal(t, true, profiling)
	assert.NoError(t, err)

	assert.NoError(t, connection.StopProfiling(ctx, "mflix"))

	profiling, err = connection.Profiling(ctx, "mflix")
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
