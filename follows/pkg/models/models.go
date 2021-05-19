package models

import (
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type FollowRequest struct {
	Id uuid.UUID         `gorm:"primaryKey"`
	Following   uuid.UUID         `gorm:"primaryKey"`
	Follower   uuid.UUID         `gorm:"primaryKey"`
	Accepted  bool

}
