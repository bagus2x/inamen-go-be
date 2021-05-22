package tournament

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
	Create(tour *entity.Tournament) error
	Read(id string) (*entity.Tournament, error)
	ReadAllByHost(hostID string) ([]*entity.Tournament, error)
	Update(id, hostID string, tour *entity.Tournament) error
	Delete(id, hostID string) error
	SearchTournamenName(text string) ([]*entity.Tournament, error)
	SearchTournamen(text string) ([]*entity.Tournament, error)
}

func NewRepo(db *mongo.Database) Repository {
	col := db.Collection("tournament")

	return &repository{
		col: col,
	}
}

func (r *repository) Create(tour *entity.Tournament) error {
	ctx, cancel := config.MongoContext()
	defer cancel()

	tour.ID = primitive.NewObjectID()
	tour.CreatedAt = time.Now().Unix()
	tour.UpdatedAt = time.Now().Unix()

	_, err := r.col.InsertOne(ctx, tour)
	if err != nil {
		log.Error(err)
		return model.ErrFailedToCreateTournament
	}

	return nil
}

func (r *repository) Read(id string) (*entity.Tournament, error) {
	ctx, cancel := config.MongoContext()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error(err)
		return nil, model.ErrInvalidIDFormat
	}

	var tournament entity.Tournament

	err = r.col.FindOne(ctx, m{"_id": objID}).Decode(&tournament)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, model.ErrTournamentNotfound
		}

		log.Error(err)
		return nil, model.ErrFailedToReadTournament
	}

	return &tournament, nil
}

func (r *repository) ReadAllByHost(hostID string) ([]*entity.Tournament, error) {
	ctx, cancel := config.MongoContext()
	defer cancel()

	cursor, err := r.col.Find(ctx, m{"host._id": hostID})
	if err != nil {
		log.Error(err)
		return nil, model.ErrFailedToReadTournament
	}

	res := make([]*entity.Tournament, 0)
	for cursor.Next(ctx) {
		var tour entity.Tournament
		err = cursor.Decode(&tour)
		if err != nil {
			log.Error(err)
			return nil, model.ErrFailedToReadTournament
		}

		res = append(res, &tour)
	}

	if len(res) < 0 {
		return nil, model.ErrTournamentNotfound
	}

	return res, nil
}

func (r *repository) Update(id, hostID string, tour *entity.Tournament) error {
	ctx, cancel := config.MongoContext()
	defer cancel()

	tour.UpdatedAt = time.Now().Unix()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error(err)
		return model.ErrInvalidIDFormat
	}

	err = r.col.FindOneAndUpdate(ctx, m{"_id": objID, "host._id": hostID}, m{"$set": m{
		"name":        tour.Name,
		"published":   tour.Published,
		"description": tour.Description,
		"visiblity":   tour.Visibilty,
		"date":        tour.Date,
		"updatedAt":   tour.UpdatedAt,
	}}, &options.FindOneAndUpdateOptions{
		ReturnDocument: options.FindOneAndUpdate().SetReturnDocument(options.After).ReturnDocument,
	}).Decode(tour)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.ErrNotFound
		}
		log.Error(err)
		return model.ErrFailedToUpdateTournament
	}

	return nil
}

func (r *repository) Delete(id, hostID string) error {
	ctx, cancel := config.MongoContext()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Error(err)
		return model.ErrInvalidIDFormat
	}

	res, err := r.col.DeleteOne(ctx, m{"_id": objID, "host._id": id})

	if err != nil {
		log.Error(err)
		return model.ErrFailedToDeleteTournament
	}

	if res.DeletedCount < 1 {
		return model.ErrTournamentNotfound
	}

	return nil
}

func (s *repository) SearchTournamenName(text string) ([]*entity.Tournament, error) {
	ctx, cancel := config.MongoContext()
	defer cancel()

	opts := options.Find().
		SetLimit(10).
		SetProjection(m{"name": 1, "score": m{"$meta": "textScore"}}).
		SetSort(m{"score": m{"$meta": "textScore"}})

	cursor, err := s.col.Find(ctx, m{"$text": m{"$search": text}}, opts)
	if err != nil {
		log.Error(err)
		return nil, model.ErrFailedToReadTournament
	}

	res := make([]*entity.Tournament, 0)
	for cursor.Next(ctx) {
		var tour entity.Tournament
		err = cursor.Decode(&tour)
		if err != nil {
			log.Error(err)
			return nil, model.ErrFailedToReadTournament
		}

		res = append(res, &tour)
	}

	return res, nil
}

func (s *repository) SearchTournamen(text string) ([]*entity.Tournament, error) {
	ctx, cancel := config.MongoContext()
	defer cancel()

	opts := options.Find().
		SetLimit(10).
		SetSort(m{"score": m{"$meta": "textScore"}})

	cursor, err := s.col.Find(ctx, m{"$text": m{"$search": text}}, opts)
	if err != nil {
		log.Error(err)
		return nil, model.ErrFailedToReadTournament
	}

	res := make([]*entity.Tournament, 0)
	for cursor.Next(ctx) {
		var tour entity.Tournament
		err = cursor.Decode(&tour)
		if err != nil {
			log.Error(err)
			return nil, model.ErrFailedToReadTournament
		}

		res = append(res, &tour)
	}

	return res, nil
}
