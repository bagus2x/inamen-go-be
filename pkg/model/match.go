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
	StartDate int64 `json:"startDate"`
	LastDate  int64 `json:"lastDate"`
}

type CreateMatchRequest struct {
	TournamentID   primitive.ObjectID   `json:"tournamentID"`
	Dates          MatchDates           `json:"dates"`
	Name           string               `json:"name"`
	Description    string               `bson:"description"`
	ParticipantIDs []primitive.ObjectID `json:"participantIDs"`
}

type CreateMatchResponse Match

type FetchMatchResponse Match

type UpdateMatchRequest struct {
	Name        string     `json:"name"`
	Description string     `bson:"description"`
	Dates       MatchDates `json:"dates"`
}

type UpdateMatchResponse Match
