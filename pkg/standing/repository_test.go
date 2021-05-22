package standing

import (
	"log"
	"testing"
	"time"

	"github.com/bagus2x/inamen-go-be/config"
	"github.com/bagus2x/inamen-go-be/pkg/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func db() *mongo.Database {
	db, err := config.MongoConnection("inamenTest", "mongodb://localhost:27017", 10, 100, 10)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func TestCreate(t *testing.T) {
	repo := NewRepo(db())
	tourID, _ := primitive.ObjectIDFromHex("60a7cabe3c52085e3d6598e7")

	data := [][]interface{}{
		{"Team Secret", 130, 280},
		{"NRG", 125, 295},
		{"OG", 115, 257},
	}

	err := repo.Create(&entity.Standing{
		ID:           primitive.NewObjectID(),
		TournamentID: tourID,
		Schema: entity.StandingSchema{
			Columns: []string{"Team Name", "Total Kills", "Total Points"},
			Data:    data,
		},
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	})

	assert.NoError(t, err)
}

func TestRead(t *testing.T) {
	repo := NewRepo(db())
	standing, err := repo.Read("60a84e776a886c88a40a0f82")
	assert.NotNil(t, standing)
	assert.NoError(t, err)
	t.Log(err)
}

func TestReadAll(t *testing.T) {
	repo := NewRepo(db())
	standings, err := repo.ReadAllByTournamentID("60a7cabe3c52085e3d6598e7")
	assert.NoError(t, err)
	assert.NotZero(t, len(standings))
	t.Log(standings)
}

func TestUpdate(t *testing.T) {
	repo := NewRepo(db())

	data := [][]interface{}{
		{"Team Secret", 130, 280},
		{"NRG", 125, 295},
		{"OG", 115, 257},
	}

	err := repo.Update("60a84e184ef1a900fa8728ca", &entity.Standing{
		Name: "Babak 4 Besar",
		Schema: entity.StandingSchema{
			Columns: []string{"Team Name", "Total Kills", "Total Points"},
			Data:    data,
		},
		UpdatedAt: time.Now().Unix(),
	})

	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	repo := NewRepo(db())

	err := repo.Delete("60a84e184ef1a900fa8728ca")
	assert.NoError(t, err)
}
