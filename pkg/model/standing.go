package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Standing struct {
	ID           primitive.ObjectID `json:"_id"`
	TournamentID primitive.ObjectID `json:"tournamentID"`
	Name         string             `json:"name"`
	Spotlight    bool               `json:"spotlight"`
	Schema       StandingSchema     `json:"schema"`
	CreatedAt    int64              `json:"createdAt"`
	UpdatedAt    int64              `json:"updatedAt"`
}

type StandingSchema struct {
	Columns []string        `json:"columns" validate:"required,dive,required"`
	Data    [][]interface{} `json:"data" validate:"required,dive,required"`
}

type CreateStandingRequest struct {
	TournamentID primitive.ObjectID `json:"tournamentID" validate:"required"`
	Name         string             `json:"name" validate:"required,lte=100"`
	Spotlight    bool               `json:"spotlight" validate:"required"`
	Schema       StandingSchema     `json:"schema" validate:"required"`
}

type CreateStandingResponse Standing

type FetchStandingResponse Standing

type UpdateStandingRequest struct {
	Name      string         `json:"name" validate:"required,lte=100"`
	Spotlight bool           `json:"spotlight" validate:"required"`
	Schema    StandingSchema `json:"schema" validate:"required"`
}

type UpdateStandingResponse Standing
