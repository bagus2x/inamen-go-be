package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Match struct {
	ID             primitive.ObjectID   `json:"_id"`
	TournamentID   primitive.ObjectID   `json:"tournamentID"`
	Name           string               `json:"name"`
	Description    string               `bson:"description"`
	Dates          MatchDates           `json:"dates"`
	ParticipantIDs []primitive.ObjectID `json:"participantIDs"`
	CreatedAt      int64                `json:"createdAt"`
	UpdatedAt      int64                `json:"updatedAt"`
}

type MatchDates struct {
	StartDate int64 `json:"startDate" validate:"required"`
	LastDate  int64 `json:"lastDate" validate:"required"`
}

type CreateMatchRequest struct {
	TournamentID   primitive.ObjectID   `json:"tournamentID" validate:"required"`
	Dates          MatchDates           `json:"dates" validate:"required"`
	Name           string               `json:"name" validate:"required,gt=5,lte=100"`
	Description    string               `bson:"description" validate:"lte=255"`
	ParticipantIDs []primitive.ObjectID `json:"participantIDs"`
}

type CreateMatchResponse Match

type FetchMatchResponse Match

type UpdateMatchRequest struct {
	Name        string     `json:"name" validate:"required"`
	Description string     `bson:"description" validate:"lte=255"`
	Dates       MatchDates `json:"dates" validate:"required"`
}

type UpdateMatchResponse Match
