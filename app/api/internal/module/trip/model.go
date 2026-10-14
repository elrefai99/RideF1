package trip

import (
	"os/user"

	"github.com/google/uuid"
)

// Ride Management
type RideStatus string

const (
	RideSearch RideStatus = "SEARCH"
	RideAccept RideStatus = "ACCEPT"
	RideCancel RideStatus = "CANCEL"
)

type RideRequest struct {
	ID              uint       `gorm:"primaryKey"`
	Uuid            uuid.UUID  `gorm:"varchar(255);not null;index:idx_riderequest_uuid_status,priority:1;index:idx_riderequest_uuid_rider_id,priority:1;"`
	Rider_id        uint       `gorm:"null;index:idx_riderequest_uuid_rider_id,priority:2;"`
	Pickup_lat      string     `gorm:"varchar(100);not null"`
	Pickup_lng      string     `gorm:"varchar(100);not null"`
	destination_lat string     `gorm:"varchar(100);not null"`
	destination_lng string     `gorm:"varchar(100);not null"`
	Status          RideStatus `gorm:"varchar(50);index:idx_riderequest_uuid_status,priority:1;default:'SEARCH'"`

	Rider user.User `gorm:"foreignKey:Rider_id;references:ID"`
}
