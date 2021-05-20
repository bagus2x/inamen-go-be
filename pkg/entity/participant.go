package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Participant struct {
	ID           primitive.ObjectID `bson:"_id"`
	TournamentID primitive.ObjectID `bson:"tournamentID"`
	Description  string             `bson:"description"`
}
