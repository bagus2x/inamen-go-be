package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Match struct {
	ID             primitive.ObjectID   `bson:"_id"`
	TournamentID   primitive.ObjectID   `bson:"tournamentID"`
	Name           string               `bson:"name"`
	Description    string               `bson:"description"`
	Dates          MatchDates           `bson:"dates"`
	ParticipantIDs []primitive.ObjectID `bson:"participantIDs"`
	CreatedAt      int64                `bson:"createdAt"`
	UpdatedAt      int64                `bson:"updatedAt"`
}

type MatchDates struct {
	StartDate int64 `bson:"startDate"`
	LastDate  int64 `bson:"lastDate"`
}
