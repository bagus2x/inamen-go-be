package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Participant struct {
	ID           primitive.ObjectID `bson:"_id"`
	TournamentID primitive.ObjectID `bson:"tournamentID"`
	TeamName     string             `bson:"teamName"`
	Players      []Player           `bson:"players"`
	Description  string             `bson:"description"`
	CreatedAt    int64              `bson:"createdAt"`
	UpdatedAt    int64              `bson:"updatedAt"`
}

type Player struct {
	ID       string `bson:"_id"`
	Username string `bson:"username"`
}
