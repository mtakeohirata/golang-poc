package models

import (
	"github.com/google/uuid"
)

var artistMocked = Artist{
	ID:   uuid.New().String(),
	Name: "TakeoMock",
}
