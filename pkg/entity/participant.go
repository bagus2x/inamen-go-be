package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Participant struct {
	ID           primitive.ObjectID `bson:"_id"`
	TournamentID primitive.ObjectID `bson:"tournamentID"`
	Players      []string           `bson:"players"`
	Description  string             `bson:"description"`
}
