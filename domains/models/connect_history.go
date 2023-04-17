package models

import "time"

type ConnectHistory struct {
	ID           string    `json:"id" bson:"_id,omitempty"`
	ConnectionID string    `json:"connection_id" bson:"connection_id"`
	ConnectedAt  time.Time `json:"connected_at" bson:"connected_at"`
}
