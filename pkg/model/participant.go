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
	ID       string `json:"_id" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type CreateParticipantRequest struct {
	TournamentID primitive.ObjectID `json:"tournamentID" validate:"required"`
	TeamName     string             `json:"teamName" validate:"required,lte=100"`
	Players      []Player           `json:"players" validate:"required,dive,required"`
	Description  string             `json:"description" validate:"lte=255"`
}

type CreateParticipantResponse Participant

type FetchParticipantResponse Participant

type UpdateParticipantRequest struct {
	TeamName    string `json:"teamName" validate:"required,lte=100"`
	Description string `json:"description" validate:"lte=255"`
}

type UpdateParticipantResponse Participant
