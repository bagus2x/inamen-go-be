package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChatRoom struct {
	ID           primitive.ObjectID `json:"_id"`
	TournamentID primitive.ObjectID `json:"tournamentID"`
	Messages     []Message          `json:"messages"`
}

type Message struct {
	ID     primitive.ObjectID `json:"_id"`
	Sender string             `json:"sender"`
	Body   string             `json:"body"`
}
