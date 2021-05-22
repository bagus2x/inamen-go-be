package tournament

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

var tour = &entity.Tournament{
	ID:   primitive.NewObjectID(),
	Name: "Turnamen lucu",
	Host: entity.Host{
		ID:       "1",
		Username: "Jaenab",
	},
	Game:              "Hago",
	Mode:              "Squad",
	Presence:          "Online",
	Published:         false,
	Visibilty:         "Private",
	TotalParticipants: 20,
	TotalTeamMembers:  4,
	Platforms:         []string{"PC", "PS5"},
}

func TestCreate(t *testing.T) {
	repo := NewRepo(db())
	err := repo.Create(tour)
	assert.NoError(t, err)
}

func TestRead(t *testing.T) {
	repo := NewRepo(db())
	standing, err := repo.Read("60a8929a9aa00cc45634fc99")
	assert.NotNil(t, standing)
	assert.NoError(t, err)
	t.Log(err)
}

func TestReadAllByHost(t *testing.T) {
	repo := NewRepo(db())
	res, err := repo.ReadAllByHost("2")
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestUpdateByHost(t *testing.T) {
	repo := NewRepo(db())

	tour := entity.Tournament{
		Name:      "Turnamen Curang",
		Published: false,
		Visibilty: "Public",
		Date: entity.TournamentDate{
			RegistrationStartDate: time.Now().Add(time.Hour * 24 * 1).Unix(),
			RegistrationLastDate:  time.Now().Add(time.Hour * 24 * 2).Unix(),
			TournamentStartDate:   time.Now().Add(time.Hour * 24 * 3).Unix(),
			TournamentLastDate:    time.Now().Add(time.Hour * 24 * 4).Unix(),
		},
		UpdatedAt: time.Now().Unix(),
	}

	err := repo.Update("60a8929a9aa00cc45634fc99", "1", &tour)

	t.Log(tour)
	assert.NoError(t, err)
}

func TestDeleteByHost(t *testing.T) {
	repo := NewRepo(db())
	err := repo.Delete("1", "60a8929a9aa00cc45634fc99")
	assert.NoError(t, err)
}

func TestSearchTournamenName(t *testing.T) {
	repo := NewRepo(db())
	tour, err := repo.SearchTournamenName("turnamen")
	assert.NoError(t, err)
	assert.NotZero(t, len(tour))
	for _, tr := range tour {
		t.Logf("%+v", tr)
	}
}

func TestSearchTournamen(t *testing.T) {
	repo := NewRepo(db())
	tour, err := repo.SearchTournamen("turnamen")
	assert.NoError(t, err)
	assert.NotZero(t, len(tour))
	for _, tr := range tour {
		t.Logf("%+v", tr)
	}
}
