package match

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
	Create(match *entity.Match) error
	Read(id string) (*entity.Match, error)
	ReadAllByTournamentID(tourID string) ([]*entity.Match, error)
	Update(id string, match *entity.Match) error
	Delete(id string) error
}

func (r *repository) Create(match *entity.Match) error {
	ctx, cancel := config.MongoContext()
	defer cancel()

	match.ID = primitive.NewObjectID()
	match.CreatedAt = time.Now().Unix()
	match.UpdatedAt = time.Now().Unix()

	_, err := r.col.InsertOne(ctx, match)
	if err != nil {
		log.Error(err)
		return model.ErrMatchNotfound
	}

	return nil
}

func (r *repository) Read(id string) (*entity.Match, error) {
	ctx, cancel := config.MongoContext()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, model.ErrInvalidIDFormat
	}

	var match entity.Match

	err = r.col.FindOne(ctx, m{"_id": objID}).Decode(&match)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, model.ErrMatchNotfound
		}

		log.Error(err)
		return nil, model.ErrFailedToReadMatch
	}

	return &match, nil
}

func (r *repository) ReadAllByTournamentID(tourID string) ([]*entity.Match, error) {
	ctx, cancel := config.MongoContext()
	defer cancel()

	tourObjID, err := primitive.ObjectIDFromHex(tourID)
	if err != nil {
		log.Error(err)
		return nil, model.ErrInvalidIDFormat
	}

	cursor, err := r.col.Find(ctx, m{"tourID": tourObjID})
	if err != nil {
		log.Error(err)
		return nil, model.ErrFailedToReadMatch
	}

	res := make([]*entity.Match, 0)
	for cursor.Next(ctx) {
		var tour entity.Match
		err = cursor.Decode(&tour)
		if err != nil {
			log.Error(err)
			return nil, model.ErrFailedToReadMatch
		}

		res = append(res, &tour)
	}

	if len(res) < 0 {
		return nil, model.ErrMatchNotfound
	}

	return res, nil
}

func (r *repository) Update(id string, match *entity.Match) error {
	ctx, cancel := config.MongoContext()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.ErrInvalidIDFormat
	}

	match.UpdatedAt = time.Now().Unix()

	err = r.col.FindOneAndUpdate(ctx, m{"_id": objID}, m{"$set": m{
		"name":        match.Name,
		"description": match.Description,
		"dates":       match.Dates,
		"updatedAt":   match.UpdatedAt,
	}}, &options.FindOneAndUpdateOptions{
		ReturnDocument: options.FindOneAndUpdate().SetReturnDocument(options.After).ReturnDocument,
	}).Decode(match)

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
		return model.ErrFailedToDeleteMatch
	}

	if res.DeletedCount < 1 {
		return model.ErrMatchNotfound
	}

	return nil
}
