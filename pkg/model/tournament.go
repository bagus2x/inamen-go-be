package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateTourRequest struct {
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	Game              string   `json:"game"`
	Mode              string   `json:"mode"`
	Presence          string   `json:"presence"`
	Visibilty         string   `json:"visiblity"`
	TotalParticipants int      `json:"totalParticipants"`
	TotalTeamMembers  int      `json:"totalTeamMembers"`
	Platforms         []string `json:"platforms"`
	Date              TourDate `json:"date"`
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
	RegistrationStartDate int64 `json:"registrationStartDate"`
	RegistrationLastDate  int64 `json:"registrationLastDate"`
	TournamentStartDate   int64 `json:"tournamentStartDate"`
	TournamentLastDate    int64 `json:"tournamentLastDate"`
}

type TourHost struct {
	ID       string `bson:"_id"`
	Username string `bson:"username"`
}

type UpdateTourRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Published   bool     `json:"published"`
	Visibilty   string   `json:"visiblity"`
	Date        TourDate `json:"date"`
}

type UpdateTourResponse FetchTourResponse
