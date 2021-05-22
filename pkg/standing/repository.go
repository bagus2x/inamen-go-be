package standing

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
	Create(standing *entity.Standing) error
	Read(id string) (*entity.Standing, error)
	ReadAllByTournamentID(tourID string) ([]*entity.Standing, error)
	Update(id string, standing *entity.Standing) error
	Delete(id string) error
}

func NewRepo(db *mongo.Database) Repository {
	col := db.Collection("standing")

	return &repository{
		col: col,
	}
}

func (r *repository) Create(standing *entity.Standing) error {
	ctx, cancel := config.MongoContext()
	defer cancel()

	standing.ID = primitive.NewObjectID()
	standing.CreatedAt = time.Now().Unix()
	standing.UpdatedAt = time.Now().Unix()

	_, err := r.col.InsertOne(ctx, standing)
	if err != nil {
		log.Error(err)
		return model.ErrFailedToCreateStanding
	}

	return nil
}

func (r *repository) Read(id string) (*entity.Standing, error) {
	ctx, cancel := config.MongoContext()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, model.ErrInvalidIDFormat
	}

	var standing entity.Standing

	err = r.col.FindOne(ctx, m{"_id": objID}).Decode(&standing)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, model.ErrStandingNotfound
		}

		log.Error(err)
		return nil, model.ErrFailedToReadStanding
	}

	return &standing, nil
}

func (r *repository) ReadAllByTournamentID(tourID string) ([]*entity.Standing, error) {
	ctx, cancel := config.MongoContext()
	defer cancel()

	tourObjID, err := primitive.ObjectIDFromHex(tourID)
	if err != nil {
		return nil, model.ErrInvalidIDFormat
	}

	cursor, err := r.col.Find(ctx, m{"tourID": tourObjID})
	if err != nil {
		log.Error(err)
		return nil, model.ErrFailedToReadStanding
	}

	res := make([]*entity.Standing, 0)
	for cursor.Next(ctx) {
		var tour entity.Standing
		err = cursor.Decode(&tour)
		if err != nil {
			log.Error(err)
			return nil, model.ErrFailedToReadStanding
		}

		res = append(res, &tour)
	}

	if len(res) < 0 {
		return nil, model.ErrStandingNotfound
	}

	return res, nil
}

func (r *repository) Update(id string, standing *entity.Standing) error {
	ctx, cancel := config.MongoContext()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.ErrInvalidIDFormat
	}

	standing.UpdatedAt = time.Now().Unix()

	err = r.col.FindOneAndUpdate(ctx, m{"_id": objID}, m{"$set": m{
		"name":      standing.Name,
		"spotlight": standing.Spotlight,
		"schema":    standing.Schema,
		"updatedAt": standing.UpdatedAt,
	}}, &options.FindOneAndUpdateOptions{
		ReturnDocument: options.FindOneAndUpdate().SetReturnDocument(options.After).ReturnDocument,
	}).Decode(standing)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.ErrStandingNotfound
		}

		log.Error(err)
		return model.ErrFailedToUpdateStanding
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
