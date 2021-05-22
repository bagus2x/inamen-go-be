package main

import (
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Manusia struct {
	Name interface{}
	No   interface{}
}

func main() {
	type test struct {
		ID primitive.ObjectID `json:"_id"`
	}

	var t test
	err := json.Unmarshal([]byte(`{"_id": "60a7c78aba454118d941fa85"}`), &t)

	fmt.Println(err)

	fmt.Println(t.ID)
}
