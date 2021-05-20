package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tournament struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
	Host string             `bson:"host"`
	Game string             `bson:"game"`
	Mode string             `bson:"mode"`
	Date TournamentDate     `bson:"date"`
}

type TournamentDate struct {
	RegistrationStartDate int64 `bson:"registrationStartDate"`
	RegistrationEndDate   int64 `bson:"registrationEndDate"`
	TournamentStartDate   int64 `bson:"tournamentStartDate"`
	TournamentEndDate     int64 `bson:"tournamentEndDate"`
}
