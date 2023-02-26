package destinationService

import "time"

type Destination struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Photo       string    `json:"photo" db:"photo"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type DestinationCreateRequest struct {
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Photo       string `json:"photo" db:"photo"`
}

type DestinationUpdateRequest struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Photo       string `json:"photo" db:"photo"`
}

type DestinationDeleteRequest struct {
	ID int `json:"id" db:"id"`
}
