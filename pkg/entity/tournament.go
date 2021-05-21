package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tournament struct {
	ID                primitive.ObjectID `bson:"_id"`
	Title             string             `bson:"title"`
	Host              string             `bson:"Host"`
	Game              string             `bson:"game"`
	Mode              string             `bson:"mode"`
	Presence          string             `bson:"presence"`
	TotalParticipants int                `bson:"totalParticipants"`
	TotalMembers      int                `bson:"totalMembers"`
	Platforms         []string           `bson:"platforms"`
	Date              TournamentDate     `bson:"date"`
}

type TournamentDate struct {
	RegistrationStartDate int64 `bson:"registrationStartDate"`
	RegistrationLastDate  int64 `bson:"registrationLastDate"`
	TournamentStartDate   int64 `bson:"tournamentStartDate"`
	TournamentLastDate    int64 `bson:"tournamentLastDate"`
}
