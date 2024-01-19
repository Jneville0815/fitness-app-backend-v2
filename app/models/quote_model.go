package models

import (
	"github.com/google/uuid"
)

type Quote struct {
	UserID    uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid"`
	Source    string    `db:"source" json:"source" validate:"required,lte=255"`
	QuoteText string    `db:"quote" json:"quote" validate:"required"`
	NumViews  int       `db:"num_views" json:"num_views" validate:"gte=0"`
}
