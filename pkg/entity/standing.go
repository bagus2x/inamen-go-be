package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Standing struct {
	ID           primitive.ObjectID `bson:"_id"`
	TournamentID primitive.ObjectID `bson:"tournamentID"`
	Name         string             `bson:"name"`
	Spotlight    bool               `bson:"spotlight"`
	Schema       StandingSchema     `bson:"schema"`
	CreatedAt    int64              `bson:"createdAt"`
	UpdatedAt    int64              `bson:"updatedAt"`
}

type StandingSchema struct {
	Columns []string        `bson:"columns"`
	Data    [][]interface{} `bson:"data"`
}
