package participant

import (
	"time"

	"github.com/bagus2x/inamen-go-be/config"
	"github.com/bagus2x/inamen-go-be/pkg/entity"
	"github.com/bagus2x/inamen-go-be/pkg/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type m bson.M

type repository struct {
	col *mongo.Collection
}

type Repository interface {
	Create(participant *entity.Participant) error
	Read(id string) (*entity.Participant, error)
	ReadAllByTournamentID(tournamentID string) ([]*entity.Participant, error)
	Update(id string, participant *entity.Participant) error
	Delete(id string) error
}

func NewRepo(db *mongo.Database) Repository {
	col := db.Collection("participant")

	return &repository{
		col: col,
	}
}

func (r *repository) Create(participant *entity.Participant) error {
	ctx, cancel := config.MongoContext()
	defer cancel()

	participant.ID = primitive.NewObjectID()
	participant.CreatedAt = time.Now().Unix()
	participant.UpdatedAt = time.Now().Unix()

	_, err := r.col.InsertOne(ctx, participant)
	if err != nil {
		log.Error(err)
		return model.ErrFailedToCreateParticipant
	}

	return nil
}

func (r *repository) Read(id string) (*entity.Participant, error) {
	ctx, cancel := config.MongoContext()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, model.ErrInvalidIDFormat
	}

	var participant entity.Participant

	err = r.col.FindOne(ctx, m{"_id": objID}).Decode(&participant)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, model.ErrParticipantNotfound
		}

		log.Error(err)
		return nil, model.ErrFailedToReadParticipant
	}

	return &participant, nil
}

func (r *repository) ReadAllByTournamentID(tournamentID string) ([]*entity.Participant, error) {
	ctx, cancel := config.MongoContext()
	defer cancel()

	tourObjID, err := primitive.ObjectIDFromHex(tournamentID)
	if err != nil {
		return nil, model.ErrInvalidIDFormat
	}

	cursor, err := r.col.Find(ctx, m{"tournamentID": tourObjID})
	if err != nil {
		log.Error(err)
		return nil, model.ErrFailedToReadParticipant
	}

	res := make([]*entity.Participant, 0)
	for cursor.Next(ctx) {
		var tour entity.Participant
		err = cursor.Decode(&tour)
		if err != nil {
			log.Error(err)
			return nil, model.ErrFailedToReadParticipant
		}

		res = append(res, &tour)
	}

	if len(res) < 0 {
		return nil, model.ErrParticipantNotfound
	}

	return res, nil
}

func (r *repository) Update(id string, participant *entity.Participant) error {
	ctx, cancel := config.MongoContext()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.ErrInvalidIDFormat
	}

	participant.UpdatedAt = time.Now().Unix()

	err = r.col.FindOneAndUpdate(ctx, m{"_id": objID}, m{"$set": m{
		"teamName":    participant.TeamName,
		"description": participant.Description,
		"updatedAt":   participant.UpdatedAt,
	}}, &options.FindOneAndUpdateOptions{
		ReturnDocument: options.FindOneAndUpdate().SetReturnDocument(options.After).ReturnDocument,
	}).Decode(participant)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.ErrParticipantNotfound
		}

		log.Error(err)
		return model.ErrFailedToUpdateParticipant
	}

	return nil
}

func (r *repository) Delete(id string) error {
	ctx, cancel := config.MongoContext()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.ErrInvalidIDFormat
	}

	res, err := r.col.DeleteOne(ctx, m{"_id": objID})
	if err != nil {
		log.Error(err)
		return model.ErrFailedToDeleteTournament
	}

	if res.DeletedCount < 1 {
		return model.ErrTournamentNotfound
	}

	return nil
}
