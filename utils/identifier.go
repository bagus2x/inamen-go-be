package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func GenerateID() primitive.ObjectID {
	return primitive.NewObjectID()
}
