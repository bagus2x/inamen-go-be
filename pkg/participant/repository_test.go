package participant

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

func TestCreateParticipant(t *testing.T) {
	repo := NewRepo(db())
	tourObjID, _ := primitive.ObjectIDFromHex("60a7c78aba454118d941fa85")
	err := repo.Create(&entity.Participant{
		ID:           primitive.NewObjectID(),
		TournamentID: tourObjID,
		TeamName:     "Team Secret",
		Players: []entity.Player{
			{ID: "1", Username: "Bagus"},
			{ID: "2", Username: "Joni"},
		},
		Description: "Lorem ipsum",
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	})

	assert.NoError(t, err)
}

func TestRead(t *testing.T) {
	repo := NewRepo(db())
	participant, err := repo.Read("60a85a67e66e971eb7c1be3d")
	assert.NoError(t, err)
	assert.NotNil(t, participant)
}

func TestReadAllByTourID(t *testing.T) {
	repo := NewRepo(db())
	participants, err := repo.ReadAllByTournamentID("60a7c78aba454118d941fa85")
	assert.NoError(t, err)
	assert.NotZero(t, len(participants))
}

func TestUpdate(t *testing.T) {
	repo := NewRepo(db())
	err := repo.Update("60a85a67e66e971eb7c1be3d", &entity.Participant{
		TeamName: "Team Bapak",
		Players: []entity.Player{
			{ID: "1", Username: "Bagus"},
			{ID: "2", Username: "Joni"},
		},
		Description: "Lorem ipsum",
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	})

	assert.NoError(t, err)
}
