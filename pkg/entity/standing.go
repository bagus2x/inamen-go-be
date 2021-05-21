package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Standing struct {
	ID           primitive.ObjectID `bson:"_id"`
	TournamentID primitive.ObjectID `bson:"tournamentID"`
	Schema       StandingSchema     `bson:"schema"`
}

type StandingSchema struct {
	Columns []string      `bson:"columns"`
	Data    []interface{} `bson:"data"`
}
