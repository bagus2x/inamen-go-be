package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Participant struct {
	ID           primitive.ObjectID `json:"_id"`
	TournamentID primitive.ObjectID `json:"tournamentID"`
	TeamName     string             `json:"teamName"`
	Players      []Player           `json:"players"`
	Description  string             `json:"description"`
	CreatedAt    int64              `json:"createdAt"`
	UpdatedAt    int64              `json:"updatedAt"`
}

type Player struct {
	ID       string `json:"_id"`
	Username string `json:"username"`
}

type CreateParticipantRequest struct {
	TournamentID primitive.ObjectID `json:"tournamentID"`
	TeamName     string             `json:"teamName"`
	Players      []Player           `json:"players"`
	Description  string             `json:"description"`
}

type CreateParticipantResponse Participant

type FetchParticipantResponse Participant

type UpdateParticipantRequest struct {
	TeamName    string `json:"teamName"`
	Description string `json:"description"`
}

type UpdateParticipantResponse Participant
