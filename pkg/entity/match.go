package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Match struct {
	ID           primitive.ObjectID `bson:"_id"`
	TournamentID primitive.ObjectID `bson:"tournamentID"`
	Dates        MatchDates         `bson:"dates"`
}

type MatchDates struct {
	StartDate int64 `bson:"startDate"`
	LastDate  int64 `bson:"lastDate"`
}
