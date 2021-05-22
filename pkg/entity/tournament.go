package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tournament struct {
	ID                primitive.ObjectID `bson:"_id"`
	Name              string             `bson:"name"`
	Description       string             `bson:"description"`
	Host              Host               `bson:"host"`
	Game              string             `bson:"game"`
	Mode              string             `bson:"mode"`
	Presence          string             `bson:"presence"`
	Published         bool               `bson:"published"`
	Visibilty         string             `bson:"visiblity"`
	TotalParticipants int                `bson:"totalParticipants"`
	TotalTeamMembers  int                `bson:"totalTeamMembers"`
	Platforms         []string           `bson:"platforms"`
	Date              TournamentDate     `bson:"date"`
	CreatedAt         int64              `bson:"createdAt"`
	UpdatedAt         int64              `bson:"updatedAt"`
}

type Host struct {
	ID       string `bson:"_id"`
	Username string `bson:"username"`
}

type TournamentDate struct {
	RegistrationStartDate int64 `bson:"registrationStartDate"`
	RegistrationLastDate  int64 `bson:"registrationLastDate"`
	TournamentStartDate   int64 `bson:"tournamentStartDate"`
	TournamentLastDate    int64 `bson:"tournamentLastDate"`
}
