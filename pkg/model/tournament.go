package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateTourRequest struct {
	Name              string   `json:"name" validate:"required,gt=5,lte=100"`
	Description       string   `json:"description" validate:"lte=255"`
	Game              string   `json:"game" validate:"required"`
	Mode              string   `json:"mode" validate:"required"`
	Presence          string   `json:"presence" validate:"required"`
	Visibilty         string   `json:"visibility" validate:"required"`
	TotalParticipants int      `json:"totalParticipants" validate:"required"`
	TotalTeamMembers  int      `json:"totalTeamMembers" validate:"required"`
	Platforms         []string `json:"platforms" validate:"required,dive,required"`
	Date              TourDate `json:"date" validate:"required"`
}

type CreateTourResponse FetchTourResponse

type FetchTourResponse struct {
	ID                primitive.ObjectID `json:"_id"`
	Name              string             `json:"name"`
	Description       string             `json:"description"`
	Host              TourHost           `json:"host"`
	Game              string             `json:"game"`
	Mode              string             `json:"mode"`
	Presence          string             `json:"presence"`
	Published         bool               `json:"published"`
	Visibilty         string             `json:"visiblity"`
	TotalParticipants int                `json:"totalParticipants"`
	TotalTeamMembers  int                `json:"totalTeamMembers"`
	Platforms         []string           `json:"platforms"`
	Date              TourDate           `json:"date"`
	CreatedAt         int64              `json:"createdAt"`
	UpdatedAt         int64              `json:"updatedAt"`
}

type FetchToursResponse []*FetchTourResponse

type TourDate struct {
	RegistrationStartDate int64 `json:"registrationStartDate" validate:"required"`
	RegistrationLastDate  int64 `json:"registrationLastDate" validate:"required"`
	TournamentStartDate   int64 `json:"tournamentStartDate" validate:"required"`
	TournamentLastDate    int64 `json:"tournamentLastDate" validate:"required"`
}

type TourHost struct {
	ID       string `bson:"_id"`
	Username string `bson:"username"`
}

type UpdateTourRequest struct {
	Name        string   `json:"name" validate:"required,gt=5,lte=100"`
	Description string   `json:"description"`
	Published   bool     `json:"published"`
	Visibilty   string   `json:"visiblity"`
	Date        TourDate `json:"date"`
}

type UpdateTourResponse FetchTourResponse

type FetchTourNameResponse struct {
	ID   primitive.ObjectID `json:"_id"`
	Name string             `json:"name"`
}
