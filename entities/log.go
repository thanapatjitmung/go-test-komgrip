package entities

import (
	"time"
)

type Log struct {
	ID        string    `bson:"_id,omitempty"`
	Action    string    `bson:"action"`
	Timestamp time.Time `bson:"timestamp"`
	Data      string    `bson:"data"`
	OldData   string    `bson:"old_data"`
	ItemId    string    `bson:"item_id"`
}
